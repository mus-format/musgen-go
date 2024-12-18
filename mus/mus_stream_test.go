package mus

import (
	"os"
	"reflect"
	"testing"

	"github.com/mus-format/musgen-go/basegen"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

const StreamPrefix = "Stream"

func TestStreamFileGenerator(t *testing.T) {

	t.Run("Test FileGenerator.AddAlias", func(t *testing.T) {
		g, err := NewFileGenerator(basegen.Conf{Package: "pkg1", Stream: true})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.IntAlias](), StreamPrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), StreamPrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-stream-alias.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.AddDTS", func(t *testing.T) {
		g, err := NewFileGenerator(basegen.Conf{Package: "pkg1", Stream: true})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStructDTSWith(reflect.TypeFor[pkg1.InterfaceImpl1](), StreamPrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStructDTSWith(reflect.TypeFor[pkg1.InterfaceImpl2](), StreamPrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-stream-dts.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.AddInterface", func(t *testing.T) {
		g, err := NewFileGenerator(basegen.Conf{Package: "pkg1", Stream: true})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddInterfaceWith(reflect.TypeFor[pkg1.Interface](), StreamPrefix,
			basegen.InterfaceOptions{
				Oneof: []reflect.Type{
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
		err = os.WriteFile("../testdata/pkg1/mus-format-stream-interface.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.AddStruct", func(t *testing.T) {
		// pkg2
		g2, err := NewFileGenerator(basegen.Conf{Package: "pkg2", Stream: true})
		if err != nil {
			t.Fatal(err)
		}
		err = g2.AddStructWith(reflect.TypeFor[pkg2.Struct](), StreamPrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g2.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg2/mus-format-stream-struct.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}

		// pkg1
		g1, err := NewFileGenerator(basegen.Conf{
			Package: "pkg1",
			Stream:  true,
			Imports: []string{
				"github.com/mus-format/musgen-go/testdata/pkg2",
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		err = g1.AddStructWith(reflect.TypeFor[pkg1.Struct](), StreamPrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		err = g1.AddStructWith(reflect.TypeFor[pkg1.ComplexStruct](), StreamPrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		bs, err = g1.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-stream-struct.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

}

func TestStreamGeneratedCode(t *testing.T) {

}
