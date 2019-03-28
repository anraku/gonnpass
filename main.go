package main

import (
	"github.com/anraku/gonnpass/domain/services"
	"github.com/anraku/gonnpass/interfaces/cli"
	"github.com/anraku/gonnpass/usecases"
)

func main() {
	// c := services.NewClient(nil)
	c := services.NewMockClient()
	requestUsecase := usecases.NewRequestIterator(c)
	app := cli.NewCommand(requestUsecase)
	app.Run()
}
