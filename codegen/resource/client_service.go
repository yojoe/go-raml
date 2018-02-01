package resource

import (
	"strings"

	"github.com/Jumpscale/go-raml/codegen/commons"
)

// ClientService defines a client service.
// client service is a module that handle
// a single root endpoint
type ClientService struct {
	Name          string
	EndpointName  string
	FilenameNoExt string
}

func NewClientService(rootEndpoint string) ClientService {
	normalizedEndpoint := commons.NormalizeIdentifier(commons.NormalizeURI(rootEndpoint))
	return ClientService{
		Name:          strings.Title(normalizedEndpoint) + "Service",
		EndpointName:  normalizedEndpoint,
		FilenameNoExt: normalizedEndpoint + "_service",
	}
}
