package musgen

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"
	"unicode"

	"github.com/mus-format/musgen-go/adesc"
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/parser"
	"github.com/mus-format/musgen-go/tdesc"
)

const SuffixMUS = "MUS"

func SerializerOf(fd tdesc.FieldDesc, gops genops.Options) string {
	ad, ok := adesc.Make(fd.Type, gops, fd.Tops)
	if ok {
		return string(ad.Name)
	}
	return serializerOf(fd.Type, fd.Tops, gops)
}

func KeySerializer(ad adesc.AnonymousDesc, gops genops.Options) string {
	var keyTops *typeops.Options
	if ad.Tops != nil {
		keyTops = ad.Tops.Key
	}
	keyAd, ok := adesc.Make(ad.KeyType, gops, keyTops)
	if ok {
		return string(keyAd.Name)
	}
	return serializerOf(ad.KeyType, keyTops, gops)
}

func ElemSerializer(ad adesc.AnonymousDesc, gops genops.Options) string {
	var elemTops *typeops.Options
	if ad.Tops != nil {
		elemTops = ad.Tops.Elem
	}
	elemAd, ok := adesc.Make(ad.ElemType, gops, elemTops)
	if ok {
		return string(elemAd.Name)
	}
	return serializerOf(ad.ElemType, elemTops, gops)
}

func PtrType(t string) (ok bool) {
	_, _, ok = parser.TypeName.ParsePtr(t)
	return
}

func SerType(td tdesc.TypeDesc) string {
	return MakeFirstLetterSmall(td.Name + SuffixMUS)
}

func SerVar(td tdesc.TypeDesc) string {
	return td.Name + SuffixMUS
}

func VarName(s string) string {
	return "v"
}

func TmpVarName(s string) string {
	return "tmp"
}

// MakeIncludeFunc creates template's include func.
func MakeIncludeFunc(tmpl *template.Template) func(string, interface{}) (string,
	error) {
	return func(name string, pipeline interface{}) (string, error) {
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, name, pipeline); err != nil {
			return "", err
		}
		return buf.String(), nil
	}
}

// MakeMinusFunc creates template's minus func.
func MakeMinusFunc() func(int, int) int {
	return func(a int, b int) int {
		return a - b
	}
}

// MakeAddFunc creates template's minus func.
func MakeAddFunc() func(int, int) int {
	return func(a, b int) int {
		return a + b
	}
}

func CurrentTypeOf(s string) string {
	r := regexp.MustCompile(`(.+)(V\d+$)`)
	match := r.FindStringSubmatch(s)
	if len(match) != 3 {
		panic(fmt.Sprintf("unexpected %v type name", s))
	}
	return match[1]
}

func FieldsLen(td tdesc.TypeDesc) (l int) {
	for i := 0; i < len(td.Fields); i++ {
		f := td.Fields[i]
		if f.Tops == nil || !f.Tops.Ignore {
			l += 1
		}
	}
	return
}

func Fields(td tdesc.TypeDesc) (fs []tdesc.FieldDesc) {
	fs = make([]tdesc.FieldDesc, 0, len(td.Fields))
	for i := 0; i < len(td.Fields); i++ {
		f := td.Fields[i]
		if f.Tops == nil || !f.Tops.Ignore {
			fs = append(fs, f)
		}
	}
	return
}

func ArrayType(t string) bool {
	_, _, ok := parser.TypeName.ParseArray(t)
	return ok
}

func MakeFirstLetterBig(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func MakeFirstLetterSmall(s string) string {
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func MakeFieldTmplPipe(td tdesc.TypeDesc, field tdesc.FieldDesc, index int,
	gops genops.Options) struct {
	VarName     string
	FieldsCount int
	Field       tdesc.FieldDesc
	Index       int
	Gops        genops.Options
} {
	return struct {
		VarName     string
		FieldsCount int
		Field       tdesc.FieldDesc
		Index       int
		Gops        genops.Options
	}{
		VarName:     VarName(td.Name),
		FieldsCount: FieldsLen(td),
		Field:       field,
		Index:       index,
		Gops:        gops,
	}
}

func numPkg(tops *typeops.Options) (pkg string) {
	if tops == nil {
		pkg = "varint"
	} else {
		switch tops.NumEncoding {
		case typeops.UndefinedNumEncoding, typeops.Varint, typeops.VarintPositive:
			pkg = "varint"
		case typeops.Raw:
			pkg = "raw"
		default:
			panic(fmt.Sprintf("unexpected NumEnccoding %v", tops.NumEncoding))
		}
	}
	return
}

func serializerOf(tp string, tops *typeops.Options, gops genops.Options) string {
	var (
		pkg     string
		serName string
	)
	switch tp {
	case "int", "int64", "int32", "int16", "int8":
		pkg = numPkg(tops)
		serName = MakeFirstLetterBig(tp)
		if tops != nil && tops.NumEncoding == typeops.VarintPositive {
			serName = "Positive" + serName
		}
	case "uint", "uint64", "uint32", "uint16", "uint8", "float64", "float32", "byte":
		pkg = numPkg(tops)
		serName = MakeFirstLetterBig(tp)
	case "bool", "string":
		pkg = "ord"
		serName = MakeFirstLetterBig(tp)
	case "[]byte", "[]uint8":
		pkg = "ord"
		serName = "ByteSlice"
	default:
		return tp + SuffixMUS
	}
	if gops.Unsafe && !streamByteSlice(tp, gops) {
		pkg = "unsafe"
	}
	return pkg + "." + serName
}

func streamByteSlice(tp string, gops genops.Options) bool {
	return gops.Stream && (tp == "[]byte" || tp == "[]uint8")
}
