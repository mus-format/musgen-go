// Code generated by musgen-go. DO NOT EDIT.

package testdata

import (
	"fmt"

	com "github.com/mus-format/common-go"
	dts "github.com/mus-format/dts-go"
	arrops "github.com/mus-format/mus-go/options/array"
	bslops "github.com/mus-format/mus-go/options/byte_slice"
	mapops "github.com/mus-format/mus-go/options/map"
	slops "github.com/mus-format/mus-go/options/slice"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/varint"
	common_testdata "github.com/mus-format/musgen-go/testdata"
)

var (
	arrayX1ltdZJBPzH2l9ΔCPTeroAΞΞ     = ord.NewValidArraySer[[3]int, int](varint.Int, arrops.WithLenSer[int](raw.Int), arrops.WithElemValidator[int](com.ValidatorFn[int](common_testdata.ValidateZeroValue[int])))
	arrayoa64Y8KzCJ7ziIbYybTebAΞΞ     = ord.NewValidArraySer[[3]int, int](raw.Int, arrops.WithElemValidator[int](com.ValidatorFn[int](common_testdata.ValidateZeroValue[int])))
	byteSlicejdpLY58ftUFXYGnakpvXgAΞΞ = ord.NewValidByteSliceSer(bslops.WithLenSer(varint.Int), bslops.WithLenValidator(com.ValidatorFn[int](common_testdata.ValidateLength20)))
	map1lW2X1YW1lK3PO8YRU2TeQΞΞ       = ord.NewValidMapSer[float32, map[MyInt][]MyStruct](varint.Float32, mapu7h3FOrvGEtyYbEzqUdSTwΞΞ, mapops.WithLenValidator[float32, map[MyInt][]MyStruct](com.ValidatorFn[int](common_testdata.ValidateLength20)), mapops.WithKeyValidator[float32, map[MyInt][]MyStruct](com.ValidatorFn[float32](common_testdata.ValidateZeroValue[float32])))
	mapu7h3FOrvGEtyYbEzqUdSTwΞΞ       = ord.NewValidMapSer[MyInt, []MyStruct](MyIntMUS, sliceJKΔDUpΣcu8IGFwO60k2TEgΞΞ, mapops.WithLenValidator[MyInt, []MyStruct](com.ValidatorFn[int](common_testdata.ValidateLength20)), mapops.WithKeyValidator[MyInt, []MyStruct](com.ValidatorFn[MyInt](common_testdata.ValidateZeroValue[MyInt])))
	ptrCΔoXuyttcvyP7bsyXSZΔzgΞΞ       = ord.NewPtrSer[[3]int](arrayX1ltdZJBPzH2l9ΔCPTeroAΞΞ)
	ptrUDAu6PhjqFE4AdrbSΔbU8wΞΞ       = ord.NewPtrSer[string](ord.String)
	ptrmUwZaFlv85r3GHTX1Δ8ΔbgΞΞ       = ord.NewPtrSer[MyStruct](MyStructMUS)
	sliceHPbgusdjLΣqyQ1S0eYAd9wΞΞ     = ord.NewValidSliceSer[MyStruct](MyStructMUS, slops.WithLenValidator[MyStruct](com.ValidatorFn[int](common_testdata.ValidateLength20)), slops.WithElemValidator[MyStruct](com.ValidatorFn[MyStruct](common_testdata.ValidateZeroValue[MyStruct])))
	sliceJKΔDUpΣcu8IGFwO60k2TEgΞΞ     = ord.NewValidSliceSer[MyStruct](MyStructMUS, slops.WithLenValidator[MyStruct](com.ValidatorFn[int](common_testdata.ValidateLength20)))
	slicegΣxd0IaODΔΣQreGt2TUu7wΞΞ     = ord.NewSliceSer[string](ord.String)
	string8wqPJ7jXpb50sidioRDFWAΞΞ    = ord.NewValidStringSer(strops.WithLenSer(raw.Int), strops.WithLenValidator(com.ValidatorFn[int](common_testdata.ValidateLength20)))
)

var MyIntMUS = myIntMUS{}

type myIntMUS struct{}

func (s myIntMUS) Marshal(v MyInt, bs []byte) (n int) {
	return varint.Int.Marshal(int(v), bs)
}

func (s myIntMUS) Unmarshal(bs []byte) (v MyInt, n int, err error) {
	tmp, n, err := varint.Int.Unmarshal(bs)
	if err != nil {
		return
	}
	v = MyInt(tmp)
	return
}

func (s myIntMUS) Size(v MyInt) (size int) {
	return varint.Int.Size(int(v))
}

func (s myIntMUS) Skip(bs []byte) (n int, err error) {
	return varint.Int.Skip(bs)
}

var MySliceMUS = mySliceMUS{}

type mySliceMUS struct{}

func (s mySliceMUS) Marshal(v MySlice, bs []byte) (n int) {
	return slicegΣxd0IaODΔΣQreGt2TUu7wΞΞ.Marshal([]string(v), bs)
}

