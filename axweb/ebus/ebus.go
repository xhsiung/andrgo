package ebus

import (
	"sync"
)

/*
//eventbus
eventMeasure := make(chan DataEvent)

//register
EvBus.Subscribe("EventMeasure", eventMeasure)

//send
EvBus.Publish("EventMeasure", "TEST Event")

//recive
for {
	select {
	case ev := <-eventMeasure:
		fmt.Println("")
	default:
		fmt.Println("")
	}
}

*/
var EvBus = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel is a channel which can accept an DataEvent
type DataChannel chan DataEvent

// DataChannelSlice is a slice of DataChannels
type DataChannelSlice []DataChannel

// EventBus stores the information about subscribers interested for a particular topic
type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	defer eb.rm.RUnlock()

	if chans, found := eb.subscribers[topic]; found {
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}

}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	defer eb.rm.Unlock()

	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
}

func (eb *EventBus) UnSubscribe(topic string) {
	eb.rm.Lock()
	defer eb.rm.Unlock()

	if _, found := eb.subscribers[topic]; found {
		delete(eb.subscribers, topic)
	}
}
