package mus

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/varint"
	"github.com/mus-format/musgen-go/testdata"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func TestGeneratedCode(t *testing.T) {

	t.Run("Test alias serializability and metadata impact", func(t *testing.T) {

		t.Run("IntAlias should be serializable", func(t *testing.T) {
			v := pkg1.IntAlias(5)
			testSerializability(v, pkg1.MarshalIntAliasMUS,
				pkg1.UnmarshalIntAliasMUS,
				pkg1.SizeIntAliasMUS,
				pkg1.SkipIntAliasMUS,
				t)
		})

		t.Run("We should be able to set Raw encoding for IntAlias", func(t *testing.T) {
			v := pkg1.IntAlias(5)
			testSerializability(v, pkg1.MarshalRawIntAliasMUS,
				func(bs []byte) (v pkg1.IntAlias, n int, err error) {
					num, n, err := raw.UnmarshalInt(bs)
					v = pkg1.IntAlias(num)
					return
				},
				pkg1.SizeRawIntAliasMUS,
				pkg1.SkipRawIntAliasMUS,
				t)
		})

		t.Run("We should be able to set VarintPositive encoding for IntAlias", func(t *testing.T) {
			v := pkg1.IntAlias(5)
			testSerializability(v, pkg1.MarshalVarintPositiveIntAliasMUS,
				func(bs []byte) (v pkg1.IntAlias, n int, err error) {
					num, n, err := varint.UnmarshalPositiveInt(bs)
					v = pkg1.IntAlias(num)
					return
				},
				pkg1.SizeVarintPositiveIntAliasMUS,
				pkg1.SkipVarintPositiveIntAliasMUS,
				t)
		})

		t.Run("We should be able to set Validator for IntAlias", func(t *testing.T) {
			v := pkg1.IntAlias(0)
			testValidation(v, pkg1.MarshalValidIntAliasMUS,
				pkg1.UnmarshalValidIntAliasMUS,
				pkg1.SizeValidIntAliasMUS,
				testdata.ErrZeroValue,
				[]int{1},
				t)
		})

		t.Run("We should be able to set Raw LenEncoding for StringAlias", func(t *testing.T) {
			v := pkg1.StringAlias("some")
			testLenEncoding(v, pkg1.MarshalLenEncodingStringAliasMUS,
				raw.UnmarshalInt,
				pkg1.SizeLenEncodingStringAliasMUS,
				len(v),
				t)
		})

		t.Run("We should be able to set LenValidator for StringAlias", func(t *testing.T) {
			v := pkg1.StringAlias("some")
			testValidation(v, pkg1.MarshalLenValidatorStringAliasMUS,
				pkg1.UnmarshalLenValidatorStringAliasMUS,
				pkg1.SizeLenValidatorStringAliasMUS,
				testdata.ErrTooLong,
				[]int{1},
				t)
		})

		t.Run("We should be able to set Validator for StringAlias", func(t *testing.T) {
			v := pkg1.StringAlias("")
			testValidation(v, pkg1.MarshalValidStringAliasMUS,
				pkg1.UnmarshalValidStringAliasMUS,
				pkg1.SizeValidStringAliasMUS,
				testdata.ErrZeroValue,
				[]int{1},
				t)
		})

		t.Run("SliceAlias should be serializable", func(t *testing.T) {
			v := pkg1.SliceAlias([]int{1, 2, 3})
			testSerializability(v, pkg1.MarshalSliceAliasMUS,
				pkg1.UnmarshalSliceAliasMUS,
				pkg1.SizeSliceAliasMUS,
				pkg1.SkipSliceAliasMUS,
				t)
		})

		t.Run("We should be able to set LenEncoding for SliceAlias", func(t *testing.T) {
			v := pkg1.SliceAlias([]int{1, 2, 3})
			testLenEncoding(v, pkg1.MarshalLenEncodingSliceAliasMUS,
				varint.UnmarshalPositiveInt,
				pkg1.SizeLenEncodingSliceAliasMUS,
				len(v),
				t)
		})

		t.Run("We should be able to set LenValidator for SliceAlias", func(t *testing.T) {
			testValidation(pkg1.SliceAlias([]int{1, 2, 3}),
				pkg1.MarshalLenValidatorSliceAliasMUS,
				pkg1.UnmarshalLenValidatorSliceAliasMUS,
				pkg1.SizeLenValidatorSliceAliasMUS,
				testdata.ErrTooLong,
				[]int{1},
				t,
			)
		})

		t.Run("We should be able to set Elem.Encoding for SliceAlias", func(t *testing.T) {
			testSerializability(pkg1.SliceAlias([]int{1, 2, 3}),
				pkg1.MarshalElemEncodingSliceAliasMUS,
				func(bs []byte) (sla pkg1.SliceAlias, n int, err error) {
					sl, n, err := ord.UnmarshalSlice[int](nil,
						mus.UnmarshallerFn[int](raw.UnmarshalInt), bs)
					sla = pkg1.SliceAlias(sl)
					return
				},
				pkg1.SizeElemEncodingSliceAliasMUS,
				pkg1.SkipElemEncodingSliceAliasMUS,
				t)
		})

		t.Run("We should be able to set Elem.Validator for SliceAlias", func(t *testing.T) {
			testValidation(
				pkg1.SliceAlias([]int{1, 0, 3}),
				pkg1.MarshalElemValidatorSliceAliasMUS,
				pkg1.UnmarshalElemValidatorSliceAliasMUS,
				pkg1.SizeElemValidatorSliceAliasMUS,
				testdata.ErrZeroValue,
				[]int{3},
				t)
		})

		t.Run("ArrayAlias should be serializable", func(t *testing.T) {
			v := pkg1.ArrayAlias([3]int{1, 2, 3})
			testSerializability(v, pkg1.MarshalArrayAliasMUS,
				pkg1.UnmarshalArrayAliasMUS,
				pkg1.SizeArrayAliasMUS,
				pkg1.SkipArrayAliasMUS,
				t)
		})

		t.Run("We should be able to set LenEncoding for ArrayAlias", func(t *testing.T) {
			v := pkg1.ArrayAlias([3]int{1, 2, 3})
			testLenEncoding(v, pkg1.MarshalLenEncodingArrayAliasMUS,
				varint.UnmarshalPositiveInt,
				pkg1.SizeLenEncodingArrayAliasMUS,
				len(v),
				t)
		})

		t.Run("We should be able to set LenValidator for ArrayAlias", func(t *testing.T) {
			testValidation(pkg1.ArrayAlias([3]int{1, 2, 3}),
				pkg1.MarshalLenValidatorArrayAliasMUS,
				pkg1.UnmarshalLenValidatorArrayAliasMUS,
				pkg1.SizeLenValidatorArrayAliasMUS,
				testdata.ErrTooLong,
				[]int{1},
				t,
			)
		})

		t.Run("We should be able to set Elem.Encoding for ArrayAlias", func(t *testing.T) {
			testSerializability(pkg1.ArrayAlias([3]int{1, 2, 3}),
				pkg1.MarshalElemEncodingArrayAliasMUS,
				func(bs []byte) (sla pkg1.ArrayAlias, n int, err error) {
					sl, n, err := ord.UnmarshalSlice[int](nil,
						mus.UnmarshallerFn[int](raw.UnmarshalInt), bs)
					sla = pkg1.ArrayAlias(sl)
					return
				},
				pkg1.SizeElemEncodingArrayAliasMUS,
				pkg1.SkipElemEncodingArrayAliasMUS,
				t)
		})

		t.Run("We should be able to set Elem.Validator for ArrayAlias", func(t *testing.T) {
			testValidation(
				pkg1.ArrayAlias([3]int{1, 0, 3}),
				pkg1.MarshalElemValidatorArrayAliasMUS,
				pkg1.UnmarshalElemValidatorArrayAliasMUS,
				pkg1.SizeElemValidatorArrayAliasMUS,
				testdata.ErrZeroValue,
				[]int{3},
				t)
		})

		t.Run("MapAlias should be serializable", func(t *testing.T) {
			v := pkg1.MapAlias(map[string]int{"some": 1, "another": 2})
			testSerializability(v, pkg1.MarshalMapAliasMUS,
				pkg1.UnmarshalMapAliasMUS,
				pkg1.SizeMapAliasMUS,
				pkg1.SkipMapAliasMUS,
				t)
		})

		t.Run("We should be able to set LenValidator for MapAlias", func(t *testing.T) {
			testValidation(
				pkg1.MapAlias(map[string]int{"some": 1, "another": 2}),
				pkg1.MarshalLenValidatorMapAliasMUS,
				pkg1.UnmarshalLenValidatorMapAliasMUS,
				pkg1.SizeLenValidatorMapAliasMUS,
				testdata.ErrTooLong,
				[]int{1},
				t)
		})

		t.Run("We should be able to set Key.Validator for MapAlias", func(t *testing.T) {
			testValidation(
				pkg1.MapAlias(map[string]int{"some": 1, "": 2}),
				pkg1.MarshalKeyValidatorMapAliasMUS,
				pkg1.UnmarshalKeyValidatorMapAliasMUS,
				pkg1.SizeKeyValidatorMapAliasMUS,
				testdata.ErrZeroValue,
				[]int{8, 2},
				t)
		})

		t.Run("We should be able to set Elem.Validator for MapAlias", func(t *testing.T) {
			testValidation(
				pkg1.MapAlias(map[string]int{"some": 0, "another": 2}),
				pkg1.MarshalElemValidatorMapAliasMUS,
				pkg1.UnmarshalElemValidatorMapAliasMUS,
				pkg1.SizeElemValidatorMapAliasMUS,
				testdata.ErrZeroValue,
				[]int{7, 16},
				t)
		})

	})

	t.Run("IntAlias should be serializable", func(t *testing.T) {
		v := pkg1.IntAlias(5)
		testSerializability(v, pkg1.MarshalIntAliasMUS,
			pkg1.UnmarshalIntAliasMUS,
			pkg1.SizeIntAliasMUS,
			pkg1.SkipIntAliasMUS,
			t)
	})

	t.Run("We should be able to set Raw encoding for IntAlias", func(t *testing.T) {
		v := pkg1.IntAlias(5)
		testSerializability(v, pkg1.MarshalRawIntAliasMUS,
			func(bs []byte) (v pkg1.IntAlias, n int, err error) {
				num, n, err := raw.UnmarshalInt(bs)
				v = pkg1.IntAlias(num)
				return
			},
			pkg1.SizeRawIntAliasMUS,
			pkg1.SkipRawIntAliasMUS,
			t)
	})

	t.Run("We should be able to set VarintPositive encoding for IntAlias", func(t *testing.T) {
		v := pkg1.IntAlias(5)
		testSerializability(v, pkg1.MarshalVarintPositiveIntAliasMUS,
			func(bs []byte) (v pkg1.IntAlias, n int, err error) {
				num, n, err := varint.UnmarshalPositiveInt(bs)
				v = pkg1.IntAlias(num)
				return
			},
			pkg1.SizeVarintPositiveIntAliasMUS,
			pkg1.SkipVarintPositiveIntAliasMUS,
			t)
	})

	t.Run("We should be able to set Validator for IntAlias", func(t *testing.T) {
		v := pkg1.IntAlias(0)
		testValidation(v, pkg1.MarshalValidIntAliasMUS,
			pkg1.UnmarshalValidIntAliasMUS,
			pkg1.SizeValidIntAliasMUS,
			testdata.ErrZeroValue,
			[]int{1},
			t)
	})

	t.Run("We should be able to set Raw LenEncoding for StringAlias", func(t *testing.T) {
		v := pkg1.StringAlias("some")
		testLenEncoding(v, pkg1.MarshalLenEncodingStringAliasMUS,
			raw.UnmarshalInt,
			pkg1.SizeLenEncodingStringAliasMUS,
			len(v),
			t)
	})

	t.Run("We should be able to set LenValidator for StringAlias", func(t *testing.T) {
		v := pkg1.StringAlias("some")
		testValidation(v, pkg1.MarshalLenValidatorStringAliasMUS,
			pkg1.UnmarshalLenValidatorStringAliasMUS,
			pkg1.SizeLenValidatorStringAliasMUS,
			testdata.ErrTooLong,
			[]int{1},
			t)
	})

	t.Run("We should be able to set Validator for StringAlias", func(t *testing.T) {
		v := pkg1.StringAlias("")
		testValidation(v, pkg1.MarshalValidStringAliasMUS,
			pkg1.UnmarshalValidStringAliasMUS,
			pkg1.SizeValidStringAliasMUS,
			testdata.ErrZeroValue,
			[]int{1},
			t)
	})

	t.Run("SliceAlias should be serializable", func(t *testing.T) {
		v := pkg1.SliceAlias([]int{1, 2, 3})
		testSerializability(v, pkg1.MarshalSliceAliasMUS,
			pkg1.UnmarshalSliceAliasMUS,
			pkg1.SizeSliceAliasMUS,
			pkg1.SkipSliceAliasMUS,
			t)
	})

	t.Run("We should be able to set LenEncoding for SliceAlias", func(t *testing.T) {
		v := pkg1.SliceAlias([]int{1, 2, 3})
		testLenEncoding(v, pkg1.MarshalLenEncodingSliceAliasMUS,
			varint.UnmarshalPositiveInt,
			pkg1.SizeLenEncodingSliceAliasMUS,
			len(v),
			t)
	})

	t.Run("We should be able to set LenValidator for SliceAlias", func(t *testing.T) {
		testValidation(pkg1.SliceAlias([]int{1, 2, 3}),
			pkg1.MarshalLenValidatorSliceAliasMUS,
			pkg1.UnmarshalLenValidatorSliceAliasMUS,
			pkg1.SizeLenValidatorSliceAliasMUS,
			testdata.ErrTooLong,
			[]int{1},
			t,
		)
	})

	t.Run("We should be able to set Elem.Encoding for SliceAlias", func(t *testing.T) {
		testSerializability(pkg1.SliceAlias([]int{1, 2, 3}),
			pkg1.MarshalElemEncodingSliceAliasMUS,
			func(bs []byte) (sla pkg1.SliceAlias, n int, err error) {
				sl, n, err := ord.UnmarshalSlice[int](nil,
					mus.UnmarshallerFn[int](raw.UnmarshalInt), bs)
				sla = pkg1.SliceAlias(sl)
				return
			},
			pkg1.SizeElemEncodingSliceAliasMUS,
			pkg1.SkipElemEncodingSliceAliasMUS,
			t)
	})

	t.Run("We should be able to set Elem.Validator for SliceAlias", func(t *testing.T) {
		testValidation(
			pkg1.SliceAlias([]int{1, 0, 3}),
			pkg1.MarshalElemValidatorSliceAliasMUS,
			pkg1.UnmarshalElemValidatorSliceAliasMUS,
			pkg1.SizeElemValidatorSliceAliasMUS,
			testdata.ErrZeroValue,
			[]int{3},
			t)
	})

	t.Run("MapAlias should be serializable", func(t *testing.T) {
		v := pkg1.MapAlias(map[string]int{"some": 1, "another": 2})
		testSerializability(v, pkg1.MarshalMapAliasMUS,
			pkg1.UnmarshalMapAliasMUS,
			pkg1.SizeMapAliasMUS,
			pkg1.SkipMapAliasMUS,
			t)
	})

	t.Run("We should be able to set LenValidator for MapAlias", func(t *testing.T) {
		testValidation(
			pkg1.MapAlias(map[string]int{"some": 1, "another": 2}),
			pkg1.MarshalLenValidatorMapAliasMUS,
			pkg1.UnmarshalLenValidatorMapAliasMUS,
			pkg1.SizeLenValidatorMapAliasMUS,
			testdata.ErrTooLong,
			[]int{1},
			t)
	})

	t.Run("We should be able to set Key.Validator for MapAlias", func(t *testing.T) {
		testValidation(
			pkg1.MapAlias(map[string]int{"some": 1, "": 2}),
			pkg1.MarshalKeyValidatorMapAliasMUS,
			pkg1.UnmarshalKeyValidatorMapAliasMUS,
			pkg1.SizeKeyValidatorMapAliasMUS,
			testdata.ErrZeroValue,
			[]int{8, 2},
			t)
	})

	t.Run("We should be able to set Elem.Validator for MapAlias", func(t *testing.T) {
		testValidation(
			pkg1.MapAlias(map[string]int{"some": 0, "another": 2}),
			pkg1.MarshalElemValidatorMapAliasMUS,
			pkg1.UnmarshalElemValidatorMapAliasMUS,
			pkg1.SizeElemValidatorMapAliasMUS,
			testdata.ErrZeroValue,
			[]int{7, 16},
			t)
	})

	t.Run("Test dts serializability", func(t *testing.T) {
		testSerializability(pkg1.Struct{Int: 9},
			pkg1.StructDTS.Marshal,
			pkg1.StructDTS.Unmarshal,
			pkg1.StructDTS.Size,
			pkg1.StructDTS.Skip,
			t)
	})

	t.Run("Test struct/alias/dts/interface/ptr serializability", func(t *testing.T) {

		t.Run("We should be able to skip a struct field", func(t *testing.T) {
			var (
				v = pkg2.Struct{
					Float32: 23.5,
					Float64: 90.005,
				}
				bs = make([]byte, pkg2.SizeSkipStructMUS(v))
			)
			pkg2.MarshalSkipStructMUS(v, bs)
			s, _, err := pkg2.UnmarshalSkipStructMUS(bs)
			if err != nil {
				t.Fatal(err)
			}
			v.Float32 = 0.0
			if !reflect.DeepEqual(s, v) {
				t.Errorf("unexpected value %v", s)
			}
		})

		t.Run("ComplexStruct should be serializable", func(t *testing.T) {
			testSerializability(makeCompplexStruct(),
				pkg1.MarshalComplexStructMUS,
				pkg1.UnmarshalComplexStructMUS,
				pkg1.SizeComplexStructMUS,
				pkg1.SkipComplexStructMUS,
				t)
		})
	})

}

