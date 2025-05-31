package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	struct_testdata "github.com/mus-format/musgen-go/testdata/struct"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestWithNotUnsafeGeneration(t *testing.T) {
	var (
		myIntType         = reflect.TypeFor[struct_testdata.MyInt]()
		mySliceType       = reflect.TypeFor[struct_testdata.MySlice]()
		myStructType      = reflect.TypeFor[struct_testdata.MyStruct]()
		complexStructType = reflect.TypeFor[struct_testdata.ComplexStruct]()
	)

	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/notunsafe"),
		genops.WithPackage("testdata"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testdata/struct",
			"struct_testdata"),
		genops.WithNotUnsafe(),
	)
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(myIntType)
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(mySliceType)
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(myStructType)
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(complexStructType)
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/notunsafe/mus-format.gen.go", bs, 0755)
	assertfatal.EqualError(err, nil, t)

}
