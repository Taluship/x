package pclient

import (
	"fmt"
	"net/url"
)

type CollectionIdOrName string

type Schema struct {
	System   bool   `json:"system"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Unique   bool   `json:"unique"`
	Options  struct {
		Min       interface{} `json:"min"`
		Max       interface{} `json:"max"`
		Pattern   string      `json:"pattern,omitempty"`
		MaxSelect int         `json:"maxSelect,omitempty"`
		MaxSize   int         `json:"maxSize,omitempty"`
		MimeTypes []string    `json:"mimeTypes,omitempty"`
		Thumbs    interface{} `json:"thumbs"`
	} `json:"options"`
}

type CollectionResponse struct {
	Id         string      `json:"id"`
	Created    string      `json:"created"`
	Updated    string      `json:"updated"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Schema     []Schema    `json:"schema"`
	ListRule   string      `json:"listRule"`
	ViewRule   string      `json:"viewRule"`
	CreateRule string      `json:"createRule"`
	UpdateRule string      `json:"updateRule"`
	DeleteRule interface{} `json:"deleteRule"`
	Options    struct{}    `json:"options"`
	Indexes    []string    `json:"indexes"`
}

type ListCollectionParams struct {
	Page      int
	PerPage   int
	Sort      string
	Filter    string
	Fields    string
	SkipTotal bool
}

type ViewCollectionParams struct {
	Fields string
}

// ListCollection list/search collection
func (c *Client) ListCollection(params ListCollectionParams) (PaginationResponse[CollectionResponse], error) {
	var data = PaginationResponse[CollectionResponse]{
		Items:      []*CollectionResponse{},
		Page:       0,
		PerPage:    0,
		TotalItems: 0,
		TotalPages: 0,
	}

	u, err := url.Parse("/api/collections/")
	if err != nil {
		return data, err
	}

	q := u.Query()
	if params.Page != 0 {
		q.Set("page", fmt.Sprintf("%d", params.Page))
	}
	if params.PerPage != 0 {
		q.Set("perPage", fmt.Sprintf("%d", params.PerPage))
	}
	if params.Sort != "" {
		q.Set("sort", params.Sort)
	}
	if params.Filter != "" {
		q.Set("filter", params.Filter)
	}
	if params.Fields != "" {
		q.Set("fields", params.Fields)
	}
	if params.SkipTotal {
		q.Set("skipTotal", fmt.Sprintf("%t", params.SkipTotal))
	}
	u.RawQuery = q.Encode()

	resp, err := c.get(u.String())
	if err != nil {
		return data, err
	}

	err = c.decode(resp, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// ViewCollection view collection
func (c *Client) ViewCollection(arg CollectionIdOrName, params ViewCollectionParams) (*CollectionResponse, error) {
	u, err := url.Parse("/api/collections/" + string(arg))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	if params.Fields != "" {
		q.Set("fields", params.Fields)
	}
	u.RawQuery = q.Encode()

	resp, err := c.get(u.String())
	if err != nil {
		return nil, err
	}

	var data = new(CollectionResponse)
	err = c.decode(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateCollection create collection
func (c *Client) CreateCollection() {

}

// UpdateCollection update collection
func (c *Client) UpdateCollection() {

}

// DeleteCollection delete collection
func (c *Client) DeleteCollection() {

}

// ImportCollection import collection
func (c *Client) ImportCollection() {

}
