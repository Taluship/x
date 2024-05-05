// Package schema contains schemas with additional fields id, created, updated form pocketbase.
package schema

type Base struct {
	ID             string `json:"id"`
	CollectionID   string `json:"collectionId"`
	CollectionName string `json:"collectionName"`
	Created        string `json:"created"`
	Updated        string `json:"updated"`
}

type Notifier struct {
	Base
	NotifyMethod   string `json:"notify_method"`
	NotifyTemplate string `json:"notify_template"`
}

type Target struct {
	Base
	Name      string `json:"name"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	Selector  string `json:"selector"`
	LastCrawl string `json:"last_crawl"`
	IsActive  bool   `json:"is_active"`
}

type Event struct {
	Status   string `json:"status"`
	Data     string `json:"data"`
	Checksum string `json:"checksum"`
}
