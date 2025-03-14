// Code generated by musgen-go. DO NOT EDIT.

package pkg2

import "github.com/mus-format/mus-go/unsafe"

var StructUnsafeMUS = structUnsafeMUS{}

type structUnsafeMUS struct{}

func (s structUnsafeMUS) Marshal(v StructUnsafe, bs []byte) (n int) {
	n = unsafe.Float32.Marshal(v.Float32, bs)
	n += unsafe.Float64.Marshal(v.Float64, bs[n:])
	n += unsafe.Uint8.Marshal(v.Byte, bs[n:])
	return n + unsafe.Bool.Marshal(v.Bool, bs[n:])
}

func (s structUnsafeMUS) Unmarshal(bs []byte) (v StructUnsafe, n int, err error) {
	v.Float32, n, err = unsafe.Float32.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	v.Float64, n1, err = unsafe.Float64.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Byte, n1, err = unsafe.Uint8.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Bool, n1, err = unsafe.Bool.Unmarshal(bs[n:])
	n += n1
	return
}

func (s structUnsafeMUS) Size(v StructUnsafe) (size int) {
	size = unsafe.Float32.Size(v.Float32)
	size += unsafe.Float64.Size(v.Float64)
	size += unsafe.Uint8.Size(v.Byte)
	return size + unsafe.Bool.Size(v.Bool)
}

func (s structUnsafeMUS) Skip(bs []byte) (n int, err error) {
	n, err = unsafe.Float32.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = unsafe.Float64.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Uint8.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Bool.Skip(bs[n:])
	n += n1
	return
}
