package main

import (
	"log"

	"gihub.com/pauloherrera/goflight/api"
	"gihub.com/pauloherrera/goflight/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
	}

	server := api.NewServer()

	server.Start(config.ServerAddress)
}
