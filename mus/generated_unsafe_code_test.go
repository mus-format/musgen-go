package musgen

import (
	"testing"

	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func TestGeneratedUnsafeCode(t *testing.T) {

	t.Run("Test struct/alias/dts/interface/ptr serializability", func(t *testing.T) {

		t.Run("ComplexStruct should be serializable", func(t *testing.T) {
			testSerializability(makeComplexStructUnsafe(),
				pkg1.ComplexStructUnsafeMUS,
				nil,
				t)
		})

	})

}

func makeComplexStructUnsafe() pkg1.ComplexStructUnsafe {
	str := "some"
	return pkg1.ComplexStructUnsafe{
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

		Alias: pkg1.MySliceUnsafe{1, 2, 3, 4},
		AnotherPkgStruct: pkg2.StructUnsafe{
			Float32: 3.0,
			Float64: 55.0,
			Bool:    false,
			Byte:    100,
		},

		Interface: pkg1.InterfaceImpl1Unsafe{},

		ByteSlice:   []byte{1, 2, 3, 4},
		StructSlice: []pkg1.SimpleStructUnsafe{{Int: 10}, {Int: 20}},

		Array: [3]int{1, 2, 3},

		PtrArray:  &[3]int{1, 1, 1},
		PtrString: &str,
		PtrStruct: &pkg1.SimpleStructUnsafe{
			Int: 100,
		},
		NilPtr: nil,

		Map: map[float32]map[pkg1.MyIntUnsafe][]pkg1.SimpleStructUnsafe{
			40.8: {
				pkg1.MyIntUnsafe(11): {
					{Int: 30},
				},
			},
		},
	}
}
