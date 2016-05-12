package raml

// Library is used to combine any collection of data type declarations,
// resource type declarations, trait declarations, and security scheme declarations
// into modular, externalized, reusable groups.
// While libraries are intended to define common declarations in external documents,
// which are then included where needed, libraries can also be defined inline.
type Library struct {
	Types           map[string]Type `yaml:"types"`
	Schemas         []map[string]string
	ResourceTypes   map[string]ResourceType     `yaml:"resourceTypes"`
	Traits          map[string]Trait            `yaml:"traits"`
	SecuritySchemes []map[string]SecurityScheme `yaml:"securitySchemes"`
	Uses            map[string]string           `yaml:"uses"`

	// Describes the content or purpose of a specific library.
	// The value is a string and MAY be formatted using markdown.
	Usage string `yaml:"usage"`

	Libraries map[string]*Library `yaml:"-"`

	ramlFile string
}

func (l *Library) PostProcess(filename string) error {
	l.ramlFile = filename

	// traits
	for name, t := range l.Traits {
		t.postProcess(name)
		l.Traits[name] = t
	}

	// resource types
	for name, rt := range l.ResourceTypes {
		rt.postProcess(name)
		l.ResourceTypes[name] = rt
	}
	return nil
}

// get type by the name
func (l *Library) getType(name string) *Type {
	for k, t := range l.Types {
		if k == name {
			return &t
		}
	}
	return nil
}

// get schema by the name
func (l *Library) getSchema(name string) *string {
	for _, schemas := range l.Schemas {
		for k, schema := range schemas {
			if k == name {
				return &schema
			}
		}
	}
	return nil
}

// get security scheme by the name
func (l *Library) getSecurityScheme(name string) *SecurityScheme {
	for _, schemes := range l.SecuritySchemes {
		for k, s := range schemes {
			if k == name {
				return &s
			}
		}
	}
	return nil
}
