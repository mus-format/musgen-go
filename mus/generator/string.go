package generator

import (
	"fmt"

	"github.com/mus-format/musgen-go/basegen"
)

func NewStringGenerator(conf basegen.Conf, opts *basegen.Options) (
	g StringGenerator) {
	g.conf = conf
	g.lenM = "nil"
	g.lenU = "nil"
	g.lenS = "nil"

	modImportName := conf.ModImportName()

	if opts != nil {
		if opts.LenEncoding != 0 {
			numG := NewNumGenerator(conf, "int", &basegen.Options{
				Encoding: opts.LenEncoding,
			})
			g.lenM = fmt.Sprintf("%s.MarshallerFn[int](%s)", modImportName,
				numG.GenerateFnName(basegen.Marshal))
			g.lenU = fmt.Sprintf("%s.UnmarshallerFn[int](%s)", modImportName,
				numG.GenerateFnName(basegen.Unmarshal))
			g.lenS = fmt.Sprintf("%s.SizerFn[int](%s)", modImportName,
				numG.GenerateFnName(basegen.Size))
		}
		if opts.LenValidator != "" {
			g.lenVl = fmt.Sprintf("com.ValidatorFn[int](%s)", opts.LenValidator)
		}
		g.validator = opts.Validator
	}
	return
}

type StringGenerator struct {
	conf      basegen.Conf
	lenM      string
	lenU      string
	lenS      string
	lenVl     string
	validator string
}

func (g StringGenerator) GenerateFnName(fnType basegen.FnType) (
	name string) {
	var (
		pkg = "ord"
		v   = ""
	)
	if g.conf.Unsafe {
		pkg = "unsafe"
	}
	if fnType == basegen.Unmarshal && g.validFnExpected() {
		v = "Valid"
	}
	return fmt.Sprintf("%s.%s%sString", pkg, string(fnType), v)
}

func (g StringGenerator) GenerateMarshalCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Marshal)
		params = fmt.Sprintf("(%s %s %s)", param(vname), param(g.lenM),
			g.conf.MarshalParam())
	)
	return name + params
}

func (g StringGenerator) GenerateUnmarshalCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Unmarshal)
		params = fmt.Sprintf("(%s %s)", param(g.lenU), g.conf.UnmarshalParam())
	)
	if g.validFnExpected() {
		params = fmt.Sprintf("(%s \n%s \n%s \n%s)", param(g.lenU), param(g.lenVl),
			param("false"), g.conf.UnmarshalParam())
	}

	return name + params
}

func (g StringGenerator) GenerateSizeCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Size)
		params = fmt.Sprintf("(%s %s)", param(vname), g.lenS)
	)
	return name + params
}

func (g StringGenerator) GenerateSkipCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Skip)
		params = fmt.Sprintf("(%s %s)", param(g.lenU), g.conf.SkipParam())
	)
	return name + params
}

func (g StringGenerator) validFnExpected() bool {
	return g.lenVl != ""
}
