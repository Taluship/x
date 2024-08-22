// Package schema contains schemas with additional fields id, created, updated form pocketbase.
package schema

import (
	"unicode"

	"github.com/mitchellh/mapstructure"
	"github.com/pocketbase/pocketbase/models/schema"
)

// collections
const (
	CollectionEvent    = "crawl_event"
	CollectionNotifier = "crawl_notifier"
	CollectionTarget   = "crawl_target"
)

type CollectionExport struct {
	Id         string               `json:"id"`
	Name       string               `json:"name"`
	Type       string               `json:"type"`
	System     bool                 `json:"system"`
	Schema     []schema.SchemaField `json:"schema"`
	Indexes    []string             `json:"indexes"`
	ListRule   interface{}          `json:"listRule"`
	ViewRule   interface{}          `json:"viewRule"`
	CreateRule interface{}          `json:"createRule"`
	UpdateRule interface{}          `json:"updateRule"`
	DeleteRule interface{}          `json:"deleteRule"`
	Options    struct {
	} `json:"options"`
}

func ToTitle(str string) string {
	// TODO separate the schema by collection prefix
	return ToPascalCase(str)
}

func ToPascalCase(str string) string {
	var result []rune

	capitalizeNext := true
	for _, r := range str {
		if r == '_' {
			capitalizeNext = true
			continue
		}
		if capitalizeNext {
			result = append(result, unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func ConvertType(typ string) string {
	// TODO handling fields like select, editor, url, date, text

	switch typ {
	default:
		return "string"
	case "select":
		return "string"
	case "bool":
		return "bool"
	case "editor":
		return "string"
	case "url":
		return "string"
	case "date":
		return "string"
	case "text":
		return "string"
	}
}

func ResponseMapToObject(response map[string]any, object any) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  object,
		TagName: "json",
	})

	if err != nil {
		return err
	}

	if err := decoder.Decode(response); err != nil {
		return err
	}
	return nil
}
