package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	struct_testdata "github.com/mus-format/musgen-go/testdata/struct"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestWithStreamUnsafeGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/stream_unsafe"),
		genops.WithPackage("testdata"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testdata/struct",
			"struct_testdata"),
		genops.WithStream(),
		genops.WithUnsafe(),
	)
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[struct_testdata.MyInt]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[struct_testdata.MySlice]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[struct_testdata.MyStruct]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDTS(reflect.TypeFor[struct_testdata.MyInt]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddInterface(reflect.TypeFor[struct_testdata.MyInterface](),
		introps.WithImpl(reflect.TypeFor[struct_testdata.MyInt]()))
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[struct_testdata.ComplexStruct]())
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/stream_unsafe/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(err, nil, t)
}
