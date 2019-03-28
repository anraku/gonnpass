package cli

import (
	"flag"
	"fmt"

	"github.com/anraku/gonnpass/usecases"
)

type Command struct {
	ru usecases.RequestUsecase
}

func NewCommand(ru usecases.RequestUsecase) *Command {
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
	events, err := c.ru.FetchEvents()
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
