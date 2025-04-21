package data

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
)

type TypeData struct {
	FullName       typename.FullName // With pkg, with generics
	SourceFullName typename.FullName
	Tops           *typeops.Options

	Fields []FieldData
	Sops   structops.Options

	Impls []typename.FullName
	Iops  introps.Options

	Gops genops.Options
}

func (d TypeData) SerializedFields() (sl []FieldData) {
	if len(d.Sops.Fields) == 0 {
		return d.Fields
	}
	sl = make([]FieldData, 0, len(d.Fields))
	for i, f := range d.Fields {
		if d.Sops.Fields[i] == nil || !d.Sops.Fields[i].Ignore {
			sl = append(sl, f)
		}
	}
	return
}
