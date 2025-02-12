package publish

import (
	"strconv"

	mqtt "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"
	"github.com/JiunMHsu/dds-go-mqtt-client/utils"
)

func PublishTemperatureFor(client mqtt.MqttClient, temprature float64) {
	message := "TEMPERATURA " + strconv.FormatFloat(temprature, 'f', 2, 64)
	client.Publish(message)
}

func PublishRandomTemperatureFor(client mqtt.MqttClient) {
	temp := utils.RandomFloat(-2.0, 2.0)
	PublishTemperatureFor(client, temp)
}

func StartTemperaturePublishingRoutine(client mqtt.MqttClient) {
	for {
		PublishRandomTemperatureFor(client)
	}
}
