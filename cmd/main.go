package main

import (
	"sync"

	mqttClient "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"
	"github.com/JiunMHsu/dds-go-mqtt-client/internal/routine"
)

func main() {

	topics := []string{
		"utn-dds-g22/heladeras/heladera-diez",
		"utn-dds-g22/heladeras/heladera-cinco",
		"utn-dds-g22/heladeras/heladera-nueve",
		"utn-dds-g22/heladeras/heladera-uno",
		"utn-dds-g22/heladeras/heladera-catorce",
		"utn-dds-g22/heladeras/heladera-once",
		"utn-dds-g22/heladeras/heladera-doce",
		"utn-dds-g22/heladeras/heladera-quince",
		"utn-dds-g22/heladeras/heladera-trece",
		"utn-dds-g22/heladeras/heladera-ocho",
		"utn-dds-g22/heladeras/heladera-siete",
		"utn-dds-g22/heladeras/heladera-dos",
		"utn-dds-g22/heladeras/heladera-seis",
		"utn-dds-g22/heladeras/heladera-tres",
	}

	var wg sync.WaitGroup

	for _, topic := range topics {
		client := mqttClient.NewClient(topic)
		wg.Add(1)
		go routine.PublishTemperatureFor(client)
	}

	wg.Wait()

	// for {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	text, _ := reader.ReadString('\n')
	// }
}
