package main

import (
	"fmt"

	"github.com/JiunMHsu/go-mqtt-publisher/client"
)

func main() {
	client := client.GetClient()
	fmt.Println(client)
}
