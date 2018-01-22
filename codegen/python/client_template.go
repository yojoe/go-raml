package python

import (
	log "github.com/Sirupsen/logrus"
)

type clientTemplate struct {
	serviceFile     string
	serviceName     string
	basicAuthFile   string
	basicAuthName   string
	passThroughFile string
	passThroughName string
	oauth2File      string
	oauth2Name      string
	initFile        string
	initName        string
	httpClientFile  string
	httpClientName  string
}

func (c *Client) initTemplates() {
	switch c.Kind {
	case clientNameRequests, clientNameGeventRequests:
		c.initTemplatesRequests()
	case clientNameAiohttp:
		c.initTemplatesAioHTTP()
	default:
		log.Fatalf("invalid client kind:%v", c.Kind)
	}
}

func (c *Client) initTemplatesRequests() {
	c.Template = clientTemplate{
		serviceFile:     "./templates/python/client_service.tmpl",
		serviceName:     "client_service",
		oauth2File:      "./templates/python/oauth2_client_python.tmpl",
		oauth2Name:      "oauth2_client_python",
		basicAuthFile:   "./templates/python/basicauth_client.tmpl",
		basicAuthName:   "basicauth_client",
		passThroughFile: "./templates/python/passthrough_auth_client.tmpl",
		passThroughName: "passthrough_auth_client",
		initFile:        "./templates/python/client_initpy_python.tmpl",
		initName:        "client_initpy_python",
		httpClientFile:  "./templates/python/httpclient_requests.tmpl",
		httpClientName:  "httpclient_requests",
	}
}

func (c *Client) initTemplatesAioHTTP() {
	c.Template = clientTemplate{
		serviceFile:     "./templates/python/client_service_aiohttp.tmpl",
		serviceName:     "client_service_aiohttp",
		oauth2File:      "./templates/python/oauth2_client_python_aiohttp.tmpl",
		oauth2Name:      "oauth2_client_python_aiohttp",
		basicAuthFile:   "./templates/python/basicauth_client.tmpl",
		basicAuthName:   "basicauth_client",
		passThroughFile: "./templates/python/passthrough_auth_client.tmpl",
		passThroughName: "passthrough_auth_client",
		initFile:        "./templates/python/client_initpy_python_aiohttp.tmpl",
		initName:        "client_initpy_python_aiohttp",
		httpClientFile:  "./templates/python/httpclient_aiohttp.tmpl",
		httpClientName:  "httpclient_aiohttp",
	}

}
