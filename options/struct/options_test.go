package structops

import (
	"testing"

	typeops "github.com/mus-format/musgen-go/options/type"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestOptions(t *testing.T) {
	var (
		o          = New()
		wantFields = []*typeops.Options{{Ignore: true}, {Ignore: true}}
		wantTops   = &typeops.Options{TimeUnit: typeops.Milli}
	)
	Apply([]SetOption{
		WithField(typeops.WithIgnore()),
		WithField(typeops.WithIgnore()),
		WithTops(typeops.WithTimeUnit(typeops.Milli)),
	}, &o)
	asserterror.EqualDeep(o.Fields, wantFields, t)
	asserterror.Equal(*o.Tops, *wantTops, t)
}
