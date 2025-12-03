package notificacoes

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"comida.app/src/infra"
)

type EmailChannel struct {
	host     string
	port     string
	sender   string
	password string
	auth     smtp.Auth
}

func NewEmailChannel() Channel {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	sender := os.Getenv("SMTP_SENDER")
	password := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", sender, password, host)

	if host == "" || port == "" || sender == "" || password == "" {
		log.Panicf("failed to create email channel: missing required env variables")
	}

	return &EmailChannel{
		host,
		port,
		sender,
		password,
		auth,
	}
}

func (c *EmailChannel) Subscriptions() map[infra.EventType]func(infra.Event) {
	res := make(map[infra.EventType]func(infra.Event))

	res[infra.OrderCreated] = c.onOrderCreated
	res[infra.OrderReadyForDelivery] = c.onOrderReadyForDelivery
	res[infra.OrderInDelivery] = c.onOrderInDeliveryRoute
	res[infra.OrderDelivered] = c.onOrderInDelivered

	return res
}

func (c *EmailChannel) onOrderCreated(evt infra.Event) {
	payload := evt.Payload.(infra.OrderCreatedPayload)

	c.sendEmail(SendEmail{
		receivers: []string{"admin@comida.app.br"},
		subject:   "Order created",
		content:   payload.OrderID.String(),
	})

	log.Println("Order created email sent for order", payload.OrderID)
}

func (c *EmailChannel) onOrderReadyForDelivery(evt infra.Event) {
	payload := evt.Payload.(infra.OrderUpdatedPayload)

	c.sendEmail(SendEmail{
		receivers: []string{"admin@comida.app.br"},
		subject:   "Order ready for delivery",
		content:   payload.OrderID.String(),
	})

	log.Println("Order ready for delivery email sent for order", payload.OrderID)
}

func (c *EmailChannel) onOrderInDeliveryRoute(evt infra.Event) {
	payload := evt.Payload.(infra.OrderUpdatedPayload)

	c.sendEmail(SendEmail{
		receivers: []string{"admin@comida.app.br"},
		subject:   "Order in delivery route",
		content:   payload.OrderID.String(),
	})

	log.Println("Order in delivery route email sent for order", payload.OrderID)
}

func (c *EmailChannel) onOrderInDelivered(evt infra.Event) {
	payload := evt.Payload.(infra.OrderUpdatedPayload)

	c.sendEmail(SendEmail{
		receivers: []string{"admin@comida.app.br"},
		subject:   "Order delivered",
		content:   payload.OrderID.String(),
	})

	log.Println("Order delivered email sent for order", payload.OrderID)
}

type SendEmail struct {
	receivers []string
	subject   string
	content   string
}

func (c *EmailChannel) sendEmail(data SendEmail) {
	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n",
		data.receivers[0], data.subject, data.content,
	))

	uri := c.host + ":" + c.port
	err := smtp.SendMail(uri, c.auth, c.sender, data.receivers, msg)

	if err != nil {
		log.Panicf("Failed to send email: %v", err)
	}
}
