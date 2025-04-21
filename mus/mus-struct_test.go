package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	struct_testdata "github.com/mus-format/musgen-go/testdata/struct"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestStructTypeGeneration(t *testing.T) {
	g, err := NewFileGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/struct"),
		genops.WithPackage("testdata"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testdata", "common_testdata"),
	)
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[struct_testdata.MyInt]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[struct_testdata.MySlice]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[struct_testdata.MyStruct]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDTS(reflect.TypeFor[struct_testdata.MyInt]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddInterface(reflect.TypeFor[struct_testdata.MyInterface](),
		introps.WithImpl(reflect.TypeFor[struct_testdata.MyInt]()))
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[struct_testdata.ComplexStruct](),
		structops.WithField(), // bool
		structops.WithField(), // byte
		structops.WithField(typeops.WithNumEncoding(typeops.Raw)),    // int8
		structops.WithField(typeops.WithNumEncoding(typeops.Varint)), // int16
		structops.WithField(), // int32

		structops.WithField(), // int64
		structops.WithField(typeops.WithNumEncoding(typeops.VarintPositive)), // uint8
		structops.WithField(), // uint16
		structops.WithField(), // uint32
		structops.WithField(), // uint64

		structops.WithField(
			typeops.WithValidator("common_testdata.ValidateZeroValue[float32]"),
		), // float32
		structops.WithField(), // float64
		structops.WithField(
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithLenValidator("common_testdata.ValidateLength20"),
			typeops.WithValidator("common_testdata.ValidateZeroValue[string]"),
		), // string
		structops.WithField(
			typeops.WithLenEncoding(typeops.Varint),
			typeops.WithLenValidator("common_testdata.ValidateLength20"),
		), // []byte
		structops.WithField(
			typeops.WithLenValidator("common_testdata.ValidateLength20"),
			typeops.WithElem(
				typeops.WithValidator("common_testdata.ValidateZeroValue[MyStruct]"),
			),
		), // []MyStruct

		structops.WithField(
			typeops.WithLenValidator("common_testdata.ValidateLength20"),
			typeops.WithElem(
				typeops.WithNumEncoding(typeops.Raw),
				typeops.WithValidator("common_testdata.ValidateZeroValue[int]")),
		), // [3]int
		structops.WithField(), // *string
		structops.WithField(), // *MyStruct
		structops.WithField(), // *string
		structops.WithField(
			typeops.WithElem(
				typeops.WithLenEncoding(typeops.Raw),
				typeops.WithElem(
					typeops.WithValidator("common_testdata.ValidateZeroValue[int]"),
				),
			),
		), // *[3]int

		structops.WithField(
			typeops.WithLenValidator("common_testdata.ValidateLength20"),
			typeops.WithKey(
				typeops.WithValidator("common_testdata.ValidateZeroValue[float32]"),
			),
			typeops.WithElem(
				typeops.WithLenValidator("common_testdata.ValidateLength20"),
				typeops.WithKey(
					typeops.WithValidator("common_testdata.ValidateZeroValue[MyInt]"),
				),
				typeops.WithElem(
					typeops.WithLenValidator("common_testdata.ValidateLength20"),
				),
			),
		), // map[float32]map[MyInt][]MyStruct
		structops.WithField(), // time.Time
		structops.WithField(), // MySlice
		structops.WithField(), // MyInterface
	)
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/struct/mus-format.gen.go", bs, 0755)
	assertfatal.EqualError(err, nil, t)
}
