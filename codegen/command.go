package codegen

const (
	langGo        = "go"
	langPython    = "python"
	langNim       = "nim"
	langTarantool = "tarantool"
)

//Command is a toplevel command to be executed by the cli's main routine
type Command interface {
	//Execute the command
	Execute() error
}
