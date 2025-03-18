package musgen

import (
	"testing"
	"time"

	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func TestGeneratedStramUnsafeCode(t *testing.T) {

	t.Run("Test struct/alias/dts/interface/ptr serializability", func(t *testing.T) {

		t.Run("ComplexStruct should be serializable", func(t *testing.T) {
			testStreamSerializability(makeComplexStructStreamUnsafe(),
				pkg1.ComplexStructStreamUnsafeMUS,
				t)
		})

		t.Run("TimeStruct should be serializable", func(t *testing.T) {
			v := pkg2.TimeStructStreamUnsafe{
				Float32: 10.3,
				Time:    time.Unix(time.Now().Unix(), 0),
				String:  "abs",
			}
			testStreamSerializability(v, pkg2.TimeStructStreamUnsafeMUS, t)
		})

	})

}

func makeComplexStructStreamUnsafe() pkg1.ComplexStructStreamUnsafe {
	str := "some"
	return pkg1.ComplexStructStreamUnsafe{
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

		Alias: pkg1.MySliceStreamUnsafe{1, 2, 3, 4},
		AnotherPkgStruct: pkg2.StructStreamUnsafe{
			Float32: 3.0,
			Float64: 55.0,
			Bool:    false,
			Byte:    100,
		},

		Interface: pkg1.InterfaceImpl1StreamUnsafe{},

		ByteSlice:   []byte{1, 2, 3, 4},
		StructSlice: []pkg1.SimpleStructStreamUnsafe{{Int: 10}, {Int: 20}},

		Array: [3]int{1, 2, 3},

		PtrArray:  &[3]int{1, 1, 1},
		PtrString: &str,
		PtrStruct: &pkg1.SimpleStructStreamUnsafe{
			Int: 100,
		},
		NilPtr: nil,

		Map: map[float32]map[pkg1.MyIntStreamUnsafe][]pkg1.SimpleStructStreamUnsafe{
			40.8: {
				pkg1.MyIntStreamUnsafe(11): {
					{Int: 30},
				},
			},
		},
	}
}
