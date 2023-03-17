package axty

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

//update
type MqttNotifyEvent2 struct {
	MqttNotifyEvent
	Serialnumber string `json:"serialnumber"`
	Modelname    string `json:"modelname"`
}
