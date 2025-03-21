package parser

import (
	"io"
	"math/big"
	"reflect"
	"testing"

	assert "github.com/ymz-ncnk/assert/error"
)

func TestParseTypedef(t *testing.T) {

	t.Run("Should work for basic types", func(t *testing.T) {
		type IntDef int
		var (
			tp                   = reflect.TypeFor[IntDef]()
			wantSourceType       = "int"
			wantErr        error = nil
		)
		sourceType, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
		assert.Equal(sourceType, wantSourceType, t)
	})

	t.Run("Should work for array source type", func(t *testing.T) {
		type ArrayDef [3]int
		var (
			tp                   = reflect.TypeFor[ArrayDef]()
			wantSourceType       = "[3]int"
			wantErr        error = nil
		)
		sourceType, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
		assert.Equal(sourceType, wantSourceType, t)
	})

	t.Run("Should work for slice source type", func(t *testing.T) {
		type IntSliceDef []int
		var (
			tp                   = reflect.TypeFor[IntSliceDef]()
			wantSourceType       = "[]int"
			wantErr        error = nil
		)
		sourceType, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
		assert.Equal(sourceType, wantSourceType, t)
	})

	t.Run("Should work for slice of structs source type", func(t *testing.T) {
		type Struct struct{}
		type StructSliceDef []Struct
		var (
			tp                   = reflect.TypeFor[StructSliceDef]()
			wantSourceType       = "[]Struct"
			wantErr        error = nil
		)
		sourceType, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
		assert.Equal(sourceType, wantSourceType, t)
	})

	t.Run("Should work for slice of structs from another package source type",
		func(t *testing.T) {
			type StructSliceDef []big.Int
			var (
				tp                   = reflect.TypeFor[StructSliceDef]()
				wantSourceType       = "[]big.Int"
				wantErr        error = nil
			)
			sourceType, err := ParseTypedef(tp)
			assert.EqualError(err, wantErr, t)
			assert.Equal(sourceType, wantSourceType, t)
		})

	t.Run("Should work for map source type", func(t *testing.T) {
		type IntStringMapDef map[int]string
		var (
			tp                   = reflect.TypeFor[IntStringMapDef]()
			wantSourceType       = "map[int]string"
			wantErr        error = nil
		)
		sourceType, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
		assert.Equal(sourceType, wantSourceType, t)
	})

	t.Run("Should work for ptr source type", func(t *testing.T) {
		type PtrIntDef *int
		var (
			tp                   = reflect.TypeFor[PtrIntDef]()
			wantSourceType       = "*int"
			wantErr        error = nil
		)
		sourceType, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
		assert.Equal(sourceType, wantSourceType, t)
	})

	t.Run("Should fail for basic type", func(t *testing.T) {
		var (
			v       int
			tp      = reflect.TypeOf(v)
			wantErr = NewUnsupportedType(tp)
		)
		_, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
	})

	t.Run("Should fail for struct type", func(t *testing.T) {
		type Struct struct{}
		var (
			tp      = reflect.TypeFor[Struct]()
			wantErr = NewUnsupportedType(tp)
		)
		_, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
	})

	t.Run("Should fail for struct source type", func(t *testing.T) {
		type Struct struct{}
		type StructDef Struct
		var (
			tp      = reflect.TypeFor[StructDef]()
			wantErr = NewUnsupportedType(tp)
		)
		_, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
	})

	t.Run("Should fail for interface type", func(t *testing.T) {
		type Interface interface {
			Print()
		}
		var (
			tp      = reflect.TypeFor[Interface]()
			wantErr = NewUnsupportedType(tp)
		)
		_, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
	})

	t.Run("Should fail for interface source type", func(t *testing.T) {
		type Interface interface {
			Print()
		}
		type InterfaceDef Interface
		var (
			tp      = reflect.TypeFor[InterfaceDef]()
			wantErr = NewUnsupportedType(tp)
		)
		_, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)
	})

	t.Run("Should fail for any type", func(t *testing.T) {
		var (
			tp      = reflect.TypeFor[any]()
			wantErr = NewUnsupportedType(tp)
		)
		_, err := ParseTypedef(tp)
		assert.EqualError(err, wantErr, t)

	})
}

func TestParseStructType(t *testing.T) {

	t.Run("Should work", func(t *testing.T) {
		type AnotherStruct struct{}
		type Interface interface {
			Print()
		}
		type IntDef int
		type Struct struct {
			Int               int
			Array             [3]int
			Slice             []int
			Map               map[int]string
			Struct            AnotherStruct
			PkgStruct         big.Int
			StructSlice       []AnotherStruct
			PkgStructSlice    []big.Int
			PtrPkgStructSlice []*big.Int
			Interface         Interface
			PkgInterface      io.Reader
			DefinedType       IntDef
			Ptr               *int
			StructPtr         *Struct
			InterfacePtr      *Interface
			DefPtr            *IntDef
			DoublePtr         **int
			Complex           map[AnotherStruct]map[IntDef][]*Interface
		}
		var (
			tp             = reflect.TypeFor[Struct]()
			wantFieldTypes = []string{
				"int",
				"[3]int",
				"[]int",
				"map[int]string",
				"AnotherStruct",
				"big.Int",
				"[]AnotherStruct",
				"[]big.Int",
				"[]*big.Int",
				"Interface",
				"io.Reader",
				"IntDef",
				"*int",
				"*Struct",
				"*Interface",
				"*IntDef",
				"**int",
				"map[AnotherStruct]map[IntDef][]*Interface",
			}
			wantErr error = nil
		)
		fieldTypes, err := ParseStructType(tp)
		assert.EqualError(err, wantErr, t)
		for i := range wantFieldTypes {
			assert.Equal(fieldTypes[i], wantFieldTypes[i], t)
		}
	})

	t.Run("Should work for sruct type definition", func(t *testing.T) {
		type Struct struct {
			Int int
		}
		type StructDef Struct
		var (
			tp                   = reflect.TypeFor[StructDef]()
			wantFieldTypes       = []string{"int"}
			wantErr        error = nil
		)
		fieldTypes, err := ParseStructType(tp)
		assert.EqualError(err, wantErr, t)
		for i := range wantFieldTypes {
			assert.Equal(fieldTypes[i], wantFieldTypes[i], t)
		}
	})

	t.Run("Should fail if a field has any type", func(t *testing.T) {
		type Struct struct {
			Any any
		}
		var (
			tp      = reflect.TypeFor[Struct]()
			wantErr = NewUnsupportedType(reflect.TypeFor[any]())
		)
		_, err := ParseStructType(tp)
		assert.EqualError(err, wantErr, t)
	})

}

func TestParseInterface(t *testing.T) {

	t.Run("Should work", func(t *testing.T) {
		type Interface interface {
			Print()
		}
		var (
			tp            = reflect.TypeFor[Interface]()
			wantErr error = nil
		)
		err := ParseInterfaceType(tp)
		assert.EqualError(err, wantErr, t)
	})

	t.Run("Should work for interface type definition", func(t *testing.T) {
		type Interface interface {
			Print()
		}
		type InterfaceDef Interface
		var (
			tp            = reflect.TypeFor[InterfaceDef]()
			wantErr error = nil
		)
		err := ParseInterfaceType(tp)
		assert.EqualError(err, wantErr, t)
	})

	t.Run("Should fail for Interface without methods", func(t *testing.T) {
		var (
			tp      = reflect.TypeFor[any]()
			wantErr = NewUnsupportedType(tp)
		)
		err := ParseInterfaceType(tp)
		assert.EqualError(err, wantErr, t)
	})

}
