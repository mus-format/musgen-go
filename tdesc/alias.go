package tdesc

import (
	"reflect"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/parser"
)

func MakeTypeDefDesc(t reflect.Type, tops *typeops.Options,
	gops genops.Options) (td TypeDesc, err error) {
	sourceType, err := parser.ParseTypedef(t)
	if err != nil {
		return
	}

	td = makeDesc(t, gops)
	td.SourceType = sourceType
	td.Fields = []FieldDesc{{Type: sourceType, Tops: tops}}
	return
}
