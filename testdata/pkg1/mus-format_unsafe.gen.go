// Code generated by musgen-go. DO NOT EDIT.

package pkg1

import (
	"fmt"

	dts "github.com/mus-format/mus-dts-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/musgen-go/testdata/pkg2"
)

var (
	arrayVVcFIJPvOaΣLVKMi652G6wΞΞ = ord.NewArraySer[[3]int, int](3, unsafe.Int)
	mapVVXjSxO6dXjΔOCΔqLsJO6QΞΞ   = ord.NewMapSer[float32, map[MyIntUnsafe][]SimpleStructUnsafe](unsafe.Float32, mapΣWU7IqDΔwΔv4DQ00uoROsQΞΞ)
	mapΣWU7IqDΔwΔv4DQ00uoROsQΞΞ   = ord.NewMapSer[MyIntUnsafe, []SimpleStructUnsafe](MyIntUnsafeMUS, sliceHE51HQEviTCic5s1ePPTRgΞΞ)
	ptrENsHnYOTUSTMKYsaG5TfAAΞΞ   = ord.NewPtrSer[[3]int](arrayVVcFIJPvOaΣLVKMi652G6wΞΞ)
	ptroNwYbYaFs3k7ΔtMmbsaP5AΞΞ   = ord.NewPtrSer[string](unsafe.String)
	ptruJYTHhBwQB9O98OsΔGQYcgΞΞ   = ord.NewPtrSer[SimpleStructUnsafe](SimpleStructUnsafeMUS)
	sliceHE51HQEviTCic5s1ePPTRgΞΞ = ord.NewSliceSer[SimpleStructUnsafe](SimpleStructUnsafeMUS)
	sliceR279MBVLlXHpndvGzXvzEQΞΞ = ord.NewSliceSer[int](unsafe.Int)
)

var MyIntUnsafeMUS = myIntUnsafeMUS{}

type myIntUnsafeMUS struct{}

func (s myIntUnsafeMUS) Marshal(v MyIntUnsafe, bs []byte) (n int) {
	return unsafe.Int.Marshal(int(v), bs)
}

func (s myIntUnsafeMUS) Unmarshal(bs []byte) (v MyIntUnsafe, n int, err error) {
	tmp, n, err := unsafe.Int.Unmarshal(bs)
	if err != nil {
		return
	}
	v = MyIntUnsafe(tmp)
	return
}

func (s myIntUnsafeMUS) Size(v MyIntUnsafe) (size int) {
	return unsafe.Int.Size(int(v))
}

func (s myIntUnsafeMUS) Skip(bs []byte) (n int, err error) {
	return unsafe.Int.Skip(bs)
}

var MySliceUnsafeMUS = mySliceUnsafeMUS{}

type mySliceUnsafeMUS struct{}

func (s mySliceUnsafeMUS) Marshal(v MySliceUnsafe, bs []byte) (n int) {
	return sliceR279MBVLlXHpndvGzXvzEQΞΞ.Marshal([]int(v), bs)
}

func (s mySliceUnsafeMUS) Unmarshal(bs []byte) (v MySliceUnsafe, n int, err error) {
	tmp, n, err := sliceR279MBVLlXHpndvGzXvzEQΞΞ.Unmarshal(bs)
	if err != nil {
		return
	}
	v = MySliceUnsafe(tmp)
	return
}

func (s mySliceUnsafeMUS) Size(v MySliceUnsafe) (size int) {
	return sliceR279MBVLlXHpndvGzXvzEQΞΞ.Size([]int(v))
}

func (s mySliceUnsafeMUS) Skip(bs []byte) (n int, err error) {
	return sliceR279MBVLlXHpndvGzXvzEQΞΞ.Skip(bs)
}

var SimpleStructUnsafeMUS = simpleStructUnsafeMUS{}

type simpleStructUnsafeMUS struct{}

func (s simpleStructUnsafeMUS) Marshal(v SimpleStructUnsafe, bs []byte) (n int) {
	return unsafe.Int.Marshal(v.Int, bs)
}

