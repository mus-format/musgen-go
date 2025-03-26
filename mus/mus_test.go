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

		err := g.AddDefinedType(reflect.TypeFor[pkg1.MyInt]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.RawMyInt](),
			typeops.WithNumEncoding(typeops.Raw),
		)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.VarintPositiveMyInt](),
			typeops.WithNumEncoding(typeops.VarintPositive),
		)
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyInt](),
			typeops.WithValidator("testdata.ValidateZeroValue"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.AllMyInt](),
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testdata.ValidateZeroValue"))
		if err != nil {
			t.Fatal(err)
		}

		// string

		err = g.AddDefinedType(reflect.TypeFor[pkg1.MyString]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.LenEncodingMyString](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.LenValidMyString](),
			typeops.WithLenValidator("testdata.ValidateLength"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyString](),
			typeops.WithValidator("testdata.ValidateZeroValue"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.AllMyString](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithLenValidator("testdata.ValidateLength"),
			typeops.WithValidator("testdata.ValidateZeroValue"))
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddStruct(reflect.TypeFor[pkg1.MyTime](),
			structops.WithSourceType(structops.Time))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.MyTimeSec](),
			structops.WithSourceType(structops.Time, typeops.WithTimeUnit(typeops.Sec)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.MyTimeMilli](),
			structops.WithSourceType(structops.Time, typeops.WithTimeUnit(typeops.Milli)))

		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.MyTimeMicro](),
			structops.WithSourceType(structops.Time, typeops.WithTimeUnit(typeops.Micro)))

		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.MyTimeNano](),
			structops.WithSourceType(structops.Time, typeops.WithTimeUnit(typeops.Nano)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.MyTimeSecUTC](),
			structops.WithSourceType(structops.Time, typeops.WithTimeUnit(typeops.SecUTC)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.MyTimeMilliUTC](),
			structops.WithSourceType(structops.Time, typeops.WithTimeUnit(typeops.MilliUTC)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.MyTimeMicroUTC](),
			structops.WithSourceType(structops.Time, typeops.WithTimeUnit(typeops.MicroUTC)))

		if err != nil {
			t.Fatal(err)
		}
		err = g.AddStruct(reflect.TypeFor[pkg1.MyTimeNanoUTC](),
			structops.WithSourceType(structops.Time, typeops.WithTimeUnit(typeops.NanoUTC)))
		if err != nil {
			t.Fatal(err)
		}

		// array

		err = g.AddDefinedType(reflect.TypeFor[pkg1.MyArray]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.LenEncodingMyArray](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ElemEncodingMyArray](),
			typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ElemValidMyArray](),
			typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyArray](),
			typeops.WithValidator("testdata.ValidateZeroValue[[3]int]"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.AllMyArray](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithValidator("testdata.ValidateZeroValue[[3]int]"),
			typeops.WithElem(
				typeops.WithNumEncoding(typeops.Raw),
				typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}

		// byte_slice

		err = g.AddDefinedType(reflect.TypeFor[pkg1.MyByteSlice]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.LenEncodingMyByteSlice](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.LenValidMyByteSlice](),
			typeops.WithLenValidator("testdata.ValidateLength"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyByteSlice](),
			typeops.WithValidator("ValidateByteSlice1"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.AllMyByteSlice](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithLenValidator("testdata.ValidateLength"),
			typeops.WithValidator("ValidateByteSlice2"))
		if err != nil {
			t.Fatal(err)
		}

		// slice

		err = g.AddDefinedType(reflect.TypeFor[pkg1.MySlice]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.LenEncodingMySlice](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.LenValidMySlice](),
			typeops.WithLenValidator("testdata.ValidateLength"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ElemEncodingMySlice](),
			typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ElemValidMySlice](),
			typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.AllMySlice](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithLenValidator("testdata.ValidateLength"),
			typeops.WithElem(
				typeops.WithNumEncoding(typeops.Raw),
				typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}

		// map

		err = g.AddDefinedType(reflect.TypeFor[pkg1.MyMap]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.LenEncodingMyMap](),
			typeops.WithLenEncoding(typeops.Raw))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.LenValidMyMap](),
			typeops.WithLenValidator("testdata.ValidateLength"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.KeyEncodingMyMap](),
			typeops.WithKey(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.KeyValidMyMap](),
			typeops.WithKey(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ElemEncodingMyMap](),
			typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ElemValidMyMap](),
			typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyMap](),
			typeops.WithValidator("ValidateMyMap1"))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.AllMyMap](),
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithLenValidator("testdata.ValidateLength"),
			typeops.WithValidator("ValidateMyMap2"),
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

		err = g.AddDefinedType(reflect.TypeFor[pkg1.MyIntPtr]())
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ElemNumEncodingMyIntPtr](),
			typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
		if err != nil {
			t.Fatal(err)
		}
		err = g.AddDefinedType(reflect.TypeFor[pkg1.ValidMyIntPtr](),
			typeops.WithValidator("testdata.ValidateZeroValue[*int]"))
		if err != nil {
			t.Fatal(err)
		}

		// struct

		err = g.AddDefinedType(reflect.TypeFor[pkg1.SimpleStructMyIntPtr]())
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

		err = g.AddDefinedType(reflect.TypeFor[pkg1.InterfaceDoubleMyIntPtr]())
		if err != nil {
			t.Fatal(err)
		}

		err = g.AddDefinedType(reflect.TypeFor[pkg1.UnsuportedChanType]())
		assert.EqualError(err,
			parser.NewUnsupportedType(reflect.TypeFor[pkg1.UnsuportedChanType]()), t)

		err = g.AddStruct(reflect.TypeFor[pkg1.UnsupportedFuncFieldStruct]())
		assert.EqualError(err, parser.NewUnsupportedTypeStr("func()"), t)

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

		// struct with time field

		err = g.AddStruct(reflect.TypeFor[pkg2.TimeStruct]())
		if err != nil {
			t.Fatal(err)
		}

		// TODO WithField()

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
	err = g.AddDefinedType(impl2)
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
