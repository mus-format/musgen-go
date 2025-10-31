// Package introps provides options for registering interface types in the code
// generator.
package introps

import "reflect"

// New creates a new Options.
func New() Options {
	return Options{Impls: []reflect.Type{}}
}

// Options contains configuration settings for the interface type.
type Options struct {
	Impls      []reflect.Type
	Marshaller bool
}

type SetOption func(o *Options)

// WithImpl returns a SetOption that registers a concrete implementation type
// for the interface being generated. The provided type must implement the
// target interface.
func WithImpl(t reflect.Type) SetOption {
	return func(o *Options) { o.Impls = append(o.Impls, t) }
}

// WithMarshaller returns a SetOption that enables serialization using the
// Marhsller interface (from the ext module).
func WithMarshaller() SetOption {
	return func(o *Options) { o.Marshaller = true }
}

func Apply(ops []SetOption, o *Options) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
