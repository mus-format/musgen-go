package testdata

import (
	"bytes"
	"testing"

	"github.com/mus-format/mus-go"
	muss "github.com/mus-format/mus-stream-go"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestComplexStructSer(v ComplexStruct, ser mus.Serializer[ComplexStruct],
	t *testing.T) {
	t.Helper()
	bs := make([]byte, ser.Size(v))
	ser.Marshal(v, bs)
	av, _, err := ser.Unmarshal(bs)
	asserterror.EqualError(err, nil, t)
	asserterror.Equal(EqualComplexStruct(v, av), true, t)
}

func TestComplexStructStreamSer(v ComplexStruct,
	ser muss.Serializer[ComplexStruct], t *testing.T) {
	t.Helper()
	var (
		bs  = make([]byte, 0, ser.Size(v))
		buf = bytes.NewBuffer(bs)
	)
	ser.Marshal(v, buf)
	av, _, err := ser.Unmarshal(buf)
	asserterror.EqualError(err, nil, t)
	asserterror.Equal(EqualComplexStruct(v, av), true, t)
}