func (s mySliceMUS) Unmarshal(bs []byte) (v MySlice, n int, err error) {
	tmp, n, err := slicegΣxd0IaODΔΣQreGt2TUu7wΞΞ.Unmarshal(bs)
	if err != nil {
		return
	}
	v = MySlice(tmp)
	return
}

func (s mySliceMUS) Size(v MySlice) (size int) {
	return slicegΣxd0IaODΔΣQreGt2TUu7wΞΞ.Size([]string(v))
}

func (s mySliceMUS) Skip(bs []byte) (n int, err error) {
	return slicegΣxd0IaODΔΣQreGt2TUu7wΞΞ.Skip(bs)
}

var MyStructMUS = myStructMUS{}

type myStructMUS struct{}

func (s myStructMUS) Marshal(v MyStruct, bs []byte) (n int) {
	n = varint.Int.Marshal(v.Int, bs)
	return n + ord.String.Marshal(v.Str, bs[n:])
}

func (s myStructMUS) Unmarshal(bs []byte) (v MyStruct, n int, err error) {
	v.Int, n, err = varint.Int.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	v.Str, n1, err = ord.String.Unmarshal(bs[n:])
	n += n1
	return
}

func (s myStructMUS) Size(v MyStruct) (size int) {
	size = varint.Int.Size(v.Int)
	return size + ord.String.Size(v.Str)
}

func (s myStructMUS) Skip(bs []byte) (n int, err error) {
	n, err = varint.Int.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Skip(bs[n:])
	n += n1
	return
}

var MyIntDTS = dts.New[MyInt](MyIntDTM, MyIntMUS)

var MyInterfaceMUS = myInterfaceMUS{}

type myInterfaceMUS struct{}

func (s myInterfaceMUS) Marshal(v MyInterface, bs []byte) (n int) {
	switch t := v.(type) {
	case MyInt:
		return MyIntDTS.Marshal(t, bs)
	default:
		panic(fmt.Sprintf("unexpected %v type", t))
	}
}

func (s myInterfaceMUS) Unmarshal(bs []byte) (v MyInterface, n int, err error) {
	dtm, n, err := dts.DTMSer.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	switch dtm {
	case MyIntDTM:
		v, n1, err = MyIntDTS.UnmarshalData(bs[n:])
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
	}
	n += n1
	return
}

func (s myInterfaceMUS) Size(v MyInterface) (size int) {
	switch t := v.(type) {
	case MyInt:
		return MyIntDTS.Size(t)
	default:
		panic(fmt.Sprintf("unexpected %v type", t))
	}
}

func (s myInterfaceMUS) Skip(bs []byte) (n int, err error) {
	dtm, n, err := dts.DTMSer.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	switch dtm {
	case MyIntDTM:
		n1, err = MyIntDTS.SkipData(bs[n:])
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
	}
	n += n1
	return
}

var ComplexStructMUS = complexStructMUS{}

type complexStructMUS struct{}

func (s complexStructMUS) Marshal(v ComplexStruct, bs []byte) (n int) {
	n = ord.Bool.Marshal(v.Bool, bs)
	n += varint.Uint8.Marshal(v.Byte, bs[n:])
	n += raw.Int8.Marshal(v.Int8, bs[n:])
	n += varint.Int16.Marshal(v.Int16, bs[n:])
	n += varint.Int32.Marshal(v.Int32, bs[n:])
	n += varint.Int64.Marshal(v.Int64, bs[n:])
	n += varint.Uint8.Marshal(v.Uint8, bs[n:])
	n += varint.Uint16.Marshal(v.Uint16, bs[n:])
	n += varint.Uint32.Marshal(v.Uint32, bs[n:])
	n += varint.Uint64.Marshal(v.Uint64, bs[n:])
	n += varint.Float32.Marshal(v.Float32, bs[n:])
	n += varint.Float64.Marshal(v.Float64, bs[n:])
	n += string8wqPJ7jXpb50sidioRDFWAΞΞ.Marshal(v.String, bs[n:])
	n += byteSlicejdpLY58ftUFXYGnakpvXgAΞΞ.Marshal(v.ByteSlice, bs[n:])
	n += sliceHPbgusdjLΣqyQ1S0eYAd9wΞΞ.Marshal(v.StructSlice, bs[n:])
	n += arrayoa64Y8KzCJ7ziIbYybTebAΞΞ.Marshal(v.Array, bs[n:])
	n += ptrUDAu6PhjqFE4AdrbSΔbU8wΞΞ.Marshal(v.PtrString, bs[n:])
	n += ptrmUwZaFlv85r3GHTX1Δ8ΔbgΞΞ.Marshal(v.PtrStruct, bs[n:])
	n += ptrUDAu6PhjqFE4AdrbSΔbU8wΞΞ.Marshal(v.NilPtr, bs[n:])
	n += ptrCΔoXuyttcvyP7bsyXSZΔzgΞΞ.Marshal(v.PtrArray, bs[n:])
	n += map1lW2X1YW1lK3PO8YRU2TeQΞΞ.Marshal(v.Map, bs[n:])
	n += raw.TimeUnix.Marshal(v.Time, bs[n:])
	n += MySliceMUS.Marshal(v.Alias, bs[n:])
	return n + MyInterfaceMUS.Marshal(v.Interface, bs[n:])
}

