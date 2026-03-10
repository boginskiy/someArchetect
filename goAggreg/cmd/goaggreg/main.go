package goaggreg

import (
	"context"
	"goAggreg/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
