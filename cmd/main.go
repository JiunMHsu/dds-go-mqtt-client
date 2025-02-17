package main

import (
	"fmt"
	"sync"

	"github.com/JiunMHsu/dds-go-mqtt-client/config"
	"github.com/JiunMHsu/dds-go-mqtt-client/internal/console"
	mqttClient "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"
	"github.com/JiunMHsu/dds-go-mqtt-client/internal/publisher"
)

func main() {

	topics := config.Env.HeladeraTopics

	if config.Env.PublishTemperature {
		for _, topic := range topics {
			fmt.Println("Publishing temperature for topic: ", topic)
			client := mqttClient.NewClient(topic)
			go publisher.StartTemperaturePublishingRoutine(client)
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go console.RunConsole(&wg)
	wg.Wait()
}