func (s simpleStructUnsafeMUS) Unmarshal(bs []byte) (v SimpleStructUnsafe, n int, err error) {
	v.Int, n, err = unsafe.Int.Unmarshal(bs)
	return
}

func (s simpleStructUnsafeMUS) Size(v SimpleStructUnsafe) (size int) {
	return unsafe.Int.Size(v.Int)
}

func (s simpleStructUnsafeMUS) Skip(bs []byte) (n int, err error) {
	n, err = unsafe.Int.Skip(bs)
	return
}

var ComplexStructUnsafeMUS = complexStructUnsafeMUS{}

type complexStructUnsafeMUS struct{}

func (s complexStructUnsafeMUS) Marshal(v ComplexStructUnsafe, bs []byte) (n int) {
	n = unsafe.Bool.Marshal(v.Bool, bs)
	n += unsafe.Uint8.Marshal(v.Byte, bs[n:])
	n += unsafe.Int8.Marshal(v.Int8, bs[n:])
	n += unsafe.Int16.Marshal(v.Int16, bs[n:])
	n += unsafe.Int32.Marshal(v.Int32, bs[n:])
	n += unsafe.Int64.Marshal(v.Int64, bs[n:])
	n += unsafe.Uint8.Marshal(v.Uint8, bs[n:])
	n += unsafe.Uint16.Marshal(v.Uint16, bs[n:])
	n += unsafe.Uint32.Marshal(v.Uint32, bs[n:])
	n += unsafe.Uint64.Marshal(v.Uint64, bs[n:])
	n += unsafe.Float32.Marshal(v.Float32, bs[n:])
	n += unsafe.Float64.Marshal(v.Float64, bs[n:])
	n += unsafe.String.Marshal(v.String, bs[n:])
	n += MySliceUnsafeMUS.Marshal(v.Alias, bs[n:])
	n += pkg2.StructUnsafeMUS.Marshal(v.AnotherPkgStruct, bs[n:])
	n += InterfaceUnsafeMUS.Marshal(v.Interface, bs[n:])
	n += unsafe.ByteSlice.Marshal(v.ByteSlice, bs[n:])
	n += sliceHE51HQEviTCic5s1ePPTRgΞΞ.Marshal(v.StructSlice, bs[n:])
	n += arrayVVcFIJPvOaΣLVKMi652G6wΞΞ.Marshal(v.Array, bs[n:])
	n += ptroNwYbYaFs3k7ΔtMmbsaP5AΞΞ.Marshal(v.PtrString, bs[n:])
	n += ptruJYTHhBwQB9O98OsΔGQYcgΞΞ.Marshal(v.PtrStruct, bs[n:])
	n += ptroNwYbYaFs3k7ΔtMmbsaP5AΞΞ.Marshal(v.NilPtr, bs[n:])
	n += ptrENsHnYOTUSTMKYsaG5TfAAΞΞ.Marshal(v.PtrArray, bs[n:])
	return n + mapVVXjSxO6dXjΔOCΔqLsJO6QΞΞ.Marshal(v.Map, bs[n:])
}

