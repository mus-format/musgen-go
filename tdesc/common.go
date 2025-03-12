package tdesc

import (
	"reflect"
	"regexp"

	genops "github.com/mus-format/musgen-go/options/generate"
)

func makeDesc(t reflect.Type, gops genops.Options) (td TypeDesc) {
	td.Gops = gops
	td.Package = pkg(t)
	td.Name = t.Name()
	td.FullName = MakeFullName(t, gops)
	return
}

func MakeFullName(t reflect.Type, gops genops.Options) (name string) {
	p := pkg(t)
	if gops.Package != p {
		return p + "." + t.Name()
	}
	return t.Name()
}

func pkg(t reflect.Type) string {
	re := regexp.MustCompile(`^(.*)\.`)
	match := re.FindStringSubmatch(t.String())
	if len(match) != 2 {
		return ""
	}
	return match[1]
}
