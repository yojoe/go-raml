package nim

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
	"github.com/Jumpscale/go-raml/raml"
)

// Client represents a Nim client
type Client struct {
	APIDef *raml.APIDefinition
	Dir    string
}

// NewClient creates a new Nim client
func NewClient(apiDef *raml.APIDefinition, dir string) Client {
	return Client{
		APIDef: apiDef,
		Dir:    dir,
	}
}

// Generate generates all Nim client files
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

	// services files
	if err := c.generateServices(rs); err != nil {
		return err
	}

	if err := c.generateSecurity(); err != nil {
		return err
	}
	// main client file
	if err := c.generateMain(); err != nil {
		return err
	}
	return nil
}

func (c *Client) generateMain() error {
	filename := filepath.Join(c.Dir, clientName(c.APIDef)+".nim")
	return commons.GenerateFile(c, "./templates/client_nim.tmpl", "client_nim", filename, true)
}

func (c *Client) generateServices(rs []resource) error {
	for _, r := range rs {
		cs := newClientService(r)
		if err := cs.generate(c.Dir); err != nil {
			return err
		}
	}
	return nil
}

// generate security related files
// it currently only supports itsyou.online oauth2
func (c *Client) generateSecurity() error {
	for _, ss := range c.APIDef.SecuritySchemes {
		if v, ok := ss.Settings["accessTokenUri"]; ok && fmt.Sprintf("%v", v) == "https://itsyou.online/v1/oauth/access_token" {
			filename := filepath.Join(c.Dir, "oauth2_client_itsyouonline.nim")
			if err := commons.GenerateFile(ss, "./templates/oauth2_client_itsyouonline.tmpl", "oauth2_client_itsyouonline", filename, true); err != nil {
				return err
			}
		}
	}
	return nil
}

// returns client name of an API definition
func clientName(apiDef *raml.APIDefinition) string {
	splt := strings.Split(apiDef.Title, " ")
	return "client_" + strings.ToLower(splt[0])
}
