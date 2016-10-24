package nim

import (
	"path/filepath"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

type Client struct {
	APIDef *raml.APIDefinition
	Dir    string
}

func (c *Client) Generate() error {
	rs := getAllResources(c.APIDef, false)

	// generate all objects from all RAML types
	if err := generateObjects(c.APIDef.Types, c.Dir); err != nil {
		return err
	}

	// generate all objects from request/response body
	if _, err := generateObjectsFromBodies(rs, c.Dir); err != nil {
		return err
	}

	// main client file
	if err := c.generateMain(); err != nil {
		return err
	}
	return nil
}

func (c *Client) generateMain() error {
	filename := filepath.Join(c.Dir, "client.nim")
	return commons.GenerateFile(c, "./templates/client_nim.tmpl", "client_nim", filename, true)
}
