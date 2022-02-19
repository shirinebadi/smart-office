package handler

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	gomqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/labstack/gommon/log"
	"github.com/shirinebadi/smart-office/internal/config"
	"github.com/shirinebadi/smart-office/internal/model"
	"github.com/shirinebadi/smart-office/pkg/mqtt"
)

type MQTTHandler struct {
	Config config.Config
	UserI  model.UserInterface
	P      mqtt.Publisher
}

func (mh *MQTTHandler) UserActivity(_ gomqtt.Client, message gomqtt.Message) {
	fmt.Println("hi")
	user, err := mh.UserI.UserSearch(hex.EncodeToString(message.Payload()))
	if err != nil {
		log.Error(err)
	}

	url := "http://127.0.0.1" + mh.Config.CentralServer.Address + "/api/office/localserver"
	fmt.Print(url)
	values := map[string]int{"id": user.Id}
	jsonValue, _ := json.Marshal(values)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	result := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result["Light"])

	//var lightResponse response.LightResponse
	if resp.StatusCode != 200 {
		log.Error("Error in Response", err)
	}

}
