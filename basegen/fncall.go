package basegen

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	Marshal   FnType = "Marshal"
	Unmarshal FnType = "Unmarshal"
	Size      FnType = "Size"
	Skip      FnType = "Skip"
)

type FnType string

type GenerateFnCall = func(conf Conf, vname string, fnType FnType, prefix,
	tp string, meta *Metadata) (call string)

type GenerateSubFn = func(conf Conf, fnType FnType, prefix, tp string,
	meta *Metadata) (m string)

func BuildGenerateFnCall(factory GeneratorFactory) GenerateFnCall {
	return func(conf Conf, vname string, fnType FnType, tp, prefix string,
		meta *Metadata) (call string) {
		var g Generator
		switch tp {
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
			g = factory.NewNumGenerator(conf, tp, meta)
		case "string":
			g = factory.NewStringGenerator(conf, meta)
		case "bool":
			g = factory.NewBoolGenerator(conf, tp, meta)
		}
		if _, ok := ParseSliceType(tp); ok {
			g = factory.NewSliceGenerator(conf, tp, prefix, meta)
		} else if _, _, ok := ParseMapType(tp); ok {
			g = factory.NewMapGenerator(conf, tp, prefix, meta)
		} else if _, _, ok := ParsePtrType(tp); ok {
			g = factory.NewPtrGenerator(conf, tp, prefix, meta)
		} else if _, _, ok := ParseArrayType(tp); ok {
			g = factory.NewSliceGenerator(conf, tp, prefix, meta)
		}
		if g == nil {
			g = factory.NewSAIGenerator(conf, tp, prefix, meta)
		}
		switch fnType {
		case Marshal:
			return g.GenerateMarshalCall(vname)
		case Unmarshal:
			return g.GenerateUnmarshalCall()
		case Size:
			return g.GenerateSizeCall(vname)
		case Skip:
			return g.GenerateSkipCall()
		default:
			panic(NewUnexpectedFnType(fnType))
		}
	}
}

func BuildGenerateSubFn(factory GeneratorFactory) GenerateSubFn {
	return func(conf Conf, fnType FnType, tp, prefix string, meta *Metadata) (
		m string) {
		var (
			vname     = "t"
			g         Generator
			arrayType bool
		)
		if tp == "string" {
			g = factory.NewStringGenerator(conf, meta)
		} else if _, ok := ParseSliceType(tp); ok {
			g = factory.NewSliceGenerator(conf, tp, prefix, meta)
		} else if _, _, ok := ParseMapType(tp); ok {
			g = factory.NewMapGenerator(conf, tp, prefix, meta)
		} else if _, _, ok := ParsePtrType(tp); ok {
			g = factory.NewPtrGenerator(conf, tp, prefix, meta)
		} else if _, _, ok := ParseArrayType(tp); ok {
			arrayType = true
			g = factory.NewSliceGenerator(conf, tp, prefix, meta)
		}
		if g != nil {
			switch fnType {
			case Marshal:
				vname1 := vname
				if arrayType {
					vname1 += "[:]"
				}
				return fmt.Sprintf("func(%s %s, %s) %s { return %s}", vname, tp,
					conf.MarshalParamSignature(),
					conf.MarshalReturnValues(),
					g.GenerateMarshalCall(vname1))
			case Unmarshal:
				fnBody := fmt.Sprintf("return %s", g.GenerateUnmarshalCall())
				if arrayType {
					tmp := vname + "a"
					fnBody = fmt.Sprintf("%s, n, err := %s\n"+
						"if err != nil { return }\n"+
						"%s = (%s)(%s)\n"+
						"return",
						tmp, g.GenerateUnmarshalCall(), vname, tp, tmp)
				}
				return fmt.Sprintf("func(%s) (%s %s, n int, err error) { %s }",
					conf.UnmarshalParamSignature(),
					vname,
					tp,
					fnBody)
			case Size:
				vname1 := vname
				if arrayType {
					vname1 += "[:]"
				}
				return fmt.Sprintf("func(%s %s) (size int) { return %s }", vname, tp,
					g.GenerateSizeCall(vname1))
			case Skip:
				return fmt.Sprintf("func(%s) (n int, err error) { return %s}",
					conf.SkipParamSignature(),
					g.GenerateSkipCall())
			default:
				panic(fmt.Errorf("unexpected %v function type", fnType))
			}
		}

		switch tp {
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
			g = factory.NewNumGenerator(conf, tp, meta)
		case "bool":
			g = factory.NewBoolGenerator(conf, tp, meta)
		default:
			g = factory.NewSAIGenerator(conf, tp, prefix, meta)
		}
		return g.GenerateFnName(fnType)
	}
}

func ParseSliceType(tp string) (elemType string, ok bool) {
	re := regexp.MustCompile(`^\[\](.+$)`)
	match := re.FindStringSubmatch(tp)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}

func ParseMapType(tp string) (keyType, elemType string, ok bool) {
	re := regexp.MustCompile(`^map\[(.+?)\](.+$)`)
	match := re.FindStringSubmatch(tp)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return
}

// ParsePtrType is a template function. If the given string represents a
// pointer type, than Valid == true. The required format is *Type.
//
// If Valid == false, Type is equal to the given type.
func ParsePtrType(t string) (stars, elemType string, ok bool) {
	re := regexp.MustCompile(`(^\*+)(.+$)`)
	match := re.FindStringSubmatch(t)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return
}

// ParseArrayType is a template function. If the given string represents an
// array type, than Valid == true. The required format is [Length]Type.
func ParseArrayType(t string) (elemType string, length int, ok bool) {
	re := regexp.MustCompile(`^\[(\d+)\](.+$)`)
	match := re.FindStringSubmatch(t)
	if len(match) == 3 {
		length, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		return match[2], length, true
	}
	return
}
