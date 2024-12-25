package mqttClient

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Client mqtt.Client = nil

type MqttClient struct {
	Client mqtt.Client
	Topic  string
}

func NewClient(topic string) MqttClient {
	if Client == nil {
		options := mqtt.NewClientOptions()
		options.AddBroker("tcp://broker.hivemq.com:1883")
		options.SetClientID("dds-tpa-heladera")
		options.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("TOPIC: %s\n", msg.Topic())
			fmt.Printf("MESSAGE: %s\n", msg.Payload())
		})
		options.SetCleanSession(true)

		Client = mqtt.NewClient(options)
		Client.Connect()
		fmt.Println("Connected to broker")
	}

	return MqttClient{Client: Client, Topic: topic}

}

func  (client MqttClient) Publish(message string) {
	client.Client.Publish(client.Topic, 0, false, message)
}