package routine

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
	"time"

	mqttClient "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"
)

func PublishTemperatureFor(client mqttClient.MqttClient) {
	for {
		temprature := randomFloat(-3.0, 3.0)
		fmt.Printf("publish for %s: %.2f\n", client.Topic, temprature)
		message := "TEMPERATURA " + strconv.FormatFloat(temprature, 'f', 2, 64)
		client.Publish(message)

		time.Sleep(2 * time.Second)
	}
}

func randomFloat(min, max float64) float64 {
	return math.Round(rand.Float64()*(max-min)*100) / 100
}
