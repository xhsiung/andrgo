package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"myapp/axty"
	"myapp/ebus"
	axmqttcli "myapp/mqttcli"
	"myapp/utils"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	axweb *AxWeb
	wsmap = make(map[string]*axty.User)
)

func init() {
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}

type AxWeb struct {
	Host         string
	Port         int
	Router       *gin.Engine
	DocumentRoot string
	Mqttchann    string
	Ssl          bool
	Cert         string
	Privatekey   string
	Mux          *sync.RWMutex
}

func New(host string, port int, docroot string, mqttchann string, ssl bool, cert string, privateky string) *AxWeb {
	axweb = &AxWeb{
		Host:         host,
		Port:         port,
		DocumentRoot: docroot,
		Mqttchann:    mqttchann,
		Ssl:          ssl,
		Cert:         cert,
		Privatekey:   privateky,
		Mux:          &sync.RWMutex{},
	}

	return axweb
}

func (slf *AxWeb) Start() {
	slf.Router = gin.Default()
	//nomsg out
	gin.SetMode(gin.ReleaseMode)

	tpl := template.Must(template.New("").Parse(slf.DocumentRoot + "/*.html"))
	slf.Router.SetHTMLTemplate(tpl)
	slf.Router.Static("/static", slf.DocumentRoot+"/static")
	slf.Router.Static("/assets", slf.DocumentRoot+"/assets")
	slf.Router.Static("/js", slf.DocumentRoot+"/js")
	slf.Router.Static("/css", slf.DocumentRoot+"/css")
	slf.Router.LoadHTMLGlob(slf.DocumentRoot + "/*.html")

	slf.Router.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "login.html", gin.H{
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	//ws
	slf.Router.GET("/ws", slf.WsContext)

	//dataEvent
	dataEvent := make(chan ebus.DataEvent)
	//register
	ebus.EvBus.Subscribe("mqdata", dataEvent)

	go func() {
		for {
			select {
			case ev := <-dataEvent:
				//fmt.Println(ev.Data)
				slf.SendWs(slf.Mqttchann, ev.Data)
			default:
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()

	if slf.Ssl {
		slf.Router.RunTLS(fmt.Sprintf(":%d", slf.Port), slf.Cert, slf.Privatekey)
	} else {
		slf.Router.Run(fmt.Sprintf(":%d", slf.Port))
	}
}

func (slf *AxWeb) WsContext(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	//action
	wsid := fmt.Sprintf("%x", &ws)
	jobj := make(map[string]interface{})

	//wsmap[wsid] = ws

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(wsmap, wsid)
			break
		}

		json.Unmarshal(msg, &jobj)
		action := jobj["action"].(string)

		//fmt.Println(string(msg))
		if action == "subscribe" {
			mulitySubscribe := axty.MulitySubscribe{}
			json.Unmarshal(msg, &mulitySubscribe)

			if user, ok := wsmap[wsid]; ok {
				for _, v := range mulitySubscribe.Multichannel {
					if !uniqueList(v.Channel, user) {
						user.Multichannel = append(user.Multichannel, v)
					}
				}

			} else {
				wsmap[wsid] = &axty.User{
					ID:           wsid,
					Multichannel: mulitySubscribe.Multichannel,
					Ws:           ws,
				}
			}

			continue
		}

		if action == "send" {
			sendObj := axty.Send{}
			json.Unmarshal(msg, &sendObj)

			for _, v1 := range sendObj.Multichannel {
				slf.SendWs(v1.Channel, []byte(sendObj.Data))
			}
			continue
		}

	}
}

func uniqueList(key string, user *axty.User) bool {
	for _, v := range user.Multichannel {
		if key == v.Channel {
			return true
		}
	}
	return false
}

func (slf *AxWeb) SendWs(channel string, data interface{}) {
	xdata, _ := json.Marshal(data)
	for _, v := range wsmap {
		for _, v2 := range v.Multichannel {
			if v2.Channel == channel {
				slf.Mux.Lock()
				v.Ws.WriteMessage(1, xdata)
				slf.Mux.Unlock()
			}
		}
	}
}

func main() {
	cfg := utils.Cfg

	host := cfg.Section("axweb").Key("host").String()
	port, _ := cfg.Section("axweb").Key("port").Int()
	wwwroot := cfg.Section("axweb").Key("www").String()

	ssl, _ := cfg.Section("axweb").Key("ssl").Bool()
	mqttchann := cfg.Section("mqttcli").Key("mqttchann").String()
	cert := cfg.Section("axweb").Key("cert").String()
	prikey := cfg.Section("axweb").Key("prikey").String()

	_ = axmqttcli.New()

	axweb = New(host, port, wwwroot, mqttchann, ssl, cert, prikey)
	axweb.Start()
}
