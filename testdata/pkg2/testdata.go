package pkg2

import "time"

const StructDTM = 3

type IgnoreFieldStruct Struct

type ValidFieldStruct Struct

type ElemStruct Struct

type Struct struct {
	Float32 float32
	Float64 float64
	Byte    byte
	Bool    bool
	Slice   [][]int
}

type TimeStructStream TimeStruct

type TimeStructUnsafe TimeStruct

type TimeStructStreamUnsafe TimeStruct

type TimeStructMilli TimeStruct

type TimeStruct struct {
	Float32 float32
	Time    time.Time
	String  string
}
