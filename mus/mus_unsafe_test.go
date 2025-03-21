package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func TestUnsafeFileGenerator(t *testing.T) {

	t.Run("Test FileGenerator.Generate pkg1", func(t *testing.T) {
		g := NewFileGenerator(genops.WithPackage("pkg1"),
			genops.WithImports([]string{
				"github.com/mus-format/musgen-go/testdata/pkg2",
			}),
			genops.WithUnsafe(),
		)

		// defined_type

		err := g.AddDefinedType(reflect.TypeFor[pkg1.MyIntUnsafe]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.MySliceUnsafe]())
		if err != nil {
			t.Fatal(err)
		}

		// struct

		err = g.AddStruct(reflect.TypeFor[pkg1.SimpleStructUnsafe]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.ComplexStructUnsafe]())
		if err != nil {
			t.Fatal(err)
		}

		// interface

		err = addInterface(reflect.TypeFor[pkg1.InterfaceUnsafe](),
			reflect.TypeFor[pkg1.InterfaceImpl1Unsafe](),
			reflect.TypeFor[pkg1.InterfaceImpl2Unsafe](),
			g,
		)
		if err != nil {
			t.Fatal(err)
		}

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format_unsafe.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.Generate pkg2", func(t *testing.T) {
		g := NewFileGenerator(genops.WithPackage("pkg2"),
			genops.WithUnsafe(),
		)

		err := g.AddStruct(reflect.TypeFor[pkg2.StructUnsafe]())
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddStruct(reflect.TypeFor[pkg2.TimeStructUnsafe]())
		if err != nil {
			t.Fatal(err)
		}

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg2/mus-format_unsafe.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})
}
