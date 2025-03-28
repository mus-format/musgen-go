// Code generated by musgen-go. DO NOT EDIT.

package pkg1

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/unsafe"
)

var StructStreamNotUnsafeMUS = structStreamNotUnsafeMUS{}

type structStreamNotUnsafeMUS struct{}

func (s structStreamNotUnsafeMUS) Marshal(v StructStreamNotUnsafe, w muss.Writer) (n int, err error) {
	n, err = ord.String.Marshal(v.String, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = unsafe.Int.Marshal(v.Int, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.TimeUnix.Marshal(v.Time, w)
	n += n1
	return
}

func (s structStreamNotUnsafeMUS) Unmarshal(r muss.Reader) (v StructStreamNotUnsafe, n int, err error) {
	v.String, n, err = ord.String.Unmarshal(r)
	if err != nil {
		return
	}
	var n1 int
	v.Int, n1, err = unsafe.Int.Unmarshal(r)
	n += n1
	if err != nil {
		return
	}
	v.Time, n1, err = unsafe.TimeUnix.Unmarshal(r)
	n += n1
	return
}

func (s structStreamNotUnsafeMUS) Size(v StructStreamNotUnsafe) (size int) {
	size = ord.String.Size(v.String)
	size += unsafe.Int.Size(v.Int)
	return size + unsafe.TimeUnix.Size(v.Time)
}

func (s structStreamNotUnsafeMUS) Skip(r muss.Reader) (n int, err error) {
	n, err = ord.String.Skip(r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = unsafe.Int.Skip(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.TimeUnix.Skip(r)
	n += n1
	return
}
