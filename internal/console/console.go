package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func RunConsole(wg *sync.WaitGroup) {
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
