package python

import (
	log "github.com/Sirupsen/logrus"
)

type clientTemplate struct {
	serviceFile string
	serviceName string
	oauth2File  string
	oauth2Name  string
	initFile    string
	initName    string
	mainFile    string
	mainName    string
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
		serviceFile: "./templates/python/client_service.tmpl",
		serviceName: "client_service",
		oauth2File:  "./templates/python/oauth2_client_python.tmpl",
		oauth2Name:  "oauth2_client_python",
		initFile:    "./templates/python/client_initpy_python.tmpl",
		initName:    "client_initpy_python",
		mainFile:    "./templates/python/client_python.tmpl",
		mainName:    "client_python",
	}
}

func (c *Client) initTemplatesAioHTTP() {
	c.Template = clientTemplate{
		serviceFile: "./templates/python/client_service_aiohttp.tmpl",
		serviceName: "client_service_aiohttp",
		oauth2File:  "./templates/python/oauth2_client_python_aiohttp.tmpl",
		oauth2Name:  "oauth2_client_python_aiohttp",
		initFile:    "./templates/python/client_initpy_python_aiohttp.tmpl",
		initName:    "client_initpy_python_aiohttp",
		mainFile:    "./templates/python/client_python_aiohttp.tmpl",
		mainName:    "client_python_aiohttp",
	}

}
