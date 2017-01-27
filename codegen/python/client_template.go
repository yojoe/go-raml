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
	case clientNameRequests:
		c.initTemplatesRequests()
	case clientNameAiohttp:
		c.initTemplatesAioHTTP()
	default:
		log.Fatalf("invalid client kind:%v", c.Kind)
	}
}

func (c *Client) initTemplatesRequests() {
	c.Template = clientTemplate{
		serviceFile: "./templates/client_service_python.tmpl",
		serviceName: "client_service_python",
		oauth2File:  "./templates/oauth2_client_python.tmpl",
		oauth2Name:  "oauth2_client_python",
		initFile:    "./templates/client_initpy_python.tmpl",
		initName:    "client_initpy_python",
		mainFile:    "./templates/client_python.tmpl",
		mainName:    "client_python",
	}
}

func (c *Client) initTemplatesAioHTTP() {
	c.Template = clientTemplate{
		serviceFile: "./templates/client_service_python_aiohttp.tmpl",
		serviceName: "client_service_python_aiohttp",
		oauth2File:  "./templates/oauth2_client_python_aiohttp.tmpl",
		oauth2Name:  "oauth2_client_python_aiohttp",
		initFile:    "./templates/client_initpy_python_aiohttp.tmpl",
		initName:    "client_initpy_python_aiohttp",
		mainFile:    "./templates/client_python_aiohttp.tmpl",
		mainName:    "client_python_aiohttp",
	}

}
