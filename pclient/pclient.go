// Package pclient pocketbase RESTful API client
package pclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	HealthPath       = "/api/health"
	ListRecordPath   = "/api/collections/{{ collection }}/records"
	ViewRecordPath   = "/api/collections/{{ collection }}/records"
	CreateRecordPath = "/api/collections/{{ collection }}/records"
	UpdateRecordPath = "/api/collections/{{ collection }}/records"
	DeleteRecordPath = "/api/collections/{{ collection }}/records"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewClient(options ...Option) *Client {
	// defaults
	c := &Client{
		httpClient: http.DefaultClient,
		baseURL:    "http://127.0.0.1:8090",
	}
	c.httpClient.Timeout = 10 * time.Second

	for _, option := range options {
		option(c)
	}
	return c
}

type Option func(*Client)

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func (c *Client) get(urlPath string) ([]byte, error) {
	requestUrl, err := url.JoinPath(c.baseURL, urlPath)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) post(urlPath string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, path.Join(c.baseURL, urlPath), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
