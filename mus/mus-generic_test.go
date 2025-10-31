package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	typeops "github.com/mus-format/musgen-go/options/type"
	generic_testdata "github.com/mus-format/musgen-go/testdata/generic"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestGenericTypeGeneration(t *testing.T) {
	g, err := NewCodeGenerator(
		genops.WithPkgPath("github.com/mus-format/musgen-go/testdata/generic"),
		genops.WithPackage("testdata"),
		genops.WithImportAlias("github.com/mus-format/musgen-go/testdata",
			"common_testdata"),
	)
	assertfatal.EqualError(err, nil, t)

	// defined type

	err = g.AddDefinedType(reflect.TypeFor[generic_testdata.MyInt]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[generic_testdata.MySlice[generic_testdata.MyInt]]())
	assertfatal.EqualError(err, nil, t)

	//
	err = g.AddDefinedType(reflect.TypeFor[generic_testdata.MyArray[int]]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddDefinedType(reflect.TypeFor[generic_testdata.MyMap[generic_testdata.MyArray[int], generic_testdata.MyInt]](),
		typeops.WithKey(
			typeops.WithValidator("common_testdata.ValidateZeroValue[MyArray[int]]"),
		),
		typeops.WithElem(
			typeops.WithValidator("common_testdata.ValidateZeroValue[MyInt]"),
		),
	)
	assertfatal.EqualError(err, nil, t)
	//

	// struct

	err = g.AddStruct(reflect.TypeFor[generic_testdata.MyStruct[generic_testdata.MySlice[generic_testdata.MyInt]]]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[generic_testdata.MyDoubleParamStruct[int, generic_testdata.MyStruct[generic_testdata.MySlice[generic_testdata.MyInt]]]]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[generic_testdata.MyTripleParamStruct[generic_testdata.MySlice[generic_testdata.MyInt], generic_testdata.MyInterface[generic_testdata.MyInt], generic_testdata.MyDoubleParamStruct[int, generic_testdata.MyStruct[generic_testdata.MySlice[generic_testdata.MyInt]]]]]())
	assertfatal.EqualError(err, nil, t)

	err = g.AddStruct(reflect.TypeFor[generic_testdata.Impl[generic_testdata.MyInt]]())
	assertfatal.EqualError(err, nil, t)

	// DTS

	tp := reflect.TypeFor[generic_testdata.Impl[generic_testdata.MyInt]]()
	err = g.AddDTS(tp)
	assertfatal.EqualError(err, nil, t)

	// interface

	err = g.AddInterface(reflect.TypeFor[generic_testdata.MyInterface[generic_testdata.MyInt]](),
		introps.WithImpl(tp),
	)
	assertfatal.EqualError(err, nil, t)

	// generate

	bs, err := g.Generate()
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("../testdata/generic/mus-format.gen.go", bs, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
