package musgen

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/varint"
	"github.com/mus-format/musgen-go/testdata"
	"github.com/mus-format/musgen-go/testdata/pkg1"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

type Unmarshaller[T any] interface {
	Unmarshal([]byte) (T, int, error)
}

type UnmarshallerFn[T any] func([]byte) (T, int, error)

func (u UnmarshallerFn[T]) Unmarshal(bs []byte) (T, int, error) {
	return u(bs)
}

func TestGeneratedCode(t *testing.T) {

	t.Run("Test alias serializability and Options impact", func(t *testing.T) {

		t.Run("IntAlias should be serializable", func(t *testing.T) {
			var (
				v                               = pkg1.IntAlias(5)
				u UnmarshallerFn[pkg1.IntAlias] = func(bs []byte) (pkg1.IntAlias, int, error) {
					num, n, err := varint.Int.Unmarshal(bs)
					return pkg1.IntAlias(num), n, err
				}
			)
			testSerializability(v, pkg1.IntAliasMUS, u, t)
		})

		t.Run("We should be able to set Raw encoding for IntAlias", func(t *testing.T) {
			var (
				v                                  = pkg1.RawIntAlias(5)
				u UnmarshallerFn[pkg1.RawIntAlias] = func(bs []byte) (pkg1.RawIntAlias, int, error) {
					num, n, err := raw.Int.Unmarshal(bs)
					return pkg1.RawIntAlias(num), n, err
				}
			)
			testSerializability[pkg1.RawIntAlias](v, pkg1.RawIntAliasMUS, u, t)
		})

		t.Run("We should be able to set VarintPositive encoding for IntAlias", func(t *testing.T) {
			var (
				v                                             = pkg1.VarintPositiveIntAlias(5)
				u UnmarshallerFn[pkg1.VarintPositiveIntAlias] = func(bs []byte) (pkg1.VarintPositiveIntAlias, int, error) {
					num, n, err := varint.PositiveInt.Unmarshal(bs)
					return pkg1.VarintPositiveIntAlias(num), n, err
				}
			)
			testSerializability[pkg1.VarintPositiveIntAlias](v,
				pkg1.VarintPositiveIntAliasMUS, u, t)
		})

		t.Run("We should be able to set Validator for IntAlias", func(t *testing.T) {
			v := pkg1.ValidIntAlias(0)
			testValidation(v, pkg1.ValidIntAliasMUS, testdata.ErrZeroValue, []int{1}, t)
		})

		t.Run("We should be able to set Raw LenEncoding for StringAlias", func(t *testing.T) {
			var (
				v                        = pkg1.LenEncodingStringAlias("some")
				lenU UnmarshallerFn[int] = func(bs []byte) (int, int, error) {
					num, n, err := raw.Int.Unmarshal(bs)
					return int(num), n, err
				}
			)
			testLenEncoding(v, pkg1.LenEncodingStringAliasMUS, lenU, len(v), t)
		})

		t.Run("We should be able to set LenValidator for StringAlias", func(t *testing.T) {
			v := pkg1.LenValidStringAlias("some")
			testValidation(v, pkg1.LenValidStringAliasMUS, testdata.ErrTooLong, []int{1}, t)
		})

		t.Run("We should be able to set Validator for StringAlias", func(t *testing.T) {
			v := pkg1.ValidStringAlias("")
			testValidation(v, pkg1.ValidStringAliasMUS, testdata.ErrZeroValue, []int{1}, t)
		})

		t.Run("SliceAlias should be serializable", func(t *testing.T) {
			var (
				v                                 = pkg1.SliceAlias([]int{1, 2, 3})
				u UnmarshallerFn[pkg1.SliceAlias] = func(bs []byte) (pkg1.SliceAlias, int, error) {
					ser := ord.NewSliceSer[int](varint.Int)
					sl, n, err := ser.Unmarshal(bs)
					return pkg1.SliceAlias(sl), n, err
				}
			)
			testSerializability(v, pkg1.SliceAliasMUS, u, t)
		})

		t.Run("We should be able to set LenEncoding for SliceAlias", func(t *testing.T) {
			var (
				v                        = pkg1.LenEncodingSliceAlias([]int{1, 2, 3})
				lenU UnmarshallerFn[int] = func(bs []byte) (int, int, error) {
					num, n, err := varint.PositiveInt.Unmarshal(bs)
					return int(num), n, err
				}
			)
			testLenEncoding(v, pkg1.LenEncodingSliceAliasMUS, lenU, len(v), t)
		})

		t.Run("We should be able to set LenValidator for SliceAlias", func(t *testing.T) {
			v := pkg1.LenValidSliceAlias([]int{1, 2, 3})
			testValidation(v, pkg1.LenValidSliceAliasMUS, testdata.ErrTooLong, []int{1}, t)
		})

		t.Run("We should be able to set Elem.Encoding for SliceAlias", func(t *testing.T) {
			var (
				v                                             = pkg1.ElemEncodingSliceAlias([]int{1, 2, 3})
				u UnmarshallerFn[pkg1.ElemEncodingSliceAlias] = func(bs []byte) (pkg1.ElemEncodingSliceAlias, int, error) {
					ser := ord.NewSliceSer[int](raw.Int)
					sl, n, err := ser.Unmarshal(bs)
					return pkg1.ElemEncodingSliceAlias(sl), n, err
				}
			)
			testSerializability(v, pkg1.ElemEncodingSliceAliasMUS, u, t)
		})

		t.Run("We should be able to set Elem.Validator for SliceAlias", func(t *testing.T) {
			v := pkg1.ElemValidSliceAlias([]int{1, 0, 3})
			testValidation(v, pkg1.ElemValidSliceAliasMUS, testdata.ErrZeroValue, []int{3}, t)
		})

		t.Run("ArrayAlias should be serializable", func(t *testing.T) {
			var (
				v                                 = pkg1.ArrayAlias([3]int{1, 2, 3})
				u UnmarshallerFn[pkg1.ArrayAlias] = func(bs []byte) (pkg1.ArrayAlias, int, error) {
					ser := ord.NewSliceSer[int](varint.Int)
					sl, n, err := ser.Unmarshal(bs)
					return pkg1.ArrayAlias(sl), n, err
				}
			)
			testSerializability(v, pkg1.ArrayAliasMUS, u, t)
		})

		t.Run("We should be able to set LenEncoding for ArrayAlias", func(t *testing.T) {
			var (
				v                        = pkg1.LenEncodingArrayAlias([3]int{1, 2, 3})
				lenU UnmarshallerFn[int] = func(bs []byte) (int, int, error) {
					num, n, err := raw.Int.Unmarshal(bs)
					return int(num), n, err
				}
			)
			testLenEncoding(v, pkg1.LenEncodingArrayAliasMUS, lenU, len(v), t)
		})

		t.Run("We should be able to set Elem.Encoding for ArrayAlias", func(t *testing.T) {
			var (
				v                                             = pkg1.ElemEncodingArrayAlias([3]int{1, 2, 3})
				u UnmarshallerFn[pkg1.ElemEncodingArrayAlias] = func(bs []byte) (pkg1.ElemEncodingArrayAlias, int, error) {
					ser := ord.NewArraySer[[3]int, int](3, raw.Int)
					arr, n, err := ser.Unmarshal(bs)
					return pkg1.ElemEncodingArrayAlias(arr), n, err
				}
			)
			testSerializability(v, pkg1.ElemEncodingArrayAliasMUS, u, t)
		})

		t.Run("We should be able to set Elem.Validator for ArrayAlias", func(t *testing.T) {
			v := pkg1.ElemValidArrayAlias([3]int{1, 0, 3})
			testValidation(v, pkg1.ElemValidArrayAliasMUS, testdata.ErrZeroValue, []int{3}, t)
		})

		t.Run("MapAlias should be serializable", func(t *testing.T) {
			var (
				v                               = pkg1.MapAlias(map[int]int{10: 1, 11: 2})
				u UnmarshallerFn[pkg1.MapAlias] = func(bs []byte) (pkg1.MapAlias, int, error) {
					ser := ord.NewMapSer[int, int](varint.Int, varint.Int)
					m, n, err := ser.Unmarshal(bs)
					return pkg1.MapAlias(m), n, err
				}
			)
			testSerializability(v, pkg1.MapAliasMUS, u, t)
		})

		t.Run("We should be able to set LenEncoding for MapAlias", func(t *testing.T) {
			var (
				v                        = pkg1.LenEncodingMapAlias(map[int]int{10: 1, 11: 2})
				lenU UnmarshallerFn[int] = func(bs []byte) (int, int, error) {
					num, n, err := raw.Int.Unmarshal(bs)
					return int(num), n, err
				}
			)
			testLenEncoding(v, pkg1.LenEncodingMapAliasMUS, lenU, len(v), t)
		})

		t.Run("We should be able to set LenValidator for MapAlias", func(t *testing.T) {
			v := pkg1.LenValidMapAlias(map[int]int{10: 1, 11: 2})
			testValidation(v, pkg1.LenValidMapAliasMUS, testdata.ErrTooLong, []int{1}, t)
		})

		t.Run("We should be able to set Key.Encoding for ArrayAlias", func(t *testing.T) {
			var (
				v                                          = pkg1.KeyEncodingMapAlias(map[int]int{10: 1, 11: 2})
				u UnmarshallerFn[pkg1.KeyEncodingMapAlias] = func(bs []byte) (pkg1.KeyEncodingMapAlias, int, error) {
					ser := ord.NewMapSer(raw.Int, varint.Int)
					m, n, err := ser.Unmarshal(bs)
					return pkg1.KeyEncodingMapAlias(m), n, err
				}
			)
			testSerializability(v, pkg1.KeyEncodingMapAliasMUS, u, t)
		})

		t.Run("We should be able to set Key.Validator for MapAlias", func(t *testing.T) {
			v := pkg1.KeyValidMapAlias(map[int]int{0: 1, 11: 2})
			testValidation(v, pkg1.KeyValidMapAliasMUS, testdata.ErrZeroValue, []int{2, 4}, t)
		})

		t.Run("We should be able to set Elem.Encoding for ArrayAlias", func(t *testing.T) {
			var (
				v                                           = pkg1.ElemEncodingMapAlias(map[int]int{10: 1, 11: 2})
				u UnmarshallerFn[pkg1.ElemEncodingMapAlias] = func(bs []byte) (pkg1.ElemEncodingMapAlias, int, error) {
					ser := ord.NewMapSer(varint.Int, raw.Int)
					m, n, err := ser.Unmarshal(bs)
					return pkg1.ElemEncodingMapAlias(m), n, err
				}
			)
			testSerializability(v, pkg1.ElemEncodingMapAliasMUS, u, t)
		})

		t.Run("We should be able to set Elem.Validator for MapAlias", func(t *testing.T) {
			v := pkg1.ElemValidMapAlias(map[int]int{10: 0, 11: 2})
			testValidation(v, pkg1.ElemValidMapAliasMUS, testdata.ErrZeroValue, []int{3, 5}, t)
		})

		// ptr

		t.Run("PtrAlias should be serializable", func(t *testing.T) {
			var (
				n                               = 5
				v                               = pkg1.PtrAlias(&n)
				u UnmarshallerFn[pkg1.PtrAlias] = func(bs []byte) (pkg1.PtrAlias, int, error) {
					ptr, n, err := ord.NewPtrSer(varint.Int).Unmarshal(bs)
					return pkg1.PtrAlias(ptr), n, err
				}
			)
			testSerializability(v, pkg1.PtrAliasMUS, u, t)
		})

		t.Run("We shoul be able to set Elem.NumEncoding for PtrAlias", func(t *testing.T) {
			var (
				num                                              = 5
				v                                                = pkg1.ElemNumEncodingPtrAlias(&num)
				u   UnmarshallerFn[pkg1.ElemNumEncodingPtrAlias] = func(bs []byte) (pkg1.ElemNumEncodingPtrAlias, int, error) {
					ptr, n, err := ord.NewPtrSer(raw.Int).Unmarshal(bs)
					return pkg1.ElemNumEncodingPtrAlias(ptr), n, err
				}
			)
			testSerializability(v, pkg1.ElemNumEncodingPtrAliasMUS, u, t)
		})

		t.Run("We should be able to serialize InterfaceDoublePtrAlias", func(t *testing.T) {
			var (
				in    pkg1.Interface               = pkg1.InterfaceImpl1{}
				inPtr                              = &in
				v     pkg1.InterfaceDoublePtrAlias = &inPtr
			)
			testSerializability(v, pkg1.InterfaceDoublePtrAliasMUS, nil, t)
		})

		// TODO ptr validation
	})

	t.Run("Test struct serializability and Options impact", func(t *testing.T) {

		t.Run("We should be able to ignore a struct field", func(t *testing.T) {
			var (
				wantFloat64 = 0.0
				v           = pkg2.IgnoreFieldStruct{
					Float32: 10.0,
					Float64: 20.0,
					Byte:    30,
					Bool:    true,
				}
				bs = make([]byte, pkg2.IgnoreFieldStructMUS.Size(v))
			)
			pkg2.IgnoreFieldStructMUS.Marshal(v, bs)
			v, _, err := pkg2.IgnoreFieldStructMUS.Unmarshal(bs)
			if err != nil {
				t.Fatal(err)
			}

			if v.Float64 != wantFloat64 {
				t.Errorf("unexpected value, want %v actual %v", wantFloat64, v.Float64)
			}
		})

		t.Run("We should be able to validate a struct field", func(t *testing.T) {
			var (
				v = pkg2.ValidFieldStruct{
					Float32: 10.0,
					Float64: 20.0,
					Byte:    0,
					Bool:    true,
				}
			)
			testValidation(v, pkg2.ValidFieldStructMUS, testdata.ErrZeroValue, []int{15}, t)
		})

		t.Run("We should be able to set options for field's inner slice", func(t *testing.T) {
			var (
				v = pkg2.ElemStruct{
					Slice: [][]int{{3}},
				}
				innerSer                                 = ord.NewSliceSer(raw.Int)
				ser                                      = ord.NewValidSliceSer(innerSer, com.ValidatorFn[int](testdata.ValidateLength1), nil)
				u        UnmarshallerFn[pkg2.ElemStruct] = func(bs []byte) (v pkg2.ElemStruct, n int, err error) {
					n, _ = varint.Float32.Skip(bs)
					var n1 int
					n1, _ = varint.Float64.Skip(bs[n:])
					n += n1
					n1, _ = varint.Byte.Skip(bs[n:])
					n += n1
					n1, _ = ord.Bool.Skip(bs[n:])
					n += n1
					v.Slice, n1, err = ser.Unmarshal(bs[n:])
					n += n1
					return
				}
			)
			testSerializability(v, pkg2.ElemStructMUS, u, t)
		})

		t.Run("We should be able to validate field's inner slice", func(t *testing.T) {
			var (
				v = pkg2.ElemStruct{
					Slice: [][]int{{1, 2, 3}},
				}
				wantErr = testdata.ErrTooLong
				wantN   = 6
			)
			testValidation(v, pkg2.ElemStructMUS, wantErr, []int{wantN}, t)
		})

		t.Run("We should be able to validate slice field", func(t *testing.T) {
			var (
				v = pkg2.ElemStruct{
					Slice: [][]int{{1, 2, 3}, {1, 2, 3}},
				}
				wantErr = testdata.ErrTooLong
				wantN   = 5
			)
			testValidation(v, pkg2.ElemStructMUS, wantErr, []int{wantN}, t)
		})

		t.Run("We should be able to serialize ComplexStruct", func(t *testing.T) {
			v := makeCompplexStruct()
			testSerializability(v, pkg1.ComplexStructMUS, nil, t)
		})

		t.Run("We should be able to serialize AnotherStruct", func(t *testing.T) {
			v := pkg1.AnotherStruct{Int: 3}
			testSerializability(v, pkg1.AnotherStructMUS, nil, t)
		})

		// t.Run("We should be able to serialize struct using DTS", func(t *testing.T) {
		// 	testSerializability(pkg1.SimpleStruct{Int: 10}, pkg1.SimpleStructDTS, nil,
		// 		t)
		// })

		// t.Run("We should be able to serialize alias using DTS", func(t *testing.T) {
		// 	testSerializability(pkg1.IntAlias(10), pkg1.IntAliasDTS, nil, t)
		// })

		// Test DTS from another package

	})

	t.Run("Test interface serializability", func(t *testing.T) {

		t.Run("We sohuld be able to serialize interface", func(t *testing.T) {
			var v pkg1.Interface = pkg1.InterfaceImpl1{Str: "hello world"}
			testSerializability(v, pkg1.InterfaceMUS, nil, t)
		})

		t.Run("We sohuld be able to serialize AnotherInterface", func(t *testing.T) {
			var v pkg1.AnotherInterface = pkg1.InterfaceImpl2(1)
			testSerializability(v, pkg1.AnotherInterfaceMUS, nil, t)
		})

	})
}

