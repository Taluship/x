package pclient

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

// ListCollection list/search collection
func (c *Client) ListCollection() {

}

// ViewCollection view collection
func (c *Client) ViewCollection() {

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
