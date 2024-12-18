package generator

import (
	"fmt"

	"github.com/mus-format/musgen-go/basegen"
)

func NewMapGenerator(conf basegen.Conf, tp, prefix string, opts *basegen.Options) (
	g MapGenerator) {
	keyType, elemType, ok := basegen.ParseMapType(tp)
	if !ok {
		panic("not a map type")
	}
	g.conf = conf
	g.opts = opts
	g.lenM = "nil"
	g.lenU = "nil"
	g.lenS = "nil"
	g.keyType = keyType
	g.elemType = elemType
	var (
		modImportName                  = conf.ModImportName()
		keyOpts       *basegen.Options = nil
		elemOpts      *basegen.Options = nil
	)
	if opts != nil {
		if opts.LenEncoding != 0 {
			numG := NewNumGenerator(conf, "int", &basegen.Options{
				Encoding: opts.LenEncoding})
			g.lenM = fmt.Sprintf("%s.MarshallerFn[int](%s)", modImportName,
				numG.GenerateFnName(basegen.Marshal))
			g.lenU = fmt.Sprintf("%s.UnmarshallerFn[int](%s)", modImportName,
				numG.GenerateFnName(basegen.Unmarshal))
			g.lenS = fmt.Sprintf("%s.SizerFn[int](%s)", modImportName,
				numG.GenerateFnName(basegen.Size))
		}
		keyOpts = opts.Key
		elemOpts = opts.Elem
		g.validator = opts.Validator

	}
	keyPrefix := basegen.Prefix(prefix, keyOpts)
	elemPrefix := basegen.Prefix(prefix, elemOpts)
	g.m1 = fmt.Sprintf("%s.MarshallerFn[%s](%s)", modImportName, keyType,
		GenerateSubFn(conf, basegen.Marshal, keyType, keyPrefix, keyOpts))
	g.u1 = fmt.Sprintf("%s.UnmarshallerFn[%s](%s)", modImportName, keyType,
		GenerateSubFn(conf, basegen.Unmarshal, keyType, keyPrefix, keyOpts))
	g.s1 = fmt.Sprintf("%s.SizerFn[%s](%s)", modImportName, keyType,
		GenerateSubFn(conf, basegen.Size, keyType, keyPrefix, keyOpts))
	g.sk1 = fmt.Sprintf("%s.SkipperFn(%s)", modImportName,
		GenerateSubFn(conf, basegen.Skip, keyType, keyPrefix, keyOpts))

	g.m2 = fmt.Sprintf("%s.MarshallerFn[%s](%s)", modImportName, elemType,
		GenerateSubFn(conf, basegen.Marshal, elemType, elemPrefix, elemOpts))
	g.u2 = fmt.Sprintf("%s.UnmarshallerFn[%s](%s)", modImportName, elemType,
		GenerateSubFn(conf, basegen.Unmarshal, elemType, elemPrefix, elemOpts))
	g.s2 = fmt.Sprintf("%s.SizerFn[%s](%s)", modImportName, elemType,
		GenerateSubFn(conf, basegen.Size, elemType, elemPrefix, elemOpts))
	g.sk2 = fmt.Sprintf("%s.SkipperFn(%s)", modImportName,
		GenerateSubFn(conf, basegen.Skip, elemType, elemPrefix, elemOpts))
	return g
}

type MapGenerator struct {
	conf      basegen.Conf
	opts      *basegen.Options
	lenM      string
	lenU      string
	lenS      string
	m1        string
	u1        string
	s1        string
	sk1       string
	m2        string
	u2        string
	s2        string
	sk2       string
	validator string
	keyType   string
	elemType  string
}

func (g MapGenerator) GenerateFnName(fnType basegen.FnType) (name string) {
	valid := ""
	if _, _, _, ok := g.validFnExpected(); fnType == basegen.Unmarshal && ok {
		valid = "Valid"
	}
	if fnType == basegen.Skip {
		return fmt.Sprintf("ord.%sMap", fnType)
	}
	return fmt.Sprintf("ord.%s%sMap[%s, %s]", fnType, valid, g.keyType, g.elemType)
}

func (g MapGenerator) GenerateMarshalCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Marshal)
		params = fmt.Sprintf("(%s %s \n%s \n%s \n%s)", param(vname), param(g.lenM),
			param(g.m1), param(g.m2), g.conf.MarshalParam())
	)
	return name + params
}

func (g MapGenerator) GenerateUnmarshalCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Unmarshal)
		params = fmt.Sprintf("(%s \n%s \n%s \n%s)", param(g.lenU), param(g.u1),
			param(g.u2), g.conf.UnmarshalParam())
	)
	if lenVl, keyVl, elemVl, ok := g.validFnExpected(); ok {
		params = fmt.Sprintf("(%s %s \n%s \n%s \n%s \n%s \n%s \n%s \n%s)",
			param(g.lenU),
			param(lenVl),
			param(g.u1),
			param(g.u2),
			param(keyVl),
			param(elemVl),
			param("nil"),
			param("nil"),
			g.conf.UnmarshalParam())
	}
	return name + params
}

func (g MapGenerator) GenerateSizeCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Size)
		params = fmt.Sprintf("(%s %s \n%s \n%s)", param(vname), param(g.lenS),
			param(g.s1), param(g.s2))
	)
	return name + params
}

func (g MapGenerator) GenerateSkipCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Skip)
		params = fmt.Sprintf("(%s \n%s \n%s \n%s)", param(g.lenU), param(g.sk1),
			param(g.sk2), g.conf.SkipParam())
	)
	return name + params
}

func (g MapGenerator) GenerateValidation() (validation string) {
	return fmt.Sprintf("if err = %s; err != nil { return }", g.validator)
}

func (g MapGenerator) validFnExpected() (lenVl, keyVl, elemVl string, ok bool) {
	if g.opts != nil {
		if g.opts.LenValidator != "" {
			lenVl = fmt.Sprintf("com.ValidatorFn[int](%s)", g.opts.LenValidator)
			ok = true
		}
		if g.opts.Key != nil && g.opts.Key.Validator != "" {
			keyVl = fmt.Sprintf("com.ValidatorFn[%s](%s)", g.keyType,
				g.opts.Key.Validator)
			ok = true
		}
		if g.opts.Elem != nil && g.opts.Elem.Validator != "" {
			elemVl = fmt.Sprintf("com.ValidatorFn[%s](%s)", g.elemType,
				g.opts.Elem.Validator)
			ok = true
		}
	}
	if ok {
		if lenVl == "" {
			lenVl = "nil"
		}
		if keyVl == "" {
			keyVl = "nil"
		}
		if elemVl == "" {
			elemVl = "nil"
		}
	}
	return
}
