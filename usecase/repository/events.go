package repository

import "github.com/anraku/gonnpass/usecase/data"

type Events interface {
	SearchEvents(input data.InputData) ([]byte, error)
}
