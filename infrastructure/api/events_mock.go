package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/anraku/gonnpass/domain/data"
	"github.com/anraku/gonnpass/domain/repository"
)

type MockClient struct{}

func NewMockClient() repository.Events {
	return &MockClient{}
}

func (c *MockClient) SearchEvents(input data.InputData) ([]byte, error) {
	p, _ := os.Getwd()
	fname := filepath.Join(p, "./infrastructure/api/mockdata/events_mock.json")
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("input file error: ", err)
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
