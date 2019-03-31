package cli

import (
	"flag"
	"fmt"

	"github.com/anraku/gonnpass/domain/data"
	"github.com/anraku/gonnpass/usecase"
)

type Command struct {
	ru usecase.EventUsecase
}

func NewCommand(ru usecase.EventUsecase) *Command {
	return &Command{
		ru: ru,
	}
}

func (c *Command) Run() error {
	var (
		and Keyword
		or  Keyword
	)

	flag.Var(&and, "and", "and search keyword")
	flag.Var(&or, "or", "or search keyword")
	updateOrder := flag.Bool("update-order", false, "update order")
	startOrder := flag.Bool("start-order", true, "update order")
	newOrder := flag.Bool("new-order", false, "update order")
	count := flag.Int("n", 10, "number of events to print")
	flag.Parse()

	order := evalOrder(*updateOrder, *startOrder, *newOrder)

	// create input data fro usecase
	input := data.InputData{
		KeywordAND: and.Values,
		KeywordOR:  or.Values,
		Order:      order,
		Count:      *count,
	}

	events, err := c.ru.SearchEvents(input)
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
