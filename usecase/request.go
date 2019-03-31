package usecase

import (
	"encoding/json"

	"github.com/anraku/gonnpass/domain/data"
	"github.com/anraku/gonnpass/domain/model"
	"github.com/anraku/gonnpass/domain/repository"
)

type RequestUsecase interface {
	SearchEvents(input data.InputData) (*model.Events, error)
}

type RequestIterator struct {
	sc repository.Events
}

func NewRequestIterator(sc repository.Events) *RequestIterator {
	return &RequestIterator{
		sc: sc,
	}
}

func (i *RequestIterator) SearchEvents(input data.InputData) (*model.Events, error) {
	body, err := i.sc.SearchEvents(input)
	if err != nil {
		return nil, err
	}

	cr := new(model.Events)
	err = json.Unmarshal(body, cr)
	return cr, err
}
