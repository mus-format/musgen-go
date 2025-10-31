package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	testdata "github.com/mus-format/musgen-go/testdata/interface_marshaller"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestInterfaceTypeWithMarshallerGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/interface_marshaller"),
		genops.WithPackage("testdata"),
	)
	assertfatal.EqualError(err, nil, t)

	tp1 := reflect.TypeFor[testdata.Impl1]()
	err = g.AddStruct(tp1)
	assertfatal.EqualError(err, nil, t)
	err = g.AddDTS(tp1)
	assertfatal.EqualError(err, nil, t)

	tp2 := reflect.TypeFor[testdata.Impl2]()
	err = g.AddStruct(tp2)
	assertfatal.EqualError(err, nil, t)
	err = g.AddDTS(tp2)
	assertfatal.EqualError(err, nil, t)

	err = g.AddInterface(reflect.TypeFor[testdata.MyInterface](),
		introps.WithImpl(tp1),
		introps.WithImpl(tp2),
		introps.WithMarshaller())
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/interface_marshaller/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(err, nil, t)
}
