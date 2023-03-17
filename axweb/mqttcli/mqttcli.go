package axmqttcli

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myapp/axty"
	"myapp/ebus"
	"myapp/utils"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	l4g "github.com/jeanphorn/log4go"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type AxMqttCli struct {
	Host     string
	Port     int
	User     string
	Passwd   string
	Reconn   int
	Ca       string
	Cert     string
	Prikey   string
	Security bool
	Topic    string
	Log      l4g.Logger
	Mux      *sync.RWMutex
	client   mqtt.Client
}

func New() *AxMqttCli {
	cfg := utils.Cfg
	host := cfg.Section("mqttcli").Key("host").String()
	port, _ := cfg.Section("mqttcli").Key("port").Int()
	user := cfg.Section("mqttcli").Key("user").String()
	passwd := cfg.Section("mqttcli").Key("passwd").String()
	topic := cfg.Section("mqttcli").Key("topic").String()

	security := cfg.Section("mqttcli").Key("security").MustBool()
	ca := cfg.Section("mqttcli").Key("ca").String()
	cert := cfg.Section("mqttcli").Key("cert").String()
	prikey := cfg.Section("mqttcli").Key("prikey").String()
	enable := cfg.Section("mqttcli").Key("enable").MustBool()
	reconn, _ := cfg.Section("mqttcli").Key("reconn").Int()

	axmqttcli := &AxMqttCli{
		Host:     host,
		Port:     port,
		User:     user,
		Passwd:   passwd,
		Reconn:   reconn,
		Security: security,
		Ca:       ca,
		Cert:     cert,
		Prikey:   prikey,
		Topic:    topic,
		Mux:      &sync.RWMutex{},
	}

	if enable {
		go axmqttcli.Start()
	}

	return axmqttcli
}

func (slf *AxMqttCli) Start() {
	fmt.Println("AxMqttCli Start")

	opts := mqtt.NewClientOptions()

	if slf.Security {
		tlsconfig, err := slf.NewTLSConfig()
		if err != nil {
			slf.Log.Debug("failed to create TLS configuration: %v", err)
		}

		//tls://a2vt5x3ve8uw69-ats.iot.ap-northeast-1.amazonaws.com:8883
		opts.AddBroker(fmt.Sprintf("tls://%s:%d", slf.Host, slf.Port))
		opts.SetClientID(fmt.Sprintf("go_mqtt_%d", time.Now().Local().UnixMilli())).SetTLSConfig(tlsconfig)

	} else {
		opts.AddBroker(fmt.Sprintf("tcp://%s:%d", slf.Host, slf.Port))
		opts.SetClientID(fmt.Sprintf("go_mqtt_%d", time.Now().Local().UnixMilli()))
	}

	//auto reconnect
	opts.SetPingTimeout(time.Duration(slf.Reconn) * time.Second)
	opts.SetKeepAlive(time.Duration(slf.Reconn) * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(time.Duration(slf.Reconn) * time.Second)

	//set user,passwd
	opts.SetUsername(slf.User)
	opts.SetPassword(slf.Passwd)

	//recieve handler
	opts.SetDefaultPublishHandler(slf.msgPubHandler)

	//conn status
	opts.OnConnect = slf.connHandler

	//conn err
	opts.OnConnectionLost = slf.connLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("token err connect:", token.Error())
		//panic(token.Error())
	}

	slf.client = client

	token := client.Subscribe(slf.Topic, 1, nil)
	token.Wait()
}

// tls
func (slf *AxMqttCli) NewTLSConfig() (config *tls.Config, err error) {
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(slf.Ca)
	if err != nil {
		return
	}
	certpool.AppendCertsFromPEM(pemCerts)

	cert, err := tls.LoadX509KeyPair(slf.Cert, slf.Prikey)
	if err != nil {
		return
	}

	config = &tls.Config{
		RootCAs:      certpool,
		ClientAuth:   tls.NoClientCert,
		ClientCAs:    nil,
		Certificates: []tls.Certificate{cert},
	}
	return config, err
}

type MqttPayloadInfo struct {
	Serialnumber string
	Modelname    string
	Payload      axty.MqttNotifyEvent
}

func (me *AxMqttCli) msgPubHandler(client mqtt.Client, msg mqtt.Message) {
	//fmt.Println(msg.Topic(), string(msg.Payload()))

	var payload axty.MqttNotifyEvent2
	json.Unmarshal(msg.Payload(), &payload)

	topicarr := strings.Split(msg.Topic(), "/")
	payload.Serialnumber = time.Now().String()
	payload.Modelname = "none"

	if len(topicarr) >= 2 {
		payload.Serialnumber = topicarr[len(topicarr)-1]
		payload.Modelname = topicarr[len(topicarr)-2]
	}

	//send data
	ebus.EvBus.Publish("mqdata", payload)
}

func (me *AxMqttCli) connHandler(client mqtt.Client) {
	fmt.Println("mqttConn success")
	client.Subscribe(me.Topic, 1, nil)
}

func (me *AxMqttCli) connLostHandler(client mqtt.Client, err error) {
	fmt.Println("mqtt connect err:", err)
}
