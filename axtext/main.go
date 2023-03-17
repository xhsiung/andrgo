package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	axweb *AxWeb
	wsmap = make(map[string]*websocket.Conn)
)

type AxWeb struct {
	Host         string
	Port         int
	Router       *gin.Engine
	DocumentRoot string
}

func New(host string, port int, docroot string) *AxWeb {
	axweb = &AxWeb{
		Host:         host,
		Port:         port,
		DocumentRoot: docroot,
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
	slf.Router.Static("/js", slf.DocumentRoot+"/js")
	slf.Router.Static("/css", slf.DocumentRoot+"/css")
	slf.Router.LoadHTMLGlob(slf.DocumentRoot + "/*.html")

	slf.Router.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "login.html", gin.H{
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	slf.Router.GET("/ws", slf.WsContext)
	slf.Router.Run(fmt.Sprintf(":%d", slf.Port))
}

//ws recieve message
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
	axweb = New("0.0.0.0", 8085, "/data/ax/www")
	axweb.Start()
}
