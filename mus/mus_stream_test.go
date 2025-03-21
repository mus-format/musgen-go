package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func TestStreamFileGenerator(t *testing.T) {

	t.Run("Test FileGenerator.Generate pkg1", func(t *testing.T) {
		g := NewFileGenerator(genops.WithPackage("pkg1"),
			genops.WithImports([]string{
				"github.com/mus-format/musgen-go/testdata",
				"github.com/mus-format/musgen-go/testdata/pkg2",
			}),
			genops.WithStream(),
		)

		// defined_type

		err := g.AddDefinedType(reflect.TypeFor[pkg1.MyIntStream]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.MySliceStream]())
		if err != nil {
			t.Fatal(err)
		}

		t.Run("Import statement should contain options packages for valid constructors",
			func(t *testing.T) {
				err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyStringStream](),
					typeops.WithLenValidator("testdata.ValidateLength"))
				if err != nil {
					t.Fatal(err)
				}
				err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyArrayStream](),
					typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
				if err != nil {
					t.Fatal(err)
				}
				err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyByteSliceStream](),
					typeops.WithLenValidator("testdata.ValidateLength"))
				if err != nil {
					t.Fatal(err)
				}
				err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMySliceStream](),
					typeops.WithLenValidator("testdata.ValidateLength"))
				if err != nil {
					t.Fatal(err)
				}
				err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyMapStream](),
					typeops.WithLenValidator("testdata.ValidateLength"))
				if err != nil {
					t.Fatal(err)
				}
			})

		// struct

		err = g.AddStruct(reflect.TypeFor[pkg1.SimpleStructStream]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.ComplexStructStream]())
		if err != nil {
			t.Fatal(err)
		}

		// interface

		err = addInterface(reflect.TypeFor[pkg1.InterfaceStream](),
			reflect.TypeFor[pkg1.InterfaceImpl1Stream](),
			reflect.TypeFor[pkg1.InterfaceImpl2Stream](),
			g)
		if err != nil {
			t.Fatal(err)
		}

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format_stream.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.Generate pkg2", func(t *testing.T) {
		g := NewFileGenerator(genops.WithPackage("pkg2"),
			genops.WithStream())

		err := g.AddStruct(reflect.TypeFor[pkg2.StructStream]())
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddStruct(reflect.TypeFor[pkg2.TimeStructStream]())
		if err != nil {
			t.Fatal(err)
		}

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg2/mus-format_stream.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

}
