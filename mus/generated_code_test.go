package musgen

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	slops "github.com/mus-format/mus-go/options/slice"
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

	t.Run("Test defined_type serializability and Options impact", func(t *testing.T) {

		t.Run("MyInt should be serializable", func(t *testing.T) {
			var (
				v                            = pkg1.MyInt(5)
				u UnmarshallerFn[pkg1.MyInt] = func(bs []byte) (pkg1.MyInt, int, error) {
					num, n, err := varint.Int.Unmarshal(bs)
					return pkg1.MyInt(num), n, err
				}
			)
			testSerializability(v, pkg1.MyIntMUS, u, t)
		})

		t.Run("We should be able to set Raw encoding for MyInt", func(t *testing.T) {
			var (
				v                               = pkg1.RawMyInt(5)
				u UnmarshallerFn[pkg1.RawMyInt] = func(bs []byte) (pkg1.RawMyInt, int, error) {
					num, n, err := raw.Int.Unmarshal(bs)
					return pkg1.RawMyInt(num), n, err
				}
			)
			testSerializability[pkg1.RawMyInt](v, pkg1.RawMyIntMUS, u, t)
		})

		t.Run("We should be able to set VarintPositive encoding for MyInt", func(t *testing.T) {
			var (
				v                                          = pkg1.VarintPositiveMyInt(5)
				u UnmarshallerFn[pkg1.VarintPositiveMyInt] = func(bs []byte) (pkg1.VarintPositiveMyInt, int, error) {
					num, n, err := varint.PositiveInt.Unmarshal(bs)
					return pkg1.VarintPositiveMyInt(num), n, err
				}
			)
			testSerializability[pkg1.VarintPositiveMyInt](v,
				pkg1.VarintPositiveMyIntMUS, u, t)
		})

		t.Run("We should be able to set Validator for MyInt", func(t *testing.T) {
			v := pkg1.ValidMyInt(0)
			testValidation(v, pkg1.ValidMyIntMUS, testdata.ErrZeroValue, []int{1}, t)
		})

		t.Run("We should be able to set Raw LenEncoding for MyString", func(t *testing.T) {
			var (
				v                        = pkg1.LenEncodingMyString("some")
				lenU UnmarshallerFn[int] = func(bs []byte) (int, int, error) {
					num, n, err := raw.Int.Unmarshal(bs)
					return int(num), n, err
				}
			)
			testLenEncoding(v, pkg1.LenEncodingMyStringMUS, lenU, len(v), t)
		})

		t.Run("We should be able to set LenValidator for MyString", func(t *testing.T) {
			v := pkg1.LenValidMyString("some")
			testValidation(v, pkg1.LenValidMyStringMUS, testdata.ErrTooLong, []int{1}, t)
		})

		t.Run("We should be able to set Validator for MyString", func(t *testing.T) {
			v := pkg1.ValidMyString("")
			testValidation(v, pkg1.ValidMyStringMUS, testdata.ErrZeroValue, []int{1}, t)
		})

		t.Run("MyTime should be serializable", func(t *testing.T) {
			var (
				v                             = pkg1.MyTime(time.Unix(time.Now().Unix(), 0))
				u UnmarshallerFn[pkg1.MyTime] = func(bs []byte) (tm pkg1.MyTime, n int, err error) {
					sec, n, err := raw.Int64.Unmarshal(bs)
					if err != nil {
						return
					}
					tm = pkg1.MyTime(time.Unix(sec, 0))
					return
				}
			)
			testSerializability(v, pkg1.MyTimeMUS, u, t)
		})

		t.Run("MyTimeSec should be serializable", func(t *testing.T) {
			var (
				v                                = pkg1.MyTimeSec(time.Unix(time.Now().Unix(), 0))
				u UnmarshallerFn[pkg1.MyTimeSec] = func(bs []byte) (tm pkg1.MyTimeSec, n int, err error) {
					sec, n, err := raw.Int64.Unmarshal(bs)
					if err != nil {
						return
					}
					tm = pkg1.MyTimeSec(time.Unix(sec, 0))
					return
				}
			)
			testSerializability(v, pkg1.MyTimeSecMUS, u, t)
		})

		t.Run("MyTimeMilli should be serializable", func(t *testing.T) {
			var (
				v                                  = pkg1.MyTimeMilli(time.Unix(time.Now().Unix(), 0))
				u UnmarshallerFn[pkg1.MyTimeMilli] = func(bs []byte) (tm pkg1.MyTimeMilli, n int, err error) {
					milli, n, err := raw.Int64.Unmarshal(bs)
					if err != nil {
						return
					}
					tm = pkg1.MyTimeMilli(time.UnixMilli(milli))
					return
				}
			)
			testSerializability(v, pkg1.MyTimeMilliMUS, u, t)
		})

		t.Run("MyTimeMicro should be serializable", func(t *testing.T) {
			var (
				v                                  = pkg1.MyTimeMicro(time.Unix(time.Now().Unix(), 0))
				u UnmarshallerFn[pkg1.MyTimeMicro] = func(bs []byte) (tm pkg1.MyTimeMicro, n int, err error) {
					micro, n, err := raw.Int64.Unmarshal(bs)
					if err != nil {
						return
					}
					tm = pkg1.MyTimeMicro(time.UnixMicro(micro))
					return
				}
			)
			testSerializability(v, pkg1.MyTimeMicroMUS, u, t)
		})

		t.Run("MyTimeNano should be serializable", func(t *testing.T) {
			var (
				v                                 = pkg1.MyTimeNano(time.Unix(time.Now().Unix(), 0))
				u UnmarshallerFn[pkg1.MyTimeNano] = func(bs []byte) (tm pkg1.MyTimeNano, n int, err error) {
					nano, n, err := raw.Int64.Unmarshal(bs)
					if err != nil {
						return
					}
					tm = pkg1.MyTimeNano(time.Unix(0, nano))
					return
				}
			)
			testSerializability(v, pkg1.MyTimeNanoMUS, u, t)
		})

		t.Run("MyTimeSecUTC should be serializable", func(t *testing.T) {
			var (
				v                                   = pkg1.MyTimeSecUTC(time.Unix(time.Now().Unix(), 0).UTC())
				u UnmarshallerFn[pkg1.MyTimeSecUTC] = func(bs []byte) (tm pkg1.MyTimeSecUTC, n int, err error) {
					sec, n, err := raw.Int64.Unmarshal(bs)
					if err != nil {
						return
					}
					tm = pkg1.MyTimeSecUTC(time.Unix(sec, 0))
					return
				}
			)
			testSerializability(v, pkg1.MyTimeSecUTCMUS, u, t)
		})

		t.Run("MyTimeMilliUTC should be serializable", func(t *testing.T) {
			var (
				v                                     = pkg1.MyTimeMilliUTC(time.Unix(time.Now().Unix(), 0).UTC())
				u UnmarshallerFn[pkg1.MyTimeMilliUTC] = func(bs []byte) (tm pkg1.MyTimeMilliUTC, n int, err error) {
					milli, n, err := raw.Int64.Unmarshal(bs)
					if err != nil {
						return
					}
					tm = pkg1.MyTimeMilliUTC(time.UnixMilli(milli))
					return
				}
			)
			testSerializability(v, pkg1.MyTimeMilliUTCMUS, u, t)
		})

		t.Run("MyTimeMicroUTC should be serializable", func(t *testing.T) {
			var (
				v                                     = pkg1.MyTimeMicroUTC(time.Unix(time.Now().Unix(), 0).UTC())
				u UnmarshallerFn[pkg1.MyTimeMicroUTC] = func(bs []byte) (tm pkg1.MyTimeMicroUTC, n int, err error) {
					milli, n, err := raw.Int64.Unmarshal(bs)
					if err != nil {
						return
					}
					tm = pkg1.MyTimeMicroUTC(time.UnixMicro(milli))
					return
				}
			)
			testSerializability(v, pkg1.MyTimeMicroUTCMUS, u, t)
		})

		t.Run("MyTimeNanoUTC should be serializable", func(t *testing.T) {
			var (
				v                                    = pkg1.MyTimeNanoUTC(time.Unix(time.Now().Unix(), 0).UTC())
				u UnmarshallerFn[pkg1.MyTimeNanoUTC] = func(bs []byte) (tm pkg1.MyTimeNanoUTC, n int, err error) {
					nano, n, err := raw.Int64.Unmarshal(bs)
					if err != nil {
						return
					}
					tm = pkg1.MyTimeNanoUTC(time.Unix(0, nano))
					return
				}
			)
			testSerializability(v, pkg1.MyTimeNanoUTCMUS, u, t)
		})

		t.Run("MySlice should be serializable", func(t *testing.T) {
			var (
				v                              = pkg1.MySlice([]int{1, 2, 3})
				u UnmarshallerFn[pkg1.MySlice] = func(bs []byte) (pkg1.MySlice, int, error) {
					ser := ord.NewSliceSer[int](varint.Int)
					sl, n, err := ser.Unmarshal(bs)
					return pkg1.MySlice(sl), n, err
				}
			)
			testSerializability(v, pkg1.MySliceMUS, u, t)
		})

		t.Run("We should be able to set LenEncoding for MySlice", func(t *testing.T) {
			var (
				v                        = pkg1.LenEncodingMySlice([]int{1, 2, 3})
				lenU UnmarshallerFn[int] = func(bs []byte) (int, int, error) {
					num, n, err := varint.PositiveInt.Unmarshal(bs)
					return int(num), n, err
				}
			)
			testLenEncoding(v, pkg1.LenEncodingMySliceMUS, lenU, len(v), t)
		})

		t.Run("We should be able to set LenValidator for MySlice", func(t *testing.T) {
			v := pkg1.LenValidMySlice([]int{1, 2, 3})
			testValidation(v, pkg1.LenValidMySliceMUS, testdata.ErrTooLong, []int{1}, t)
		})

		t.Run("We should be able to set Elem.Encoding for MySlice", func(t *testing.T) {
			var (
				v                                          = pkg1.ElemEncodingMySlice([]int{1, 2, 3})
				u UnmarshallerFn[pkg1.ElemEncodingMySlice] = func(bs []byte) (pkg1.ElemEncodingMySlice, int, error) {
					ser := ord.NewSliceSer[int](raw.Int)
					sl, n, err := ser.Unmarshal(bs)
					return pkg1.ElemEncodingMySlice(sl), n, err
				}
			)
			testSerializability(v, pkg1.ElemEncodingMySliceMUS, u, t)
		})

		t.Run("We should be able to set Elem.Validator for MySlice", func(t *testing.T) {
			v := pkg1.ElemValidMySlice([]int{1, 0, 3})
			testValidation(v, pkg1.ElemValidMySliceMUS, testdata.ErrZeroValue, []int{3}, t)
		})

		t.Run("MyArray should be serializable", func(t *testing.T) {
			var (
				v                              = pkg1.MyArray([3]int{1, 2, 3})
				u UnmarshallerFn[pkg1.MyArray] = func(bs []byte) (pkg1.MyArray, int, error) {
					ser := ord.NewSliceSer[int](varint.Int)
					sl, n, err := ser.Unmarshal(bs)
					return pkg1.MyArray(sl), n, err
				}
			)
			testSerializability(v, pkg1.MyArrayMUS, u, t)
		})

		t.Run("We should be able to set LenEncoding for MyArray", func(t *testing.T) {
			var (
				v                        = pkg1.LenEncodingMyArray([3]int{1, 2, 3})
				lenU UnmarshallerFn[int] = func(bs []byte) (int, int, error) {
					num, n, err := raw.Int.Unmarshal(bs)
					return int(num), n, err
				}
			)
			testLenEncoding(v, pkg1.LenEncodingMyArrayMUS, lenU, len(v), t)
		})

		t.Run("We should be able to set Elem.Encoding for MyArray", func(t *testing.T) {
			var (
				v                                          = pkg1.ElemEncodingMyArray([3]int{1, 2, 3})
				u UnmarshallerFn[pkg1.ElemEncodingMyArray] = func(bs []byte) (pkg1.ElemEncodingMyArray, int, error) {
					ser := ord.NewArraySer[[3]int, int](3, raw.Int)
					arr, n, err := ser.Unmarshal(bs)
					return pkg1.ElemEncodingMyArray(arr), n, err
				}
			)
			testSerializability(v, pkg1.ElemEncodingMyArrayMUS, u, t)
		})

		t.Run("We should be able to set Elem.Validator for MyArray", func(t *testing.T) {
			v := pkg1.ElemValidMyArray([3]int{1, 0, 3})
			testValidation(v, pkg1.ElemValidMyArrayMUS, testdata.ErrZeroValue, []int{3}, t)
		})

		t.Run("MyMap should be serializable", func(t *testing.T) {
			var (
				v                            = pkg1.MyMap(map[int]int{10: 1, 11: 2})
				u UnmarshallerFn[pkg1.MyMap] = func(bs []byte) (pkg1.MyMap, int, error) {
					ser := ord.NewMapSer[int, int](varint.Int, varint.Int)
					m, n, err := ser.Unmarshal(bs)
					return pkg1.MyMap(m), n, err
				}
			)
			testSerializability(v, pkg1.MyMapMUS, u, t)
		})

		t.Run("We should be able to set LenEncoding for MyMap", func(t *testing.T) {
			var (
				v                        = pkg1.LenEncodingMyMap(map[int]int{10: 1, 11: 2})
				lenU UnmarshallerFn[int] = func(bs []byte) (int, int, error) {
					num, n, err := raw.Int.Unmarshal(bs)
					return int(num), n, err
				}
			)
			testLenEncoding(v, pkg1.LenEncodingMyMapMUS, lenU, len(v), t)
		})

		t.Run("We should be able to set LenValidator for MyMap", func(t *testing.T) {
			v := pkg1.LenValidMyMap(map[int]int{10: 1, 11: 2})
			testValidation(v, pkg1.LenValidMyMapMUS, testdata.ErrTooLong, []int{1}, t)
		})

		t.Run("We should be able to set Key.Encoding for MyArray", func(t *testing.T) {
			var (
				v                                       = pkg1.KeyEncodingMyMap(map[int]int{10: 1, 11: 2})
				u UnmarshallerFn[pkg1.KeyEncodingMyMap] = func(bs []byte) (pkg1.KeyEncodingMyMap, int, error) {
					ser := ord.NewMapSer(raw.Int, varint.Int)
					m, n, err := ser.Unmarshal(bs)
					return pkg1.KeyEncodingMyMap(m), n, err
				}
			)
			testSerializability(v, pkg1.KeyEncodingMyMapMUS, u, t)
		})

		t.Run("We should be able to set Key.Validator for MyMap", func(t *testing.T) {
			v := pkg1.KeyValidMyMap(map[int]int{0: 1, 11: 2})
			testValidation(v, pkg1.KeyValidMyMapMUS, testdata.ErrZeroValue, []int{2, 4}, t)
		})

		t.Run("We should be able to set Elem.Encoding for MyArray", func(t *testing.T) {
			var (
				v                                        = pkg1.ElemEncodingMyMap(map[int]int{10: 1, 11: 2})
				u UnmarshallerFn[pkg1.ElemEncodingMyMap] = func(bs []byte) (pkg1.ElemEncodingMyMap, int, error) {
					ser := ord.NewMapSer(varint.Int, raw.Int)
					m, n, err := ser.Unmarshal(bs)
					return pkg1.ElemEncodingMyMap(m), n, err
				}
			)
			testSerializability(v, pkg1.ElemEncodingMyMapMUS, u, t)
		})

		t.Run("We should be able to set Elem.Validator for MyMap", func(t *testing.T) {
			v := pkg1.ElemValidMyMap(map[int]int{10: 0, 11: 2})
			testValidation(v, pkg1.ElemValidMyMapMUS, testdata.ErrZeroValue, []int{3, 5}, t)
		})

		// ptr

		t.Run("MyIntPtr should be serializable", func(t *testing.T) {
			var (
				n                               = 5
				v                               = pkg1.MyIntPtr(&n)
				u UnmarshallerFn[pkg1.MyIntPtr] = func(bs []byte) (pkg1.MyIntPtr, int, error) {
					ptr, n, err := ord.NewPtrSer(varint.Int).Unmarshal(bs)
					return pkg1.MyIntPtr(ptr), n, err
				}
			)
			testSerializability(v, pkg1.MyIntPtrMUS, u, t)
		})

		t.Run("We shoul be able to set Elem.NumEncoding for MyIntPtr", func(t *testing.T) {
			var (
				num                                              = 5
				v                                                = pkg1.ElemNumEncodingMyIntPtr(&num)
				u   UnmarshallerFn[pkg1.ElemNumEncodingMyIntPtr] = func(bs []byte) (pkg1.ElemNumEncodingMyIntPtr, int, error) {
					ptr, n, err := ord.NewPtrSer(raw.Int).Unmarshal(bs)
					return pkg1.ElemNumEncodingMyIntPtr(ptr), n, err
				}
			)
			testSerializability(v, pkg1.ElemNumEncodingMyIntPtrMUS, u, t)
		})

		t.Run("We should be able to serialize InterfaceDoubleMyIntPtr", func(t *testing.T) {
			var (
				in    pkg1.Interface               = pkg1.InterfaceImpl1{}
				inPtr                              = &in
				v     pkg1.InterfaceDoubleMyIntPtr = &inPtr
			)
			testSerializability(v, pkg1.InterfaceDoubleMyIntPtrMUS, nil, t)
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
			os.Setenv("TZ", "UTC")
			var (
				v = pkg2.ElemStruct{
					Slice: [][]int{{3}},
				}
				innerSer = ord.NewSliceSer(raw.Int)
				ser      = ord.NewValidSliceSer(innerSer,
					slops.WithLenValidator[[]int](com.ValidatorFn[int](testdata.ValidateLength1)))
				u UnmarshallerFn[pkg2.ElemStruct] = func(bs []byte) (v pkg2.ElemStruct, n int, err error) {
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
					n1, _ = raw.TimeUnixNanoUTC.Skip(bs[n:])
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

		t.Run("We should be able to serialize TimeStruct", func(t *testing.T) {
			v := pkg2.TimeStruct{
				Float32: 10.3,
				Time:    time.Unix(time.Now().Unix(), 0),
				String:  "abs",
			}
			testSerializability(v, pkg2.TimeStructMUS, nil, t)
		})

		// t.Run("We should be able to serialize struct using DTS", func(t *testing.T) {
		// 	testSerializability(pkg1.SimpleStruct{Int: 10}, pkg1.SimpleStructDTS, nil,
		// 		t)
		// })

		// t.Run("We should be able to serialize alias using DTS", func(t *testing.T) {
		// 	testSerializability(pkg1.MyInt(10), pkg1.MyIntDTS, nil, t)
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

func testSerializability[T any](v T, ser mus.Serializer[T], u UnmarshallerFn[T],
	t *testing.T) {
	opt := cmp.FilterValues(func(x1, x2 any) bool {
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
		_, ok1 := x1.(pkg1.TimeTypedef)
		_, ok2 := x2.(pkg1.TimeTypedef)
		if ok1 && ok2 {
			return true
		}
		return false
	}, cmp.Comparer(func(i1, i2 any) bool {
		_, ok1 := i1.(pkg1.InterfaceImpl1)
		_, ok2 := i2.(pkg1.InterfaceImpl1)
		if ok1 && ok2 {
			return true
		}
		var (
			tm1 pkg1.TimeTypedef
			tm2 pkg1.TimeTypedef
		)
		tm1, ok1 = i1.(pkg1.TimeTypedef)
		tm2, ok2 = i2.(pkg1.TimeTypedef)
		if ok1 && ok2 {
			return tm1.Time().Equal(tm2.Time())
		}
		return fmt.Sprint(i1) == fmt.Sprint(i2)
	}))
	equal := func(v, v1 T) bool {
		return cmp.Equal(v, v1, opt)
	}
	testGeneralSerializability(v, ser, u, equal, t)
}

func testGeneralSerializability[T any](v T, ser mus.Serializer[T],
	u UnmarshallerFn[T], equal func(v, v1 T) bool, t *testing.T) {
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
	if !equal(v, v1) {
		t.Errorf("unexpected value, want %v actual %v", v, v1)
	}
	if n1 != n2 {
		t.Errorf("unexpected n, want %v actual %v", n1, n2)
	}

	if u != nil {
		v1, n1, err = u(bs)
		if err != nil {
			t.Fatal(err)
		}
		if !equal(v, v1) {
			t.Errorf("unexpected value, want %v actual %v", v, v1)
		}
		if n1 != n2 {
			t.Errorf("unexpected n, want %v actual %v", n1, n2)
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

		Alias: pkg1.MySlice{1, 2, 3, 4},
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

		Map: map[float32]map[pkg1.MyInt][]pkg1.SimpleStruct{
			40.8: {
				pkg1.MyInt(11): {
					{Int: 30},
				},
			},
		},
	}
}
