package tdesc

import (
	"reflect"

	genops "github.com/mus-format/musgen-go/options/generate"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/parser"
)

func MakeStructDesc(t reflect.Type, sops structops.Options,
	gops genops.Options) (td TypeDesc, err error) {
	fieldTypes, err := parser.ParseStructType(t)
	if err != nil {
		return
	}
	if sops.Fields != nil && len(sops.Fields) != len(fieldTypes) {
		err = ErrWrongOptionsAmount
		return
	}
	td = makeDesc(t, gops)
	td.Sops = sops
	td.Fields = buildStructFields(t, fieldTypes, sops)
	return
}

func buildStructFields(t reflect.Type, types []string, sops structops.Options) (
	fields []FieldDesc) {
	fields = make([]FieldDesc, len(types))
	for i := range fields {
		fields[i] = FieldDesc{
			Name: t.Field(i).Name,
			Type: types[i],
		}
		if len(sops.Fields) > 0 {
			fields[i].Tops = (*typeops.Options)(sops.Fields[i])
		}
	}
	return
}
