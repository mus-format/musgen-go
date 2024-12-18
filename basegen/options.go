package basegen

import (
	"fmt"
	"reflect"
)

type NumEncoding int

const (
	UndefinedNumEncoding NumEncoding = iota
	Varint
	VarintPositive
	Raw
)

func (e NumEncoding) Package() (pkg string) {
	switch e {
	case UndefinedNumEncoding, Varint, VarintPositive:
		return "varint"
	case Raw:
		return "raw"
	default:
		panic(fmt.Errorf("undefined %d NumEncoding", e))
	}
}

func (e NumEncoding) FuncName(fnType FnType) (name string) {
	if e == VarintPositive {
		return string(fnType) + "Positive"
	}
	return string(fnType)
}

type PtrType int

func (t PtrType) PackageName() string {
	switch t {
	case UndefinedPtrType, Ordinary:
		return "ord"
	case PM:
		return "pm"
	default:
		panic(fmt.Errorf("unexpected %v encoding", int(t)))
	}
}

const (
	UndefinedPtrType = iota
	Ordinary
	PM
)

type Options struct {
	Prefix       string
	Ignore       bool
	Encoding     NumEncoding
	Validator    string
	LenEncoding  NumEncoding
	LenValidator string
	Key          *Options
	Elem         *Options
	Oneof        []reflect.Type
}
