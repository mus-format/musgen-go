// Code generated by fvar. DO NOT EDIT.

package basegen

var templates map[string]string

func init() {
	templates = make(map[string]string)
	templates["alias_marshal.tmpl"] = `{{- $f := index .Fields 0 }}
{{- $r := Receiver .Name }}
{{- $a := ArrayType $f.Type }}
{{- $v := print .AliasOf "(" $r ")" }}
{{- if $a }}
	{{ $r }}a := {{ $v }}
	{{- $v = print $r "a" "[:]" }}
{{- end }}
return {{ GenerateFnCall .Conf $v "Marshal" $f.Type .Prefix $f.Options }}`
	templates["alias_size.tmpl"] = `{{- $f := index .Fields 0 }}
{{- $r := Receiver .Name }}
{{- $a := ArrayType $f.Type }}
{{- $v := print .AliasOf "(" $r ")" }}
{{- if $a }}
	{{ $r }}a := {{ $v }}
	{{- $v = print $r "a" "[:]" }}
{{- end }}
return {{ GenerateFnCall .Conf $v "Size" $f.Type .Prefix $f.Options }}`
	templates["alias_skip.tmpl"] = `{{- $f := index .Fields 0 }}
return {{ GenerateFnCall .Conf "" "Skip" $f.Type .Prefix $f.Options }}`
	templates["alias_unmarshal.tmpl"] = `{{- $f := index .Fields 0 }}
{{- $r := Receiver .Name }}
{{- $v := print $r "a" }}
{{- $a := ArrayType $f.Type }}
{{ $v }}, n, err := {{ GenerateFnCall .Conf "" "Unmarshal" $f.Type .Prefix $f.Options }}
if err != nil {
	return
}
{{- if and $f.Options (ne $f.Options.Validator "") }}
	if err = {{ $f.Options.Validator }}({{ $v }}); err != nil {
		return
	}
{{- end }}
{{- if $a }}
	{{ $r }} = {{ .Name }}(({{ .AliasOf }})({{ $v }}))
{{- else }}
	{{ $r }} = {{ .Name }}({{ $v }})
{{- end }}
return`
	templates["interface_marshal.tmpl"] = `{{- $c := .Conf}}
{{- $p := .Prefix }}
switch tp := {{ Receiver .Name }}.(type) {
	{{- range $i, $oneOf := .Oneof }}
		case {{ $oneOf }}:
			return {{ $p }}{{ $oneOf }}DTS.Marshal(tp, {{ $c.MarshalParam }})
	{{- end }}
		default:
			panic(fmt.Errorf("unexpected %v type", tp))
}`
	templates["interface_size.tmpl"] = `{{- $p := .Prefix }}
switch tp := {{ Receiver .Name }}.(type) {
{{- range $index, $oneOf := .Oneof }}
	case {{ $oneOf }}:
		return {{ $p }}{{ $oneOf }}DTS.Size(tp)
{{- end }}
	default:
		panic(fmt.Errorf("unexpected %v type", tp))
}`
	templates["interface_skip.tmpl"] = `{{- $c := .Conf }}
{{- $p := .Prefix }}
dtm, n, err := dts.UnmarshalDTM({{ $c.UnmarshalParam }})
if err != nil {
	return
}
var n1 int
switch dtm {
{{- range $index, $oneOf := .Oneof }}
	case {{ $oneOf }}DTM:
		n1, err = {{ $p }}{{ $oneOf }}DTS.SkipData({{ $c.UnmarshalParam }})
{{- end }}
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
}
n += n1
return`
	templates["interface_unmarshal.tmpl"] = `{{- $r := Receiver .Name }}
{{- $c := .Conf }}
{{- $p := .Prefix }}
dtm, n, err := dts.UnmarshalDTM({{ $c.UnmarshalParam }})
if err != nil {
	return
}
var n1 int
switch dtm {
{{- range $index, $oneOf := .Oneof }}
	case {{ $oneOf }}DTM:
		{{ $r }}, n1, err = {{ $p }}{{ $oneOf }}DTS.UnmarshalData({{ $c.UnmarshalParam }})
{{- end }}
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
}
n += n1
return`
	templates["main_alias_struct.tmpl"] = `{{- /* TypeDesc */ -}}
{{ include "serializer.tmpl" . }}
`
	templates["main_dts.tmpl"] = `{{ include "serializer.tmpl" . }}
{{- $n := print .Prefix .Name }}
var {{ $n }}DTS = dts.New[{{ .Name }}]({{ .Name }}DTM, 
	{{ .Conf.ModImportName }}.MarshallerFn[{{ .Name }}](Marshal{{ $n }}MUS),  
	{{ .Conf.ModImportName }}.UnmarshallerFn[{{ .Name }}](Unmarshal{{ $n }}MUS), 
	{{ .Conf.ModImportName }}.SizerFn[{{ .Name }}](Size{{ $n }}MUS),
	{{ .Conf.ModImportName }}.SkipperFn(Skip{{ $n }}MUS))`
	templates["main_dvs.tmpl"] = `var registry = com.NewRegistry(
	[]com.TypeVersion{
	{{- range $index, $tp := . }}
		{{- $ctp := CurrentTypeOf $tp }}
		dvs.Version[{{ $tp }}, {{ $ctp }}]{
			DTS: {{ $tp }}DTS,
			MigrateOld: Migrate{{ $tp }},
			MigrateCurrent: MigrateTo{{ $tp }},
		},
	{{- end }}
	},
)

var (
	{{- range $index, $tp := . }}
		{{- $ctp := CurrentTypeOf $tp}}
		{{ $ctp }}DVS = dvs.New[{{ $ctp }}](registry)
	{{- end }}
)`
	templates["main_interface.tmpl"] = `{{ include "serializer.tmpl" . }}
`
	templates["serializer.tmpl"] = `func Marshal{{ .Prefix }}{{ .Name }}MUS({{ Receiver .Name }} {{ .Name }}, {{ .Conf.MarshalParamSignature }}) (n int {{ if .Conf.Stream }} , err error {{ end }}) {
	{{- if ne .AliasOf "" }}
		{{- include "alias_marshal.tmpl" . }}
	{{- else if gt (len .Oneof) 0 }}
		{{- include "interface_marshal.tmpl" . }}
	{{- else }}
		{{- include "struct_marshal.tmpl" . }}
	{{- end }}
}

func Unmarshal{{ .Prefix }}{{ .Name }}MUS({{ .Conf.UnmarshalParamSignature }}) ({{ Receiver .Name }} {{ .Name }}, n int, err error) {
	{{- if ne .AliasOf "" }}
		{{- include "alias_unmarshal.tmpl" . }}
	{{- else if gt (len .Oneof) 0 }}
		{{- include "interface_unmarshal.tmpl" . }}
	{{- else }}
		{{- include "struct_unmarshal.tmpl" . }}
	{{- end }}
}

func Size{{ .Prefix }}{{ .Name }}MUS({{ Receiver .Name }} {{ .Name }}) (size int) {
	{{- if ne .AliasOf "" }}
		{{- include "alias_size.tmpl" . }}
	{{- else if gt (len .Oneof) 0 }}
		{{- include "interface_size.tmpl" . }}
	{{- else }}
		{{- include "struct_size.tmpl" . }}
	{{- end }}
}

func Skip{{ .Prefix }}{{ .Name }}MUS({{ .Conf.SkipParamSignature }}) (n int, err error) {
	{{- if ne .AliasOf "" }}
		{{- include "alias_skip.tmpl" . }}
	{{- else if gt (len .Oneof) 0 }}
		{{- include "interface_skip.tmpl" . }}
	{{- else }}
		{{- include "struct_skip.tmpl" . }}
	{{- end }}
}
`
	templates["struct_marshal.tmpl"] = `{{- $l := FieldsLen . }}
{{- $r := Receiver .Name }}
{{- $p := .Prefix }}
{{- if eq $l 0 }}
	return
{{- else }}
	{{- $c := .Conf }}
	{{- range $i, $f := Fields . }}
		{{- $v := print $r "." $f.Name }}
		{{- if ArrayType $f.Type }}{{ $v = print $v "[:]"}}{{ end }}
		{{- $fnCall := GenerateFnCall $c $v "Marshal" $f.Type $p $f.Options }}
		{{- if or (not $f.Options) (not $f.Options.Ignore) }}
			{{- if eq $l 1 }}
				return {{ $fnCall }}
			{{- else }}
				{{- if $c.Stream }}

					{{- if eq $i 0 }}
						n, err = {{ $fnCall }}
					{{- end }}
					{{- if eq $i 1 }}
						var n1 int
					{{- end }}
					{{- if gt $i 0 }}
						n1, err = {{ $fnCall }}
						n += n1
					{{- end }}
					{{- if (lt $i (minus $l 1)) }}
						if err != nil {
							return
						}
					{{- end }}
					{{- if eq $i (minus $l 1) }}
						return
					{{- end }}

				{{- else }}

					{{- if eq $i 0 }}
						n = {{ $fnCall }}	
					{{- end }}
					{{- if and (gt $i 0) (lt $i (minus $l 1)) }}
						n += {{ $fnCall }}
					{{- end }}
					{{- if eq $i (minus $l 1) }}
						return n + {{ $fnCall }}
					{{- end }}

				{{- end }}
			{{- end }}
		{{- end }}
	{{- end }}
{{- end }}`
	templates["struct_size.tmpl"] = `{{- $l := FieldsLen . }}
{{- $r := Receiver .Name }}
{{- $p := .Prefix }}
{{- if eq $l 0 }}
	return
{{- else }}
	{{- $c := .Conf }}
	{{- range $i, $f := Fields . }}
		{{- $v := print $r "." $f.Name }}
		{{- if ArrayType $f.Type }}{{ $v = print $v "[:]"}}{{ end }}
		{{- $fnCall := GenerateFnCall $c $v "Size" $f.Type $p $f.Options }}
		{{- if or (not $f.Options) (not $f.Options.Ignore) }}
			{{- if eq $l 1 }}
				return {{ $fnCall }}
			{{- else }}
				{{- if eq $i 0 }}
					size = {{ $fnCall }}
				{{- end }}
				{{- if and (gt $i 0) (lt $i (minus $l 1)) }}
					size += {{ $fnCall }}
				{{- end }}
				{{- if eq $i (minus $l 1) }}
					return size + {{ $fnCall }}
				{{- end }}
			{{- end }}
		{{- end }}
	{{- end }}
{{- end }}	`
	templates["struct_skip.tmpl"] = `{{- $l := FieldsLen . }}
{{- $p := .Prefix }}
{{- if eq $l 0 }}
	return
{{- else }}
	{{- $c := .Conf }}
	{{- range $i, $f := Fields . }}
		{{- $fnCall := GenerateFnCall $c "" "Skip" $f.Type $p $f.Options }}
		{{- if or (not $f.Options) (not $f.Options.Ignore) }}
			{{- if eq $l 1 }}
				return {{ $fnCall }}
			{{- else }}
				{{- if eq $i 0 }}
					n, err = {{ $fnCall }}
				{{- end }}
				{{- if eq $i 1 }}
					var n1 int
				{{- end }}
				{{- if ge $i 1 }}
					n1, err = {{ $fnCall }}
					n += n1
				{{- end }}
				{{- if lt $i (minus $l 1) }}
					if err != nil {
						return
					}
				{{- end }}
				{{- if eq $i (minus $l 1) }}
					return
				{{- end }}
			{{- end }}
		{{- end }}
	{{- end }}
{{- end }}	`
	templates["struct_unmarshal.tmpl"] = `{{- $l := FieldsLen . }}
{{- $r := Receiver .Name }}
{{- $p := .Prefix }}
{{- if eq $l 0 }}
	return
{{- else }}
	{{- $c := .Conf }}
	{{- range $i, $f := Fields . }}
		{{- $a := ArrayType $f.Type}}
		{{- $v := print $r "." $f.Name }}
		{{- $fnCall := GenerateFnCall $c $v "Unmarshal" $f.Type $p $f.Options }}
		{{- if or (not $f.Options) (not $f.Options.Ignore) }}

{{- /* if only one field */}}

			{{- if eq $l 1 }}
				{{- if $a }}
					{{ $r }}{{ $f.Name }}, n, err := {{ $fnCall }}
					if err != nil {
						return
					}
					{{ $v }} = ({{ $f.Type }})({{ $r }}{{ $f.Name }})
				{{- else }}
					{{ $v }}, n, err = {{ $fnCall }}
				{{- end }}

				{{- if and $f.Options (ne $f.Options.Validator "") }}
					{{- if not $a }}
						if err != nil {
							return
						}
					{{- end }}
					err = {{ $f.Options.Validator }}({{ $v }})
				{{- end }}
				return
			{{- else }}

{{- /* if first field */}}

				{{- if eq $i 0 }}
					{{- if $a }}
						{{ $r }}{{ $f.Name }}, n, err := {{ $fnCall }}
						if err != nil {
							return
						}
						{{ $v }} = ({{ $f.Type }})({{ $r }}{{ $f.Name }})
					{{- else }}
						{{ $v }}, n, err = {{ $fnCall }}
					{{- end }}
				{{- end }}

{{- /* if second field */}}

				{{- if eq $i 1 }}
					var n1 int
				{{- end }}

{{- /* if > second field */}}

				{{- if ge $i 1 }}
					{{- if $a }}
						{{ $r }}{{ $f.Name }}, n1, err := {{ $fnCall }}
						n += n1
						if err != nil {
							return
						}
						{{ $v }} = ({{ $f.Type }})({{ $r }}{{ $f.Name }})
					{{- else }}
						{{ $v }}, n1, err = {{ $fnCall }}
						n += n1
					{{- end}}
				{{- end }}

{{- /* if not the last one field */}}

				{{- if lt $i (minus $l 1) }}
					{{- if not $a }}
						if err != nil {
							return
						}
					{{- end }}
					{{- if and $f.Options (ne $f.Options.Validator "") }}
						if err = {{ $f.Options.Validator }}({{ $v }}); err != nil {
							return
						}
					{{- end }}
				{{- end }}

{{- /* if the last one field */}}

				{{- if eq $i (minus $l 1) }}
					{{- if and $f.Options (ne $f.Options.Validator "") }}
						{{- if not $a }}
							if err != nil {
								return
							}
						{{- end }}
						err = {{ $f.Options.Validator }}({{ $v }})
					{{- end }}
					return
				{{- end }}

			{{- end }}
		{{- end }}
	{{- end }}
{{- end }}`
}