func (s complexStructMUS) Unmarshal(bs []byte) (v ComplexStruct, n int, err error) {
	v.Bool, n, err = ord.Bool.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	v.Byte, n1, err = varint.Uint8.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int8, n1, err = raw.Int8.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int16, n1, err = varint.Int16.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int32, n1, err = varint.Int32.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Int64, n1, err = varint.Int64.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint8, n1, err = varint.Uint8.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint16, n1, err = varint.Uint16.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint32, n1, err = varint.Uint32.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Uint64, n1, err = varint.Uint64.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Float32, n1, err = varint.Float32.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	err = common_testdata.ValidateZeroValue[float32](v.Float32)
	if err != nil {
		return
	}
	v.Float64, n1, err = varint.Float64.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.String, n1, err = string8wqPJ7jXpb50sidioRDFWAΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	err = common_testdata.ValidateZeroValue[string](v.String)
	if err != nil {
		return
	}
	v.ByteSlice, n1, err = byteSlicejdpLY58ftUFXYGnakpvXgAΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.StructSlice, n1, err = sliceHPbgusdjLΣqyQ1S0eYAd9wΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Array, n1, err = arrayoa64Y8KzCJ7ziIbYybTebAΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.PtrString, n1, err = ptrUDAu6PhjqFE4AdrbSΔbU8wΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.PtrStruct, n1, err = ptrmUwZaFlv85r3GHTX1Δ8ΔbgΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.NilPtr, n1, err = ptrUDAu6PhjqFE4AdrbSΔbU8wΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.PtrArray, n1, err = ptrCΔoXuyttcvyP7bsyXSZΔzgΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Map, n1, err = map1lW2X1YW1lK3PO8YRU2TeQΞΞ.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Time, n1, err = raw.TimeUnix.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Alias, n1, err = MySliceMUS.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Interface, n1, err = MyInterfaceMUS.Unmarshal(bs[n:])
	n += n1
	return
}

func (s complexStructMUS) Size(v ComplexStruct) (size int) {
	size = ord.Bool.Size(v.Bool)
	size += varint.Uint8.Size(v.Byte)
	size += raw.Int8.Size(v.Int8)
	size += varint.Int16.Size(v.Int16)
	size += varint.Int32.Size(v.Int32)
	size += varint.Int64.Size(v.Int64)
	size += varint.Uint8.Size(v.Uint8)
	size += varint.Uint16.Size(v.Uint16)
	size += varint.Uint32.Size(v.Uint32)
	size += varint.Uint64.Size(v.Uint64)
	size += varint.Float32.Size(v.Float32)
	size += varint.Float64.Size(v.Float64)
	size += string8wqPJ7jXpb50sidioRDFWAΞΞ.Size(v.String)
	size += byteSlicejdpLY58ftUFXYGnakpvXgAΞΞ.Size(v.ByteSlice)
	size += sliceHPbgusdjLΣqyQ1S0eYAd9wΞΞ.Size(v.StructSlice)
	size += arrayoa64Y8KzCJ7ziIbYybTebAΞΞ.Size(v.Array)
	size += ptrUDAu6PhjqFE4AdrbSΔbU8wΞΞ.Size(v.PtrString)
	size += ptrmUwZaFlv85r3GHTX1Δ8ΔbgΞΞ.Size(v.PtrStruct)
	size += ptrUDAu6PhjqFE4AdrbSΔbU8wΞΞ.Size(v.NilPtr)
	size += ptrCΔoXuyttcvyP7bsyXSZΔzgΞΞ.Size(v.PtrArray)
	size += map1lW2X1YW1lK3PO8YRU2TeQΞΞ.Size(v.Map)
	size += raw.TimeUnix.Size(v.Time)
	size += MySliceMUS.Size(v.Alias)
	return size + MyInterfaceMUS.Size(v.Interface)
}

func (s complexStructMUS) Skip(bs []byte) (n int, err error) {
	n, err = ord.Bool.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = varint.Uint8.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = raw.Int8.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.Int16.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.Int32.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.Int64.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.Uint8.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.Uint16.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.Uint32.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.Uint64.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.Float32.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = varint.Float64.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = string8wqPJ7jXpb50sidioRDFWAΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = byteSlicejdpLY58ftUFXYGnakpvXgAΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = sliceHPbgusdjLΣqyQ1S0eYAd9wΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = arrayoa64Y8KzCJ7ziIbYybTebAΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ptrUDAu6PhjqFE4AdrbSΔbU8wΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ptrmUwZaFlv85r3GHTX1Δ8ΔbgΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ptrUDAu6PhjqFE4AdrbSΔbU8wΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = ptrCΔoXuyttcvyP7bsyXSZΔzgΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = map1lW2X1YW1lK3PO8YRU2TeQΞΞ.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = raw.TimeUnix.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = MySliceMUS.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = MyInterfaceMUS.Skip(bs[n:])
	n += n1
	return
}
