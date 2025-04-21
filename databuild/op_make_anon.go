package databuild

import (
	"errors"

	"github.com/mus-format/musgen-go/data"
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/scanner"
	"github.com/mus-format/musgen-go/typename"
)

var ErrStop = errors.New("stop")

func NewMakeAnonOp(typeBuilder TypeDataBuilder, gops genops.Options) *MakeAnonDataOp {
	m := map[data.AnonSerName]data.AnonData{}
	return &MakeAnonDataOp{CollectAnonOp: NewCollectAnonOp(m, typeBuilder, gops)}
}

type MakeAnonDataOp struct {
	*CollectAnonOp
	d  data.AnonData
	ok bool
}

func (o *MakeAnonDataOp) ProcessType(t scanner.Type[typename.FullName],
	tops *typeops.Options) (err error) {
	if err = o.CollectAnonOp.ProcessType(t, tops); err != nil {
		return
	}
	if d := o.CollectAnonOp.FirstAnonData(); d != nil {
		o.d = *d
		o.ok = true
	}
	return ErrStop
}

func (o *MakeAnonDataOp) AnonData() (d data.AnonData, ok bool) {
	return o.d, o.ok
}
