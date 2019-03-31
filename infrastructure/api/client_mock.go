package api

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
	fname := filepath.Join(p, "./infrastructures/api/mockdatas/mock_response.txt")
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("input file error: ", err)
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
