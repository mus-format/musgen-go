package pkg1

import (
	"fmt"

	"github.com/mus-format/musgen-go/testdata/pkg2"
)

type MyIntStream int

type MySliceStream []int
type ValidMyStringStream string
type ValidMyArrayStream [3]int
type ValidMyByteSliceStream []byte
type ValidMySliceStream []int
type ValidMyMapStream map[int]int

type SimpleStructStream struct {
	Int int
}

type ComplexStructStream struct {
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

	Alias            MySliceStream
	AnotherPkgStruct pkg2.StructStream
	Interface        InterfaceStream

	ByteSlice   []byte
	StructSlice []SimpleStructStream

	Array [3]int

	PtrString *string
	PtrStruct *SimpleStructStream
	NilPtr    *string
	PtrArray  *[3]int

	Map map[float32]map[MyIntStream][]SimpleStructStream
}

// -----------------------------------------------------------------------------

const (
	InterfaceImpl1StreamDTM = 11
	InterfaceImpl2StreamDTM = 12
	// SimpleStructStreamDTM   = 12
	// MyIntStreamDTM       = 14
)

type InterfaceStream interface {
	Print()
}

type InterfaceImpl1Stream struct{}

func (i InterfaceImpl1Stream) Print() {
	fmt.Println("impl1")
}

type InterfaceImpl2Stream int

func (i InterfaceImpl2Stream) Print() {
	fmt.Println("impl2")
}
