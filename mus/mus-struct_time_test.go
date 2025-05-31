package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	sttime_testdata "github.com/mus-format/musgen-go/testdata/struct_time"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestStructTimeGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/struct_time"),
		genops.WithPackage("testdata"),
	)
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[sttime_testdata.MyMicroTime](),
		structops.WithTops(
			typeops.WithSourceType(typeops.Time),
			typeops.WithTimeUnit(typeops.Micro),
		))
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[sttime_testdata.MyDefaultTime](),
		structops.WithTops(
			typeops.WithSourceType(typeops.Time),
		),
	)
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/struct_time/mus-format.gen.go", bs, 0755)
	assertfatal.EqualError(err, nil, t)

}
