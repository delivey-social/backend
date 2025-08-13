package eventbus

import "comida.app/src/infra"

type EventBus struct {
	handlers map[infra.Event][]func(infra.Event)
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[infra.Event][]func(infra.Event)),
	}
}

func (eb *EventBus) Subscribe(event infra.Event, handler func(infra.Event)) {
	eb.handlers[event] = append(eb.handlers[event], handler)
}

func (eb *EventBus) Publish(event infra.Event) {
	for _, handler := range eb.handlers[event] {
		handler(event)
	}
}