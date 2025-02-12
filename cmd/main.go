package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	mqttClient "github.com/JiunMHsu/dds-go-mqtt-client/internal/mqtt-client"
	"github.com/JiunMHsu/dds-go-mqtt-client/internal/publish"
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
		go publish.StartTemperaturePublishingRoutine(client)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go runConsole(&wg)
	wg.Wait()
}

func runConsole(wg *sync.WaitGroup) {
	defer wg.Done()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Console -- Type 'exit' to quit")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "exit":
			fmt.Println("Exiting console...")
			return
		default:
			fmt.Println("Unknown command:", input)
		}
	}
}
