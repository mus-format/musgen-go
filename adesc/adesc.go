package adesc

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
)

type AnonymousName string

type AnonymousDesc struct {
	Name      AnonymousName
	Kind      string
	Type      string
	ArrLength int
	LenSer    string
	LenVl     string
	KeyType   string
	KeyVl     string
	ElemType  string
	ElemVl    string
	Tops      *typeops.Options
}

func Make(t string, gops genops.Options, tops *typeops.Options) (
	d AnonymousDesc, ok bool) {
	return Collect(t, tops, gops, map[AnonymousName]AnonymousDesc{})
}

func Collect(t string, tops *typeops.Options, gops genops.Options,
	m map[AnonymousName]AnonymousDesc) (d AnonymousDesc, ok bool) {
	d, ok = makeStringDesc(t, gops, tops, m)
	if ok {
		return
	}
	d, ok = makeArrayDesc(t, tops, gops, m)
	if ok {
		return
	}
	d, ok = makeSliceDesc(t, tops, gops, m)
	if ok {
		return
	}
	d, ok = makeMapDesc(t, tops, gops, m)
	if ok {
		return
	}
	d, ok = makePtrDesc(t, tops, gops, m)
	return
}
