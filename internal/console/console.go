package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/JiunMHsu/dds-go-mqtt-client/internal/publisher"
)

func RunConsole(wg *sync.WaitGroup) {
	defer wg.Done()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Console -- Type 'quit' to quit")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "quit":
			fmt.Println("Quitting console")
			return
		default:
			handlePublishCommand(input)
		}
	}
}

func handlePublishCommand(input string) {
	command := strings.Split(input, " ")

	if len(command) < 2 {
		fmt.Println("Invalid command")
		return
	}

	switch command[0] {
	case "temp":
		handleTemperatureCommand(command)
	case "fraud":
		handleFraudCommand(command)
	case "open-for":
		handleOpeningRequestCommand(command)
	default:
		fmt.Println("Unknown command")
	}
}

func handleTemperatureCommand(command []string) {
	if len(command) == 2 {
		publisher.PublishRandomTemperatureFor(command[1])
		fmt.Println("Random temperature published for topic", command[1])
		return
	}

	if len(command) == 3 {
		temp, err := strconv.ParseFloat(command[1], 64)
		if err != nil {
			fmt.Println("Invalid temperature")
			return
		}

		publisher.PublishTemperatureFor(command[2], temp)
		fmt.Println("Temperature", temp, "published for topic", command[2])
		return
	}

	fmt.Println("Invalid command")
}

func handleFraudCommand(command []string) {
	if len(command) != 2 {
		fmt.Println("Invalid command")
		return
	}

	publisher.PublishFraudFor(command[1])
	fmt.Println("Fraud published for topic", command[1])
}

func handleOpeningRequestCommand(command []string) {
	if len(command) != 3 {
		fmt.Println("Invalid command")
		return
	}

	publisher.PublishOpeningRequest(command[2], command[1])
	fmt.Println("Opening request published for topic", command[1])
}
