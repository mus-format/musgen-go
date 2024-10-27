// Code generated by musgen-go. DO NOT EDIT.\n\npackage

package pkg1

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/varint"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

func MarshalStreamStructMUS(v Struct, w muss.Writer) (n int, err error) {
	return varint.MarshalInt(v.Int, w)
}

func UnmarshalStreamStructMUS(r muss.Reader) (v Struct, n int, err error) {
	v.Int, n, err = varint.UnmarshalInt(r)
	return
}

func SizeStreamStructMUS(v Struct) (size int) {
	return varint.SizeInt(v.Int)
}

func SkipStreamStructMUS(r muss.Reader) (n int, err error) {
	return varint.SkipInt(r)
}

func MarshalStreamComplexStructMUS(v ComplexStruct, w muss.Writer) (n int, err error) {
	n, err = ord.MarshalBool(v.Bool, w)
	if err != nil {
		return
	}
	var n1 int
	n1, err = varint.MarshalUint8(v.Byte, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalInt8(v.Int8, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalInt16(v.Int16, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalInt32(v.Int32, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalInt64(v.Int64, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalUint8(v.Uint8, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalUint16(v.Uint16, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalUint32(v.Uint32, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalUint64(v.Uint64, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalFloat32(v.Float32, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.MarshalFloat64(v.Float64, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.MarshalString(v.String, nil, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = MarshalStreamSliceAliasMUS(v.Alias, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.MarshalPtr[Struct](v.Ptr, muss.MarshallerFn[Struct](MarshalStreamStructMUS), w)
	n += n1
	if err != nil {
		return
	}
	n1, err = pkg2.MarshalStreamStructMUS(v.AnotherPkgStruct, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = MarshalStreamInterfaceMUS(v.Interface, w)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.MarshalSlice[uint8](v.SliceByte,
		nil,
		muss.MarshallerFn[uint8](varint.MarshalUint8),
		w)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.MarshalSlice[Struct](v.SliceStruct,
		nil,
		muss.MarshallerFn[Struct](MarshalStreamStructMUS),
		w)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.MarshalMap[float32, map[IntAlias][]Struct](v.Map, nil,
		muss.MarshallerFn[float32](varint.MarshalFloat32),
		muss.MarshallerFn[map[IntAlias][]Struct](func(t map[IntAlias][]Struct, w muss.Writer) (n int, err error) {
			return ord.MarshalMap[IntAlias, []Struct](t, nil,
				muss.MarshallerFn[IntAlias](MarshalStreamIntAliasMUS),
				muss.MarshallerFn[[]Struct](func(t []Struct, w muss.Writer) (n int, err error) {
					return ord.MarshalSlice[Struct](t,
						nil,
						muss.MarshallerFn[Struct](MarshalStreamStructMUS),
						w)
				}),
				w)
		}),
		w)
	n += n1
	return
}

func UnmarshalStreamComplexStructMUS(r muss.Reader) (v ComplexStruct, n int, err error) {
	v.Bool, n, err = ord.UnmarshalBool(r)
	if err != nil {
		return
	}
	var n1 int
	v.Byte, n1, err = varint.UnmarshalUint8(r)
	n += n1
	if err != nil {
		return
	}
	v.Int8, n1, err = varint.UnmarshalInt8(r)
	n += n1
	if err != nil {
		return
	}
	v.Int16, n1, err = varint.UnmarshalInt16(r)
	n += n1
	if err != nil {
		return
	}
	v.Int32, n1, err = varint.UnmarshalInt32(r)
	n += n1
	if err != nil {
		return
	}
	v.Int64, n1, err = varint.UnmarshalInt64(r)
	n += n1
	if err != nil {
		return
	}
	v.Uint8, n1, err = varint.UnmarshalUint8(r)
	n += n1
	if err != nil {
		return
	}
	v.Uint16, n1, err = varint.UnmarshalUint16(r)
	n += n1
	if err != nil {
		return
	}
	v.Uint32, n1, err = varint.UnmarshalUint32(r)
	n += n1
	if err != nil {
		return
	}
	v.Uint64, n1, err = varint.UnmarshalUint64(r)
	n += n1
	if err != nil {
		return
	}
	v.Float32, n1, err = varint.UnmarshalFloat32(r)
	n += n1
	if err != nil {
		return
	}
	v.Float64, n1, err = varint.UnmarshalFloat64(r)
	n += n1
	if err != nil {
		return
	}
	v.String, n1, err = ord.UnmarshalString(nil, r)
	n += n1
	if err != nil {
		return
	}
	v.Alias, n1, err = UnmarshalStreamSliceAliasMUS(r)
	n += n1
	if err != nil {
		return
	}
	v.Ptr, n1, err = ord.UnmarshalPtr[Struct](muss.UnmarshallerFn[Struct](UnmarshalStreamStructMUS), r)
	n += n1
	if err != nil {
		return
	}
	v.AnotherPkgStruct, n1, err = pkg2.UnmarshalStreamStructMUS(r)
	n += n1
	if err != nil {
		return
	}
	v.Interface, n1, err = UnmarshalStreamInterfaceMUS(r)
	n += n1
	if err != nil {
		return
	}
	v.SliceByte, n1, err = ord.UnmarshalSlice[uint8](nil,
		muss.UnmarshallerFn[uint8](varint.UnmarshalUint8),
		r)
	n += n1
	if err != nil {
		return
	}
	v.SliceStruct, n1, err = ord.UnmarshalSlice[Struct](nil,
		muss.UnmarshallerFn[Struct](UnmarshalStreamStructMUS),
		r)
	n += n1
	if err != nil {
		return
	}
	v.Map, n1, err = ord.UnmarshalMap[float32, map[IntAlias][]Struct](nil,
		muss.UnmarshallerFn[float32](varint.UnmarshalFloat32),
		muss.UnmarshallerFn[map[IntAlias][]Struct](func(r muss.Reader) (t map[IntAlias][]Struct, n int, err error) {
			return ord.UnmarshalMap[IntAlias, []Struct](nil,
				muss.UnmarshallerFn[IntAlias](UnmarshalStreamIntAliasMUS),
				muss.UnmarshallerFn[[]Struct](func(r muss.Reader) (t []Struct, n int, err error) {
					return ord.UnmarshalSlice[Struct](nil,
						muss.UnmarshallerFn[Struct](UnmarshalStreamStructMUS),
						r)
				}),
				r)
		}),
		r)
	n += n1
	return
}

func SizeStreamComplexStructMUS(v ComplexStruct) (size int) {
	size = ord.SizeBool(v.Bool)
	size += varint.SizeUint8(v.Byte)
	size += varint.SizeInt8(v.Int8)
	size += varint.SizeInt16(v.Int16)
	size += varint.SizeInt32(v.Int32)
	size += varint.SizeInt64(v.Int64)
	size += varint.SizeUint8(v.Uint8)
	size += varint.SizeUint16(v.Uint16)
	size += varint.SizeUint32(v.Uint32)
	size += varint.SizeUint64(v.Uint64)
	size += varint.SizeFloat32(v.Float32)
	size += varint.SizeFloat64(v.Float64)
	size += ord.SizeString(v.String, nil)
	size += SizeStreamSliceAliasMUS(v.Alias)
	size += ord.SizePtr[Struct](v.Ptr, muss.SizerFn[Struct](SizeStreamStructMUS))
	size += pkg2.SizeStreamStructMUS(v.AnotherPkgStruct)
	size += SizeStreamInterfaceMUS(v.Interface)
	size += ord.SizeSlice[uint8](v.SliceByte,
		nil,
		muss.SizerFn[uint8](varint.SizeUint8))
	size += ord.SizeSlice[Struct](v.SliceStruct,
		nil,
		muss.SizerFn[Struct](SizeStreamStructMUS))
	return size + ord.SizeMap[float32, map[IntAlias][]Struct](v.Map, nil,
		muss.SizerFn[float32](varint.SizeFloat32),
		muss.SizerFn[map[IntAlias][]Struct](func(t map[IntAlias][]Struct) (size int) {
			return ord.SizeMap[IntAlias, []Struct](t, nil,
				muss.SizerFn[IntAlias](SizeStreamIntAliasMUS),
				muss.SizerFn[[]Struct](func(t []Struct) (size int) {
					return ord.SizeSlice[Struct](t,
						nil,
						muss.SizerFn[Struct](SizeStreamStructMUS))
				}))
		}))
}

func SkipStreamComplexStructMUS(r muss.Reader) (n int, err error) {
	n, err = ord.SkipBool(r)
	if err != nil {
		return
	}
	var n1 int
	n1, err = varint.SkipUint8(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipInt8(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipInt16(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipInt32(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipInt64(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipUint8(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipUint16(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipUint32(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipUint64(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipFloat32(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.SkipFloat64(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipString(nil, r)
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipStreamSliceAliasMUS(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipPtr(muss.SkipperFn(SkipStreamStructMUS), r)
	n += n1
	if err != nil {
		return
	}
	n1, err = pkg2.SkipStreamStructMUS(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipStreamInterfaceMUS(r)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipSlice(nil,
		muss.SkipperFn(varint.SkipUint8),
		r)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipSlice(nil,
		muss.SkipperFn(SkipStreamStructMUS),
		r)
	n += n1
	if err != nil {
		return
	}
	n1, err = ord.SkipMap(nil,
		muss.SkipperFn(varint.SkipFloat32),
		muss.SkipperFn(func(r muss.Reader) (n int, err error) {
			return ord.SkipMap(nil,
				muss.SkipperFn(SkipStreamIntAliasMUS),
				muss.SkipperFn(func(r muss.Reader) (n int, err error) {
					return ord.SkipSlice(nil,
						muss.SkipperFn(SkipStreamStructMUS),
						r)
				}),
				r)
		}),
		r)
	n += n1
	return
}
