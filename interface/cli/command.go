package cli

import (
	"flag"
	"fmt"

	"github.com/anraku/gonnpass/domain/data"
	"github.com/anraku/gonnpass/usecase"
)

type Command struct {
	ru usecase.RequestUsecase
}

func NewCommand(ru usecase.RequestUsecase) *Command {
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
	flag.Parse()

	fmt.Printf("daimori: and: %+v\n", and)
	fmt.Printf("daimori: or: %+v\n", or)
	// create input data fro usecase
	input := data.InputData{
		KeywordAND: and.Values,
		KeywordOR:  or.Values,
	}

	events, err := c.ru.SearchEvents(input)
	if err != nil {
		return err
	}

	fmt.Println(events)
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
