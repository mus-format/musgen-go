package generator

import (
	"fmt"

	"github.com/mus-format/musgen-go/basegen"
)

func NewByteSliceGenerator(conf basegen.Conf, tp, prefix string,
	opts *basegen.Options) (g ByteSliceGenerator) {
	elemType, ok := basegen.ParseSliceType(tp)
	if !ok {
		elemType, _, ok = basegen.ParseArrayType(tp)
		if !ok {
			panic("not a slice or array type")
		}
	}

	g.conf = conf
	g.opts = opts
	g.lenM = "nil"
	g.lenU = "nil"
	g.lenS = "nil"
	g.elemType = elemType
	var (
		modImportName = conf.ModImportName()
		// elemOpts      *basegen.Options = nil
	)
	if opts != nil {
		if opts.LenEncoding != 0 {
			numG := NewNumGenerator(conf, "int", &basegen.Options{
				Encoding: opts.LenEncoding})
			g.lenM = modImportName + ".MarshallerFn[int](" +
				numG.GenerateFnName(basegen.Marshal) + ")"
			g.lenU = modImportName + ".UnmarshallerFn[int](" +
				numG.GenerateFnName(basegen.Unmarshal) + ")"
			g.lenS = modImportName + ".SizerFn[int](" +
				numG.GenerateFnName(basegen.Size) + ")"
		}
		// elemOpts = opts.Elem
		g.validator = opts.Validator
	}
	// elemPrefix := basegen.Prefix(prefix, elemOpts)
	// g.m = fmt.Sprintf("%s.MarshallerFn[%s](%s)", modImportName, elemType,
	// 	GenerateSubFn(conf, basegen.Marshal, elemType, elemPrefix, elemOpts))
	// g.u = fmt.Sprintf("%s.UnmarshallerFn[%s](%s)", modImportName, elemType,
	// 	GenerateSubFn(conf, basegen.Unmarshal, elemType, elemPrefix, elemOpts))
	// g.s = fmt.Sprintf("%s.SizerFn[%s](%s)", modImportName, elemType,
	// 	GenerateSubFn(conf, basegen.Size, elemType, elemPrefix, elemOpts))
	// g.sk = fmt.Sprintf("%s.SkipperFn(%s)", modImportName,
	// 	GenerateSubFn(conf, basegen.Skip, elemType, elemPrefix, elemOpts))
	return g
}

type ByteSliceGenerator struct {
	conf basegen.Conf
	opts *basegen.Options
	lenM string
	lenU string
	lenS string
	// m         string
	// u         string
	// s         string
	// sk        string
	validator string
	elemType  string
}

func (g ByteSliceGenerator) GenerateFnName(fnType basegen.FnType) (name string) {
	var (
		pkg = "ord"
	)
	if g.conf.Unsafe && !g.conf.Stream {
		pkg = "unsafe"
	}
	valid := ""
	if _, ok := g.validFnExpected(); fnType == basegen.Unmarshal && ok {
		valid = "Valid"
	}
	if fnType == basegen.Skip {
		return fmt.Sprintf("%s.%sByteSlice", pkg, fnType)
	}
	return fmt.Sprintf("%s.%s%sByteSlice", pkg, fnType, valid)
}

func (g ByteSliceGenerator) GenerateMarshalCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Marshal)
		params = fmt.Sprintf("(%s \n%s \n%s)", param(vname), param(g.lenM),
			g.conf.MarshalParam())
	)
	return name + params
}

func (g ByteSliceGenerator) GenerateUnmarshalCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Unmarshal)
		params = fmt.Sprintf("(%s \n%s)", param(g.lenU), g.conf.UnmarshalParam())
	)
	if lenVl, ok := g.validFnExpected(); ok {
		params = fmt.Sprintf("(%s \n%s \n%s \n%s)", param(g.lenU),
			param(lenVl),
			param("false"),
			g.conf.UnmarshalParam())
	}
	return name + params
}

func (g ByteSliceGenerator) GenerateSizeCall(vname string) (call string) {
	var (
		name   = g.GenerateFnName(basegen.Size)
		params = fmt.Sprintf("(%s \n%s)", param(vname), param(g.lenS))
	)
	return name + params
}

func (g ByteSliceGenerator) GenerateSkipCall() (call string) {
	var (
		name   = g.GenerateFnName(basegen.Skip)
		params = fmt.Sprintf("(%s \n%s)", param(g.lenU), g.conf.SkipParam())
	)
	return name + params
}

func (g ByteSliceGenerator) GenerateValidation() (validation string) {
	return fmt.Sprintf("if err = %s; err != nil { return }", g.validator)
}

func (g ByteSliceGenerator) validFnExpected() (lenVl string, ok bool) {
	if g.opts != nil {
		if g.opts.LenValidator != "" {
			lenVl = fmt.Sprintf("com.ValidatorFn[int](%s)", g.opts.LenValidator)
			ok = true
		}
		// if g.opts.Elem != nil && g.opts.Elem.Validator != "" {
		// 	elemVl = fmt.Sprintf("com.ValidatorFn[%s](%s)", g.elemType,
		// 		g.opts.Elem.Validator)
		// 	ok = true
		// }
	}
	if ok {
		if lenVl == "" {
			lenVl = "nil"
		}
		// if elemVl == "" {
		// 	elemVl = "nil"
		// }
	}
	return
}