func (s complexStructUnsafeMUS) Unmarshal(bs []byte) (v ComplexStructUnsafe, n int, err error) {
	v.Bool, n, err = unsafe.Bool.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	v.Byte, n1, err = unsafe.Uint8.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int8, n1, err = unsafe.Int8.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int16, n1, err = unsafe.Int16.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int32, n1, err = unsafe.Int32.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int64, n1, err = unsafe.Int64.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint8, n1, err = unsafe.Uint8.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint16, n1, err = unsafe.Uint16.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint32, n1, err = unsafe.Uint32.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint64, n1, err = unsafe.Uint64.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Float32, n1, err = unsafe.Float32.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Float64, n1, err = unsafe.Float64.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.String, n1, err = unsafe.String.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Alias, n1, err = MySliceUnsafeMUS.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.AnotherPkgStruct, n1, err = pkg2.StructUnsafeMUS.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Interface, n1, err = InterfaceUnsafeMUS.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.ByteSlice, n1, err = unsafe.ByteSlice.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.StructSlice, n1, err = sliceHE51HQEviTCic5s1ePPTRgΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Array, n1, err = arrayVVcFIJPvOaΣLVKMi652G6wΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.PtrString, n1, err = ptroNwYbYaFs3k7ΔtMmbsaP5AΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.PtrStruct, n1, err = ptruJYTHhBwQB9O98OsΔGQYcgΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.NilPtr, n1, err = ptroNwYbYaFs3k7ΔtMmbsaP5AΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.PtrArray, n1, err = ptrENsHnYOTUSTMKYsaG5TfAAΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Map, n1, err = mapVVXjSxO6dXjΔOCΔqLsJO6QΞΞ.Unmarshal(bs[n:])
	n += n1
	return
}

func (s complexStructUnsafeMUS) Size(v ComplexStructUnsafe) (size int) {
	size = unsafe.Bool.Size(v.Bool)
	size += unsafe.Uint8.Size(v.Byte)
	size += unsafe.Int8.Size(v.Int8)
	size += unsafe.Int16.Size(v.Int16)
	size += unsafe.Int32.Size(v.Int32)
	size += unsafe.Int64.Size(v.Int64)
	size += unsafe.Uint8.Size(v.Uint8)
	size += unsafe.Uint16.Size(v.Uint16)
	size += unsafe.Uint32.Size(v.Uint32)
	size += unsafe.Uint64.Size(v.Uint64)
	size += unsafe.Float32.Size(v.Float32)
	size += unsafe.Float64.Size(v.Float64)
	size += unsafe.String.Size(v.String)
	size += MySliceUnsafeMUS.Size(v.Alias)
	size += pkg2.StructUnsafeMUS.Size(v.AnotherPkgStruct)
	size += InterfaceUnsafeMUS.Size(v.Interface)
	size += unsafe.ByteSlice.Size(v.ByteSlice)
	size += sliceHE51HQEviTCic5s1ePPTRgΞΞ.Size(v.StructSlice)
	size += arrayVVcFIJPvOaΣLVKMi652G6wΞΞ.Size(v.Array)
	size += ptroNwYbYaFs3k7ΔtMmbsaP5AΞΞ.Size(v.PtrString)
	size += ptruJYTHhBwQB9O98OsΔGQYcgΞΞ.Size(v.PtrStruct)
	size += ptroNwYbYaFs3k7ΔtMmbsaP5AΞΞ.Size(v.NilPtr)
	size += ptrENsHnYOTUSTMKYsaG5TfAAΞΞ.Size(v.PtrArray)
	return size + mapVVXjSxO6dXjΔOCΔqLsJO6QΞΞ.Size(v.Map)
}

func (s complexStructUnsafeMUS) Skip(bs []byte) (n int, err error) {
	n, err = unsafe.Bool.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = unsafe.Uint8.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Int8.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Int16.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Int32.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Int64.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Uint8.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Uint16.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Uint32.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Uint64.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Float32.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.Float64.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.String.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = MySliceUnsafeMUS.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = pkg2.StructUnsafeMUS.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = InterfaceUnsafeMUS.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = unsafe.ByteSlice.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = sliceHE51HQEviTCic5s1ePPTRgΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = arrayVVcFIJPvOaΣLVKMi652G6wΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ptroNwYbYaFs3k7ΔtMmbsaP5AΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ptruJYTHhBwQB9O98OsΔGQYcgΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ptroNwYbYaFs3k7ΔtMmbsaP5AΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ptrENsHnYOTUSTMKYsaG5TfAAΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = mapVVXjSxO6dXjΔOCΔqLsJO6QΞΞ.Skip(bs[n:])
	n += n1
	return
}

var InterfaceImpl1UnsafeMUS = interfaceImpl1UnsafeMUS{}

type interfaceImpl1UnsafeMUS struct{}

