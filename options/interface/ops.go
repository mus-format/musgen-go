package introps

import "reflect"

func NewOptions(t reflect.Type) Options {
	return Options{
		t:         t,
		ImplTypes: []reflect.Type{},
	}
}

type Options struct {
	t          reflect.Type
	ImplTypes  []reflect.Type
	Marshaller bool
}

func (o Options) ImplTypeStr(i int) string {
	if o.t.PkgPath() != o.ImplTypes[i].PkgPath() {
		return o.ImplTypes[i].String()
	} else {
		return o.ImplTypes[i].Name()
	}
}

type SetOption func(o *Options)

func WithImplType(impl reflect.Type) SetOption {
	return func(o *Options) {
		o.ImplTypes = append(o.ImplTypes, impl)
	}
}

func WithMarshaller() SetOption {
	return func(o *Options) {
		o.Marshaller = true
	}
}

func Apply(ops []SetOption, o *Options) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
