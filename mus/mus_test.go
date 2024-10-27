package mus

import (
	"os"
	"reflect"
	"testing"

	"github.com/mus-format/musgen-go/basegen"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func TestFileGenerator(t *testing.T) {

	t.Run("Test FileGenerator.AddAlias", func(t *testing.T) {
		g, err := NewFileGenerator(basegen.Conf{
			Package: "pkg1",
			Imports: []string{"github.com/mus-format/musgen-go/testdata"},
		})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAlias(reflect.TypeFor[pkg1.IntAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.IntAlias](), "Raw",
			basegen.NumMetadata{Encoding: basegen.Raw})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.IntAlias](), "VarintPositive",
			basegen.NumMetadata{Encoding: basegen.VarintPositive})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.IntAlias](), "Valid",
			basegen.NumMetadata{Validator: "testdata.ValidateZeroValue[int]"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAlias(reflect.TypeFor[pkg1.StringAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.StringAlias](), "LenEncoding",
			basegen.StringMetadata{LenEncoding: basegen.Raw})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.StringAlias](), "LenValidator",
			basegen.StringMetadata{LenValidator: "testdata.ValidateLength"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.StringAlias](), "Valid",
			basegen.StringMetadata{Validator: "testdata.ValidateZeroValue[string]"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAlias(reflect.TypeFor[pkg1.SliceAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), "LenEncoding",
			basegen.SliceMetadata{LenEncoding: basegen.Raw})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), "LenValidator",
			basegen.SliceMetadata{LenValidator: "testdata.ValidateLength"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), "ElemEncoding",
			basegen.SliceMetadata{Elem: basegen.NumMetadata{Encoding: basegen.Raw}})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), "ElemValidator",
			basegen.SliceMetadata{Elem: basegen.NumMetadata{
				Validator: "testdata.ValidateZeroValue[int]"},
			})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAlias(reflect.TypeFor[pkg1.MapAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.MapAlias](), "LenEncoding",
			basegen.MapMetadata{LenEncoding: basegen.Raw})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.MapAlias](), "LenValidator",
			basegen.MapMetadata{LenValidator: "testdata.ValidateLength"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.MapAlias](), "KeyValidator",
			basegen.MapMetadata{
				Key: basegen.StringMetadata{Validator: "testdata.ValidateZeroValue[string]"},
			})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.MapAlias](), "ElemValidator",
			basegen.MapMetadata{
				Elem: basegen.NumMetadata{Validator: "testdata.ValidateZeroValue[int]"},
			})
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-alias.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.AddDTS", func(t *testing.T) {
		g, err := NewFileGenerator(basegen.Conf{Package: "pkg1"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStructDTS(reflect.TypeFor[pkg1.InterfaceImpl1]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStructDTS(reflect.TypeFor[pkg1.InterfaceImpl2]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStructDTS(reflect.TypeFor[pkg1.Struct]())
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-dts.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.AddInterface", func(t *testing.T) {
		g, err := NewFileGenerator(basegen.Conf{Package: "pkg1"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddInterface(reflect.TypeFor[pkg1.Interface](),
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
		err = os.WriteFile("../testdata/pkg1/mus-format-interface.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.AddStruct", func(t *testing.T) {
		// pkg2
		g2, err := NewFileGenerator(basegen.Conf{
			Package: "pkg2",
			Imports: []string{
				"github.com/mus-format/musgen-go/testdata",
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		err = g2.AddStruct(reflect.TypeFor[pkg2.Struct]())
		if err != nil {
			t.Fatal(err)
		}
		// There is no test function for this case.
		err = g2.AddStructWith(reflect.TypeFor[pkg2.Struct](), "ValidAllFields",
			basegen.StructMetadata{
				basegen.NumFieldMetadata{
					NumMetadata: basegen.NumMetadata{Validator: "testdata.ValidateZeroValue[float32]"},
				},
				basegen.NumFieldMetadata{
					NumMetadata: basegen.NumMetadata{Validator: "testdata.ValidateZeroValue[float64]"},
				},
				nil,
				basegen.BoolFieldMetadata{
					BoolMetadata: basegen.BoolMetadata{Validator: "testdata.ValidateZeroValue[bool]"},
				},
			},
		)
		if err != nil {
			t.Fatal(err)
		}
		// There is no test function for this case.
		err = g2.AddStructWith(reflect.TypeFor[pkg2.Struct](), "ValidateFirstField",
			basegen.StructMetadata{
				basegen.NumFieldMetadata{
					NumMetadata: basegen.NumMetadata{Validator: "testdata.ValidateZeroValue[float32]"},
				},
				nil,
				nil,
				nil,
			})
		if err != nil {
			t.Fatal(err)
		}
		// There is no test function for this case.
		err = g2.AddStructWith(reflect.TypeFor[pkg2.Struct](), "ValidateLastField",
			basegen.StructMetadata{
				nil,
				nil,
				nil,
				basegen.BoolFieldMetadata{
					BoolMetadata: basegen.BoolMetadata{Validator: "testdata.ValidateZeroValue[bool]"},
				},
			})
		if err != nil {
			t.Fatal(err)
		}
		// There is no test function for this case.
		err = g2.AddStructWith(reflect.TypeFor[pkg2.AnotherStruct](), "Prefix",
			basegen.StructMetadata{
				basegen.CustomTypeFieldMetadata{Prefix: "Skip"},
				basegen.CustomTypeFieldMetadata{Prefix: basegen.EmptyPrefix},
			},
		)
		if err != nil {
			t.Fatal(err)
		}
		err = g2.AddStructWith(reflect.TypeFor[pkg2.Struct](), "Skip",
			basegen.StructMetadata{
				basegen.NumFieldMetadata{Ignore: true},
				nil,
				nil,
				nil})
		if err != nil {
			t.Fatal(err)
		}
		bs, err := g2.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg2/mus-format-struct.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}

		// pkg1
		g1, err := NewFileGenerator(basegen.Conf{
			Package: "pkg1",
			Imports: []string{
				"github.com/mus-format/musgen-go/testdata/pkg2",
			}},
		)
		if err != nil {
			t.Fatal(err)
		}
		err = g1.AddStruct(reflect.TypeFor[pkg1.ComplexStruct]())
		if err != nil {
			t.Fatal(err)
		}
		bs, err = g1.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format-struct.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

}
