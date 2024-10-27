package basegen

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"
)

const (
	EmptyPrefix = "-"
)

func Receiver(s string) string {
	return "v"
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

func CurrentTypeOf(s string) string {
	r := regexp.MustCompile(`(.+)(V\d+$)`)
	match := r.FindStringSubmatch(s)
	if len(match) != 3 {
		panic(fmt.Sprintf("unexpected %v type name", s))
	}
	return match[1]
}

func FieldsLen(td TypeDesc) (l int) {
	for i := 0; i < len(td.Fields); i++ {
		f := td.Fields[i]
		if f.Metadata == nil || !f.Metadata.Ignore {
			l += 1
		}
	}
	return
}

func Fields(td TypeDesc) (fs []FieldDesc) {
	fs = make([]FieldDesc, 0, len(td.Fields))
	for i := 0; i < len(td.Fields); i++ {
		f := td.Fields[i]
		if f.Metadata == nil || !f.Metadata.Ignore {
			fs = append(fs, f)
		}
	}
	return
}

func Prefix(prefix string, meta *Metadata) string {
	if meta != nil {
		if meta.Prefix == EmptyPrefix {
			return ""
		}
		if meta.Prefix != "" {
			return meta.Prefix
		}
	}
	if prefix != "" {
		return prefix
	}
	return ""
}
