package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/shirinebadi/smart-office/internal/config"
)

func CreateMQTTConnection(cfg config.MQTT) mqtt.Client {
	server := fmt.Sprintf("tcp://%s:1883", cfg.Host)
	fmt.Println(server)
	opts := mqtt.NewClientOptions().AddBroker(server)
	opts.SetClientID("go_mqtt_client")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}
