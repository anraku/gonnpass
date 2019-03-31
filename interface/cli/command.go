package cli

import (
	"fmt"
	"os"

	"github.com/anraku/gonnpass/domain/data"
	"github.com/anraku/gonnpass/usecase"
	"github.com/urfave/cli"
)

type Command struct {
	ru usecase.EventUsecase
}

func NewCommand(ru usecase.EventUsecase) *Command {
	return &Command{
		ru: ru,
	}
}

const EVENTS_COUNT = 10

func (cmd *Command) Run() error {
	app := cli.NewApp()
	app.Name = "gonnpass"
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "and",
			Usage: "AND is search keyword",
		},
		cli.StringSliceFlag{
			Name:  "or",
			Usage: "OR for search keyword",
		},
		cli.BoolFlag{
			Name:  "update-order",
			Usage: "Sort events order by update_at",
		},
		cli.BoolFlag{
			Name:  "start-order",
			Usage: "Sort events order by start_at",
		},
		cli.BoolFlag{
			Name:  "new-order",
			Usage: "Sort events order by new_at",
		},
		cli.IntFlag{
			Name:  "n",
			Value: EVENTS_COUNT,
			Usage: "number of events to print",
		},
	}
	app.Action = func(c *cli.Context) error {
		order := evalOrder(c.Bool("update-order"), c.Bool("start-order"), c.Bool("new-order"))

		// create input data fro usecase
		input := data.InputData{
			KeywordAND: c.StringSlice("and"),
			KeywordOR:  c.StringSlice("or"),
			Order:      order,
			Count:      c.Int("n"),
		}

		events, err := cmd.ru.SearchEvents(input)
		if err != nil {
			return err
		}

		layout := "2006/01/02 15:04:05"
		for _, v := range events.Events {
			start := v.StartedAt.Format(layout)
			end := v.EndedAt.Format(layout)
			fmt.Println(v.Title)
			fmt.Printf("    %s - %s\n", start, end)
			fmt.Printf("    %s\n", v.EventURL)
		}
		return nil
	}

	return app.Run(os.Args)
}

type Keyword struct {
	Values []string
}

func (k *Keyword) String() string {
	return fmt.Sprintf("%v", k.Values)
}

func (k *Keyword) Set(v string) error {
	k.Values = append(k.Values, v)
	return nil
}

func evalOrder(update, start, newOrder bool) data.Order {
	if update {
		return data.UpdateOrder
	} else if newOrder {
		return data.NewOrder
	} else {
		return data.StartOrder
	}
}
