package basegen

import "reflect"

type InterfaceMetadataBuilder interface {
	BuildInterfaceMetadata() (Metadata, error)
}

type InterfaceMetadata struct {
	OneOf []reflect.Type
}

func (m InterfaceMetadata) BuildInterfaceMetadata() (meta Metadata, err error) {
	if len(m.OneOf) < 1 {
		err = ErrEmptyOneOf
		return
	}
	meta = Metadata{
		OneOf: m.OneOf,
	}
	return
}
