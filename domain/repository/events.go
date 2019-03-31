package repository

import "github.com/anraku/gonnpass/domain/data"

type Events interface {
	SearchEvents(input data.InputData) ([]byte, error)
}
