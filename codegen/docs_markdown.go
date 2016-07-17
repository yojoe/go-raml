package codegen

import (
	"fmt"
	"github.com/Jumpscale/go-raml/raml"
	"log"
)

type markdownDocs struct {
	api    *raml.APIDefinition
	output string
}

func (d *markdownDocs) flatten(resources map[string]*raml.Resource, base string, out map[string]*raml.Resource) {
	for name, r := range resources {
		fullName := fmt.Sprintf("%s%s", base, name)
		if len(r.Methods) > 0 {
			out[fullName] = r
		}

		d.flatten(r.Nested, fullName, out)
	}
}

func (d *markdownDocs) generate() error {
	//d.api.BaseURI
	for _, v := range d.api.Resources {
		log.Printf("Resource: %s", v.DisplayName)

		for _, m := range v.Methods {
			log.Printf("%s:%s", m.Name, m.DisplayName)
		}
	}

	resources := make(map[string]*raml.Resource)
	for k, v := range d.api.Resources {
		func(resource raml.Resource) {
			resources[k] = &resource
		}(v)
	}
	flat := make(map[string]*raml.Resource)
	d.flatten(resources, "", flat)
	log.Printf("Flat resource: %s", flat)
	return generateFile(map[string]interface{}{"Api": d.api, "Resources": flat}, "./templates/docs_markdown.tmpl", "docs_markdown", d.output, true)
}
