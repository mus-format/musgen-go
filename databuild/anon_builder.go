package databuild

import (
	"github.com/mus-format/musgen-go/data"
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/scanner"
	"github.com/mus-format/musgen-go/typename"
)

func NewAnonDataBuilder(typeBuilder TypeDataBuilder,
	gops genops.Options) AnonDataBuilder {
	return AnonDataBuilder{typeBuilder, gops}
}

type AnonDataBuilder struct {
	typeBuilder TypeDataBuilder
	gops        genops.Options
}

func (b AnonDataBuilder) Build(name typename.FullName, tops *typeops.Options) (
	d data.AnonData, ok bool, err error) {
	op := NewMakeAnonOp(b.typeBuilder, b.gops)

	if err = scanner.Scan(name, op, tops); err != nil {
		if err != ErrStop {
			return
		}
		err = nil
	}
	d, ok = op.AnonData()
	return
}

func (b AnonDataBuilder) Collect(name typename.FullName,
	m map[data.AnonSerName]data.AnonData, tops *typeops.Options) (err error) {
	op := NewCollectAnonOp(m, b.typeBuilder, b.gops)
	return scanner.Scan(name, op, tops)
}
