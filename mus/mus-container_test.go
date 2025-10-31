package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	testdata "github.com/mus-format/musgen-go/testdata/container"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestContainerTypeGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/container"),
		genops.WithPackage("testdata"),
		genops.WithImport("github.com/mus-format/musgen-go/testdata"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testdata/generic",
			"generic_testdata"),
	)
	assertfatal.EqualError(err, nil, t)

	// array

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyArray]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.LenEncodingMyArray](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ElemEncodingMyArray](),
		typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ElemValidMyArray](),
		typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ValidMyArray](),
		typeops.WithValidator("testdata.ValidateZeroValue[ValidMyArray]"))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.AllMyArray](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testdata.ValidateLength"),
		typeops.WithElem(
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testdata.ValidateZeroValue[int]")),
		typeops.WithValidator("testdata.ValidateZeroValue[AllMyArray]"),
	)
	assertfatal.EqualError(err, nil, t)
	// byte_slice

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyByteSlice]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.LenEncodingMyByteSlice](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.LenValidMyByteSlice](),
		typeops.WithLenValidator("testdata.ValidateLength"))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ValidMyByteSlice](),
		typeops.WithValidator("ValidateByteSlice1"))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.AllMyByteSlice](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testdata.ValidateLength"),
		typeops.WithValidator("ValidateByteSlice2"))
	assertfatal.EqualError(err, nil, t)

	// slice

	err = g.AddDefinedType(reflect.TypeFor[testdata.MySlice]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.LenEncodingMySlice](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.LenValidMySlice](),
		typeops.WithLenValidator("testdata.ValidateLength"))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ElemEncodingMySlice](),
		typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ElemValidMySlice](),
		typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ValidMySlice](),
		typeops.WithValidator("ValidateMySlice1"))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.AllMySlice](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testdata.ValidateLength"),
		typeops.WithElem(
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testdata.ValidateZeroValue[int]")),
		typeops.WithValidator("ValidateMySlice2"))
	assertfatal.EqualError(err, nil, t)

	// map

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyMap]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.LenEncodingMyMap](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.LenValidMyMap](),
		typeops.WithLenValidator("testdata.ValidateLength"))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.KeyEncodingMyMap](),
		typeops.WithKey(typeops.WithNumEncoding(typeops.Raw)))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.KeyValidMyMap](),
		typeops.WithKey(typeops.WithValidator("testdata.ValidateZeroValue[int]")))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ElemEncodingMyMap](),
		typeops.WithElem(typeops.WithNumEncoding(typeops.Raw)))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ElemValidMyMap](),
		typeops.WithElem(typeops.WithValidator("testdata.ValidateZeroValue[string]")))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ValidMyMap](),
		typeops.WithValidator("ValidateMyMap1"))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.AllMyMap](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testdata.ValidateLength"),
		typeops.WithValidator("ValidateMyMap2"),
		typeops.WithKey(
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testdata.ValidateZeroValue[int]")),
		typeops.WithElem(
			typeops.WithNumEncoding(typeops.Raw),
			typeops.WithValidator("testdata.ValidateZeroValue[string]")))
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[testdata.ComplexMap](),
		typeops.WithKey(
			typeops.WithLenEncoding(typeops.Raw),
			typeops.WithValidator("testdata.ValidateZeroValue[[3]int]"),
		),
		typeops.WithElem(
			typeops.WithElem(
				typeops.WithLenEncoding(typeops.Raw),
			),
		),
	)
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/container/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(err, nil, t)
}
