package axty

import "github.com/gorilla/websocket"

// for mqtt
type MqttNotifyEvent struct {
	Type      string `json:"type"`
	Condition string `json:"condition"`
	Object    []struct {
		ID    string    `json:"id"`
		Value []float64 `json:"value"`
		Unit  string    `json:"unit"`
	} `json:"object"`
	Class  string `json:"class"`
	ID     string `json:"id"`
	Source string `json:"source"`
}

// for mqtt
type MqttNotifyEvent2 struct {
	MqttNotifyEvent
	Serialnumber string `json:"serialnumber"`
	Modelname    string `json:"modelname"`
}

// for ws
type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Multichannel []struct {
		Channel string `json:"channel"`
	} `json:"multichannel"`
	Ws *websocket.Conn
}

// for ws
type MulitySubscribe struct {
	ID           string `json:"id"`
	Action       string `json:"action"`
	Multichannel []struct {
		Channel string `json:"channel"`
	} `json:"multichannel"`
}

// for ws
type Send struct {
	Action       string `json:"action"`
	Multichannel []struct {
		Channel string `json:"channel"`
	} `json:"multichannel"`
	Data string `json:"data"`
}
