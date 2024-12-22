package main

import (
	"fmt"

	"github.com/JiunMHsu/dds-go-mqtt-client/client"
)

func main() {
	client := client.GetClient()
	fmt.Println(client)
}
