package mus

import (
	"testing"

	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
	"github.com/mus-format/musgen-go/testdata/pkg1"
)

func TestGeneratedUnsafeCode(t *testing.T) {

	t.Run("Test alias serializability", func(t *testing.T) {

		testSerializability(pkg1.BoolAlias(true),
			pkg1.MarshalUnsafeBoolAliasMUS,
			func(bs []byte) (v pkg1.BoolAlias, n int, err error) {
				_, _, err = unsafe.UnmarshalBool(bs)
				if err != nil {
					return
				}
				return pkg1.UnmarshalUnsafeBoolAliasMUS(bs)
			},
			pkg1.SizeUnsafeBoolAliasMUS,
			pkg1.SkipUnsafeBoolAliasMUS,
			t)
		testSerializability(pkg1.ByteAlias(10),
			pkg1.MarshalUnsafeByteAliasMUS,
			func(bs []byte) (v pkg1.ByteAlias, n int, err error) {
				_, _, err = unsafe.UnmarshalByte(bs)
				if err != nil {
					return
				}
				return pkg1.UnmarshalUnsafeByteAliasMUS(bs)
			},
			pkg1.SizeUnsafeByteAliasMUS,
			pkg1.SkipUnsafeByteAliasMUS,
			t)
		testSerializability(pkg1.IntAlias(10),
			pkg1.MarshalUnsafeIntAliasMUS,
			func(bs []byte) (v pkg1.IntAlias, n int, err error) {
				_, _, err = unsafe.UnmarshalInt(bs)
				if err != nil {
					return
				}
				return pkg1.UnmarshalUnsafeIntAliasMUS(bs)
			},
			pkg1.SizeUnsafeIntAliasMUS,
			pkg1.SkipUnsafeIntAliasMUS,
			t)
		testSerializability(pkg1.StringAlias("some"),
			pkg1.MarshalUnsafeStringAliasMUS,
			func(bs []byte) (v pkg1.StringAlias, n int, err error) {
				_, _, err = unsafe.UnmarshalString(nil, bs)
				if err != nil {
					return
				}
				return pkg1.UnmarshalUnsafeStringAliasMUS(bs)
			},
			pkg1.SizeUnsafeStringAliasMUS,
			pkg1.SkipUnsafeStringAliasMUS,
			t)
		testSerializability(pkg1.Float32Alias(20.002),
			pkg1.MarshalUnsafeFloat32AliasMUS,
			func(bs []byte) (v pkg1.Float32Alias, n int, err error) {
				_, _, err = unsafe.UnmarshalFloat32(bs)
				if err != nil {
					return
				}
				return pkg1.UnmarshalUnsafeFloat32AliasMUS(bs)
			},
			pkg1.SizeUnsafeFloat32AliasMUS,
			pkg1.SkipUnsafeFloat32AliasMUS,
			t)
		testSerializability(pkg1.SliceAlias([]int{10, 20}),
			pkg1.MarshalUnsafeSliceAliasMUS,
			func(bs []byte) (v pkg1.SliceAlias, n int, err error) {
				_, n, err = varint.UnmarshalInt(bs)
				if err != nil {
					return
				}
				_, _, err = unsafe.UnmarshalInt(bs[n:])
				if err != nil {
					return
				}
				return pkg1.UnmarshalUnsafeSliceAliasMUS(bs)
			},
			pkg1.SizeUnsafeSliceAliasMUS,
			pkg1.SkipUnsafeSliceAliasMUS,
			t)
		testSerializability(pkg1.MapAlias(map[string]int{"some": 100}),
			pkg1.MarshalUnsafeMapAliasMUS,
			func(bs []byte) (v pkg1.MapAlias, n int, err error) {
				_, n, err = varint.UnmarshalInt(bs)
				if err != nil {
					return
				}
				var n1 int
				_, n1, err = unsafe.UnmarshalString(nil, bs[n:])
				n += n1
				if err != nil {
					return
				}
				_, _, err = unsafe.UnmarshalInt(bs[n:])
				if err != nil {
					return
				}
				return pkg1.UnmarshalUnsafeMapAliasMUS(bs)
			},
			pkg1.SizeUnsafeMapAliasMUS,
			pkg1.SkipUnsafeMapAliasMUS,
			t)

	})

	t.Run("Test struct/alias/dts/interface/ptr serializability", func(t *testing.T) {

		t.Run("ComplexStruct should be serializable", func(t *testing.T) {
			testSerializability(makeCompplexStruct(),
				pkg1.MarshalUnsafeComplexStructMUS,
				pkg1.UnmarshalUnsafeComplexStructMUS,
				pkg1.SizeUnsafeComplexStructMUS,
				pkg1.SkipUnsafeComplexStructMUS,
				t)
		})

	})

}
