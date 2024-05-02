// Package pclient pocketbase RESTful API client
package pclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

type Err struct {
	Code    int
	Message string
}

func (e Err) Error() string {
	return e.Message
}

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

func NewTokenClient(token string) *Client {
	return NewClient(WithAuthToken(token))
}

type Option func(*Client)

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func WithAuthToken(token string) Option {
	return func(c *Client) {
		if c.httpClient == nil {
			c.httpClient = http.DefaultClient
		}
		transport := c.httpClient.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		c.httpClient.Transport = roundTripperFunc(
			func(req *http.Request) (*http.Response, error) {
				req = req.Clone(req.Context())
				req.Header.Set("Authorization", fmt.Sprintf("%s", token))
				return transport.RoundTrip(req)
			},
		)
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

	if resp.StatusCode > 299 {
		return nil, Err{Code: resp.StatusCode, Message: http.StatusText(resp.StatusCode)}
	}

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

	if resp.StatusCode > 299 {
		return nil, Err{Code: resp.StatusCode, Message: http.StatusText(resp.StatusCode)}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// roundTripperFunc creates a RoundTripper (transport)
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}
