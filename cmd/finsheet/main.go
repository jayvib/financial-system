package main

import (
	"os"
	"log"
	fcli "financial-system/internal/cli"
)

func main() {
	app := fcli.NewApp(fcli.FlagInitializer)
	fcli.FlagInitializer(app)
	fcli.ActionInitializer(app)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
