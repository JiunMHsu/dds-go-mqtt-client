package publisher

import (
	"fmt"
	"time"

	"github.com/JiunMHsu/dds-go-mqtt-client/config"
	mqtt "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"
	"github.com/JiunMHsu/dds-go-mqtt-client/utils"
)

func PublishTemperatureFor(topic string, temprature float64) {
	message := fmt.Sprintf("TEMPERATURA %.2f", temprature)
	client := mqtt.GetClient(topic)
	client.Publish(message)
}

func PublishRandomTemperatureFor(topic string) {
	temp := utils.RandomFloat(-2.0, 2.0)
	PublishTemperatureFor(topic, temp)
}

func StartTemperaturePublishingRoutine(topic string) {
	for {
		PublishRandomTemperatureFor(topic)
		time.Sleep(time.Duration(config.Env.PublishTemperatureInterval) * time.Second)
	}
}
