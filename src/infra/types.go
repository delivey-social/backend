package infra

import "github.com/google/uuid"

type EventType string

const (
	OrderCreated          EventType = "ORDER CREATED"
	OrderReadyForDelivery EventType = "ORDER READY FOR DELIVERY"
	OrderInDelivery       EventType = "ORDER IN DELIVERY"
	OrderDelivered        EventType = "ORDER DELIVERED"
)

type Event struct {
	Type    EventType
	Payload OrderPayload
}

type OrderPayload struct {
	OrderID uuid.UUID
}
