package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type MockClient struct{}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (c *MockClient) Get(query string, param map[string]string) ([]byte, error) {
	p, _ := os.Getwd()
	fname := filepath.Join(p, "./domain/services/mockdatas/mock_response.txt")
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("input file error: ", err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}
