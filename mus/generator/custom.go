package generator

import (
	"fmt"
	"strings"

	"github.com/mus-format/musgen-go/basegen"
)

func NewSAIGenerator(conf basegen.Conf, tp, prefix string, opts *basegen.Options) (
	g CustomGenerator) {
	return CustomGenerator{conf, tp, basegen.Prefix(prefix, opts)}
}

type CustomGenerator struct {
	conf   basegen.Conf
	tp     string
	prefix string
}

func (g CustomGenerator) GenerateFnName(fnType basegen.FnType) (name string) {
	if sl := strings.Split(g.tp, "."); len(sl) == 2 {
		var (
			pkg = sl[0]
			tp  = sl[1]
		)
		return fmt.Sprintf("%s.%s%s%sMUS", pkg, fnType, g.prefix, tp)
	}
	return fmt.Sprintf("%s%s%sMUS", fnType, g.prefix, g.tp)
}

func (g CustomGenerator) GenerateMarshalCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Marshal)
		params = fmt.Sprintf("(%s %s)", param(vname), g.conf.MarshalParam())
	)
	return name + params
}

func (g CustomGenerator) GenerateUnmarshalCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Unmarshal)
		params = fmt.Sprintf("(%s)", g.conf.UnmarshalParam())
	)
	return name + params
}

func (g CustomGenerator) GenerateSizeCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Size)
		params = fmt.Sprintf("(%s)", vname)
	)
	return name + params
}

func (g CustomGenerator) GenerateSkipCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Skip)
		params = fmt.Sprintf("(%s)", g.conf.SkipParam())
	)
	return name + params
}
