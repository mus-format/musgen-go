package generator

import (
	"fmt"
	"strings"

	"github.com/mus-format/musgen-go/basegen"
)

func NewNumGenerator(conf basegen.Conf, tp string, opts *basegen.Options) (
	g NumGenerator) {
	g.conf = conf
	g.tp = tp
	if opts != nil {
		g.enc = opts.Encoding
	}
	return
}

type NumGenerator struct {
	conf basegen.Conf
	tp   string
	enc  basegen.NumEncoding
}

func (g NumGenerator) GenerateFnName(fnType basegen.FnType) (name string) {
	pkg := g.enc.Package()
	if g.conf.Unsafe && g.enc == basegen.UndefinedNumEncoding {
		pkg = "unsafe"
	}
	typeName := strings.ToUpper(g.tp[:1]) + g.tp[1:]
	return fmt.Sprintf("%s.%s%s", pkg, g.enc.FuncName(fnType), typeName)
}

func (g NumGenerator) GenerateMarshalCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Marshal)
		params = fmt.Sprintf("(%s %s)", param(vname), g.conf.MarshalParam())
	)
	return name + params
}

func (g NumGenerator) GenerateUnmarshalCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Unmarshal)
		params = fmt.Sprintf("(%s)", g.conf.UnmarshalParam())
	)
	return name + params
}

func (g NumGenerator) GenerateSizeCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Size)
		params = fmt.Sprintf("(%s)", vname)
	)
	return name + params
}

func (g NumGenerator) GenerateSkipCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Skip)
		params = fmt.Sprintf("(%s)", g.conf.SkipParam())
	)
	return name + params
}
