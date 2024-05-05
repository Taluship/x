package crawler

import (
	"testing"

	"github.com/pluja/pocketbase"
)

func TestCrawler_APIClient(t *testing.T) {
	c := pocketbase.NewClient("http://127.0.0.1:8090", pocketbase.WithAdminEmailPassword("zcharyma@gmail.com", "Admin@1234"))

	list2, err := c.List("crawl_targets", pocketbase.ParamsList{})
	if err != nil {
		return
	}

	t.Logf("List: %+v", list2)
}
