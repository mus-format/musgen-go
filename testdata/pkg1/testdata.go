package pkg1

import (
	"fmt"

	"github.com/mus-format/musgen-go/testdata/pkg2"
)

const (
	InterfaceImpl1DTM = 1
	InterfaceImpl2DTM = 2
	StructDTM         = 3
)

type BoolAlias bool

type ByteAlias byte

type IntAlias int

type StringAlias string

type Float32Alias float32

type SliceAlias []int

type ArrayAlias [3]int

type MapAlias map[string]int

type Struct struct {
	Int int
}

type StructAlias Struct

type Interface interface {
	Print()
}

type InterfaceImpl1 struct{}

func (i InterfaceImpl1) Print() {
	fmt.Println("impl1")
}

type InterfaceImpl2 struct{}

func (i InterfaceImpl2) Print() {
	fmt.Println("impl2")
}

type ComplexStruct struct {
	Bool bool
	Byte byte

	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64

	Float32 float32
	Float64 float64

	String string

	Alias            SliceAlias
	Ptr              *Struct
	AnotherPkgStruct pkg2.Struct
	Interface        Interface

	SliceByte   []byte
	SliceStruct []Struct

	Array [3]int

	Map map[float32]map[IntAlias][]Struct
}

// func ValidateNum[T comparable](t T) (err error) {
// 	if t == *new(T) {
// 		err = errs.ErrZeroValue
// 	}
// 	return
// }

// func ValidateString(str string) (err error) {
// 	if str == "" {
// 		err = errs.ErrZeroValue
// 	}
// 	return
// }

// func ValidateLength(l int) (err error) {
// 	if l > 0 {
// 		err = errs.ErrTooLong
// 	}
// 	return
// }

// func ValidateMap(m map[string]int) (err error) {
// 	return errs.ErrNotValid
// }

// func ValidateSimpleStructPtr(p *Struct) (err error) {
// 	return errs.ErrNotValid
// }