func (s interfaceImpl1UnsafeMUS) Marshal(v InterfaceImpl1Unsafe, bs []byte) (n int) {
	return
}

func (s interfaceImpl1UnsafeMUS) Unmarshal(bs []byte) (v InterfaceImpl1Unsafe, n int, err error) {
	return
}

func (s interfaceImpl1UnsafeMUS) Size(v InterfaceImpl1Unsafe) (size int) {
	return
}

func (s interfaceImpl1UnsafeMUS) Skip(bs []byte) (n int, err error) {
	return
}

var InterfaceImpl2UnsafeMUS = interfaceImpl2UnsafeMUS{}

type interfaceImpl2UnsafeMUS struct{}

func (s interfaceImpl2UnsafeMUS) Marshal(v InterfaceImpl2Unsafe, bs []byte) (n int) {
	return unsafe.Int.Marshal(int(v), bs)
}

func (s interfaceImpl2UnsafeMUS) Unmarshal(bs []byte) (v InterfaceImpl2Unsafe, n int, err error) {
	tmp, n, err := unsafe.Int.Unmarshal(bs)
	if err != nil {
		return
	}
	v = InterfaceImpl2Unsafe(tmp)
	return
}

func (s interfaceImpl2UnsafeMUS) Size(v InterfaceImpl2Unsafe) (size int) {
	return unsafe.Int.Size(int(v))
}

func (s interfaceImpl2UnsafeMUS) Skip(bs []byte) (n int, err error) {
	return unsafe.Int.Skip(bs)
}

var InterfaceImpl1UnsafeDTS = dts.New[InterfaceImpl1Unsafe](InterfaceImpl1UnsafeDTM, InterfaceImpl1UnsafeMUS)
var InterfaceImpl2UnsafeDTS = dts.New[InterfaceImpl2Unsafe](InterfaceImpl2UnsafeDTM, InterfaceImpl2UnsafeMUS)

var InterfaceUnsafeMUS = interfaceUnsafeMUS{}

type interfaceUnsafeMUS struct{}

func (s interfaceUnsafeMUS) Marshal(v InterfaceUnsafe, bs []byte) (n int) {
	switch t := v.(type) {
	case InterfaceImpl1Unsafe:
		return InterfaceImpl1UnsafeDTS.Marshal(t, bs)
	case InterfaceImpl2Unsafe:
		return InterfaceImpl2UnsafeDTS.Marshal(t, bs)
	default:
		panic(fmt.Sprintf("unexpected %v type", t))
	}
}

func (s interfaceUnsafeMUS) Unmarshal(bs []byte) (v InterfaceUnsafe, n int, err error) {
	dtm, n, err := dts.DTMSer.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	switch dtm {
	case InterfaceImpl1UnsafeDTM:
		v, n1, err = InterfaceImpl1UnsafeDTS.UnmarshalData(bs[n:])
	case InterfaceImpl2UnsafeDTM:
		v, n1, err = InterfaceImpl2UnsafeDTS.UnmarshalData(bs[n:])
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
	}
	n += n1
	return
}

func (s interfaceUnsafeMUS) Size(v InterfaceUnsafe) (size int) {
	switch t := v.(type) {
	case InterfaceImpl1Unsafe:
		return InterfaceImpl1UnsafeDTS.Size(t)
	case InterfaceImpl2Unsafe:
		return InterfaceImpl2UnsafeDTS.Size(t)
	default:
		panic(fmt.Sprintf("unexpected %v type", t))
	}
}

func (s interfaceUnsafeMUS) Skip(bs []byte) (n int, err error) {
	dtm, n, err := dts.DTMSer.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	switch dtm {
	case InterfaceImpl1UnsafeDTM:
		n1, err = InterfaceImpl1UnsafeDTS.SkipData(bs[n:])
	case InterfaceImpl2UnsafeDTM:
		n1, err = InterfaceImpl2UnsafeDTS.SkipData(bs[n:])
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
	}
	n += n1
	return
}
