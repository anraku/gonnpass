package services

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://connpass.com/api/v1/event/"
	userAgent      = "gonnpass"
)

type Client interface {
	Get(query string, param map[string]string) ([]byte, error)
}

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

func (c *HTTPClient) Get(query string, param map[string]string) ([]byte, error) {
	resp, err := c.client.Get(c.BaseURL.String() + query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
