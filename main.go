package main

import (
	"flag"
	"fmt"

	"github.com/anraku/gonnpass/domain/services"
	"github.com/anraku/gonnpass/usecases"
)

const (
	BASE_URL = "https://connpass.com/api/v1/event/"
)

func main() {
	var (
		i = flag.Int("int", 0, "int flag")
		s = flag.String("str", "default", "string flag")
		b = flag.Bool("bool", false, "bool flag")
	)
	flag.Parse()

	//c := services.NewClient(nil)
	c := services.NewMockClient()
	ur := usecases.NewRequestIterator(c)
	events := ur.FetchEvents()

	fmt.Println(events)
	fmt.Println(*i, *s, *b)
}
