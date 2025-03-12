package tdesc

import (
	typeops "github.com/mus-format/musgen-go/options/type"
)

type FieldDesc struct {
	Name string
	Type string
	Tops *typeops.Options
}
