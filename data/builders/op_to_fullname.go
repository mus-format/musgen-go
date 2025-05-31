package builders

import (
	"fmt"
	"log"
	"strings"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/scanner"
	"github.com/mus-format/musgen-go/typename"
)

func NewToFullNameOp(knownPkgs map[typename.Package]typename.PkgPath,
	gops genops.Options) *ToFullNameOp {
	return &ToFullNameOp{&strings.Builder{}, knownPkgs, gops}
}

type ToFullNameOp struct {
	b         *strings.Builder
	knownPkgs map[typename.Package]typename.PkgPath
	gops      genops.Options
}

func (o *ToFullNameOp) ProcessType(t scanner.Type[typename.CompleteName],
	tops *typeops.Options) (
	err error) {
	var pkg typename.Package
	o.b.WriteString(t.Stars)
	switch t.Kind {
	case scanner.Defined:
		if pkg, err = o.choosePkg(t); err != nil {
			return
		}
		o.b.WriteString(string(pkg))
		o.b.WriteString(".")
		o.b.WriteString(string(t.Name))
	case scanner.Array:
		o.b.WriteString("[")
		o.b.WriteString(t.ArrLength)
		o.b.WriteString("]")
	case scanner.Slice:
		o.b.WriteString("[]")
	case scanner.Map:
		o.b.WriteString("map")
	case scanner.Prim:
		o.b.WriteString(string(t.Name))
	default:
		return fmt.Errorf("unexpected %v kind", t.Kind)
	}
	return
}

func (o *ToFullNameOp) ProcessLeftSquare() {
	o.b.WriteString("[")
}

func (o *ToFullNameOp) ProcessComma() {
	o.b.WriteString(",")
}

func (o *ToFullNameOp) ProcessRightSquare() {
	o.b.WriteString("]")
}

func (o *ToFullNameOp) FullName() typename.FullName {
	return typename.FullName(o.b.String())
}

func (o *ToFullNameOp) choosePkg(t scanner.Type[typename.CompleteName]) (
	pkg typename.Package, err error) {
	if alias, pst := o.gops.ImportAliases()[genops.ImportPath(t.PkgPath)]; pst {
		pkg = typename.Package(alias)
	} else {
		pkg = t.Package
		if t.Position == scanner.Param {
			log.Printf("WARNING: no alias for '%v' in musgen.CodeGenerator options\n", t.PkgPath)
		}
	}
	if err = o.checkPkg(pkg, t); err != nil {
		return
	}
	o.knownPkgs[pkg] = t.PkgPath
	return
}

func (o *ToFullNameOp) checkPkg(pkg typename.Package,
	t scanner.Type[typename.CompleteName]) (err error) {
	if pkgPath, pst := o.knownPkgs[pkg]; pst && pkgPath != t.PkgPath {
		err = NewTwoPathsSameAliasError(pkgPath, t.PkgPath, pkg)
	}
	return
}
