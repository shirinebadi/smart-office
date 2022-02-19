package mqtt_worker

import (
	"log"

	"github.com/shirinebadi/smart-office/internal/config"
	"github.com/shirinebadi/smart-office/internal/handler"
	"github.com/shirinebadi/smart-office/pkg/data/db"
	"github.com/shirinebadi/smart-office/pkg/mqtt"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	infinity := make(chan bool)
	client := mqtt.CreateMQTTConnection(cfg.MQTT)
	subscriber := mqtt.NewSubscriber(client)
	publisher := mqtt.NewPublisher(client)

	myDB, err := db.Init()
	if err != nil {
		log.Fatal("Failed to setup db: ", err.Error())
	}

	userI := &db.Mydb{DB: myDB}

	mqttHandler := &handler.MQTTHandler{Config: cfg, UserI: userI, P: *publisher}
	subscriber.Subscribe("smart-office/lightSetTime", mqttHandler.UserActivity)
	<-infinity
}

// Register registers mqtt-subscriber command for smart-home binary.
func Register(root *cobra.Command, cfg config.Config) {
	publish := &cobra.Command{
		Use:   "mqtt",
		Short: "mqtt worker",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(publish)
}
