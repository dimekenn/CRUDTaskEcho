package main

import (
	"context"
	"crudTaskEcho/configs"
	"crudTaskEcho/internal/app"
	"embed"
	"encoding/json"
	"log"
)

//go:embed configs.json
var fs embed.FS

const configName = "configs.json"

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//reading json file for configs
	data, readErr := fs.ReadFile(configName)
	if readErr != nil {
		log.Fatal(readErr)
	}
	//creating config entity to deserialize configs.json
	cfg := configs.NewConfig()
	if unmErr := json.Unmarshal(data, &cfg); unmErr != nil {
		log.Fatal(unmErr)
	}

	errCh := make(chan error, 1)

	go app.StartHTTPServer(ctx, errCh, cfg)

	log.Fatalf("%v", <-errCh)
}
