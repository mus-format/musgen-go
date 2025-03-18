package structops

import (
	typeops "github.com/mus-format/musgen-go/options/type"
)

func NewOptions() Options {
	return Options{
		Fields: []*typeops.Options{},
	}
}

type Options struct {
	Type   *TypeOps
	Fields []*typeops.Options
}

type TypeOps struct {
	SourceType SourceType
	Ops        *typeops.Options
}

type SetOption func(o Options) Options

func WithField(ops ...typeops.SetOption) SetOption {
	return func(o Options) Options {
		if len(ops) > 0 {
			fo := &typeops.Options{}
			typeops.Apply(ops, fo)
			o.Fields = append(o.Fields, fo)
		}
		return o
	}
}

func WithNil() SetOption {
	return func(o Options) Options {
		o.Fields = append(o.Fields, nil)
		return o
	}
}

func WithSourceType(t SourceType, ops ...typeops.SetOption) SetOption {
	return func(o Options) Options {
		o.Type = &TypeOps{
			SourceType: t,
		}
		if len(ops) > 0 {
			o.Type.Ops = &typeops.Options{}
			typeops.Apply(ops, o.Type.Ops)
		}
		return o
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
