package pkg1

import (
	"fmt"

	"github.com/mus-format/musgen-go/testdata/pkg2"
)

type MyIntUnsafe int

type MySliceUnsafe []int

type SimpleStructUnsafe struct {
	Int int
}

type ComplexStructUnsafe struct {
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

	Alias            MySliceUnsafe
	AnotherPkgStruct pkg2.StructUnsafe
	Interface        InterfaceUnsafe

	ByteSlice   []byte
	StructSlice []SimpleStructUnsafe

	Array [3]int

	PtrString *string
	PtrStruct *SimpleStructUnsafe
	NilPtr    *string
	PtrArray  *[3]int

	Map map[float32]map[MyIntUnsafe][]SimpleStructUnsafe
}

// -----------------------------------------------------------------------------

const (
	InterfaceImpl1UnsafeDTM = 11
	InterfaceImpl2UnsafeDTM = 12
	// SimpleStructUnsafeDTM   = 12
	// MyIntUnsafeDTM       = 14
)

type InterfaceUnsafe interface {
	Print()
}

type InterfaceImpl1Unsafe struct{}

func (i InterfaceImpl1Unsafe) Print() {
	fmt.Println("impl1")
}

type InterfaceImpl2Unsafe int

func (i InterfaceImpl2Unsafe) Print() {
	fmt.Println("impl2")
}
