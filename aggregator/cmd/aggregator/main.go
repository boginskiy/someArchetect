package main

import (
	"aggregatorProject/internal/app"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
