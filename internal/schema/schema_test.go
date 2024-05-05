package schema

import (
	"testing"

	"github.com/pluja/pocketbase"
	"github.com/stretchr/testify/assert"
)

func TestTargetCreationAndQueries(t *testing.T) {
	c := pocketbase.NewClient("http://127.0.0.1:8090", pocketbase.WithAdminEmailPassword("zcharyma@gmail.com", "Admin@1234"))

	newTarget := Target{
		Name:      "Test Target",
		Type:      "anime",
		URL:       "http://test.com",
		Selector:  "Test Selector",
		LastCrawl: "2022-01-01T00:00:00Z",
		IsActive:  true,
	}

	response, err := c.Create("crawl_targets", newTarget)
	assert.NoError(t, err)

	list, err := c.List("crawl_targets", pocketbase.ParamsList{})
	assert.NoError(t, err)

	t.Logf("List: %+v", list)

	for _, itemMap := range list.Items {
		var queriedTarget = Target{
			Base: Base{
				ID:             itemMap["id"].(string),
				CollectionID:   itemMap["collectionId"].(string),
				CollectionName: itemMap["collectionName"].(string),
				Created:        itemMap["created"].(string),
				Updated:        itemMap["updated"].(string),
			},
			Name:      itemMap["name"].(string),
			Type:      itemMap["type"].(string),
			URL:       itemMap["url"].(string),
			Selector:  itemMap["selector"].(string),
			LastCrawl: itemMap["last_crawl"].(string),
			IsActive:  itemMap["is_active"].(bool),
		}
		if queriedTarget.ID == response.ID {
			assert.Equal(t, newTarget.Name, queriedTarget.Name)
			assert.Equal(t, newTarget.Type, queriedTarget.Type)
			assert.Equal(t, newTarget.URL, queriedTarget.URL)
			assert.Equal(t, newTarget.Selector, queriedTarget.Selector)
			assert.Equal(t, newTarget.IsActive, queriedTarget.IsActive)
		}
	}
}
