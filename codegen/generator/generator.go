package generator

// Server represent server code generator
type Server interface {
	// Generate generates code
	Generate() error

	// APIDocsDir returns directory of APIDocs
	APIDocsDir() string
}
