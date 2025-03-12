package introps

import "reflect"

func NewOptions() Options {
	return Options{
		ImplTypes: []reflect.Type{},
	}
}

type Options struct {
	ImplTypes []reflect.Type
}

type SetOption func(o *Options)

func WithImplType(impl reflect.Type) SetOption {
	return func(o *Options) {
		o.ImplTypes = append(o.ImplTypes, impl)
	}
}

func Apply(ops []SetOption, o *Options) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
