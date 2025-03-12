package pkg1

import (
	"fmt"

	"github.com/mus-format/musgen-go/testdata/pkg2"
)

type IntAliasStreamUnsafe int

type SliceAliasStreamUnsafe []int

type SimpleStructStreamUnsafe struct {
	Int int
}

type ComplexStructStreamUnsafe struct {
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

	Alias            SliceAliasStreamUnsafe
	AnotherPkgStruct pkg2.StructStreamUnsafe
	Interface        InterfaceStreamUnsafe

	ByteSlice   []byte
	StructSlice []SimpleStructStreamUnsafe

	Array [3]int

	PtrString *string
	PtrStruct *SimpleStructStreamUnsafe
	NilPtr    *string
	PtrArray  *[3]int

	Map map[float32]map[IntAliasStreamUnsafe][]SimpleStructStreamUnsafe
}

// -----------------------------------------------------------------------------

const (
	InterfaceImpl1StreamUnsafeDTM = 11
	InterfaceImpl2StreamUnsafeDTM = 12
	// SimpleStructStreamUnsafeDTM   = 12
	// IntAliasStreamUnsafeDTM       = 14
)

type InterfaceStreamUnsafe interface {
	Print()
}

type InterfaceImpl1StreamUnsafe struct{}

func (i InterfaceImpl1StreamUnsafe) Print() {
	fmt.Println("impl1")
}

type InterfaceImpl2StreamUnsafe int

func (i InterfaceImpl2StreamUnsafe) Print() {
	fmt.Println("impl2")
}
