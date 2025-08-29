package notificacoes

import (
	"comida.app/src/infra"
)

type EventSubscriber interface {
	Subscribe(event infra.EventType, handler func(infra.Event))
}

type Channel interface {
	Subscriptions() map[infra.EventType]func(infra.Event)
}

func NewNotificacoesService(subscriber EventSubscriber)  {
	channels := []Channel{
		NewLoggerChannel(),
	}

	for _, ch := range channels {
		for evt, handler := range ch.Subscriptions() {
			subscriber.Subscribe(evt, handler)
		}
	}
}
