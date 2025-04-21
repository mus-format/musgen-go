package databuild

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	"github.com/mus-format/musgen-go/scanner"
	"github.com/mus-format/musgen-go/typename"
)

type TypeNameConvertor interface {
	ConvertToFullName(cname typename.CompleteName) (fname typename.FullName, err error)
	ConvertToRelName(fname typename.FullName) (rname typename.RelName, err error)
}

func NewConverter(gops genops.Options) Converter {
	return Converter{
		knownPkgs: map[typename.Pkg]typename.PkgPath{gops.Package(): gops.PkgPath},
		gops:      gops}
}

type Converter struct {
	knownPkgs map[typename.Pkg]typename.PkgPath
	gops      genops.Options
}

func (c Converter) ConvertToFullName(cname typename.CompleteName) (
	fname typename.FullName, err error) {
	op := NewToFullNameOp(c.knownPkgs, c.gops)
	if err = scanner.Scan(cname, op, nil); err != nil {
		return
	}
	fname = op.FullName()
	return
}

func (c Converter) ConvertToRelName(fname typename.FullName) (
	rname typename.RelName, err error) {
	op := NewToRelNameOp(c.gops)
	if err = scanner.Scan(fname, op, nil); err != nil {
		return
	}
	rname = op.RelName()
	return
}
