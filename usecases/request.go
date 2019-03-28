package usecases

import (
	"encoding/json"

	"github.com/anraku/gonnpass/domain/models"
	"github.com/anraku/gonnpass/domain/services"
)

type RequestUsecase interface {
	FetchEvents()
}

type RequestIterator struct {
	sc services.Client
}

func NewRequestIterator(sc services.Client) *RequestIterator {
	return &RequestIterator{
		sc: sc,
	}
}

func (i *RequestIterator) FetchEvents() *models.ConnpassResponse {
	body, err := i.sc.Get("?keyword=python", nil)
	if err != nil {
		panic(err)
	}

	cr := new(models.ConnpassResponse)
	err = json.Unmarshal(body, cr)
	if err != nil {
		panic(err)
	}

	return cr
}
