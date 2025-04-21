package databuild

import (
	"fmt"
	"strings"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/scanner"
	"github.com/mus-format/musgen-go/typename"
)

func NewToRelNameOp(gops genops.Options) *ToRelNameOp {
	return &ToRelNameOp{&strings.Builder{}, gops}
}

type ToRelNameOp struct {
	b    *strings.Builder
	gops genops.Options
}

func (o *ToRelNameOp) ProcessType(t scanner.Type[typename.FullName],
	tops *typeops.Options) (
	err error) {
	var pkg typename.Pkg
	o.b.WriteString(t.Stars)
	switch t.Kind {
	case scanner.Defined:
		if pkg, err = o.choosePkg(t); err != nil {
			return
		}
		if pkg != "" {
			o.b.WriteString(string(pkg))
			o.b.WriteString(".")
		}
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

func (o *ToRelNameOp) ProcessLeftSquare() {
	o.b.WriteString("[")
}

func (o *ToRelNameOp) ProcessComma() {
	o.b.WriteString(",")
}

func (o *ToRelNameOp) ProcessRightSquare() {
	o.b.WriteString("]")
}

func (o *ToRelNameOp) RelName() typename.RelName {
	return typename.RelName(o.b.String())
}

func (o *ToRelNameOp) choosePkg(t scanner.Type[typename.FullName]) (
	pkg typename.Pkg, err error) {
	if t.Pkg == o.gops.Package() {
		pkg = ""
	} else {
		pkg = t.Pkg
	}
	return
}
