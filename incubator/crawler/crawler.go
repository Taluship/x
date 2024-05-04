package crawler

import (
	"net/http"
)

type Crawler interface {
	Crawl()
}

type Client struct {
	client *http.Client
}

func NewCrawlerClient() *Client {
	c := &Client{}

	return c
}
