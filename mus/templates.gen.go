// Code generated by fvar. DO NOT EDIT.

package musgen

var templates map[string]string

func init() {
	templates = make(map[string]string)
	templates["anonymous_definitions.tmpl"] = `{{/* {Map: map[adesc.AnonymousName]adesc.AnonymousDesc, Ops: genops.Options} */}}
{{- if gt (len .Map)  0 }}
	{{- $gops := .Ops }}
	var (
		{{- $constructorName := "" }}
		{{- range $name, $ad := .Map }}
{{- /* string type */}}
			{{- if eq $ad.Kind "string" }}
				{{- if ne $ad.LenVl "nil" }}
					{{- $constructorName = "NewValidStringSer" }}
				{{- else }}
					{{- $constructorName = "NewStringSer" }}
				{{- end }}
				{{ $ad.Name }} = ord.{{ $constructorName }}({{ MakeStringOps $ad }})
			{{- end }}
{{- /* array type */}}
			{{- if eq $ad.Kind "array" }}
				{{- $elemSer := ElemSerializer $ad $gops }}
				{{- if ne $ad.ElemVl "nil" }}
					{{- $constructorName = "NewValidArraySer" }}
				{{- else }}
					{{- $constructorName = "NewArraySer" }}
				{{- end }}
				{{ $ad.Name }} = ord.{{ $constructorName }}[{{ $ad.Type }}, {{ $ad.ElemType }}]({{ $ad.ArrLength }}, {{ $elemSer }}, {{ MakeArrayOps $ad }})
			{{- end }}
{{- /* byte slice type */}}
			{{- if eq $ad.Kind "byte_slice" }}
				{{- if ne $ad.LenVl "nil" }}
					{{- $constructorName = "NewValidByteSliceSer" }}
				{{- else }}
					{{- $constructorName = "NewByteSliceSer" }}
				{{- end }}
				{{ $ad.Name }} = ord.{{ $constructorName }}({{ MakeByteSliceOps $ad }})
			{{- end }}
{{- /* slice type */}}
			{{- if eq $ad.Kind "slice" }}
				{{- $elemSer := ElemSerializer $ad $gops }}
				{{- if or (ne $ad.LenVl "nil") (ne $ad.ElemVl "nil") }}
					{{- $constructorName = "NewValidSliceSer" }}
				{{- else }}
					{{- $constructorName = "NewSliceSer" }}
				{{- end }}
				{{ $ad.Name }} = ord.{{ $constructorName }}[{{ $ad.ElemType }}]({{ $elemSer }}, {{ MakeSliceOps $ad }})
			{{- end }}
{{- /* map type */}}
			{{- if eq $ad.Kind "map" }}
				{{- $keySer := KeySerializer $ad $gops }}
				{{- $elemSer := ElemSerializer $ad $gops }}
				{{- if or (ne $ad.LenVl "nil") (ne $ad.KeyVl "nil") (ne $ad.ElemVl "nil") }}
					{{- $constructorName = "NewValidMapSer" }}
				{{- else}}
					{{- $constructorName = "NewMapSer" }}
				{{- end }}
				{{ $ad.Name }} = ord.{{ $constructorName }}[{{ $ad.KeyType }}, {{ $ad.ElemType }}]({{ $keySer}}, {{ $elemSer }}, {{ MakeMapOps $ad }})
			{{- end }}
{{- /* ptr type */}}
			{{- if eq $ad.Kind "ptr" }}
				{{- $elemSer := ElemSerializer $ad $gops }}
				{{ $ad.Name }} = ord.NewPtrSer[{{ $ad.ElemType }}]({{ $elemSer }})
			{{- end }}
		{{- end }}
	)
{{- end }}`
	templates["defined_type_ser.tmpl"] = `{{/* tdesc.TypeDesc */}}
{{- $serType := SerType . }}
{{- $serVar := SerVar . }}
{{- $f := index .Fields 0 }}
{{- $fieldSer := SerializerOf $f .Gops }}
{{- $v := VarName .Name }}
{{- $Type := .FullName}}
{{- $tmp := TmpVarName .Name }}
{{- $mslp := .Gops.MarshalSignatureLastParam }}
{{- $mlp := .Gops.MarshalLastParam true }}
{{- $uslp := .Gops.UnmarshalSignatureLastParam }}
{{- $ulp := .Gops.UnmarshalLastParam true}}
{{- $gops := .Gops }}

{{- $ft := $f.Type }}
{{- if PtrType $f.Type }}
  {{- $ft = print "(" $f.Type ")" }}
{{- end }}

var {{ $serVar }} = {{ $serType }}{}

type {{ $serType }} struct{}

func (s {{ $serType }}) Marshal({{ $v }} {{ $Type }}, {{ $mslp }}) (n int {{ if $gops.Stream }} , err error {{ end }}) {
  return {{ $fieldSer }}.Marshal({{ $ft }}({{ $v }}), {{ $mlp }})
}

func (s {{ $serType }}) Unmarshal({{ $uslp }}) ({{ $v }} {{ $Type }}, n int, err error) {
  {{ $tmp }}, n, err := {{ $fieldSer }}.Unmarshal({{ $ulp }})
    if err != nil {
      return
    }
  {{ $v }} = {{ $Type }}({{ $tmp }})
  {{- if and $f.Tops (ne $f.Tops.Validator "") }}
    err = {{ $f.Tops.Validator }}({{ $v }})
  {{- end }}
  return
}

func (s {{ $serType }}) Size({{ $v }} {{ $Type }}) (size int) {
  return {{ $fieldSer }}.Size({{ $ft }}({{ $v }}))
}

func (s {{ $serType }}) Skip({{ $uslp }}) (n int, err error) {
  return {{ $fieldSer }}.Skip({{ $ulp }})
}`
	templates["dts.tmpl"] = `{{/* Name string, FullName string */}}
var {{ .Name }}DTS = dts.New[{{ .FullName }}]({{ .Name }}DTM, {{ .Name }}MUS)`
	templates["field_marshal.tmpl"] = `{{- /* {VarName string, FieldsCount int, Field tdesc.FieldDesc, Index int, Gops genops.Options} */}}
{{- $mlp := .Gops.MarshalLastParam false }}
{{- if eq .Index 0 }}
	{{- $mlp = .Gops.MarshalLastParam true }}
{{- end}}

{{- $vf := print .VarName "." .Field.Name }}
{{- $fieldSer := SerializerOf .Field .Gops }}

{{- $nVar := "n" }}
{{- if gt .Index 0 }}
	{{- $nVar = "n1" }}
{{- end }}

{{- $call := print $fieldSer ".Marshal(" $vf ", " $mlp ")" }}
{{- if eq .FieldsCount 1 }}
	return {{ $call }}
{{- else }}
	{{- if .Gops.Stream }}
		{{- if eq .Index 1 }}
			var n1 int
		{{- end }}
		{{ $nVar }}, err = {{ $call }}
		{{- if ge .Index 1 }}
			n += n1
		{{- end }}
		{{- if eq .Index (minus .FieldsCount 1) }}
			return
		{{- else }}
			if err != nil {
				return
			}		
		{{- end }}
	{{- else }}
		{{- if eq .Index 0 }}
			n = {{ $call }}
		{{- else if eq .Index (minus .FieldsCount 1) }}
			return n + {{ $call }}
		{{- else }}
			n += {{ $call }}
		{{- end }}
	{{- end }}
{{- end }}`
	templates["field_size.tmpl"] = `{{- /* {VarName string, FieldsCount int, Field tdesc.FieldDesc, Index int, Gops genops.Options} */}}
{{- $vf := print .VarName "." .Field.Name }}
{{- $fieldSer := SerializerOf .Field .Gops }}

{{- $call := print $fieldSer ".Size(" $vf ")" }}
{{- if eq .FieldsCount 1 }}
	return {{ $call }}
{{- else }}
	{{- if eq .Index 0 }}
		size = {{ $call }}
	{{- else if eq .Index (minus .FieldsCount 1) }}
		return size + {{ $call }}
	{{- else }}
		size += {{ $call }}
	{{- end }}
{{- end }}`
	templates["field_skip.tmpl"] = `{{- /* {VarName string, FieldsCount int, Field tdesc.FieldDesc, Index int, Gops genops.Options} */}}
{{- $ulp := .Gops.UnmarshalLastParam false }}
{{- if eq .Index 0 }}
	{{- $ulp = .Gops.UnmarshalLastParam true }}
{{- end}}

{{- $vf := print .VarName "." .Field.Name }}
{{- $fieldSer := SerializerOf .Field .Gops }}
{{- if eq .Index 1 }}
	var n1 int
{{- end }}
{{- $nVar := "n" }}
{{- if gt .Index 0 }}
	{{- $nVar = "n1" }}
{{- end }}
{{ $nVar }}, err = {{ $fieldSer }}.Skip({{ $ulp }})
{{- if ge .Index 1 }}
	n += n1
{{- end }}		
{{- if and (ne .FieldsCount 1) (ne .Index (minus .FieldsCount 1))}}
	if err != nil {
		return
	}
{{- end }}`
	templates["field_unmarshal.tmpl"] = `{{- /* {VarName string, FieldsCount int, Field tdesc.FieldDesc, Index int, Gops genops.Options} */}}
{{- $ulp := .Gops.UnmarshalLastParam false }}
{{- if eq .Index 0 }}
	{{- $ulp = .Gops.UnmarshalLastParam true }}
{{- end}}

{{- $vf := print .VarName "." .Field.Name }}
{{- $fieldSer := SerializerOf .Field .Gops }}
{{- if eq .Index 1 }}
	var n1 int
{{- end }}
{{- $nVar := "n" }}
{{- if gt .Index 0 }}
	{{- $nVar = "n1" }}
{{- end }}
{{ $vf }}, {{ $nVar }}, err = {{ $fieldSer }}.Unmarshal({{ $ulp }})
{{- if ge .Index 1 }}
	n += n1
{{- end }}		
{{- if and .Field.Tops (ne .Field.Tops.Validator "") }}
	if err != nil {
		return
	}
	err = {{ .Field.Tops.Validator }}({{ $vf }})
{{- end }}
{{- if and (ne .FieldsCount 1) (ne .Index (minus .FieldsCount 1))}}
	if err != nil {
		return
	}
{{- end }}`
	templates["imports.tmpl"] = `{{/* genops.Options */}}
import (
	com "github.com/mus-format/common-go"
	{{- if .Stream }}
		muss "github.com/mus-format/mus-stream-go"
		dts "github.com/mus-format/mus-stream-dts-go"
		"github.com/mus-format/mus-stream-go/ord"
		"github.com/mus-format/mus-stream-go/raw"
		"github.com/mus-format/mus-stream-go/unsafe"
		"github.com/mus-format/mus-stream-go/varint"
		arrops "github.com/mus-format/mus-stream-go/options/array"
		bslops "github.com/mus-format/mus-stream-go/options/byte_slice"
		mapops "github.com/mus-format/mus-stream-go/options/map"
		slops "github.com/mus-format/mus-stream-go/options/slice"
		strops "github.com/mus-format/mus-stream-go/options/string"
	{{- else }}
		arrops "github.com/mus-format/mus-go/options/array"
		bslops "github.com/mus-format/mus-go/options/byte_slice"
		mapops "github.com/mus-format/mus-go/options/map"
		slops "github.com/mus-format/mus-go/options/slice"
		strops "github.com/mus-format/mus-go/options/string"
	{{- end }}
	{{- if .Imports }}
		{{- range $i, $imp := .Imports }}
			"{{ $imp }}"
		{{- end }}
	{{- end }}
)
`
	templates["interface_marshal.tmpl"] = `{{/* tdesc.TypeDesc */}}
{{- $mlp := .Gops.MarshalLastParam true }}
{{- $v := VarName .Name }}
switch t := {{ $v }}.(type) {
	{{- range $i, $oneOf := .Oneof }}
		case {{ $oneOf }}:
			return {{ $oneOf }}DTS.Marshal(t, {{ $mlp }})
	{{- end }}
		default:
			panic(fmt.Sprintf("unexpected %v type", t))
}`
	templates["interface_ser.tmpl"] = `{{/* tdesc.TypeDesc */}}
{{- $serVar := SerVar . }}
{{- $serType := SerType . }}
{{- $v := VarName .Name }}
{{- $Type := .FullName}}
{{- $mslp := .Gops.MarshalSignatureLastParam }}
{{- $uslp := .Gops.UnmarshalSignatureLastParam }}
{{- $gops := .Gops }}

var {{ $serVar }} = {{ $serType }}{}

type {{ $serType }} struct{}

func (s {{ $serType }}) Marshal({{ $v }} {{ $Type }}, {{ $mslp }}) (n int {{ if $gops.Stream }} , err error {{ end }}) {
	{{- include "interface_marshal.tmpl" . }}
}

func (s {{ $serType }}) Unmarshal({{ $uslp }}) ({{ $v }} {{ $Type }}, n int, err error) {
	{{- include "interface_unmarshal.tmpl" . }}
}

func (s {{ $serType }}) Size({{ $v }} {{ $Type }}) (size int) {
	{{- include "interface_size.tmpl" . }}
}

func (s {{ $serType }}) Skip({{ $uslp }}) (n int, err error) {
	{{- include "interface_skip.tmpl" . }}
}`
	templates["interface_size.tmpl"] = `{{/* tdesc.TypeDesc */}}
{{- $v := VarName .Name }}
switch t := {{ $v }}.(type) {
{{- range $index, $oneOf := .Oneof }}
	case {{ $oneOf }}:
		return {{ $oneOf }}DTS.Size(t)
{{- end }}
	default:
		panic(fmt.Sprintf("unexpected %v type", t))
}`
	templates["interface_skip.tmpl"] = `{{/* tdesc.TypeDesc */}}
{{- $fulp := .Gops.UnmarshalLastParam true }}
{{- $ulp := .Gops.UnmarshalLastParam false }}
{{- $v := VarName .Name }}
dtm, n, err := dts.DTMSer.Unmarshal({{ $fulp }})
if err != nil {
	return
}
var n1 int
switch dtm {
{{- range $index, $oneOf := .Oneof }}
	case {{ $oneOf }}DTM:
		n1, err = {{ $oneOf }}DTS.SkipData({{ $ulp }})
{{- end }}
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
}
n += n1
return`
	templates["interface_unmarshal.tmpl"] = `{{/* tdesc.TypeDesc */}}
{{- $fulp := .Gops.UnmarshalLastParam true }}
{{- $ulp := .Gops.UnmarshalLastParam false }}
{{- $v := VarName .Name }}
dtm, n, err := dts.DTMSer.Unmarshal({{ $fulp }})
if err != nil {
	return
}
var n1 int
switch dtm {
{{- range $index, $oneOf := .Oneof }}
	case {{ $oneOf }}DTM:
		{{ $v }}, n1, err = {{ $oneOf }}DTS.UnmarshalData({{ $ulp }})
{{- end }}
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
}
n += n1
return`
	templates["package.tmpl"] = `{{/* genops.Options */}}
// Code generated by musgen-go. DO NOT EDIT.

package {{ .Package }}
`
	templates["struct_ser.tmpl"] = `{{/* tdesc.TypeDesc */}}
{{- $serVar := SerVar . }}
{{- $serType := SerType . }}

{{- $v := VarName .Name }}
{{- $Type := .FullName}}

{{- $mslp := .Gops.MarshalSignatureLastParam }}
{{- $uslp := .Gops.UnmarshalSignatureLastParam }}

{{- $gops := .Gops }}
{{- $td := . }}

var {{ $serVar }} = {{ $serType }}{}

type {{ $serType }} struct{}

func (s {{ $serType }}) Marshal({{ $v }} {{ $Type }}, {{ $mslp }}) (n int {{ if $gops.Stream }} , err error {{ end }}) {
	{{- if eq (len .Fields) 0}}
		return
	{{- else }}
		{{- range $i, $f := Fields . }}
			{{- include "field_marshal.tmpl" (MakeFieldTmplPipe $td $f $i $gops) }}		
		{{- end }}
	{{- end }}
}

func (s {{ $serType }}) Unmarshal({{ $uslp }}) ({{ $v }} {{ $Type }}, n int, err error) {
	{{- range $i, $f := Fields . }}
			{{- include "field_unmarshal.tmpl" (MakeFieldTmplPipe $td $f $i $gops) }}		
	{{- end }}
	return
}

func (s {{ $serType }}) Size({{ $v }} {{ $Type }}) (size int) {
	{{- if eq (len .Fields) 0}}
		return
	{{- else }}
		{{- range $i, $f := Fields . }}
			{{- include "field_size.tmpl" (MakeFieldTmplPipe $td $f $i $gops) }}		
		{{- end }}
	{{- end }}
}

func (s {{ $serType }}) Skip({{ $uslp }}) (n int, err error) {
	{{- range $i, $f := Fields . }}
		{{- include "field_skip.tmpl" (MakeFieldTmplPipe $td $f $i $gops) }}		
	{{- end }}
	return
}`
}
