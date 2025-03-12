package structops

import (
	typeops "github.com/mus-format/musgen-go/options/type"
)

func NewOptions() Options {
	return []*typeops.Options{}
}

type Options []*typeops.Options

type SetOption func(o Options) Options

func WithField(ops ...typeops.SetOption) SetOption {
	return func(fo Options) Options {
		o := &typeops.Options{}
		typeops.Apply(ops, o)
		return append(fo, o)
	}
}

func WithNil() SetOption {
	return func(o Options) Options {
		return append(o, nil)
	}
}

func Apply(ops []SetOption, o Options) Options {
	for i := range ops {
		if ops[i] != nil {
			o = ops[i](o)
		}
	}
	return o
}
