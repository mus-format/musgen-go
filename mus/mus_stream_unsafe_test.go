package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func TestStreamUnsafeFileGenerator(t *testing.T) {

	t.Run("Test FileGenerator.Generate pkg1", func(t *testing.T) {
		g := NewFileGenerator(genops.WithPackage("pkg1"),
			genops.WithImports([]string{
				"github.com/mus-format/musgen-go/testdata/pkg2",
			}),
			genops.WithUnsafe(),
			genops.WithStream(),
		)

		// typedef

		err := g.AddTypedef(reflect.TypeFor[pkg1.IntAliasStreamUnsafe]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.SliceAliasStreamUnsafe]())
		if err != nil {
			t.Fatal(err)
		}

		// struct

		err = g.AddStruct(reflect.TypeFor[pkg1.SimpleStructStreamUnsafe]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.ComplexStructStreamUnsafe]())
		if err != nil {
			t.Fatal(err)
		}

		// interface

		err = addInterface(reflect.TypeFor[pkg1.InterfaceStreamUnsafe](),
			reflect.TypeFor[pkg1.InterfaceImpl1StreamUnsafe](),
			reflect.TypeFor[pkg1.InterfaceImpl2StreamUnsafe](),
			g,
		)
		if err != nil {
			t.Fatal(err)
		}

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format_stream_unsafe.gen.go", bs,
			0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.Generate pkg2", func(t *testing.T) {
		g := NewFileGenerator(genops.WithPackage("pkg2"),
			genops.WithUnsafe(),
			genops.WithStream(),
		)

		err := g.AddStruct(reflect.TypeFor[pkg2.StructStreamUnsafe]())
		if err != nil {
			t.Fatal(err)
		}

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg2/mus-format_stream_unsafe.gen.go", bs,
			0755)
		if err != nil {
			t.Fatal(err)
		}
	})
}
