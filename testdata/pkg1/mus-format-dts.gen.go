// Code generated by musgen-go. DO NOT EDIT.

package pkg1

import (
	dts "github.com/mus-format/mus-dts-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

func MarshalInterfaceImpl1MUS(v InterfaceImpl1, bs []byte) (n int) {
	return
}

func UnmarshalInterfaceImpl1MUS(bs []byte) (v InterfaceImpl1, n int, err error) {
	return
}

func SizeInterfaceImpl1MUS(v InterfaceImpl1) (size int) {
	return
}

func SkipInterfaceImpl1MUS(bs []byte) (n int, err error) {
	return
}

var InterfaceImpl1DTS = dts.New[InterfaceImpl1](InterfaceImpl1DTM,
	mus.MarshallerFn[InterfaceImpl1](MarshalInterfaceImpl1MUS),
	mus.UnmarshallerFn[InterfaceImpl1](UnmarshalInterfaceImpl1MUS),
	mus.SizerFn[InterfaceImpl1](SizeInterfaceImpl1MUS),
	mus.SkipperFn(SkipInterfaceImpl1MUS))

func MarshalInterfaceImpl2MUS(v InterfaceImpl2, bs []byte) (n int) {
	return
}

func UnmarshalInterfaceImpl2MUS(bs []byte) (v InterfaceImpl2, n int, err error) {
	return
}

func SizeInterfaceImpl2MUS(v InterfaceImpl2) (size int) {
	return
}

func SkipInterfaceImpl2MUS(bs []byte) (n int, err error) {
	return
}

var InterfaceImpl2DTS = dts.New[InterfaceImpl2](InterfaceImpl2DTM,
	mus.MarshallerFn[InterfaceImpl2](MarshalInterfaceImpl2MUS),
	mus.UnmarshallerFn[InterfaceImpl2](UnmarshalInterfaceImpl2MUS),
	mus.SizerFn[InterfaceImpl2](SizeInterfaceImpl2MUS),
	mus.SkipperFn(SkipInterfaceImpl2MUS))

func MarshalStructMUS(v Struct, bs []byte) (n int) {
	return varint.MarshalInt(v.Int, bs[n:])
}

func UnmarshalStructMUS(bs []byte) (v Struct, n int, err error) {
	v.Int, n, err = varint.UnmarshalInt(bs[n:])
	return
}

func SizeStructMUS(v Struct) (size int) {
	return varint.SizeInt(v.Int)
}

func SkipStructMUS(bs []byte) (n int, err error) {
	return varint.SkipInt(bs[n:])
}

var StructDTS = dts.New[Struct](StructDTM,
	mus.MarshallerFn[Struct](MarshalStructMUS),
	mus.UnmarshallerFn[Struct](UnmarshalStructMUS),
	mus.SizerFn[Struct](SizeStructMUS),
	mus.SkipperFn(SkipStructMUS))
