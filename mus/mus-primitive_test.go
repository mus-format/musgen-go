package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	testdata "github.com/mus-format/musgen-go/testdata/primitive"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestPrimitiveTypesGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/primitive"),
		genops.WithPackage("testdata"),
		genops.WithImport("github.com/mus-format/musgen-go/testdata"),
	)
	assertfatal.EqualError(err, nil, t)

	// bool

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyBool]())
	assertfatal.EqualError(err, nil, t)

	// byte

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyByte]())
	assertfatal.EqualError(err, nil, t)

	// float32

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyFloat32]())
	assertfatal.EqualError(err, nil, t)

	// float64

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyFloat64]())
	assertfatal.EqualError(err, nil, t)

	// int

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyInt]())
	assertfatal.EqualError(err, nil, t)
	err = g.AddDefinedType(reflect.TypeFor[testdata.RawMyInt](),
		typeops.WithNumEncoding(typeops.Raw),
	)
	assertfatal.EqualError(err, nil, t)
	err = g.AddDefinedType(reflect.TypeFor[testdata.VarintPositiveMyInt](),
		typeops.WithNumEncoding(typeops.VarintPositive),
	)
	assertfatal.EqualError(err, nil, t)
	err = g.AddDefinedType(reflect.TypeFor[testdata.ValidMyInt](),
		typeops.WithValidator("testdata.ValidateZeroValue"))
	assertfatal.EqualError(err, nil, t)
	err = g.AddDefinedType(reflect.TypeFor[testdata.AllMyInt](),
		typeops.WithNumEncoding(typeops.Raw),
		typeops.WithValidator("testdata.ValidateZeroValue"))
	assertfatal.EqualError(err, nil, t)

	// string

	err = g.AddDefinedType(reflect.TypeFor[testdata.MyString]())
	assertfatal.EqualError(err, nil, t)
	err = g.AddDefinedType(reflect.TypeFor[testdata.LenEncodingMyString](),
		typeops.WithLenEncoding(typeops.Raw))
	assertfatal.EqualError(err, nil, t)
	err = g.AddDefinedType(reflect.TypeFor[testdata.LenValidMyString](),
		typeops.WithLenValidator("testdata.ValidateLength3"))
	assertfatal.EqualError(err, nil, t)
	err = g.AddDefinedType(reflect.TypeFor[testdata.ValidMyString](),
		typeops.WithValidator("testdata.ValidateZeroValue"))
	assertfatal.EqualError(err, nil, t)
	err = g.AddDefinedType(reflect.TypeFor[testdata.AllMyString](),
		typeops.WithLenEncoding(typeops.Raw),
		typeops.WithLenValidator("testdata.ValidateLength3"),
		typeops.WithValidator("testdata.ValidateZeroValue"))
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	assertfatal.EqualError(err, nil, t)
	err = os.WriteFile("../testdata/primitive/mus-format.gen.go", bs, 0644)
	assertfatal.EqualError(err, nil, t)
}
