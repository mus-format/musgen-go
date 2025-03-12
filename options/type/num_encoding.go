package typeops

import (
	"fmt"
)

const (
	UndefinedNumEncoding NumEncoding = iota
	Varint
	VarintPositive
	Raw
)

type NumEncoding int

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

func (e NumEncoding) LenSer() string {
	switch e {
	case UndefinedNumEncoding, VarintPositive:
		return "varint.PositiveInt"
	case Varint:
		return "varint.Int"
	case Raw:
		return "raw.Int"
	default:
		panic(fmt.Errorf("undefined %d NumEncoding", e))
	}
}
