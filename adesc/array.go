package adesc

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/parser"
)

func makeArrayDesc(t string, tops *typeops.Options, gops genops.Options,
	m map[AnonymousName]AnonymousDesc) (d AnonymousDesc, ok bool) {
	var (
		elemType string
		length   int
	)
	if elemType, length, ok = parser.TypeName.ParseArray(t); ok {
		var (
			lenSer  = "nil"
			elemVl  = "nil"
			elemOps *typeops.Options
		)
		if tops != nil {
			if tops.LenEncoding != typeops.UndefinedNumEncoding {
				lenSer = tops.LenSer()
			}
			if tops.Elem != nil {
				if tops.Elem.Validator != "" {
					elemVl = validatorStr(elemType, tops.Elem.Validator)
				}
				elemOps = tops.Elem
			}
		}
		subMake(elemType, elemOps, gops, m)
		d = AnonymousDesc{
			Name:      anonymousName("array", t, gops, tops),
			Kind:      "array",
			Type:      t,
			ArrLength: length,
			LenSer:    lenSer,
			ElemType:  elemType,
			ElemVl:    elemVl,
			Tops:      tops,
		}
		m[d.Name] = d
	}
	return
}
