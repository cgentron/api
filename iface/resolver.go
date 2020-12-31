package iface

import (
	"context"
	"reflect"
)

// Manifest The plugin manifest.
type Manifest struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Import      string `yaml:"import"`
	Description string `yaml:"summary"`
}

// Constructor ...
type Constructor func() (ResolverHandler, error)

// ResolverHandler is implementing the resolution of a input message
// to an output messsage.
type ResolverHandler interface {
	// Resolve is resolving an input message to an output message.
	Resolve(context.Context, interface{}, interface{}) error
}

// Symbols variable stores the map of stdlib symbols per package.
var Symbols = map[string]map[string]reflect.Value{}

func init() {
	Symbols["github.com/cgentron/proto/iface"] = map[string]reflect.Value{
		"ResolverHandler":  reflect.ValueOf((*ResolverHandler)(nil)),
		"_ResolverHandler": reflect.ValueOf((*_iface_resolver_Handler)(nil)),
	}
}

// _iface_resolver_Handler is an interface wrapper for Handler type
type _iface_resolver_Handler struct {
	WResolve func(a0 context.Context, a1 interface{}) (interface{}, error)
}

func (W _iface_resolver_Handler) Resolve(a0 context.Context, a1 interface{}) (interface{}, error) {
	return W.WResolve(a0, a1)
}
