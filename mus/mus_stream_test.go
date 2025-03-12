package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func TestStreamFileGenerator(t *testing.T) {

	t.Run("Test FileGenerator.Generate pkg1", func(t *testing.T) {
		g := NewFileGenerator(genops.WithPackage("pkg1"),
			genops.WithImports([]string{
				"github.com/mus-format/musgen-go/testdata/pkg2",
			}),
			genops.WithStream(),
		)

		// typedef

		err := g.AddTypedef(reflect.TypeFor[pkg1.IntAliasStream]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.SliceAliasStream]())
		if err != nil {
			t.Fatal(err)
		}

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
			g,
		)
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
			genops.WithStream(),
		)

		err := g.AddStruct(reflect.TypeFor[pkg2.StructStream]())
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

	// t.Run("Test FileGenerator.AddTypedef", func(t *testing.T) {
	// 	g := NewFileGenerator(genops.WithPackage("pkg1"),
	// 		genops.WithStream(),
	// 	)
	// 	err := g.AddTypedef(reflect.TypeFor[pkg1.IntAliasStream]())
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = g.AddTypedefWith(reflect.TypeFor[pkg1.ByteSliceAlias](), StreamPrefix, nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = g.AddTypedefWith(reflect.TypeFor[pkg1.SliceAlias](), StreamPrefix, nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	bs, err := g.Generate()
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = os.WriteFile("../testdata/pkg1/mus-format-stream-alias.gen.go", bs, 0755)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// })

	// t.Run("Test FileGenerator.AddDTS", func(t *testing.T) {
	// 	g, err := NewFileGenerator(basegen.Conf{Package: "pkg1", Stream: true})
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = g.AddStructDTSWith(reflect.TypeFor[pkg1.InterfaceImpl1](), StreamPrefix, nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = g.AddStructDTSWith(reflect.TypeFor[pkg1.InterfaceImpl2](), StreamPrefix, nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	bs, err := g.Generate()
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = os.WriteFile("../testdata/pkg1/mus-format-stream-dts.gen.go", bs, 0755)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// })

	// t.Run("Test FileGenerator.AddInterface", func(t *testing.T) {
	// 	g, err := NewFileGenerator(basegen.Conf{Package: "pkg1", Stream: true})
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = g.AddInterfaceWith(reflect.TypeFor[pkg1.Interface](), StreamPrefix,
	// 		basegen.InterfaceOptions{
	// 			Oneof: []reflect.Type{
	// 				reflect.TypeFor[pkg1.InterfaceImpl1](),
	// 				reflect.TypeFor[pkg1.InterfaceImpl2](),
	// 			},
	// 		})
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	bs, err := g.Generate()
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = os.WriteFile("../testdata/pkg1/mus-format-stream-interface.gen.go", bs, 0755)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// })

	// t.Run("Test FileGenerator.AddStruct", func(t *testing.T) {
	// 	// pkg2
	// 	g2, err := NewFileGenerator(basegen.Conf{Package: "pkg2", Stream: true})
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = g2.AddStructWith(reflect.TypeFor[pkg2.Struct](), StreamPrefix, nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	bs, err := g2.Generate()
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = os.WriteFile("../testdata/pkg2/mus-format-stream-struct.gen.go", bs, 0755)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	// pkg1
	// 	g1, err := NewFileGenerator(basegen.Conf{
	// 		Package: "pkg1",
	// 		Stream:  true,
	// 		Imports: []string{
	// 			"github.com/mus-format/musgen-go/testdata/pkg2",
	// 		},
	// 	})
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = g1.AddStructWith(reflect.TypeFor[pkg1.Struct](), StreamPrefix, nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = g1.AddStructWith(reflect.TypeFor[pkg1.ComplexStruct](), StreamPrefix, nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	bs, err = g1.Generate()
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	err = os.WriteFile("../testdata/pkg1/mus-format-stream-struct.gen.go", bs, 0755)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// })

}
