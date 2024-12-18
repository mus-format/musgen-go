package basegen

import "reflect"

type InterfaceOptionsBuilder interface {
	BuildInterfaceOptions() (Options, error)
}

type InterfaceOptions struct {
	Oneof []reflect.Type
}

func (m InterfaceOptions) BuildInterfaceOptions() (opts Options, err error) {
	if len(m.Oneof) < 1 {
		err = ErrEmptyOneof
		return
	}
	opts = Options{
		Oneof: m.Oneof,
	}
	return
}
