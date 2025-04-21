package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	testdata "github.com/mus-format/musgen-go/testdata/crossgen"
	pkg "github.com/mus-format/musgen-go/testdata/crossgen/pkg"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestCrossGeneration(t *testing.T) {

	t.Run("pkg", func(t *testing.T) {
		g, err := NewFileGenerator(
			genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/crossgen/pkg"),
		)
		assertfatal.EqualError(err, nil, t)

		// defined type

		err = g.AddDefinedType(reflect.TypeFor[pkg.MyInt]())
		assertfatal.EqualError(err, nil, t)

		// generate

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/crossgen/pkg/mus-format.gen.go", bs, 0755)
		assertfatal.EqualError(err, nil, t)
	})

	g, err := NewFileGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/crossgen"),
		genops.WithPackage("testdata"),
		genops.WithImport("github.com/mus-format/musgen-go/testdata/crossgen/pkg"),
	)
	assertfatal.EqualError(err, nil, t)

	// defined type

	// err = g.AddDefinedType(reflect.TypeFor[pkg.MyInt]())
	// assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[pkg.MySlice]())
	assertfatal.EqualError(err, nil, t)

	// struct

	tp := reflect.TypeFor[pkg.MyStruct]()
	err = g.AddStruct(tp)
	assertfatal.EqualError(err, nil, t)

	// defined type

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyMap]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[pkg.MyArray[testdata.MyMap]]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[pkg.MyAnotherArray[pkg.MyInt]]())
	assertfatal.EqualError(err, nil, t)

	// struct
	err = g.AddStruct(reflect.TypeFor[testdata.MyStructWithCrossgen]())
	assertfatal.EqualError(err, nil, t)

	// interface

	err = g.AddDTS(tp)
	assertfatal.EqualError(err, nil, t)

	err = g.AddInterface(reflect.TypeFor[pkg.MyInterface](),
		introps.WithImpl(tp),
	)
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("../testdata/crossgen/mus-format.gen.go", bs, 0755)
	assertfatal.EqualError(err, nil, t)
}
