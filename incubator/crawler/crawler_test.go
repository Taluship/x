package crawler

import (
	"testing"
)

func TestCrawler_APIClient(t *testing.T) {
	c := NewCrawlerClient()

	targets, err := c.FetchTargets()
	if err != nil {
		t.Error(err)
		return
	}
	for _, target := range targets {
		t.Logf("Target: %+v", target)
	}
}
