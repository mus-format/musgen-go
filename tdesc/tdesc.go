package tdesc

import (
	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
)

type TypeDesc struct {
	Name       string
	FullName   string
	SourceType string
	Fields     []FieldDesc
	Package    string
	Gops       genops.Options
	Sops       structops.Options
	Iops       introps.Options
}
