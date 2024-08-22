package crawler

import (
	"net/http"

	"github.com/taluship/x/cloud"
	"github.com/taluship/x/internal/schema"
	"github.com/zcharym/pocketbase-client"
)

type Crawler interface {
	Crawl()
	FetchTargets() ([]*schema.CrawlTarget, error)
}

type Client struct {
	cloud.BaseClient
	httpClient *http.Client
}

func (c Client) Crawl() {
	// TODO implement me
	panic("implement me")
}

func (c Client) FetchTargets() ([]*schema.CrawlTarget, error) {
	var targets []*schema.CrawlTarget
	resp, err := c.Pocketbase.FullList(schema.CollectionTarget, pocketbase.ParamsList{})
	if err != nil {
		return targets, err
	}

	for _, itemMap := range resp.Items {
		target := &schema.CrawlTarget{}
		if err := schema.ResponseMapToObject(itemMap, target); err != nil {
			return targets, err
		}
		targets = append(targets, target)
	}

	return targets, nil
}

func NewCrawlerClient() *Client {
	c := &Client{
		BaseClient: *cloud.NewClient(),
	}

	return c
}
