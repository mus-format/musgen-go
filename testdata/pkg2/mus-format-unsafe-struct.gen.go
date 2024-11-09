// Code generated by musgen-go. DO NOT EDIT.

package pkg2

import "github.com/mus-format/mus-go/unsafe"

func MarshalUnsafeStructMUS(v Struct, bs []byte) (n int) {
	n = unsafe.MarshalFloat32(v.Float32, bs[n:])
	n += unsafe.MarshalFloat64(v.Float64, bs[n:])
	n += unsafe.MarshalUint8(v.Byte, bs[n:])
	return n + unsafe.MarshalBool(v.Bool, bs[n:])
}

func UnmarshalUnsafeStructMUS(bs []byte) (v Struct, n int, err error) {
	v.Float32, n, err = unsafe.UnmarshalFloat32(bs[n:])
	if err != nil {
		return
	}
	var n1 int
	v.Float64, n1, err = unsafe.UnmarshalFloat64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Byte, n1, err = unsafe.UnmarshalUint8(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Bool, n1, err = unsafe.UnmarshalBool(bs[n:])
	n += n1
	return
}

func SizeUnsafeStructMUS(v Struct) (size int) {
	size = unsafe.SizeFloat32(v.Float32)
	size += unsafe.SizeFloat64(v.Float64)
	size += unsafe.SizeUint8(v.Byte)
	return size + unsafe.SizeBool(v.Bool)
}

func SkipUnsafeStructMUS(bs []byte) (n int, err error) {
	n, err = unsafe.SkipFloat32(bs[n:])
	if err != nil {
		return
	}
	var n1 int
	n1, err = unsafe.SkipFloat64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipUint8(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipBool(bs[n:])
	n += n1
	return
}
