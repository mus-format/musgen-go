package data

import (
	"fmt"

	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
)

const (
	UndefinedAnonSerKind AnonSerKind = iota
	AnonString
	AnonArray
	AnonByteSlice
	AnonSlice
	AnonMap
	AnonPtr
)

const (
	AnonStringPrefix    = "string"
	AnonArrayPrefix     = "array"
	AnonByteSlicePrefix = "byteSlice"
	AnonSlicePrefix     = "slice"
	AnonMapPrefix       = "map"
	AnonPtrPrefix       = "ptr"
)

type AnonSerName string

type AnonData struct {
	AnonSerName AnonSerName
	Kind        AnonSerKind

	ArrType   typename.FullName
	ArrLength string

	LenSer string
	LenVl  string

	KeyType typename.FullName
	KeyVl   string

	ElemType typename.FullName
	ElemVl   string

	Tops *typeops.Options
}

type AnonSerKind int

func (k AnonSerKind) String() string {
	switch k {
	case AnonString:
		return AnonStringPrefix
	case AnonArray:
		return AnonArrayPrefix
	case AnonByteSlice:
		return AnonByteSlicePrefix
	case AnonSlice:
		return AnonSlicePrefix
	case AnonMap:
		return AnonMapPrefix
	case AnonPtr:
		return AnonPtrPrefix
	default:
		panic(fmt.Sprintf("unexpected %v AnonSerKind", int(k)))
	}
}
