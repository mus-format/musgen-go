// Code generated by musgen-go. DO NOT EDIT.\n\npackage

package pkg1

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/unsafe"
)

func MarshalUnsafeStreamBoolAliasMUS(v BoolAlias, w muss.Writer) (n int, err error) {
	return unsafe.MarshalBool(bool(v), w)
}

func UnmarshalUnsafeStreamBoolAliasMUS(r muss.Reader) (v BoolAlias, n int, err error) {
	va, n, err := unsafe.UnmarshalBool(r)
	if err != nil {
		return
	}
	v = BoolAlias(va)
	return
}

func SizeUnsafeStreamBoolAliasMUS(v BoolAlias) (size int) {
	return unsafe.SizeBool(bool(v))
}

func SkipUnsafeStreamBoolAliasMUS(r muss.Reader) (n int, err error) {
	return unsafe.SkipBool(r)
}
