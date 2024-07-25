package main

import (
	"codeZone/internal/app"
	"context"
	"flag"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "env-file", ".env", "path to .env file")

	flag.Parse()
}

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("error initializing application: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("error running application: %v", err)
	}
}
