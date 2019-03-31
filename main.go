package main

import (
	"fmt"
	"os"

	"github.com/anraku/gonnpass/infrastructures/api"
	"github.com/anraku/gonnpass/interfaces/cli"
	"github.com/anraku/gonnpass/usecase"
)

func main() {
	c := api.NewClient(nil)
	//c := api.NewMockClient()
	requestUsecase := usecase.NewRequestIterator(c)
	app := cli.NewCommand(requestUsecase)
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
