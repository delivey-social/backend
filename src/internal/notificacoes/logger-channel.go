package notificacoes

import (
	"fmt"

	"comida.app/src/infra"
)

type LoggerChannel struct{}

func NewLoggerChannel() Channel {
	return &LoggerChannel{}
}

func (c *LoggerChannel) Subscriptions() map[infra.EventType]func(infra.Event) {
	res := make(map[infra.EventType]func(infra.Event))

	res[infra.OrderCreated] = c.onOrderCreated
	res[infra.OrderReadyForDelivery] = c.onOrderReadyForDelivery

	return res
}

func (c *LoggerChannel) onOrderCreated(evt infra.Event) {
	payload := evt.Payload.(infra.OrderCreatedPayload)

	fmt.Println("ORDER CREATED", payload.OrderID)
}

func (c *LoggerChannel) onOrderReadyForDelivery(evt infra.Event) {
	payload := evt.Payload.(infra.OrderUpdatedPayload)

	fmt.Println("ORDER READY FOR DELIVERY", payload)
}
