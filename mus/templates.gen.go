// Code generated by fvar. DO NOT EDIT.

package mus

var templates map[string]string

func init() {
	templates = make(map[string]string)
	templates["file_header.tmpl"] = `{{/* Conf */}}
// Code generated by musgen-go. DO NOT EDIT.

package {{ .Package }}

import (
	com "github.com/mus-format/common-go"
	{{- if .Stream }}
		muss "github.com/mus-format/mus-stream-go"
		dts "github.com/mus-format/mus-stream-dts-go"
		"github.com/mus-format/mus-stream-go/ord"
		"github.com/mus-format/mus-stream-go/raw"
		"github.com/mus-format/mus-stream-go/unsafe"
		"github.com/mus-format/mus-stream-go/varint"
	{{- end }}
	{{- if .Imports }}
		{{- range $i, $imp := .Imports }}
			"{{ $imp }}"
		{{- end }}
	{{- end }}
)
`
}
