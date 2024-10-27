package mus

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/musgen-go/testdata/pkg1"
)

func TestGeneratedStreamCode(t *testing.T) {

	t.Run("Test struct/alias/dts/interface/ptr serializability", func(t *testing.T) {

		t.Run("ComplexStruct should be serializable", func(t *testing.T) {
			testStreamSerializability(makeCompplexStruct(),
				pkg1.MarshalStreamComplexStructMUS,
				pkg1.UnmarshalStreamComplexStructMUS,
				pkg1.SizeStreamComplexStructMUS,
				pkg1.SkipStreamComplexStructMUS,
				t)
		})

	})

}

func testStreamSerializability[T any](v T, m muss.MarshallerFn[T],
	u muss.UnmarshallerFn[T],
	s muss.SizerFn[T],
	sk muss.SkipperFn,
	t *testing.T,
) {
	buf := bytes.NewBuffer(make([]byte, 0, s(v)))
	m(v, buf)
	v1, n1, err := u(buf)
	if err != nil {
		t.Fatal(err)
	}
	buf = bytes.NewBuffer(make([]byte, 0, s(v)))
	m(v, buf)
	n2, err := sk(buf)
	if err != nil {
		t.Fatal(err)
	}
	opt := cmp.FilterValues(func(x1, x2 interface{}) bool {
		var (
			t1        = reflect.TypeOf(x1)
			t2        = reflect.TypeOf(x2)
			intrImpl1 = reflect.TypeFor[pkg1.InterfaceImpl1]()
		)
		if t1 == intrImpl1 && t2 == intrImpl1 {
			return true
		}
		if (t1.Kind() == reflect.Slice && t2.Kind() == reflect.Slice) ||
			(t1.Kind() == reflect.Map && t2.Kind() == reflect.Map) {
			return true
		}
		return false
	}, cmp.Comparer(func(i1, i2 interface{}) bool {
		_, ok1 := i1.(pkg1.InterfaceImpl1)
		_, ok2 := i2.(pkg1.InterfaceImpl1)
		if ok1 && ok2 {
			return true
		}
		return fmt.Sprint(i1) == fmt.Sprint(i2)
	}))
	if !cmp.Equal(v, v1, opt) {
		t.Errorf("unexpected value %v", v1)
	}
	if n1 != n2 {
		t.Errorf("unexpected value %v", n2)
	}
}
