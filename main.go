package main

import (
	"encoding/json"
	"fmt"

	domain "github.com/anraku/gonnpass/domain/services"
	"github.com/anraku/gonnpass/interfaces"
)

const (
	BASE_URL = "https://connpass.com/api/v1/event/"
)

func main() {
	//c := domain.NewClient(nil)
	c := domain.NewMockClient()
	body, err := c.Get("?keyword=python", nil)
	if err != nil {
		panic(err)
	}

	cr := new(interfaces.ConnpassResponse)
	err = json.Unmarshal(body, cr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", cr)
}
