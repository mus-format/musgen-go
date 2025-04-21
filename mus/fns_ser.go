package musgen

import (
	"fmt"
	"unicode"

	"github.com/mus-format/musgen-go/data"
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
)

func NewSerFns(keywordFns KeywordFns, utilFns UtilFns,
	anonBuilder AnonSerDataBuilder,
	typeBuilder TypeDataBuilder,
) SerFns {
	return SerFns{keywordFns, utilFns, anonBuilder, typeBuilder}
}

type SerFns struct {
	keywordFns  KeywordFns
	utilFns     UtilFns
	anonBuilder AnonSerDataBuilder
	typeBuilder TypeDataBuilder
}

func (f SerFns) SerOf(name typename.FullName, tops *typeops.Options,
	gops genops.Options) string {
	anonData, ok, err := f.anonBuilder.Build(name, tops)
	if err != nil {
		panic(fmt.Sprintf("can't get serializer for the %v, cause %v", name, err))
	}
	if ok {
		return string(anonData.AnonSerName)
	}
	var (
		pkg     string
		serName string
	)
	switch name {
	case "int", "int64", "int32", "int16", "int8":
		pkg = f.numPkg(tops)
		serName = f.capitalize(string(name))
		if tops != nil && tops.NumEncoding == typeops.VarintPositive {
			serName = "Positive" + serName
		}
	case "uint", "uint64", "uint32", "uint16", "uint8", "float64", "float32", "byte":
		pkg = f.numPkg(tops)
		serName = f.capitalize(string(name))
	case "bool", "string":
		pkg = "ord"
		serName = f.capitalize(string(name))
	case "[]byte", "[]uint8":
		pkg = "ord"
		serName = "ByteSlice"
	case "time.Time":
		pkg = "raw"
		serName = f.utilFns.TimeSer(tops)
	default:
		return f.keywordFns.serVar(name)
	}
	if gops.NotUnsafe && !f.utilFns.ByteSliceStream(name, gops) && name != "string" {
		pkg = "unsafe"
	}
	if gops.Unsafe && !f.utilFns.ByteSliceStream(name, gops) {
		pkg = "unsafe"
	}
	return pkg + "." + serName
}

func (f SerFns) AnonKeySer(anonData data.AnonData, gops genops.Options) string {
	var keyTops *typeops.Options
	if anonData.Tops != nil {
		keyTops = anonData.Tops.Key
	}
	return f.SerOf(anonData.KeyType, keyTops, gops)
}

func (f SerFns) AnonElemSer(anonData data.AnonData, gops genops.Options) string {
	var elemTops *typeops.Options
	if anonData.Tops != nil {
		elemTops = anonData.Tops.Elem
	}
	return f.SerOf(anonData.ElemType, elemTops, gops)
}

func (f SerFns) RelName(name typename.FullName, gops genops.Options) string {
	return string(f.typeBuilder.ToRelName(name))
}

func (f SerFns) capitalize(str string) string {
	r := []rune(str)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func (f SerFns) numPkg(tops *typeops.Options) string {
	if tops == nil {
		return "varint"
	}
	switch tops.NumEncoding {
	case typeops.UndefinedNumEncoding, typeops.Varint, typeops.VarintPositive:
		return "varint"
	case typeops.Raw:
		return "raw"
	default:
		panic(fmt.Sprintf("unexpected NumEnccoding %v", tops.NumEncoding))
	}
}

func (f SerFns) RegisterItself(m map[string]any) {
	m["SerOf"] = func(fname typename.FullName, tops *typeops.Options,
		gops genops.Options) string {
		return f.SerOf(fname, tops, gops)
	}
	m["AnonKeySer"] = func(anonData data.AnonData, gops genops.Options) string {
		return f.AnonKeySer(anonData, gops)
	}
	m["AnonElemSer"] = func(anonData data.AnonData, gops genops.Options) string {
		return f.AnonElemSer(anonData, gops)
	}
	m["RelName"] = func(fname typename.FullName, gops genops.Options) string {
		return f.RelName(fname, gops)
	}
}