func testSerializability[T any](v T, m mus.MarshallerFn[T],
	u mus.UnmarshallerFn[T],
	s mus.SizerFn[T],
	sk mus.SkipperFn,
	t *testing.T,
) {
	bs := make([]byte, s(v))
	m(v, bs)
	v1, n1, err := u(bs)
	if err != nil {
		t.Fatal(err)
	}
	n2, err := sk(bs)
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

func testValidation[T any](v T, m mus.MarshallerFn[T],
	u mus.UnmarshallerFn[T],
	s mus.SizerFn[T],
	wantErr error,
	wantN []int,
	t *testing.T,
) {
	bs := make([]byte, s(v))
	m(v, bs)
	_, n, err := u(bs)
	if err != wantErr {
		t.Errorf("unexpected error %v", err)
	}
	var found bool
	for i := 0; i < len(wantN); i++ {
		if n == wantN[i] {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("unexpected value %v", n)
	}
}

func testLenEncoding[T any](v T, m mus.MarshallerFn[T],
	u mus.UnmarshallerFn[int],
	s mus.SizerFn[T],
	wantLength int,
	t *testing.T,
) {
	bs := make([]byte, s(v))
	m(v, bs)
	l, _, err := u(bs)
	if err != nil {
		t.Fatal(err)
	}
	if l != wantLength {
		t.Errorf("unexpected value %v", l)
	}
}

func makeCompplexStruct() pkg1.ComplexStruct {
	return pkg1.ComplexStruct{
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

		Alias: pkg1.SliceAlias{1, 2, 3, 4},
		Ptr: &pkg1.Struct{
			Int: 100,
		},
		AnotherPkgStruct: pkg2.Struct{
			Float32: 3.0,
			Float64: 55.0,
			Bool:    false,
			Byte:    100,
		},
		Interface: pkg1.InterfaceImpl1{},

		SliceByte:   []byte{1, 2, 3, 4},
		SliceStruct: []pkg1.Struct{{Int: 10}, {Int: 20}},

		Array: [3]int{1, 2, 3},

		Map: map[float32]map[pkg1.IntAlias][]pkg1.Struct{
			40.8: {
				pkg1.IntAlias(11): {
					{Int: 30},
				},
			},
		},
	}
}
