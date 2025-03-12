package adesc

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/parser"
)

func makePtrDesc(t string, tops *typeops.Options, gops genops.Options,
	m map[AnonymousName]AnonymousDesc,
) (d AnonymousDesc, ok bool) {
	var elemType string
	if _, elemType, ok = parser.TypeName.ParsePtr(t); ok {
		var elemOps *typeops.Options
		if tops != nil {
			elemOps = tops.Elem
		}
		subMake(elemType, elemOps, gops, m)
		d = AnonymousDesc{
			Name:     anonymousName("ptr", t, gops, tops),
			Kind:     "ptr",
			ElemType: elemType,
			Tops:     tops,
		}
		m[d.Name] = d
	}
	return
}
