package mus

import (
	"os"
	"reflect"
	"testing"

	"github.com/mus-format/musgen-go/basegen"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

const UnsafePrefix = "Unsafe"

func TestUnsafeFileGenerator(t *testing.T) {

	t.Run("Test FileGenerator.AddAlias", func(t *testing.T) {
		g, err := NewFileGenerator(basegen.Conf{Package: "pkg1", Unsafe: true})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.BoolAlias](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.ByteAlias](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.IntAlias](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.StringAlias](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.Float32Alias](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.MapAlias](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-unsafe-alias.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.AddDTS", func(t *testing.T) {
		g, err := NewFileGenerator(basegen.Conf{Package: "pkg1", Unsafe: true})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStructDTSWith(reflect.TypeFor[pkg1.InterfaceImpl1](), UnsafePrefix,
			nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStructDTSWith(reflect.TypeFor[pkg1.InterfaceImpl2](), UnsafePrefix,
			nil)
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-unsafe-dts.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.AddInterface", func(t *testing.T) {
		g, err := NewFileGenerator(basegen.Conf{Package: "pkg1", Unsafe: true})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddInterfaceWith(reflect.TypeFor[pkg1.Interface](), UnsafePrefix,
			basegen.InterfaceMetadata{
				OneOf: []reflect.Type{
					reflect.TypeFor[pkg1.InterfaceImpl1](),
					reflect.TypeFor[pkg1.InterfaceImpl2](),
				},
			})
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-unsafe-interface.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.AddStruct", func(t *testing.T) {
		// pkg2
		g2, err := NewFileGenerator(basegen.Conf{Package: "pkg2", Unsafe: true})
		if err != nil {
			t.Fatal(err)
		}
		err = g2.AddStructWith(reflect.TypeFor[pkg2.Struct](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g2.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg2/mus-format-unsafe-struct.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}

		// pkg1
		g1, err := NewFileGenerator(basegen.Conf{
			Package: "pkg1",
			Unsafe:  true,
			Imports: []string{
				"github.com/mus-format/musgen-go/testdata/pkg2",
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		err = g1.AddStructWith(reflect.TypeFor[pkg1.Struct](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g1.AddStructWith(reflect.TypeFor[pkg1.ComplexStruct](), UnsafePrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		bs, err = g1.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-unsafe-struct.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

}
