package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	typeops "github.com/mus-format/musgen-go/options/type"

	testdata "github.com/mus-format/musgen-go/testdata/register_interface"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestInterfaceTypeRegistration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/register_interface"),
		genops.WithPackage("testdata"),
	)
	assertfatal.EqualError(err, nil, t)

	err = g.RegisterInterface(reflect.TypeFor[testdata.MyInterface](),
		introps.WithStructImpl(reflect.TypeFor[testdata.Impl1]()),
		introps.WithDefinedTypeImpl(reflect.TypeFor[testdata.Impl2](),
			typeops.WithNumEncoding(typeops.Raw),
		),
	)
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/register_interface/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(err, nil, t)
}