func testSerializability[T any](v T, ser mus.Serializer[T], u UnmarshallerFn[T], t *testing.T) {
	bs := make([]byte, ser.Size(v))
	ser.Marshal(v, bs)
	v1, n1, err := ser.Unmarshal(bs)
	if err != nil {
		t.Fatal(err)
	}
	n2, err := ser.Skip(bs)
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

	if u != nil {
		v1, n1, err = u(bs)
		if err != nil {
			t.Fatal(err)
		}
		if !cmp.Equal(v, v1, opt) {
			t.Errorf("unexpected value %v", v1)
		}
		if n1 != n2 {
			t.Errorf("unexpected value %v", n2)
		}
	}
}

// wantN is an array because of unpredictable map key/value pairs order.
func testValidation[T any](v T, ser mus.Serializer[T], wantErr error,
	wantN []int, t *testing.T) {
	bs := make([]byte, ser.Size(v))
	ser.Marshal(v, bs)
	_, n, err := ser.Unmarshal(bs)
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

func testLenEncoding[T any](v T, ser mus.Serializer[T], lenU Unmarshaller[int],
	wantLength int, t *testing.T) {
	bs := make([]byte, ser.Size(v))
	ser.Marshal(v, bs)
	l, _, err := lenU.Unmarshal(bs)
	if err != nil {
		t.Fatal(err)
	}
	if l != wantLength {
		t.Errorf("unexpected value %v", l)
	}
}

func makeCompplexStruct() pkg1.ComplexStruct {
	str := "some"
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
		AnotherPkgStruct: pkg2.Struct{
			Float32: 3.0,
			Float64: 55.0,
			Bool:    false,
			Byte:    100,
		},

		Interface: pkg1.InterfaceImpl1{},

		ByteSlice:   []byte{1, 2, 3, 4},
		StructSlice: []pkg1.SimpleStruct{{Int: 10}, {Int: 20}},

		Array: [3]int{1, 2, 3},

		PtrArray:  &[3]int{1, 1, 1},
		PtrString: &str,
		PtrStruct: &pkg1.SimpleStruct{
			Int: 100,
		},
		NilPtr: nil,

		Map: map[float32]map[pkg1.IntAlias][]pkg1.SimpleStruct{
			40.8: {
				pkg1.IntAlias(11): {
					{Int: 30},
				},
			},
		},
	}
}
