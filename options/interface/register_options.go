package introps

import (
	"reflect"

	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
)

// NewRegisterOptions creates a new RegisterOptions.
func NewRegisterOptions() RegisterOptions {
	return RegisterOptions{
		StructImpls:      []StructImpl{},
		DefinedTypeImpls: []DefinedTypeImpl{},
	}
}

type RegisterOptions struct {
	StructImpls      []StructImpl
	DefinedTypeImpls []DefinedTypeImpl
	Marshaller       bool
}

type StructImpl struct {
	Type reflect.Type
	Ops  []structops.SetOption
}

type DefinedTypeImpl struct {
	Type reflect.Type
	Ops  []typeops.SetOption
}

type SetRegisterOption func(o *RegisterOptions)

// WithStructImpl returns a SetRegisterOption that registers a concrete
// implementation type for the interface being generated. The provided type
// must implement the target interface.
func WithStructImpl(t reflect.Type, ops ...structops.SetOption) SetRegisterOption {
	impl := StructImpl{Type: t, Ops: ops}
	return func(o *RegisterOptions) { o.StructImpls = append(o.StructImpls, impl) }
}

// WithDefinedTypeImpl returns a SetRegisterOption that registers a concrete
// implementation type for the interface being generated. The provided type
// must implement the target interface.
func WithDefinedTypeImpl(t reflect.Type, ops ...typeops.SetOption) SetRegisterOption {
	impl := DefinedTypeImpl{Type: t, Ops: ops}
	return func(o *RegisterOptions) {
		o.DefinedTypeImpls = append(o.DefinedTypeImpls,
			impl)
	}
}

// WithRegisterMarshaller returns a SetOption that enables serialization using
// the Marhsller interface (from the ext module).
func WithRegisterMarshaller() SetRegisterOption {
	return func(o *RegisterOptions) { o.Marshaller = true }
}

func ApplyRegister(ops []SetRegisterOption, o *RegisterOptions) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
