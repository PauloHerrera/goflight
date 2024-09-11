package main

import (
	"log"

	"gihub.com/pauloherrera/goflight/api"
	"gihub.com/pauloherrera/goflight/storage"
	"gihub.com/pauloherrera/goflight/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to load configurations:", err)
	}

	worker := storage.NewWorker(config.DatabaseUri, config.DatabaseName)

	server := api.NewServer(worker)

	server.Start(config.ServerAddress)
}
