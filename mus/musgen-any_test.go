package musgen

import (
	"os"
	"reflect"
	"testing"

	"github.com/mus-format/musgen-go/data/builders"
	genops "github.com/mus-format/musgen-go/options/generate"
	testdata "github.com/mus-format/musgen-go/testdata/any"
	"github.com/mus-format/musgen-go/typename"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestAnyGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/any"),
		genops.WithPackage("testdata"))
	assertfatal.EqualError(err, nil, t)

	tp := reflect.TypeFor[testdata.MyAny]()
	err = g.AddDefinedType(tp)
	assertfatal.EqualError(err, builders.NewUnsupportedTypeError(tp), t)

	anyType := reflect.TypeFor[any]()

	tp = reflect.TypeFor[testdata.MyAnySlice]()
	err = g.AddDefinedType(tp)
	assertfatal.EqualError(err, typename.NewUnsupportedTypeError(
		reflect.TypeFor[any]()), t)

	err = g.AddStruct(reflect.TypeFor[testdata.MyAnyStruct]())
	assertfatal.EqualError(err, typename.NewUnsupportedTypeError(anyType), t)

	tp = reflect.TypeFor[testdata.MyAnyGenericSlice[any]]()
	err = g.AddDefinedType(tp)
	assertfatal.EqualError(err, typename.NewUnsupportedTypeError(
		reflect.TypeFor[any]()), t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/any/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(err, nil, t)
}
