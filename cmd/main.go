package main

import (
	"sync"

	"github.com/JiunMHsu/dds-go-mqtt-client/internal/console"
	mqttClient "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"
	"github.com/JiunMHsu/dds-go-mqtt-client/internal/publisher"
)

func main() {

	topics := []string{
		"utn-dds-g22/heladeras/heladera-plaza-de-mayo",
		"utn-dds-g22/heladeras/heladera-ferro",
		"utn-dds-g22/heladeras/heladera-atlanta",
		"utn-dds-g22/heladeras/heladera-utn-lugano",
		"utn-dds-g22/heladeras/heladera-obelisco",
		"utn-dds-g22/heladeras/heladera-caminito-de-la-boca",
		"utn-dds-g22/heladeras/heladera-plaza-italia",
		"utn-dds-g22/heladeras/heladera-barrio-chino",
		"utn-dds-g22/heladeras/heladera-facultad-de-derecho",
		"utn-dds-g22/heladeras/heladera-hospital-piniero",
		"utn-dds-g22/heladeras/heladera-plaza-serrano",
		"utn-dds-g22/heladeras/heladera-abasto-shopping",
		"utn-dds-g22/heladeras/heladera-guerrin",
		"utn-dds-g22/heladeras/heladera-linea-d",
		"utn-dds-g22/heladeras/heladera-utn-medrano",
	}

	for _, topic := range topics {
		client := mqttClient.NewClient(topic)
		go publisher.StartTemperaturePublishingRoutine(client)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go console.RunConsole(&wg)
	wg.Wait()
}
