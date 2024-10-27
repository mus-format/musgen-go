package parser

import (
	"io"
	"math/big"
	"reflect"
	"testing"
)

// TODO Test intr, _, _, err := Parse()

func TestParsePrimitiveType(t *testing.T) {
	var (
		v    int
		want = ErrUnsupportedType
	)
	_, _, _, err := Parse(reflect.TypeOf(v))
	if err == nil || err.Error() != want.Error() {
		t.Errorf("want '%v', actual '%v'", want, err)
	}
}

func TestParsePtrType(t *testing.T) {

	t.Run("Unsupported simple pointer type", func(t *testing.T) {
		var (
			v    *int
			want = ErrUnsupportedType
		)
		_, _, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

	t.Run("Unsupported simple pointer type alias", func(t *testing.T) {
		type IntPtrAlias *int
		var (
			v    IntPtrAlias
			want = ErrUnsupportedType
		)
		_, _, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

	t.Run("Unsupported atruct pointer", func(t *testing.T) {
		type Struct struct{}
		var (
			v    *Struct
			want = ErrUnsupportedType
		)
		_, _, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

	t.Run("Unsupported struct pointer alias", func(t *testing.T) {
		type Struct struct{}
		type StructAlias *Struct
		var (
			v    StructAlias
			want = ErrUnsupportedType
		)
		_, _, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

}

func TestParseInterfaceAlias(t *testing.T) {

	t.Run("Interface alias as nil type", func(t *testing.T) {
		type InterfaceAlias io.Reader
		var (
			v    InterfaceAlias
			want = ErrUnsupportedType
		)
		_, _, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

	t.Run("Unsupported interface pointer alias", func(t *testing.T) {
		type InterfacePtrAlias **io.Reader
		var (
			v    InterfacePtrAlias
			want = ErrUnsupportedType
		)
		_, _, _, err := Parse(reflect.TypeOf(v))
		if err == nil || err.Error() != want.Error() {
			t.Errorf("want '%v', actual '%v'", want, err)
		}
	})

}

func TestParseTypeFromAnotherPkg(t *testing.T) {
	type MyStruct struct {
		big **big.Int
	}
	var (
		v           MyStruct
		wantIntr    = false
		wantAliasOf = ""
		wantFields  = []string{"**big.Int"}
	)
	test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
	_ = v.big
}

func TestParsePrimitiveTypeAlias(t *testing.T) {

	t.Run("Bool alias", func(t *testing.T) {
		type BoolAlias bool
		var (
			v           BoolAlias
			wantIntr             = false
			wantAliasOf          = "bool"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
	})

	t.Run("Uint64 alias", func(t *testing.T) {
		type Uint64Alias uint64
		var (
			v           Uint64Alias
			wantIntr             = false
			wantAliasOf          = "uint64"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
	})

	t.Run("Uint32 alias", func(t *testing.T) {
		type Uint32Alias uint32
		var (
			v           Uint32Alias
			wantIntr             = false
			wantAliasOf          = "uint32"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
	})

	t.Run("Uint16 alias", func(t *testing.T) {
		type Uint16Alias uint16
		var (
			v           Uint16Alias
			wantIntr             = false
			wantAliasOf          = "uint16"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
	})

	t.Run("Uint8 alias", func(t *testing.T) {
		type Uint8Alias uint8
		var (
			v           Uint8Alias
			wantIntr             = false
			wantAliasOf          = "uint8"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
	})

	t.Run("Uint alias", func(t *testing.T) {
		type UintAlias uint
		var (
			v           UintAlias
			wantIntr             = false
			wantAliasOf          = "uint"
			wantFields  []string = nil
		)
		test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
	})

}

func TestParseArrayAlias(t *testing.T) {
	type ArrayAlias [3]int
	var (
		v           ArrayAlias
		wantIntr             = false
		wantAliasOf          = "[3]int"
		wantFields  []string = nil
	)
	test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
}

func TestParseSliceAlias(t *testing.T) {
	type SliceAlias []string
	var (
		v           SliceAlias
		wantIntr             = false
		wantAliasOf          = "[]string"
		wantFields  []string = nil
	)
	test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
}

func TestParseMapAlias(t *testing.T) {
	type MapAlias map[int32]float64
	var (
		v           MapAlias
		wantIntr             = false
		wantAliasOf          = "map[int32]float64"
		wantFields  []string = nil
	)
	test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
}

func TestParseStruct(t *testing.T) {
	type Struct struct {
		R int
		T float64
	}
	var (
		v           Struct
		wantIntr             = false
		wantAliasOf          = ""
		wantFields  []string = []string{"int", "float64"}
	)
	test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
}

func TestParseStructAlias(t *testing.T) {
	type Struct struct {
		uint8Field uint8
		F          float32
	}
	type StructAlias Struct
	var (
		v           StructAlias
		wantIntr             = false
		wantAliasOf          = ""
		wantFields  []string = []string{"uint8", "float32"}
	)
	test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
	_ = v.uint8Field
}

func TestParseNestedMapAlias(t *testing.T) {
	type NestedMapAlias map[*map[int64]map[uint8]uint32]**map[string]map[int]**string
	var (
		v           NestedMapAlias
		wantIntr             = false
		wantAliasOf          = "map[*map[int64]map[uint8]uint32]**map[string]map[int]**string"
		wantFields  []string = nil
	)
	test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
}

func TestParseTrickyStructAlias(t *testing.T) {
	type AnotherStruct struct {
		uint8Field uint8
	}
	type ArrayAlias [2]int
	type Struct struct {
		boolField       bool
		ArrayField      [3]*[]**[3]map[int]string
		StructField     AnotherStruct
		StructPtrField  *AnotherStruct
		MapField        map[AnotherStruct]map[ArrayAlias][]*AnotherStruct
		ArrayAliasField ArrayAlias
	}
	type FirstStructAlias Struct
	type SecondStructAlias FirstStructAlias
	var (
		v           SecondStructAlias
		wantIntr             = false
		wantAliasOf          = ""
		wantFields  []string = []string{
			"bool",
			"[3]*[]**[3]map[int]string",
			"AnotherStruct",
			"*AnotherStruct",
			"map[AnotherStruct]map[ArrayAlias][]*AnotherStruct",
			"ArrayAlias",
		}
	)
	test(reflect.TypeOf(v), wantIntr, wantAliasOf, wantFields, t)
	_ = v.boolField
	_ = v.StructField.uint8Field
}

func test(tp reflect.Type, wantIntr bool, wantAliasOf string, wantFields []string,
	t *testing.T) {
	intr, aliasOf, fieldsTypes, err := Parse(tp)
	if err != nil {
		t.Error(err)
	}
	if intr != wantIntr {
		t.Errorf("want '%v', actual '%v'", wantIntr, intr)
	}
	if aliasOf != wantAliasOf {
		t.Errorf("want '%v', actual '%v'", wantAliasOf, aliasOf)
	}
	if !reflect.DeepEqual(fieldsTypes, wantFields) {
		t.Errorf("want '%v', actual '%v'", wantFields, fieldsTypes)
	}
}
