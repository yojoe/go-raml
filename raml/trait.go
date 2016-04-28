package raml

// A Trait is a partial method definition that, like a method, can provide
// method-level properties such as description, headers, query string
// parameters, and responses. Methods that use one or more traits inherit
// those traits' properties.
type Trait struct {
	Name string
	// TODO: Fill this during the post-processing phase

	// The usage property of a resource type or trait is used to describe how
	// the resource type or trait should be used
	Usage string

	// Briefly describes what the method does to the resource
	Description string

	// As in Method.
	Bodies Bodies `yaml:"body"`

	// As in Method.
	Headers map[HTTPHeader]Header `yaml:"headers"`

	// As in Method.
	Responses map[HTTPCode]Response `yaml:"responses"`

	// As in Method.
	QueryParameters map[string]NamedParameter `yaml:"queryParameters"`

	// As in Method.
	Protocols []string `yaml:"protocols"`

	// When defining resource types and traits, it can be useful to capture
	// patterns that manifest several levels below the inheriting resource or
	// method, without requiring the creation of the intermediate levels.
	// For example, a resource type definition may describe a body parameter
	// that will be used if the API defines a post method for that resource,
	// but the processing application should not create the post method itself.
	//
	// This optional structure key indicates that the value of the property
	// should be applied if the property name itself (without the question
	// mark) is already defined (whether explicitly or implicitly) at the
	// corresponding level in that resource or method.
	OptionalBodies          Bodies                    `yaml:"body?"`
	OptionalHeaders         map[HTTPHeader]Header     `yaml:"headers?"`
	OptionalResponses       map[HTTPCode]Response     `yaml:"responses?"`
	OptionalQueryParameters map[string]NamedParameter `yaml:"queryParameters?"`
}
