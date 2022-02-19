package mqtt

import (
	"github.com/labstack/gommon/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscriber struct {
	client mqtt.Client
}

func NewSubscriber(client mqtt.Client) *Subscriber {
	return &Subscriber{
		client: client,
	}
}

func (s *Subscriber) Subscribe(topic string, callback mqtt.MessageHandler) {
	token := s.client.Subscribe(topic, 0, callback)
	token.Wait()
	if token.Error() != nil {
		log.Error(token.Error())
	}
}
