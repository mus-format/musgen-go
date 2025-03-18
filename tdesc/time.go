package tdesc

import (
	"reflect"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
)

func MakeTimeDesc(t reflect.Type, tops *typeops.Options, gops genops.Options) (
	td TypeDesc) {
	sourceType := "time.Time"
	td = makeDesc(t, gops)
	td.SourceType = sourceType
	td.Fields = []FieldDesc{{Type: sourceType, Tops: tops}}
	return
}
