package usecase

import (
	"encoding/json"

	"github.com/anraku/gonnpass/domain/model"
	"github.com/anraku/gonnpass/usecase/data"
	"github.com/anraku/gonnpass/usecase/repository"
)

type EventUsecase interface {
	SearchEvents(input data.InputData) (*model.Events, error)
}

type EventIterator struct {
	sc repository.Events
}

func NewEventIterator(sc repository.Events) *EventIterator {
	return &EventIterator{
		sc: sc,
	}
}

func (i *EventIterator) SearchEvents(input data.InputData) (*model.Events, error) {
	body, err := i.sc.SearchEvents(input)
	if err != nil {
		return nil, err
	}

	cr := new(model.Events)
	err = json.Unmarshal(body, cr)
	return cr, err
}
