package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	testdata "github.com/mus-format/musgen-go/testdata/pointer"
	"github.com/mus-format/musgen-go/typename"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestPointerGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/pointer"),
		genops.WithPackage("testdata"),
	)
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyIntPtr]())
	assertfatal.EqualError(err, nil, t)

	tp := reflect.TypeFor[testdata.MyDoubleIntPtr]()
	err = g.AddDefinedType(tp)
	assertfatal.EqualError(err, typename.NewMultiPointerError(tp), t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.MySlicePtr]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[testdata.MyStruct]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyStructPtr]())
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/pointer/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(err, nil, t)
}
