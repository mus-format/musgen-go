package musgen

import (
	"unicode"

	"github.com/mus-format/musgen-go/data"
	genops "github.com/mus-format/musgen-go/options/generate"
	"github.com/mus-format/musgen-go/typename"
)

const (
	SuffixMUS = "MUS"
	SuffixDTS = "DTS"
	SuffixDTM = "DTM"
)

func NewKeywordFns(typeBuilder TypeDataBuilder,
	crossgenTypes map[typename.FullName]struct{}, gops genops.Options) KeywordFns {
	serNames := map[typename.FullName]string{}
	var fullName typename.FullName
	for t, serName := range gops.SerNames { // serName could be with pkg
		fullName = typeBuilder.ToFullName(t)
		serNames[fullName] = serName
	}
	return KeywordFns{serNames, typeBuilder, crossgenTypes}
}

type KeywordFns struct {
	serNames      map[typename.FullName]string
	typeBuilder   TypeDataBuilder
	crossgenTypes map[typename.FullName]struct{}
}

func (f KeywordFns) SerVar(fullName typename.FullName) string {
	return f.serVar(fullName)
}

func (f KeywordFns) SerType(fullName typename.FullName) string {
	return f.uncapitalize(f.SerVar(fullName))
}

func (f KeywordFns) DTSVar(fullName typename.FullName) string {
	return string(f.typeName(fullName)) + SuffixDTS
}

func (f KeywordFns) DTMVar(fullName typename.FullName) string {
	return string(f.typeName(fullName)) + SuffixDTM
}

func (f KeywordFns) SerReceiver(d data.TypeData) string {
	return "v"
}

func (f KeywordFns) TmpVar(d data.TypeData) string {
	return "tmp"
}

func (f KeywordFns) serVar(name typename.FullName) string {
	return string(f.typeName(name)) + SuffixMUS
}

func (f KeywordFns) typeName(name typename.FullName) string {
	if serName, pst := f.serNames[name]; pst {
		return serName
	}
	var rname typename.RelName
	if _, pst := f.crossgenTypes[name]; pst {
		rname = typename.RelName(name.TypeName())
	} else {
		rname = f.typeBuilder.ToRelName(name)
	}
	return rname.WithoutSquares()
}

func (f KeywordFns) uncapitalize(str string) string {
	r := []rune(str)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func (f KeywordFns) RegisterItself(m map[string]any) {
	m["SerVar"] = func(fullName typename.FullName) string {
		return f.SerVar(fullName)
	}
	m["SerType"] = func(fullName typename.FullName) string {
		return f.SerType(fullName)
	}
	m["DTSVar"] = func(fullName typename.FullName) string {
		return f.DTSVar(fullName)
	}
	m["DTMVar"] = func(fullName typename.FullName) string {
		return f.DTMVar(fullName)
	}
	m["SerReceiver"] = func(d data.TypeData) string {
		return f.SerReceiver(d)
	}
	m["TmpVar"] = func(d data.TypeData) string {
		return f.TmpVar(d)
	}
}
