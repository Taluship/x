//go:build ignore

// generates the struct structs.go which contains several structs based on format and constraint from schema.json file

package main

import (
	"encoding/json"
	"os"
	"text/template"

	"github.com/taluship/x/internal/schema"
	"github.com/taluship/x/log"
)

var fnMap = template.FuncMap{
	"title":        schema.ToTitle,
	"pascal_case":  schema.ToPascalCase,
	"convert_type": schema.ConvertType,
}

var schemaTmpl = template.Must(template.New("schema").Funcs(fnMap).Parse(`// Package schema
// Automatically generated by gen_structs.go, DO NOT EDIT manually
package schema

type Base struct {
	ID             string ` + "`json:\"id\"`" + `
	CollectionID   string ` + "`json:\"collectionId\"`" + `
	CollectionName string ` + "`json:\"collectionName\"`" + `
	Created        string ` + "`json:\"created\"`" + `
	Updated        string ` + "`json:\"updated\"`" + `
}
{{range .}}
type {{.Name | title}} struct {
	Base {{range .Schema}}
	{{.Name | pascal_case}} {{.Type | convert_type}} ` + "`json:\"{{.Name}}\"`" + `{{end}}
}
{{end}}

`))

func main() {
	log.Infof("generating structs...")
	defer log.Infof("generation done.")

	data, err := os.ReadFile("schema.json")
	if err != nil {
		log.Errorf("failed to read schema.json: %v\n", err)
		return
	}
	// Unmarshal the JSON into the Go structs
	var collections []schema.CollectionExport
	err = json.Unmarshal(data, &collections)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create("structs.go")
	if err != nil {
		log.Errorf("failed to create structs.go: %v", err)
		return
	}
	defer outFile.Close()

	if err := schemaTmpl.Execute(outFile, collections); err != nil {
		log.Errorf("failed to execute template: %v", err)
		// _ = os.Remove("structs.go")
		return
	}
}
