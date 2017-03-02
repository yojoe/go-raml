package angular

type clientTemplate struct {
	serviceFile string
	serviceName string
	mainFile    string
	mainName    string
	staticDir   string
}

func (c *Client) initTemplates() {
	c.Template = clientTemplate{
		serviceFile: "./templates/angular/templates/client_service_angular.tmpl",
		serviceName: "client_service_angular",
		mainFile:    "./templates/angular/templates/app_module_angular.tmpl",
		mainName:    "app_module_angular",
		staticDir:   "templates/angular/static",
	}
}
