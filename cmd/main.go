package main

import (
	"fmt"
	"log"
	"user-api/config"

	netchi "user-api/transport"
)

func main() {
	config, configError := config.ReadConfigFromFile("config.json")
	if configError != nil {
		log.Fatal(configError)
	}

	fmt.Println("Port in config:", config.Port)

	handler := netchi.Initialize(config.Port)

	fmt.Println("Handler port:", handler.Port)
}
