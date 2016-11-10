package capnp

var (
	allTypes = map[string]struct{}{} // all registered types
)

func registerType(name string) {
	allTypes[name] = struct{}{}
}

func typesRegistered(name string) bool {
	_, ok := allTypes[name]
	return ok
}
