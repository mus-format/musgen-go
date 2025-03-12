package tdesc

import (
	"reflect"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	"github.com/mus-format/musgen-go/parser"
)

func MakeInterfaceDesc(t reflect.Type, iops introps.Options,
	gops genops.Options) (td TypeDesc, err error) {
	if len(iops.ImplTypes) == 0 {
		err = ErrNoOneImplTypeSpecified
		return
	}
	for i := range iops.ImplTypes {
		if err = parser.SupportedType(iops.ImplTypes[i]); err != nil {
			return
		}
	}

	err = parser.ParseInterfaceType(t)
	if err != nil {
		return
	}
	td = makeDesc(t, gops)
	if err = checkOneof(iops); err != nil {
		return
	}
	td.Oneof = buildOneof(t, iops)
	return
}

func buildOneof(t reflect.Type, iops introps.Options) (oneof []string) {
	oneof = make([]string, len(iops.ImplTypes))
	for i, oneOf := range iops.ImplTypes {
		if t.PkgPath() != oneOf.PkgPath() {
			oneof[i] = oneOf.String()
		} else {
			oneof[i] = oneOf.Name()
		}
	}
	return
}

func checkOneof(iops introps.Options) (err error) {
	for _, implType := range iops.ImplTypes {
		if err = parser.SupportedType(implType); err != nil {
			return
		}
	}
	return
}
