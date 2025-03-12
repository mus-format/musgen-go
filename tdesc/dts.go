package tdesc

import (
	"reflect"

	genops "github.com/mus-format/musgen-go/options/generate"
	"github.com/mus-format/musgen-go/parser"
)

func MakeDTSDesc(t reflect.Type, gops genops.Options) (td TypeDesc, err error) {
	err = parser.SupportedType(t)
	if err != nil {
		return
	}
	td = makeDesc(t, gops)
	return
}
