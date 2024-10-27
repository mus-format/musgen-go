package pkg2

const StructDTM = 3

type Struct struct {
	Float32 float32
	Float64 float64
	Byte    byte
	Bool    bool
}

type AnotherStruct struct {
	Struct1 Struct
	Struct2 Struct
}
