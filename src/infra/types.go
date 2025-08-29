package infra

import "github.com/google/uuid"

type EventType int

const (
	OrderCreated EventType = iota
	OrderReadyForDelivery
	OrderInDelivery
	OrderDelivered
)

type Event struct {
	Type    EventType
	Payload any
}

type OrderCreatedPayload struct {
	OrderID uuid.UUID
}

type OrderUpdatedPayload struct {
	OrderID uuid.UUID
}
