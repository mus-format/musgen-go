package adesc

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/parser"
)

func makeSliceDesc(t string, tops *typeops.Options, gops genops.Options,
	m map[AnonymousName]AnonymousDesc,
) (d AnonymousDesc, ok bool) {
	var elemType string
	if elemType, ok = parser.TypeName.ParseSlice(t); ok {
		if elemType == "byte" || elemType == "uint8" {
			return makeByteSliceDesc(t, tops, gops, m)
		}
		var (
			lenSer  = "nil"
			lenVl   = "nil"
			elemVl  = "nil"
			elemOps *typeops.Options
		)
		if tops != nil {
			if tops.LenEncoding != typeops.UndefinedNumEncoding {
				lenSer = tops.LenSer()
			}
			if tops.LenValidator != "" {
				lenVl = validatorStr("int", tops.LenValidator)
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
			Name:     anonymousName("slice", t, gops, tops),
			Kind:     "slice",
			LenSer:   lenSer,
			LenVl:    lenVl,
			ElemType: elemType,
			ElemVl:   elemVl,
			Tops:     tops,
		}
		m[d.Name] = d
	}
	return
}

func makeByteSliceDesc(t string, tops *typeops.Options, gops genops.Options,
	m map[AnonymousName]AnonymousDesc,
) (d AnonymousDesc, ok bool) {
	if tops != nil &&
		(tops.LenEncoding != typeops.UndefinedNumEncoding || tops.LenValidator != "") {
		var (
			lenSer = "nil"
			lenVl  = "nil"
		)
		if tops.LenEncoding != typeops.UndefinedNumEncoding {
			lenSer = tops.LenSer()
		}
		if tops.LenValidator != "" {
			lenVl = validatorStr("int", tops.LenValidator)
		}
		d = AnonymousDesc{
			Name:   anonymousName("byte_slice", t, gops, tops),
			Kind:   "byte_slice",
			LenSer: lenSer,
			LenVl:  lenVl,
			Tops:   tops,
		}
		m[d.Name] = d
		ok = true
	}
	return
}
