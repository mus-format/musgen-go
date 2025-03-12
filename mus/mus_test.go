package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/parser"
	"github.com/mus-format/musgen-go/tdesc"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
	assert "github.com/ymz-ncnk/assert/error"
)

func TestFileGenerator(t *testing.T) {

	t.Run("Test FileGenerator.Generate pkg1", func(t *testing.T) {
		g := NewFileGenerator(
			genops.WithPackage("pkg1"),
			genops.WithImports([]string{
				"github.com/mus-format/musgen-go/testdata",
				"github.com/mus-format/musgen-go/testdata/pkg2",
			}),
		)

		// int

		err := g.AddTypedef(reflect.TypeFor[pkg1.IntAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.RawIntAlias](),
			typeops.WithNumEncoding(typeops.Raw),
		)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.VarintPositiveIntAlias](),
			typeops.WithNumEncoding(typeops.VarintPositive),
		)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ValidIntAlias](),
			typeops.WithValidator("testdata.ValidateZeroValue"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.AllIntAlias](),
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testdata.ValidateZeroValue"))
		if err != nil {
			t.Fatal(err)
		}

		// string

		err = g.AddTypedef(reflect.TypeFor[pkg1.StringAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.LenEncodingStringAlias](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.LenValidStringAlias](),
			typeops.WithLenValidator("testdata.ValidateLength"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ValidStringAlias](),
			typeops.WithValidator("testdata.ValidateZeroValue"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.AllStringAlias](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithLenValidator("testdata.ValidateLength"),
			typeops.WithValidator("testdata.ValidateZeroValue"))
		if err != nil {
			t.Fatal(err)
		}

		// byte_slice

		err = g.AddTypedef(reflect.TypeFor[pkg1.ByteSliceAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.LenEncodingByteSliceAlias](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.LenValidByteSliceAlias](),
			typeops.WithLenValidator("testdata.ValidateLength"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ValidByteSliceAlias](),
			typeops.WithValidator("ValidateByteSlice1"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.AllByteSliceAlias](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithLenValidator("testdata.ValidateLength"),
			typeops.WithValidator("ValidateByteSlice2"))
		if err != nil {
			t.Fatal(err)
		}

		// slice

		err = g.AddTypedef(reflect.TypeFor[pkg1.SliceAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.LenEncodingSliceAlias](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.LenValidSliceAlias](),
			typeops.WithLenValidator("testdata.ValidateLength"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ElemEncodingSliceAlias](),
			typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ElemValidSliceAlias](),
			typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.AllSliceAlias](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithLenValidator("testdata.ValidateLength"),
			typeops.WithElem(
				typeops.WithNumEncoding(typeops.Raw),
				typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}

		// array

		err = g.AddTypedef(reflect.TypeFor[pkg1.ArrayAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.LenEncodingArrayAlias](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ElemEncodingArrayAlias](),
			typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ElemValidArrayAlias](),
			typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ValidArrayAlias](),
			typeops.WithValidator("testdata.ValidateZeroValue[[3]int]"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.AllArrayAlias](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithValidator("testdata.ValidateZeroValue[[3]int]"),
			typeops.WithElem(
				typeops.WithNumEncoding(typeops.Raw),
				typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}

		// map

		err = g.AddTypedef(reflect.TypeFor[pkg1.MapAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.LenEncodingMapAlias](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.LenValidMapAlias](),
			typeops.WithLenValidator("testdata.ValidateLength"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.KeyEncodingMapAlias](),
			typeops.WithKey(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.KeyValidMapAlias](),
			typeops.WithKey(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ElemEncodingMapAlias](),
			typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ElemValidMapAlias](),
			typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ValidMapAlias](),
			typeops.WithValidator("ValidateMapAlias1"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.AllMapAlias](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithLenValidator("testdata.ValidateLength"),
			typeops.WithValidator("ValidateMapAlias2"),
			typeops.WithKey(
				typeops.WithNumEncoding(typeops.Raw),
				typeops.WithValidator("testdata.ValidateZeroValue[int]")),
			typeops.WithElem(
				typeops.WithNumEncoding(typeops.Raw),
				typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}

		// ptr

		err = g.AddTypedef(reflect.TypeFor[pkg1.PtrAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ElemNumEncodingPtrAlias](),
			typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddTypedef(reflect.TypeFor[pkg1.ValidPtrAlias](),
			typeops.WithValidator("testdata.ValidateZeroValue[*int]"))
		if err != nil {
			t.Fatal(err)
		}

		// struct

		err = g.AddTypedef(reflect.TypeFor[pkg1.SimpleStructPtrAlias]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.SimpleStruct]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.AnotherStruct]())
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddStruct(reflect.TypeFor[pkg1.ComplexStruct]())
		if err != nil {
			t.Fatal(err)
		}

		// struct, alias, dts, interface
		err = addInterface(reflect.TypeFor[pkg1.Interface](),
			reflect.TypeFor[pkg1.InterfaceImpl1](),
			reflect.TypeFor[pkg1.InterfaceImpl2](),
			g,
		)
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddInterface(reflect.TypeFor[pkg1.AnotherInterface](),
			introps.WithImplType(reflect.TypeFor[pkg1.InterfaceImpl1]()),
			introps.WithImplType(reflect.TypeFor[pkg1.InterfaceImpl2]()))
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddInterface(reflect.TypeFor[pkg1.Interface]())
		if err != tdesc.ErrNoOneImplTypeSpecified {
			t.Errorf("unexpected error, want %v actual %v", tdesc.ErrNoOneImplTypeSpecified,
				err)
		}

		err = g.AddTypedef(reflect.TypeFor[pkg1.InterfaceDoublePtrAlias]())
		if err != nil {
			t.Fatal(err)
		}

		// generate

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("AddDTS should fail if receives unsupported type", func(t *testing.T) {
		var (
			tp      = reflect.TypeFor[int]()
			wantErr = parser.NewUnsupportedType(tp)
			g       = NewFileGenerator(genops.WithPackage("pkg2"))
		)
		err := g.AddDTS(tp)
		assert.EqualError(err, wantErr, t)
	})

	t.Run("AddInterface should fail if receives no one ImplType", func(t *testing.T) {
		var (
			tp      = reflect.TypeFor[pkg1.Interface]()
			wantErr = tdesc.ErrNoOneImplTypeSpecified
			g       = NewFileGenerator(genops.WithPackage("pkg2"))
		)
		err := g.AddInterface(tp)
		assert.EqualError(err, wantErr, t)
	})

	t.Run("Test FileGenerator.Generate pkg2", func(t *testing.T) {
		g := NewFileGenerator(
			genops.WithPackage("pkg2"),
			genops.WithImports([]string{"github.com/mus-format/musgen-go/testdata"}),
		)

		err := g.AddStruct(reflect.TypeFor[pkg2.Struct]())
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddStruct(reflect.TypeFor[pkg2.IgnoreFieldStruct](),
			structops.WithNil(),
			structops.WithField(typeops.WithIgnore()),
			structops.WithNil(),
			structops.WithNil(),
			structops.WithNil(),
		)
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddStruct(reflect.TypeFor[pkg2.ValidFieldStruct](),
			structops.WithNil(),
			structops.WithNil(),
			structops.WithField(
				typeops.WithValidator("testdata.ValidateZeroValue[byte]"),
			),
			structops.WithNil(),
			structops.WithNil(),
		)
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddStruct(reflect.TypeFor[pkg2.ElemStruct](),
			structops.WithNil(),
			structops.WithNil(),
			structops.WithNil(),
			structops.WithNil(),
			structops.WithField(
				typeops.WithElem(
					typeops.WithElem(
						typeops.WithNumEncoding(typeops.Raw),
					),
					typeops.WithLenValidator("testdata.ValidateLength1"),
				),
				typeops.WithLenValidator("testdata.ValidateLength1"),
			),
		)
		if err != nil {
			t.Fatal(err)
		}

		// TODo WithField()

		// generate

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg2/mus-format.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Test FileGenerator.Generate pkg3", func(t *testing.T) {
		g := NewFileGenerator(
			genops.WithPackage("pkg3"),
			genops.WithImports([]string{
				"github.com/mus-format/musgen-go/testdata",
				"github.com/mus-format/musgen-go/testdata/pkg1",
			}),
		)

		// interface from another pkg

		err := addInterface(reflect.TypeFor[pkg1.Interface](),
			reflect.TypeFor[pkg1.InterfaceImpl1](),
			reflect.TypeFor[pkg1.InterfaceImpl2](),
			g,
		)
		if err != nil {
			t.Fatal(err)
		}

		// generate

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg3/mus-format.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}
	})

}

func addInterface(intr, impl1, impl2 reflect.Type, g *FileGenerator) (err error) {
	err = g.AddStruct(impl1)
	if err != nil {
		return
	}
	err = g.AddTypedef(impl2)
	if err != nil {
		return
	}
	err = g.AddDTS(impl1)
	if err != nil {
		return
	}
	err = g.AddDTS(impl2)
	if err != nil {
		return
	}
	err = g.AddInterface(intr,
		introps.WithImplType(impl1),
		introps.WithImplType(impl2),
	)
	return
}
