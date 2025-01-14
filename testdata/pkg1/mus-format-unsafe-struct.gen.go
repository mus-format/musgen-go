// Code generated by musgen-go. DO NOT EDIT.

package pkg1

import (
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func MarshalUnsafeStructMUS(v Struct, bs []byte) (n int) {
	return unsafe.MarshalInt(v.Int, bs[n:])
}

func UnmarshalUnsafeStructMUS(bs []byte) (v Struct, n int, err error) {
	v.Int, n, err = unsafe.UnmarshalInt(bs[n:])
	return
}

func SizeUnsafeStructMUS(v Struct) (size int) {
	return unsafe.SizeInt(v.Int)
}

func SkipUnsafeStructMUS(bs []byte) (n int, err error) {
	return unsafe.SkipInt(bs[n:])
}

func MarshalUnsafeComplexStructMUS(v ComplexStruct, bs []byte) (n int) {
	n = unsafe.MarshalBool(v.Bool, bs[n:])
	n += unsafe.MarshalUint8(v.Byte, bs[n:])
	n += unsafe.MarshalInt8(v.Int8, bs[n:])
	n += unsafe.MarshalInt16(v.Int16, bs[n:])
	n += unsafe.MarshalInt32(v.Int32, bs[n:])
	n += unsafe.MarshalInt64(v.Int64, bs[n:])
	n += unsafe.MarshalUint8(v.Uint8, bs[n:])
	n += unsafe.MarshalUint16(v.Uint16, bs[n:])
	n += unsafe.MarshalUint32(v.Uint32, bs[n:])
	n += unsafe.MarshalUint64(v.Uint64, bs[n:])
	n += unsafe.MarshalFloat32(v.Float32, bs[n:])
	n += unsafe.MarshalFloat64(v.Float64, bs[n:])
	n += unsafe.MarshalString(v.String, nil, bs[n:])
	n += ord.MarshalPtr[string](v.PtrString, mus.MarshallerFn[string](func(t string, bs []byte) (n int) { return unsafe.MarshalString(t, nil, bs[n:]) }), bs[n:])
	n += MarshalUnsafeSliceAliasMUS(v.Alias, bs[n:])
	n += ord.MarshalPtr[Struct](v.Ptr, mus.MarshallerFn[Struct](MarshalUnsafeStructMUS), bs[n:])
	n += ord.MarshalPtr[string](v.NilPtr, mus.MarshallerFn[string](func(t string, bs []byte) (n int) { return unsafe.MarshalString(t, nil, bs[n:]) }), bs[n:])
	n += pkg2.MarshalUnsafeStructMUS(v.AnotherPkgStruct, bs[n:])
	n += MarshalUnsafeInterfaceMUS(v.Interface, bs[n:])
	n += unsafe.MarshalByteSlice(v.SliceByte,
		nil,
		bs[n:])
	n += ord.MarshalSlice[Struct](v.SliceStruct,
		nil,
		mus.MarshallerFn[Struct](MarshalUnsafeStructMUS),
		bs[n:])
	n += ord.MarshalSlice[int](v.Array[:],
		nil,
		mus.MarshallerFn[int](unsafe.MarshalInt),
		bs[n:])
	n += ord.MarshalPtr[[3]int](v.PtrArray, mus.MarshallerFn[[3]int](func(t [3]int, bs []byte) (n int) {
		return ord.MarshalSlice[int](t[:],
			nil,
			mus.MarshallerFn[int](unsafe.MarshalInt),
			bs[n:])
	}), bs[n:])
	return n + ord.MarshalMap[float32, map[IntAlias][]Struct](v.Map, nil,
		mus.MarshallerFn[float32](unsafe.MarshalFloat32),
		mus.MarshallerFn[map[IntAlias][]Struct](func(t map[IntAlias][]Struct, bs []byte) (n int) {
			return ord.MarshalMap[IntAlias, []Struct](t, nil,
				mus.MarshallerFn[IntAlias](MarshalUnsafeIntAliasMUS),
				mus.MarshallerFn[[]Struct](func(t []Struct, bs []byte) (n int) {
					return ord.MarshalSlice[Struct](t,
						nil,
						mus.MarshallerFn[Struct](MarshalUnsafeStructMUS),
						bs[n:])
				}),
				bs[n:])
		}),
		bs[n:])
}

func UnmarshalUnsafeComplexStructMUS(bs []byte) (v ComplexStruct, n int, err error) {
	v.Bool, n, err = unsafe.UnmarshalBool(bs[n:])
	if err != nil {
		return
	}
	var n1 int
	v.Byte, n1, err = unsafe.UnmarshalUint8(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int8, n1, err = unsafe.UnmarshalInt8(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int16, n1, err = unsafe.UnmarshalInt16(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int32, n1, err = unsafe.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int64, n1, err = unsafe.UnmarshalInt64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint8, n1, err = unsafe.UnmarshalUint8(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint16, n1, err = unsafe.UnmarshalUint16(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint32, n1, err = unsafe.UnmarshalUint32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint64, n1, err = unsafe.UnmarshalUint64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Float32, n1, err = unsafe.UnmarshalFloat32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Float64, n1, err = unsafe.UnmarshalFloat64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.String, n1, err = unsafe.UnmarshalString(nil, bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.PtrString, n1, err = ord.UnmarshalPtr[string](mus.UnmarshallerFn[string](func(bs []byte) (t string, n int, err error) { return unsafe.UnmarshalString(nil, bs[n:]) }), bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Alias, n1, err = UnmarshalUnsafeSliceAliasMUS(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Ptr, n1, err = ord.UnmarshalPtr[Struct](mus.UnmarshallerFn[Struct](UnmarshalUnsafeStructMUS), bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.NilPtr, n1, err = ord.UnmarshalPtr[string](mus.UnmarshallerFn[string](func(bs []byte) (t string, n int, err error) { return unsafe.UnmarshalString(nil, bs[n:]) }), bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.AnotherPkgStruct, n1, err = pkg2.UnmarshalUnsafeStructMUS(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Interface, n1, err = UnmarshalUnsafeInterfaceMUS(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.SliceByte, n1, err = unsafe.UnmarshalByteSlice(nil,
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.SliceStruct, n1, err = ord.UnmarshalSlice[Struct](nil,
		mus.UnmarshallerFn[Struct](UnmarshalUnsafeStructMUS),
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	vArray, n1, err := ord.UnmarshalSlice[int](nil,
		mus.UnmarshallerFn[int](unsafe.UnmarshalInt),
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Array = ([3]int)(vArray)
	v.PtrArray, n1, err = ord.UnmarshalPtr[[3]int](mus.UnmarshallerFn[[3]int](func(bs []byte) (t [3]int, n int, err error) {
		ta, n, err := ord.UnmarshalSlice[int](nil,
			mus.UnmarshallerFn[int](unsafe.UnmarshalInt),
			bs[n:])
		if err != nil {
			return
		}
		t = ([3]int)(ta)
		return
	}), bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Map, n1, err = ord.UnmarshalMap[float32, map[IntAlias][]Struct](nil,
		mus.UnmarshallerFn[float32](unsafe.UnmarshalFloat32),
		mus.UnmarshallerFn[map[IntAlias][]Struct](func(bs []byte) (t map[IntAlias][]Struct, n int, err error) {
			return ord.UnmarshalMap[IntAlias, []Struct](nil,
				mus.UnmarshallerFn[IntAlias](UnmarshalUnsafeIntAliasMUS),
				mus.UnmarshallerFn[[]Struct](func(bs []byte) (t []Struct, n int, err error) {
					return ord.UnmarshalSlice[Struct](nil,
						mus.UnmarshallerFn[Struct](UnmarshalUnsafeStructMUS),
						bs[n:])
				}),
				bs[n:])
		}),
		bs[n:])
	n += n1
	return
}

func SizeUnsafeComplexStructMUS(v ComplexStruct) (size int) {
	size = unsafe.SizeBool(v.Bool)
	size += unsafe.SizeUint8(v.Byte)
	size += unsafe.SizeInt8(v.Int8)
	size += unsafe.SizeInt16(v.Int16)
	size += unsafe.SizeInt32(v.Int32)
	size += unsafe.SizeInt64(v.Int64)
	size += unsafe.SizeUint8(v.Uint8)
	size += unsafe.SizeUint16(v.Uint16)
	size += unsafe.SizeUint32(v.Uint32)
	size += unsafe.SizeUint64(v.Uint64)
	size += unsafe.SizeFloat32(v.Float32)
	size += unsafe.SizeFloat64(v.Float64)
	size += unsafe.SizeString(v.String, nil)
	size += ord.SizePtr[string](v.PtrString, mus.SizerFn[string](func(t string) (size int) { return unsafe.SizeString(t, nil) }))
	size += SizeUnsafeSliceAliasMUS(v.Alias)
	size += ord.SizePtr[Struct](v.Ptr, mus.SizerFn[Struct](SizeUnsafeStructMUS))
	size += ord.SizePtr[string](v.NilPtr, mus.SizerFn[string](func(t string) (size int) { return unsafe.SizeString(t, nil) }))
	size += pkg2.SizeUnsafeStructMUS(v.AnotherPkgStruct)
	size += SizeUnsafeInterfaceMUS(v.Interface)
	size += unsafe.SizeByteSlice(v.SliceByte,
		nil)
	size += ord.SizeSlice[Struct](v.SliceStruct,
		nil,
		mus.SizerFn[Struct](SizeUnsafeStructMUS))
	size += ord.SizeSlice[int](v.Array[:],
		nil,
		mus.SizerFn[int](unsafe.SizeInt))
	size += ord.SizePtr[[3]int](v.PtrArray, mus.SizerFn[[3]int](func(t [3]int) (size int) {
		return ord.SizeSlice[int](t[:],
			nil,
			mus.SizerFn[int](unsafe.SizeInt))
	}))
	return size + ord.SizeMap[float32, map[IntAlias][]Struct](v.Map, nil,
		mus.SizerFn[float32](unsafe.SizeFloat32),
		mus.SizerFn[map[IntAlias][]Struct](func(t map[IntAlias][]Struct) (size int) {
			return ord.SizeMap[IntAlias, []Struct](t, nil,
				mus.SizerFn[IntAlias](SizeUnsafeIntAliasMUS),
				mus.SizerFn[[]Struct](func(t []Struct) (size int) {
					return ord.SizeSlice[Struct](t,
						nil,
						mus.SizerFn[Struct](SizeUnsafeStructMUS))
				}))
		}))
}

func SkipUnsafeComplexStructMUS(bs []byte) (n int, err error) {
	n, err = unsafe.SkipBool(bs[n:])
	if err != nil {
		return
	}
	var n1 int
	n1, err = unsafe.SkipUint8(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipInt8(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipInt16(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipInt64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipUint8(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipUint16(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipUint32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipUint64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipFloat32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipFloat64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipString(nil, bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipPtr(mus.SkipperFn(func(bs []byte) (n int, err error) { return unsafe.SkipString(nil, bs[n:]) }), bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipUnsafeSliceAliasMUS(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipPtr(mus.SkipperFn(SkipUnsafeStructMUS), bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipPtr(mus.SkipperFn(func(bs []byte) (n int, err error) { return unsafe.SkipString(nil, bs[n:]) }), bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = pkg2.SkipUnsafeStructMUS(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipUnsafeInterfaceMUS(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.SkipByteSlice(nil,
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipSlice(nil,
		mus.SkipperFn(SkipUnsafeStructMUS),
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipSlice(nil,
		mus.SkipperFn(unsafe.SkipInt),
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipPtr(mus.SkipperFn(func(bs []byte) (n int, err error) {
		return ord.SkipSlice(nil,
			mus.SkipperFn(unsafe.SkipInt),
			bs[n:])
	}), bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipMap(nil,
		mus.SkipperFn(unsafe.SkipFloat32),
		mus.SkipperFn(func(bs []byte) (n int, err error) {
			return ord.SkipMap(nil,
				mus.SkipperFn(SkipUnsafeIntAliasMUS),
				mus.SkipperFn(func(bs []byte) (n int, err error) {
					return ord.SkipSlice(nil,
						mus.SkipperFn(SkipUnsafeStructMUS),
						bs[n:])
				}),
				bs[n:])
		}),
		bs[n:])
	n += n1
	return
}
