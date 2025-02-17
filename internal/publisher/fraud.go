package publisher

import mqtt "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"

func PublishFraudFor(topic string) {
	message := "FRAUDE"
	client := mqtt.NewClient(topic)
	client.Publish(message)
}
