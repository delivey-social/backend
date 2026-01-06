package notificacoes

import (
	"log/slog"

	"comida.app/src/infra"
)

type LoggerChannel struct {
	logger *slog.Logger
}

func NewLoggerChannel() Channel {
	return &LoggerChannel{
		logger: slog.Default(),
	}
}

func (c *LoggerChannel) Subscriptions() map[infra.EventType]func(infra.Event) {
	res := make(map[infra.EventType]func(infra.Event))

	res[infra.OrderCreated] = c.log
	res[infra.OrderReadyForDelivery] = c.log
	res[infra.OrderInDelivery] = c.log
	res[infra.OrderDelivered] = c.log

	return res
}

func (c *LoggerChannel) log(evt infra.Event) {
	c.logger.Info("[NOTIFICATION]", "type", evt.Type, "id", evt.Payload.OrderID)
}
