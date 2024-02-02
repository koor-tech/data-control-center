package modules

type NewModuleFn = func() error

var Modules = map[string]NewModuleFn{}

type Module interface {
	Run() error
}
