package adesc

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/parser"
)

func makeMapDesc(t string, tops *typeops.Options, gops genops.Options,
	m map[AnonymousName]AnonymousDesc,
) (d AnonymousDesc, ok bool) {
	var keyType, elemType string
	if keyType, elemType, ok = parser.TypeName.ParseMap(t); ok {
		var (
			lenSer  = "nil"
			lenVl   = "nil"
			keyVl   = "nil"
			elemVl  = "nil"
			keyOps  *typeops.Options
			elemOps *typeops.Options
		)
		if tops != nil {
			if tops.LenEncoding != typeops.UndefinedNumEncoding {
				lenSer = tops.LenSer()
			}
			if tops.LenValidator != "" {
				lenVl = validatorStr("int", tops.LenValidator)
			}
			if tops.Key != nil {
				if tops.Key.Validator != "" {
					keyVl = validatorStr(keyType, tops.Key.Validator)
				}
				keyOps = tops.Key
			}
			if tops.Elem != nil {
				if tops.Elem.Validator != "" {
					elemVl = validatorStr(elemType, tops.Elem.Validator)
				}
				elemOps = tops.Elem
			}
		}
		subMake(keyType, keyOps, gops, m)
		subMake(elemType, elemOps, gops, m)
		d = AnonymousDesc{
			Name:     anonymousName("map", t, gops, tops),
			Kind:     "map",
			LenSer:   lenSer,
			LenVl:    lenVl,
			KeyType:  keyType,
			KeyVl:    keyVl,
			ElemType: elemType,
			ElemVl:   elemVl,
			Tops:     tops,
		}
		m[d.Name] = d
	}
	return
}
