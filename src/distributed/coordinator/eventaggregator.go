package coordinator

import "time"

type EventRaiser interface {
	AddListener(eventName string, f func(interface{}))
}

type EventAggrigator struct {
	listeners map[string][]func(interface{})
}

func NewEventAggregator() *EventAggrigator {
	ea := EventAggrigator{
		listeners: make(map[string][]func(interface{})),
	}

	return &ea
}

func (ea *EventAggrigator) AddListener(name string, f func(interface{})) {
	ea.listeners[name] = append(ea.listeners[name], f)
}

func (ea *EventAggrigator) PublishEvent(name string, eventData interface{}) {
	if ea.listeners[name] != nil {
		for _, r := range ea.listeners[name] {
			r(eventData)
		}
	}
}

type EventData struct {
	Name      string
	Value     float64
	Timestamp time.Time
}
