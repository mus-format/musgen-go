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
			basegen.NumOptions{Encoding: basegen.Raw})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.IntAlias](), "VarintPositive",
			basegen.NumOptions{Encoding: basegen.VarintPositive})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.IntAlias](), "Valid",
			basegen.NumOptions{Validator: "testdata.ValidateZeroValue[int]"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAlias(reflect.TypeFor[pkg1.StringAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.StringAlias](), "LenEncoding",
			basegen.StringOptions{LenEncoding: basegen.Raw})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.StringAlias](), "LenValidator",
			basegen.StringOptions{LenValidator: "testdata.ValidateLength"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.StringAlias](), "Valid",
			basegen.StringOptions{Validator: "testdata.ValidateZeroValue[string]"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAlias(reflect.TypeFor[pkg1.ByteSliceAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAlias(reflect.TypeFor[pkg1.SliceAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), "LenEncoding",
			basegen.SliceOptions{LenEncoding: basegen.Raw})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), "LenValidator",
			basegen.SliceOptions{LenValidator: "testdata.ValidateLength"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), "ElemEncoding",
			basegen.SliceOptions{Elem: basegen.NumOptions{Encoding: basegen.Raw}})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.SliceAlias](), "ElemValidator",
			basegen.SliceOptions{Elem: basegen.NumOptions{
				Validator: "testdata.ValidateZeroValue[int]"},
			})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAlias(reflect.TypeFor[pkg1.ArrayAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.ArrayAlias](), "LenEncoding",
			basegen.SliceOptions{LenEncoding: basegen.Raw})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.ArrayAlias](), "ElemEncoding",
			basegen.ArrayOptions{Elem: basegen.NumOptions{Encoding: basegen.Raw}})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.ArrayAlias](), "ElemValidator",
			basegen.ArrayOptions{Elem: basegen.NumOptions{
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
			basegen.MapOptions{LenEncoding: basegen.Raw})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.MapAlias](), "LenValidator",
			basegen.MapOptions{LenValidator: "testdata.ValidateLength"})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.MapAlias](), "KeyValidator",
			basegen.MapOptions{
				Key: basegen.StringOptions{Validator: "testdata.ValidateZeroValue[string]"},
			})
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddAliasWith(reflect.TypeFor[pkg1.MapAlias](), "ElemValidator",
			basegen.MapOptions{
				Elem: basegen.NumOptions{Validator: "testdata.ValidateZeroValue[int]"},
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
			basegen.StructOptions{
				basegen.NumFieldOptions{
					NumOptions: basegen.NumOptions{Validator: "testdata.ValidateZeroValue[float32]"},
				},
				basegen.NumFieldOptions{
					NumOptions: basegen.NumOptions{Validator: "testdata.ValidateZeroValue[float64]"},
				},
				nil,
				basegen.BoolFieldOptions{
					BoolOptions: basegen.BoolOptions{Validator: "testdata.ValidateZeroValue[bool]"},
				},
			},
		)
		if err != nil {
			t.Fatal(err)
		}
		// There is no test function for this case.
		err = g2.AddStructWith(reflect.TypeFor[pkg2.Struct](), "ValidateFirstField",
			basegen.StructOptions{
				basegen.NumFieldOptions{
					NumOptions: basegen.NumOptions{Validator: "testdata.ValidateZeroValue[float32]"},
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
			basegen.StructOptions{
				nil,
				nil,
				nil,
				basegen.BoolFieldOptions{
					BoolOptions: basegen.BoolOptions{Validator: "testdata.ValidateZeroValue[bool]"},
				},
			})
		if err != nil {
			t.Fatal(err)
		}
		// There is no test function for this case.
		err = g2.AddStructWith(reflect.TypeFor[pkg2.AnotherStruct](), "Prefix",
			basegen.StructOptions{
				basegen.CustomTypeFieldOptions{Prefix: "Skip"},
				basegen.CustomTypeFieldOptions{Prefix: basegen.EmptyPrefix},
			},
		)
		if err != nil {
			t.Fatal(err)
		}
		err = g2.AddStructWith(reflect.TypeFor[pkg2.Struct](), "Skip",
			basegen.StructOptions{
				basegen.NumFieldOptions{Ignore: true},
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
