package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/hooks/auth"
	"github.com/mochi-co/mqtt/v2/listeners"
	ini "gopkg.in/ini.v1"
)

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("./conf/conf.ini")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}

func main() {
	// Create the new MQTT Server.
	server := mqtt.New(nil)

	// Allow all connections.
	_ = server.AddHook(new(auth.AllowHook), nil)

	host := cfg.Section("mqttsrv").Key("host").String()
	port, _ := cfg.Section("mqttsrv").Key("port").Int()
	hostport := fmt.Sprintf("%s:%d", host, port)

	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP("axmqttsrv", hostport, nil)
	err := server.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Second * 1)
	}
}
