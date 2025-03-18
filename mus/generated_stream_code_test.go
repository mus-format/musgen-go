package musgen

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func TestGeneratedStreamCode(t *testing.T) {

	t.Run("Test struct/alias/dts/interface/ptr serializability", func(t *testing.T) {

		t.Run("ComplexStruct should be serializable", func(t *testing.T) {
			testStreamSerializability(makeCompplexStructStream(),
				pkg1.ComplexStructStreamMUS,
				t)
		})

		t.Run("TimeStruct should be serializable", func(t *testing.T) {
			v := pkg2.TimeStructStream{
				Float32: 10.3,
				Time:    time.Unix(time.Now().Unix(), 0),
				String:  "abs",
			}
			testStreamSerializability(v, pkg2.TimeStructStreamMUS, t)
		})

	})

}

func testStreamSerializability[T any](v T, ser muss.Serializer[T],
	t *testing.T,
) {
	buf := bytes.NewBuffer(make([]byte, 0, ser.Size(v)))
	ser.Marshal(v, buf)
	v1, n1, err := ser.Unmarshal(buf)
	if err != nil {
		t.Fatal(err)
	}
	buf = bytes.NewBuffer(make([]byte, 0, ser.Size(v)))
	ser.Marshal(v, buf)
	n2, err := ser.Skip(buf)
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

func makeCompplexStructStream() pkg1.ComplexStructStream {
	str := "some"
	return pkg1.ComplexStructStream{
		Bool: true,
		Byte: 111,

		Int8:  8,
		Int16: 16,
		Int32: 32,
		Int64: 64,

		Uint8:  8,
		Uint16: 16,
		Uint32: 32,
		Uint64: 64,

		Float32: 32.5,
		Float64: 64.5,

		String: "some",

		Alias: pkg1.MySliceStream{1, 2, 3, 4},
		AnotherPkgStruct: pkg2.StructStream{
			Float32: 3.0,
			Float64: 55.0,
			Bool:    false,
			Byte:    100,
		},

		Interface: pkg1.InterfaceImpl1Stream{},

		ByteSlice:   []byte{1, 2, 3, 4},
		StructSlice: []pkg1.SimpleStructStream{{Int: 10}, {Int: 20}},

		Array: [3]int{1, 2, 3},

		PtrArray:  &[3]int{1, 1, 1},
		PtrString: &str,
		PtrStruct: &pkg1.SimpleStructStream{
			Int: 100,
		},
		NilPtr: nil,

		Map: map[float32]map[pkg1.MyIntStream][]pkg1.SimpleStructStream{
			40.8: {
				pkg1.MyIntStream(11): {
					{Int: 30},
				},
			},
		},
	}
}
