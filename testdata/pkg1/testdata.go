package pkg1

import (
	"fmt"
	"time"

	"github.com/mus-format/musgen-go/testdata/pkg2"
)

const (
	InterfaceImpl1DTM = 1
	InterfaceImpl2DTM = 2
	SimpleStructDTM   = 3
	MyIntDTM          = 4
)

type TimeTypedef interface {
	Time() time.Time
}

type MyBool bool
type MyByte byte
type MyFloat32 float32

type MyInt int
type RawMyInt int
type VarintPositiveMyInt int
type ValidMyInt int
type AllMyInt int

type MyString string
type LenEncodingMyString string
type LenValidMyString string
type ValidMyString string
type AllMyString string

type MyTime time.Time

func (t MyTime) Time() time.Time {
	return time.Time(t)
}

type MyTimeSec time.Time

func (t MyTimeSec) Time() time.Time {
	return time.Time(t)
}

type MyTimeMilli time.Time

func (t MyTimeMilli) Time() time.Time {
	return time.Time(t)
}

type MyTimeMicro time.Time

func (t MyTimeMicro) Time() time.Time {
	return time.Time(t)
}

type MyTimeNano time.Time

func (t MyTimeNano) Time() time.Time {
	return time.Time(t)
}

type MyTimeSecUTC time.Time

func (t MyTimeSecUTC) Time() time.Time {
	return time.Time(t)
}

type MyTimeMilliUTC time.Time

func (t MyTimeMilliUTC) Time() time.Time {
	return time.Time(t)
}

type MyTimeMicroUTC time.Time

func (t MyTimeMicroUTC) Time() time.Time {
	return time.Time(t)
}

type MyTimeNanoUTC time.Time

func (t MyTimeNanoUTC) Time() time.Time {
	return time.Time(t)
}

type MyArray [3]int
type LenEncodingMyArray [3]int
type ElemEncodingMyArray [3]int
type ElemValidMyArray [3]int
type ValidMyArray [3]int
type AllMyArray [3]int

type MyByteSlice []byte
type LenEncodingMyByteSlice []byte
type LenValidMyByteSlice []byte
type ValidMyByteSlice []byte
type AllMyByteSlice []byte

type MySlice []int
type LenEncodingMySlice []int
type LenValidMySlice []int
type ElemEncodingMySlice []int
type ElemValidMySlice []int
type ValidMySlice []int
type AllMySlice []int

type MyMap map[int]int
type LenEncodingMyMap map[int]int
type LenValidMyMap map[int]int
type KeyEncodingMyMap map[int]int
type KeyValidMyMap map[int]int
type ElemEncodingMyMap map[int]int
type ElemValidMyMap map[int]int
type ValidMyMap map[int]int
type AllMyMap map[int]int

type MyIntPtr *int
type ElemNumEncodingMyIntPtr *int

type ValidMyIntPtr *int

type SimpleStructMyIntPtr *SimpleStruct

type InterfaceDoubleMyIntPtr **Interface

type StructAlias SimpleStruct

type InterfaceAlias Interface

type AnotherStruct SimpleStruct
type AnotherInterface Interface

type UnsuportedChanType chan int

// -----------------------------------------------------------------------------

type UnsupportedFuncFieldStruct struct {
	Fn func()
}

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

	Alias            MySlice
	AnotherPkgStruct pkg2.Struct
	Interface        Interface

	ByteSlice   []byte
	StructSlice []SimpleStruct

	Array [3]int

	PtrString *string
	PtrStruct *SimpleStruct
	NilPtr    *string
	PtrArray  *[3]int

	Map map[float32]map[MyInt][]SimpleStruct
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
