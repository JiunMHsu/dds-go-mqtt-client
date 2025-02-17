package publisher

import (
	"fmt"

	mqtt "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"
)

// publicar una solicitud de apertura debería implicar
// también esperar una respuesta de autorización
// Pero en este caso simplemente se publica la solicitud
func PublishOpeningRequest(topic, card string) {
	message := fmt.Sprintf("SOLICITUD_APERTURA %s", card)
	client := mqtt.NewClient(topic)
	client.Publish(message)
}
