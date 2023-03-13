package main

import (
	"fmt"
	"log"
	"net/http"
	"user-api/config"

	netchi "user-api/transport"
)

func main() {
	config, configError := config.ReadConfigFromFile("config.json")
	if configError != nil {
		log.Fatal(configError)
	}

	handler := netchi.Initialize(config.Port)

	fmt.Printf("server is running on port: %d\n", handler.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), handler.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
