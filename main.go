package main

import (
	"fmt"
	"os"

	"github.com/anraku/gonnpass/infrastructure/api"
	"github.com/anraku/gonnpass/interface/cli"
	"github.com/anraku/gonnpass/usecase"
)

func main() {
	c := api.NewClient(nil)
	// c := api.NewMockClient()
	requestUsecase := usecase.NewEventIterator(c)
	app := cli.NewCommand(requestUsecase)
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
