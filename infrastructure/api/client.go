package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/anraku/gonnpass/domain/data"
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

func NewClient(httpClient *http.Client) *HTTPClient {
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
	fmt.Printf("input: %+v", input)
	s := base.String() + "?keyword="
	for _, v := range input.KeywordAND {
		s += v
		s += "&keyword="
	}
	fmt.Println("daimori and: ", strings.LastIndex(s, "&keyword="))
	s = s[:strings.LastIndex(s, "&keyword=")]

	s += "&keyword_or="
	for _, v := range input.KeywordOR {
		s += v
		s += "&keyword_or="
	}
	fmt.Println("daimori or: ", strings.LastIndex(s, "&keyword_or="))
	s = s[:strings.LastIndex(s, "&keyword_or=")]

	fmt.Println("daimori: URL: ", s)
	return s
}
