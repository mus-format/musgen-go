package generator

import (
	"fmt"
	"regexp"

	"github.com/mus-format/musgen-go/basegen"
)

func NewPtrGenerator(conf basegen.Conf, tp, prefix string, opts *basegen.Options) (
	g PtrGenerator) {
	stars, elemType, ok := ParsePtrType(tp)
	if !ok {
		panic("not a pointer type")
	}
	if len(stars) != 1 {
		panic("too many stars")
	}
	g.conf = conf
	g.elemType = elemType
	var (
		modImportName                  = conf.ModImportName()
		elemOpts      *basegen.Options = nil
	)
	if opts != nil {
		if opts.Elem != nil {
			elemOpts = opts.Elem
		}
		g.validator = opts.Validator
	}
	g.m = fmt.Sprintf("%s.MarshallerFn[%s](%s)", modImportName, elemType,
		GenerateSubFn(conf, basegen.Marshal, elemType, prefix, elemOpts))
	g.u = fmt.Sprintf("%s.UnmarshallerFn[%s](%s)", modImportName, elemType,
		GenerateSubFn(conf, basegen.Unmarshal, elemType, prefix, elemOpts))
	g.s = fmt.Sprintf("%s.SizerFn[%s](%s)", modImportName, elemType,
		GenerateSubFn(conf, basegen.Size, elemType, prefix, elemOpts))
	g.sk = fmt.Sprintf("%s.SkipperFn(%s)", modImportName,
		GenerateSubFn(conf, basegen.Skip, elemType, prefix, elemOpts))
	return
}

type PtrGenerator struct {
	conf      basegen.Conf
	m         string
	u         string
	s         string
	sk        string
	ptrType   basegen.PtrType
	validator string
	elemType  string
}

func (g PtrGenerator) GenerateFnName(fnType basegen.FnType) (name string) {
	if fnType == basegen.Skip {
		return fmt.Sprintf("%s.%sPtr", g.ptrType.PackageName(), fnType)
	}
	return fmt.Sprintf("%s.%sPtr[%s]", g.ptrType.PackageName(), fnType,
		g.elemType)
}

func (g PtrGenerator) GenerateMarshalCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Marshal)
		params = fmt.Sprintf("(%s %s %s)", param(vname), param(g.m),
			g.conf.MarshalParam())
		// params = "(" + param(vname) + param(g.m) + "bs[n:])"
	)
	return name + params
}

func (g PtrGenerator) GenerateUnmarshalCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Unmarshal)
		params = fmt.Sprintf("(%s %s)", param(g.u), g.conf.UnmarshalParam())
	)
	return name + params
}

func (g PtrGenerator) GenerateSizeCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Size)
		params = fmt.Sprintf("(%s %s)", param(vname), param(g.s))
	)
	return name + params
}

func (g PtrGenerator) GenerateSkipCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Skip)
		params = fmt.Sprintf("(%s %s)", param(g.sk), g.conf.SkipParam())
	)
	return name + params
}

func (g PtrGenerator) GenerateValidation() (validation string) {
	return fmt.Sprintf("if err = %s; err != nil { return }", g.validator)
}

func ParsePtrType(tp string) (stars, baseType string, ok bool) {
	re := regexp.MustCompile(`(^\*+)(.+$)`)
	match := re.FindStringSubmatch(tp)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return "", "", false
}
