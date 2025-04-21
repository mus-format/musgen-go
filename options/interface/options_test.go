package introps

import (
	"reflect"
	"testing"

	prim_testdata "github.com/mus-format/musgen-go/testdata/primitive"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestOptions(t *testing.T) {
	var (
		o         = New()
		wantImpls = []reflect.Type{reflect.TypeFor[prim_testdata.MyInt](),
			reflect.TypeFor[prim_testdata.MyInt]()}
		wantMarshaller = true
	)
	Apply([]SetOption{
		WithImpl(reflect.TypeFor[prim_testdata.MyInt]()),
		WithImpl(reflect.TypeFor[prim_testdata.MyInt]()),
		WithMarshaller(),
	}, &o)
	asserterror.EqualDeep(o.Impls, wantImpls, t)
	asserterror.Equal(o.Marshaller, wantMarshaller, t)
}
