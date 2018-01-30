package resource

const (
	// go-raml extension to express Catch-All URL params
	// for example /path/{path:*} is going to match:
	//  - /path/a    -> path = a
	//  - /path/a/b  -> path = a/b
	//  - /path/a/b/c -> path = a/b/c
	CatchAllRoute = "{path:*}"
)
