package musgen

import (
	"fmt"
	"strings"

	"github.com/mus-format/musgen-go/data"
)

func NewOptionFns(typeBuilder TypeDataBuilder) OptionFns {
	return OptionFns{typeBuilder}
}

type OptionFns struct {
	typeBuilder TypeDataBuilder
}

func (f OptionFns) StringOps(anonData data.AnonData) string {
	b := strings.Builder{}
	if anonData.LenSer != "" && anonData.LenSer != "nil" {
		b.WriteString(fmt.Sprintf("strops.WithLenSer(%v)", anonData.LenSer))
	}
	if anonData.LenVl != "" && anonData.LenVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("strops.WithLenValidator(%v)", anonData.LenVl))
	}
	return b.String()
}

func (f OptionFns) ArrayOps(anonData data.AnonData) string {
	var (
		b        = strings.Builder{}
		elemType = f.typeBuilder.ToRelName(anonData.ElemType)
	)
	if anonData.LenSer != "" && anonData.LenSer != "nil" {
		b.WriteString(fmt.Sprintf("arrops.WithLenSer[%v](%v)", elemType,
			anonData.LenSer))
	}
	if anonData.ElemVl != "" && anonData.ElemVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("arrops.WithElemValidator[%v](%v)", elemType,
			anonData.ElemVl))
	}
	return b.String()
}

func (f OptionFns) ByteSliceOps(anonData data.AnonData) string {
	b := strings.Builder{}
	if anonData.LenSer != "" && anonData.LenSer != "nil" {
		b.WriteString(fmt.Sprintf("bslops.WithLenSer(%v)", anonData.LenSer))
	}
	if anonData.LenVl != "" && anonData.LenVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("bslops.WithLenValidator(%v)", anonData.LenVl))
	}
	return b.String()
}

func (f OptionFns) SliceOps(anonData data.AnonData) string {
	var (
		b        = strings.Builder{}
		elemType = f.typeBuilder.ToRelName(anonData.ElemType)
	)
	if anonData.LenSer != "" && anonData.LenSer != "nil" {
		b.WriteString(fmt.Sprintf("slops.WithLenSer[%v](%v)", elemType,
			anonData.LenSer))
	}
	if anonData.LenVl != "" && anonData.LenVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("slops.WithLenValidator[%v](%v)", elemType,
			anonData.LenVl))
	}
	if anonData.ElemVl != "" && anonData.ElemVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("slops.WithElemValidator[%v](%v)", elemType,
			anonData.ElemVl))
	}
	return b.String()
}

func (f OptionFns) MapOps(anonData data.AnonData) string {
	var (
		b        = strings.Builder{}
		keyType  = f.typeBuilder.ToRelName(anonData.KeyType)
		elemType = f.typeBuilder.ToRelName(anonData.ElemType)
	)
	if anonData.LenSer != "" && anonData.LenSer != "nil" {
		b.WriteString(fmt.Sprintf("mapops.WithLenSer[%v, %v](%v)", keyType,
			elemType, anonData.LenSer))
	}
	if anonData.LenVl != "" && anonData.LenVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("mapops.WithLenValidator[%v, %v](%v)", keyType,
			elemType, anonData.LenVl))
	}
	if anonData.KeyVl != "" && anonData.KeyVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("mapops.WithKeyValidator[%v, %v](%v)", keyType,
			elemType, anonData.KeyVl))
	}
	if anonData.ElemVl != "" && anonData.ElemVl != "nil" {
		if b.Len() > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("mapops.WithValueValidator[%v, %v](%v)", keyType,
			elemType, anonData.ElemVl))
	}
	return b.String()
}

func (f OptionFns) RegisterItself(m map[string]any) {
	m["StringOps"] = func(anonData data.AnonData) string {
		return f.StringOps(anonData)
	}
	m["ArrayOps"] = func(anonData data.AnonData) string {
		return f.ArrayOps(anonData)
	}
	m["ByteSliceOps"] = func(anonData data.AnonData) string {
		return f.ByteSliceOps(anonData)
	}
	m["SliceOps"] = func(anonData data.AnonData) string {
		return f.SliceOps(anonData)
	}
	m["MapOps"] = func(anonData data.AnonData) string {
		return f.MapOps(anonData)
	}
}
