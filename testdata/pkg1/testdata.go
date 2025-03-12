package pkg1

import (
	"fmt"

	"github.com/mus-format/musgen-go/testdata/pkg2"
)

const (
	InterfaceImpl1DTM = 1
	InterfaceImpl2DTM = 2
	SimpleStructDTM   = 3
	IntAliasDTM       = 4
)

type BoolAlias bool
type ByteAlias byte
type Float32Alias float32

type IntAlias int
type RawIntAlias int
type VarintPositiveIntAlias int
type ValidIntAlias int
type AllIntAlias int

type StringAlias string
type LenEncodingStringAlias string
type LenValidStringAlias string
type ValidStringAlias string
type AllStringAlias string

type ByteSliceAlias []byte
type LenEncodingByteSliceAlias []byte
type LenValidByteSliceAlias []byte
type ValidByteSliceAlias []byte
type AllByteSliceAlias []byte

type SliceAlias []int
type LenEncodingSliceAlias []int
type LenValidSliceAlias []int
type ElemEncodingSliceAlias []int
type ElemValidSliceAlias []int
type ValidSliceAlias []int
type AllSliceAlias []int

type ArrayAlias [3]int
type LenEncodingArrayAlias [3]int
type ElemEncodingArrayAlias [3]int
type ElemValidArrayAlias [3]int
type ValidArrayAlias [3]int
type AllArrayAlias [3]int

type MapAlias map[int]int
type LenEncodingMapAlias map[int]int
type LenValidMapAlias map[int]int
type KeyEncodingMapAlias map[int]int
type KeyValidMapAlias map[int]int
type ElemEncodingMapAlias map[int]int
type ElemValidMapAlias map[int]int
type ValidMapAlias map[int]int
type AllMapAlias map[int]int

type PtrAlias *int
type ElemNumEncodingPtrAlias *int

type ValidPtrAlias *int

type SimpleStructPtrAlias *SimpleStruct

type InterfaceDoublePtrAlias **Interface

type StructAlias SimpleStruct

type InterfaceAlias Interface

type AnotherStruct SimpleStruct
type AnotherInterface Interface

// -----------------------------------------------------------------------------

type SimpleStruct struct {
	Int int
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
	AnotherPkgStruct pkg2.Struct
	Interface        Interface

	ByteSlice   []byte
	StructSlice []SimpleStruct

	Array [3]int

	PtrString *string
	PtrStruct *SimpleStruct
	NilPtr    *string
	PtrArray  *[3]int

	Map map[float32]map[IntAlias][]SimpleStruct
}

// -----------------------------------------------------------------------------

type Interface interface {
	Print()
}

type InterfaceImpl1 struct {
	Str string
}

func (i InterfaceImpl1) Print() {
	fmt.Println("impl1")
}

type InterfaceImpl2 int

func (i InterfaceImpl2) Print() {
	fmt.Println("impl2")
}

// ----------------------------------------------------------------------------
