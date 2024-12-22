package client

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Client mqtt.Client = nil

func GetClient() mqtt.Client {
	if Client != nil {
		return Client
	}

	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://broker.hivemq.com:1883")
	options.SetClientID("dds-tpa-heladera")
	options.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MESSAGE: %s\n", msg.Payload())
	})
	options.SetCleanSession(true)

	Client = mqtt.NewClient(options)
	return Client
}
