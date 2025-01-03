// Code generated by musgen-go. DO NOT EDIT.

package pkg2

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/varint"
)

func MarshalStreamStructMUS(v Struct, w muss.Writer) (n int, err error) {
	n, err = varint.MarshalFloat32(v.Float32, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = varint.MarshalFloat64(v.Float64, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalUint8(v.Byte, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.MarshalBool(v.Bool, w)
	n += n1
	return
}

func UnmarshalStreamStructMUS(r muss.Reader) (v Struct, n int, err error) {
	v.Float32, n, err = varint.UnmarshalFloat32(r)
	if err != nil {
		return
	}
	var n1 int
	v.Float64, n1, err = varint.UnmarshalFloat64(r)
	n += n1
	if err != nil {
		return
	}
	v.Byte, n1, err = varint.UnmarshalUint8(r)
	n += n1
	if err != nil {
		return
	}
	v.Bool, n1, err = ord.UnmarshalBool(r)
	n += n1
	return
}

func SizeStreamStructMUS(v Struct) (size int) {
	size = varint.SizeFloat32(v.Float32)
	size += varint.SizeFloat64(v.Float64)
	size += varint.SizeUint8(v.Byte)
	return size + ord.SizeBool(v.Bool)
}

func SkipStreamStructMUS(r muss.Reader) (n int, err error) {
	n, err = varint.SkipFloat32(r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = varint.SkipFloat64(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipUint8(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipBool(r)
	n += n1
	return
}
