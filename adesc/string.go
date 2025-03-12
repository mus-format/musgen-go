package adesc

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
)

func makeStringDesc(t string, gops genops.Options, tops *typeops.Options,
	m map[AnonymousName]AnonymousDesc,
) (d AnonymousDesc, ok bool) {
	if t == "string" && tops != nil {
		lenSer := "nil"
		if tops.LenEncoding != typeops.UndefinedNumEncoding &&
			tops.LenEncoding != typeops.VarintPositive {
			lenSer = tops.LenSer()
			ok = true
		}
		lenVl := "nil"
		if tops.LenValidator != "" {
			lenVl = validatorStr("int", tops.LenValidator)
			ok = true
		}
		if ok {
			d = AnonymousDesc{
				Name:   anonymousName("string", t, gops, tops),
				Kind:   "string",
				LenSer: lenSer,
				LenVl:  lenVl,
				Tops:   tops,
			}
			m[d.Name] = d
		}
	}
	return
}
