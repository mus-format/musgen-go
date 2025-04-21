package musgen

import (
	"bytes"
	"regexp"
	"text/template"

	"github.com/mus-format/musgen-go/data"
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
)

const (
	PtrPattern   = `(^\*)(.+$)`
	ArrayPattern = `[\d+](.+$)`
)

var (
	ptrRe   = regexp.MustCompile(PtrPattern)
	arrayRe = regexp.MustCompile(ArrayPattern)
)

func NewUtilFns(keywordFns KeywordFns) UtilFns {
	return UtilFns{keywordFns}
}

type UtilFns struct {
	keywordFns KeywordFns
}

func (f UtilFns) PtrType(name typename.FullName) (ok bool) {
	return ptrRe.MatchString(string(name))
}

func (f UtilFns) ArrayType(name typename.FullName) (ok bool) {
	return arrayRe.MatchString(string(name))
}

func (f UtilFns) WithComma(str string) string {
	if str == "" {
		return str
	}
	return ", " + str
}

func (f UtilFns) MinusFunc() func(int, int) int {
	return func(a int, b int) int { return a - b }
}

func (f UtilFns) ByteSliceStream(name typename.FullName,
	gops genops.Options) bool {
	return gops.Stream && (name == "[]byte" || name == "[]uint8")
}

func (f UtilFns) TimeSer(tops *typeops.Options) string {
	if tops == nil {
		return "TimeUnix"
	}
	return tops.TimeUnit.Ser()
}

func (f UtilFns) IncludeFunc(tmpl *template.Template) func(name string, pipeline any) (str string, err error) {
	return func(name string, pipeline any) (str string, err error) {
		var buf bytes.Buffer
		if err = tmpl.ExecuteTemplate(&buf, name, pipeline); err != nil {
			return
		}
		str = buf.String()
		return
	}
}

func (f UtilFns) FieldTmplPipe(typeData data.TypeData,
	fieldData data.FieldData, index int, gops genops.Options) FieldTmplPipe {
	result := struct {
		SerReceiver string
		FieldsCount int
		Field       data.FieldData
		Index       int
		Tops        *typeops.Options
		Gops        genops.Options
	}{
		SerReceiver: f.keywordFns.SerReceiver(typeData),
		FieldsCount: len(typeData.SerializedFields()),
		Field:       fieldData,
		Index:       index,
		Gops:        gops,
	}
	if len(typeData.Sops.Fields) > 0 {
		result.Tops = typeData.Sops.Fields[index]
	}
	return result
}

func (f UtilFns) RegisterItself(m map[string]any, tmpl *template.Template) {
	m["PtrType"] = func(name typename.FullName) bool {
		return f.PtrType(name)
	}
	m["ArrayType"] = func(name typename.FullName) bool {
		return f.ArrayType(name)
	}
	m["WithComma"] = func(str string) string {
		return f.WithComma(str)
	}
	m["minus"] = func(n1 int, n2 int) int {
		return f.MinusFunc()(n1, n2)
	}
	m["ByteSliceStream"] = func(name typename.FullName, gops genops.Options) bool {
		return f.ByteSliceStream(name, gops)
	}
	m["TimeSer"] = func(tops *typeops.Options) string {
		return f.TimeSer(tops)
	}
	m["include"] = func(name string, pipeline any) (str string, err error) {
		return f.IncludeFunc(tmpl)(name, pipeline)
	}
	m["FieldTmplPipe"] = func(typeData data.TypeData,
		fieldData data.FieldData, index int, gops genops.Options) FieldTmplPipe {
		return f.FieldTmplPipe(typeData, fieldData, index, gops)
	}
}

type FieldTmplPipe struct {
	SerReceiver string
	FieldsCount int
	Field       data.FieldData
	Index       int
	Tops        *typeops.Options
	Gops        genops.Options
}
