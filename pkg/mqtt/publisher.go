package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/labstack/gommon/log"
)

type Publisher struct {
	client mqtt.Client
}

func NewPublisher(client mqtt.Client) *Publisher {
	return &Publisher{client: client}
}

func (p *Publisher) Publish(topic, msg string) {
	token := p.client.Publish(topic, 1, false, msg)
	token.Wait()
	if token.Error() != nil {
		log.Error("Error in Token", token.Error())
	}
}

func (p *Publisher) PublishSetLights(time string) {
	p.Publish("smart-office/lightSetTime", time)
}
