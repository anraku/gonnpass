package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/anraku/gonnpass/usecase/data"
	"github.com/anraku/gonnpass/usecase/repository"
)

const (
	defaultBaseURL = "https://connpass.com/api/v1/event/"
	userAgent      = "gonnpass"
)

type HTTPClient struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
}

func NewClient(httpClient *http.Client) repository.Events {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &HTTPClient{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	return c
}

func (c *HTTPClient) SearchEvents(input data.InputData) ([]byte, error) {
	url := generateURL(c.BaseURL, input)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func generateURL(base *url.URL, input data.InputData) string {
	const (
		and   = "keyword="
		or    = "keyword_or="
		order = "order="
		count = "count="
		start = "ymd="
	)

	s := base.String() + "?"

	for _, v := range input.KeywordAND {
		s += and + v + "&"
	}

	for _, v := range input.KeywordOR {
		s += or + v + "&"
	}

	s += order + fmt.Sprintf("%d", input.Order) + "&"
	s += count + fmt.Sprintf("%d", input.Count) + "&"

	return s
}
