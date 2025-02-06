package routine

import (
	"math"
	"math/rand/v2"
	"strconv"
	"time"

	mqttClient "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"
)

func PublishTemperatureFor(client mqttClient.MqttClient) {
	for {
		temprature := randomFloat(-2.0, 2.0)
		message := "TEMPERATURA " + strconv.FormatFloat(temprature, 'f', 2, 64)
		// fmt.Printf("Publishing message for %s: %s\n", client.Topic, message)
		client.Publish(message)

		time.Sleep(5 * time.Second)
	}
}

func randomFloat(min, max float64) float64 {
	return math.Round((rand.Float64()*(max-min)+min)*100) / 100
}
