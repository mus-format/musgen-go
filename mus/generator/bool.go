package generator

import (
	"fmt"
	"strings"

	"github.com/mus-format/musgen-go/basegen"
)

func NewBoolGenerator(conf basegen.Conf, tp string, opts *basegen.Options) (
	g BoolGenerator) {
	g.conf = conf
	g.tp = tp
	return
}

type BoolGenerator struct {
	conf basegen.Conf
	tp   string
}

func (g BoolGenerator) GenerateFnName(fnType basegen.FnType) (name string) {
	var (
		pkg      = "ord"
		typeName = strings.ToUpper(g.tp[:1]) + g.tp[1:]
	)
	if g.conf.Unsafe {
		pkg = "unsafe"
	}
	return fmt.Sprintf("%s.%s%s", pkg, string(fnType), typeName)
}

func (g BoolGenerator) GenerateMarshalCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Marshal)
		params = fmt.Sprintf("(%s %s)", param(vname), g.conf.MarshalParam())
	)
	return name + params
}

func (g BoolGenerator) GenerateUnmarshalCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Unmarshal)
		params = fmt.Sprintf("(%s)", g.conf.UnmarshalParam())
	)
	return name + params
}

func (g BoolGenerator) GenerateSizeCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Size)
		params = fmt.Sprintf("(%s)", vname)
	)
	return name + params
}

func (g BoolGenerator) GenerateSkipCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Skip)
		params = fmt.Sprintf("(%s)", g.conf.SkipParam())
	)
	return name + params
}
