package tdesc

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	structops "github.com/mus-format/musgen-go/options/struct"
)

type TypeDesc struct {
	Name       string
	FullName   string
	SourceType string
	Oneof      []string
	Fields     []FieldDesc
	Package    string
	Gops       genops.Options
	Sops       structops.Options
}
