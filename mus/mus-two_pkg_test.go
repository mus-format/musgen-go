package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	testdata "github.com/mus-format/musgen-go/testdata/two_pkg"
	pkg "github.com/mus-format/musgen-go/testdata/two_pkg/pkg"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestTwoPkgGeneration(t *testing.T) {

	t.Run("First pkg", func(t *testing.T) {
		tp := reflect.TypeFor[pkg.MyIntSerName]()
		g, err := NewCodeGenerator(
			genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/two_pkg/pkg"),
			genops.WithSerName(tp, "MyAwesomeInt"),
		)
		assertfatal.EqualError(err, nil, t)

		err = g.AddDefinedType(reflect.TypeFor[pkg.MyInt]())
		assertfatal.EqualError(err, nil, t)

		err = g.AddDefinedType(reflect.TypeFor[pkg.MySlice]())
		assertfatal.EqualError(err, nil, t)

		err = g.AddDefinedType(tp)
		assertfatal.EqualError(err, nil, t)

		// generate

		bs, err := g.Generate()
		assertfatal.EqualError(err, nil, t)
		err = os.WriteFile("../testdata/two_pkg/pkg/mus-format.gen.go", bs, 0644)
		assertfatal.EqualError(err, nil, t)
	})

	t.Run("Second pkg", func(t *testing.T) {
		tp := reflect.TypeFor[pkg.MyIntSerName]()
		g, err := NewCodeGenerator(
			genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/two_pkg"),
			genops.WithPackage("testdata"),
			genops.WithImport("github.com/mus-format/musgen-go/testdata/two_pkg/pkg"),
			genops.WithSerName(tp, "pkg.MyAwesomeInt"),
		)
		assertfatal.EqualError(err, nil, t)

		err = g.AddDefinedType(reflect.TypeFor[testdata.MySlice]())
		assertfatal.EqualError(err, nil, t)

		err = g.AddDefinedType(reflect.TypeFor[testdata.MySliceSerName]())
		assertfatal.EqualError(err, nil, t)

		// generate

		bs, err := g.Generate()
		assertfatal.EqualError(err, nil, t)
		err = os.WriteFile("../testdata/two_pkg/mus-format.gen.go", bs, 0644)
		assertfatal.EqualError(err, nil, t)
	})

}
