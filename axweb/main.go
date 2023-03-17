package main

import (
	"fmt"
	"html/template"
	"log"
	"myapp/ebus"
	"myapp/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	axweb *AxWeb
	wsmap = make(map[string]*websocket.Conn)
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

	//dataEvent
	dataEvent := make(chan ebus.DataEvent)
	//register
	ebus.EvBus.Subscribe("mqdata", dataEvent)

	go func() {
		for {
			select {
			case ev := <-dataEvent:
				fmt.Println(ev.Data)

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
	wsmap[wsid] = ws

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(wsmap, wsid)
			break
		}
		fmt.Println(string(msg))
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

	axweb = New(host, port, wwwroot, mqttchann, ssl, cert, prikey)
	axweb.Start()
}
