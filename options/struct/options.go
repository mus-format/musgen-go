// Package structops provides options for registering struct types in the
// code generator.
package structops

import (
	typeops "github.com/mus-format/musgen-go/options/type"
)

// New creates a new Options.
func New() Options {
	return Options{Fields: []*typeops.Options{}}
}

// Options contains configuration settings for the struct type. It allows
// specifying both top-level type options and per-field customizations.
type Options struct {
	Fields []*typeops.Options
	Tops   *typeops.Options
}

type SetOption func(o *Options)

// WithField returns a SetOption that appends field-specific configuration to
// the Options. Use WithField without params if there are no options for the
// field.
func WithField(ops ...typeops.SetOption) SetOption {
	return func(o *Options) {
		if len(ops) == 0 {
			o.Fields = append(o.Fields, nil)
			return
		}
		fo := &typeops.Options{}
		typeops.Apply(ops, fo)
		o.Fields = append(o.Fields, fo)
	}
}

// WithTops returns a SetOption that configures struct-specific serialization.
func WithTops(ops ...typeops.SetOption) SetOption {
	return func(o *Options) {
		if len(ops) == 0 {
			return
		}
		o.Tops = &typeops.Options{}
		typeops.Apply(ops, o.Tops)
	}
}

func Apply(ops []SetOption, o *Options) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
