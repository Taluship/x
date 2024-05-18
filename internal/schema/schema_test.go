package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestTargetCreationAndQueries(t *testing.T) {
// 	c := pocketbase.NewClient("http://127.0.0.1:8090", pocketbase.WithAdminEmailPassword("zcharyma@gmail.com", "Admin@1234"))
//
// 	newTarget := Target{
// 		Name:      "Test Target",
// 		Type:      "anime",
// 		URL:       "http://test.com",
// 		Selector:  "Test Selector",
// 		LastCrawl: "2022-01-01T00:00:00Z",
// 		IsActive:  true,
// 	}
//
// 	response, err := c.Create("crawl_targets", newTarget)
// 	assert.NoError(t, err)
//
// 	list, err := c.List("crawl_targets", pocketbase.ParamsList{})
// 	assert.NoError(t, err)
//
// 	t.Logf("List: %+v", list)
//
// 	for _, itemMap := range list.Items {
// 		var queriedTarget = Target{
// 			Base: Base{
// 				ID:             itemMap["id"].(string),
// 				CollectionID:   itemMap["collectionId"].(string),
// 				CollectionName: itemMap["collectionName"].(string),
// 				Created:        itemMap["created"].(string),
// 				Updated:        itemMap["updated"].(string),
// 			},
// 			Name:      itemMap["name"].(string),
// 			Type:      itemMap["type"].(string),
// 			URL:       itemMap["url"].(string),
// 			Selector:  itemMap["selector"].(string),
// 			LastCrawl: itemMap["last_crawl"].(string),
// 			IsActive:  itemMap["is_active"].(bool),
// 		}
// 		if queriedTarget.ID == response.ID {
// 			assert.Equal(t, newTarget.Name, queriedTarget.Name)
// 			assert.Equal(t, newTarget.Type, queriedTarget.Type)
// 			assert.Equal(t, newTarget.URL, queriedTarget.URL)
// 			assert.Equal(t, newTarget.Selector, queriedTarget.Selector)
// 			assert.Equal(t, newTarget.IsActive, queriedTarget.IsActive)
// 		}
// 	}
// }

func TestParseResponseMapToStruct(t *testing.T) {
	data := map[string]interface{}{
		"collectionId":   "v6xo9qcrupn1pn6",
		"collectionName": "crawl_target",
		"created":        "2024-05-05 03:59:52.816Z",
		"id":             "gqzat6k3q8xzwif",
		"is_active":      true,
		"last_crawl":     "",
		"name":           "name",
		"selector":       "fasfds",
		"type":           "anime",
		"updated":        "2024-05-05 03:59:52.816Z",
		"url":            "http://example.com",
	}

	target := &CrawlTarget{}
	if err := ParseResponseMapToStruct(data, target); err != nil {
		t.Error(err)
		return
	}

	// ID type check
	assert.Equal(t, "v6xo9qcrupn1pn6", target.CollectionID)
	assert.Equal(t, "gqzat6k3q8xzwif", target.ID)

	// string type value check
	assert.Equal(t, "crawl_target", target.CollectionName)
	assert.Equal(t, "name", target.Name)
	assert.Equal(t, "fasfds", target.Selector)
	assert.Equal(t, "anime", target.Type)

	// date type value check
	assert.Equal(t, "2024-05-05 03:59:52.816Z", target.Created)
	assert.Equal(t, "", target.LastCrawl)
	assert.Equal(t, "2024-05-05 03:59:52.816Z", target.Updated)

	// boolean type value check
	assert.Equal(t, true, target.IsActive)

	// url type value check
	assert.Equal(t, "http://example.com", target.Url)
}
