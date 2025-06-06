package builders

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mus-format/musgen-go/data"
	"github.com/mus-format/musgen-go/data/builders/testdata/mock"
	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	container_testdata "github.com/mus-format/musgen-go/testdata/container"
	intr_testdata "github.com/mus-format/musgen-go/testdata/interface"
	prim_testdata "github.com/mus-format/musgen-go/testdata/primitive"
	struct_testdata "github.com/mus-format/musgen-go/testdata/struct"
	time_testdata "github.com/mus-format/musgen-go/testdata/time"
	"github.com/mus-format/musgen-go/typename"
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

func TestTypeDataBuilder(t *testing.T) {

	t.Run("BuildDefinedTypeData", func(t *testing.T) {

		testBuildDefinedTypeData := func(tp reflect.Type, conv TypeNameConvertor,
			tops *typeops.Options,
			gops genops.Options,
			wantTypeData data.TypeData,
			wantErr error,
			mocks []*mok.Mock,
			t *testing.T,
		) {
			t.Helper()
			b := NewTypeDataBuilder(conv, gops)
			d, err := b.BuildDefinedTypeData(tp, tops)
			asserterror.EqualError(err, wantErr, t)
			asserterror.EqualDeep(d, wantTypeData, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		}

		t.Run("Should work for primitive type", func(t *testing.T) {
			var (
				tp           = reflect.TypeFor[prim_testdata.MyInt]()
				gops         = genops.New()
				wantTypeData = data.TypeData{
					FullName:       "testdata.MyInt",
					SourceFullName: "int",

					Fields: []data.FieldData{
						{FullName: "int"},
					},
					Gops: gops,
				}
				wantErr error = nil
				conv          = mock.NewTypeNameConvertor().RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "github.com/mus-format/musgen-go/testdata/primitive/testdata.MyInt", t)
						return wantTypeData.FullName, nil
					}).RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "int", t)
						return "int", nil
					})
				mocks = []*mok.Mock{conv.Mock}
			)
			testBuildDefinedTypeData(tp, conv, nil, gops, wantTypeData, wantErr,
				mocks, t)
		})

		t.Run("Should work for slice ", func(t *testing.T) {
			var (
				tp           = reflect.TypeFor[container_testdata.MySlice]()
				gops         = genops.New()
				wantTypeData = data.TypeData{
					FullName:       "testdata.MySlice",
					SourceFullName: "int",

					Fields: []data.FieldData{
						{FullName: "int"},
					},
					Gops: gops,
				}
				wantErr error = nil
				conv          = mock.NewTypeNameConvertor().RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "github.com/mus-format/musgen-go/testdata/container/testdata.MySlice", t)
						return wantTypeData.FullName, nil
					}).RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "[]int", t)
						return "int", nil
					})
				mocks = []*mok.Mock{conv.Mock}
			)
			testBuildDefinedTypeData(tp, conv, nil, gops, wantTypeData, wantErr,
				mocks, t)
		})

		t.Run("If Converter.ConvertToFullName fails with an error BuildDefinedTypeData should return it",
			func(t *testing.T) {
				var (
					tp           = reflect.TypeFor[prim_testdata.MyInt]()
					gops         = genops.New()
					wantTypeData = data.TypeData{}
					wantErr      = errors.New("Converter.ConvertToFullName error")
					conv         = mock.NewTypeNameConvertor().RegisterConvertToFullName(
						func(cname typename.CompleteName) (fname typename.FullName, err error) {
							err = wantErr
							return
						})
					mocks = []*mok.Mock{conv.Mock}
				)
				testBuildDefinedTypeData(tp, conv, nil, gops, wantTypeData, wantErr,
					mocks, t)
			})

		t.Run("If second Converter.ConvertToFullName fails with an error BuildDefinedTypeData should return it",
			func(t *testing.T) {
				var (
					tp           = reflect.TypeFor[prim_testdata.MyInt]()
					gops         = genops.New()
					wantTypeData = data.TypeData{}
					wantErr      = errors.New("Converter.ConvertToFullName error")

					conv = mock.NewTypeNameConvertor().RegisterConvertToFullName(
						func(cname typename.CompleteName) (fname typename.FullName, err error) {
							return "testdata.MyInt", nil
						}).RegisterConvertToFullName(
						func(cname typename.CompleteName) (fname typename.FullName, err error) {
							err = wantErr
							return
						})
					mocks = []*mok.Mock{conv.Mock}
				)
				testBuildDefinedTypeData(tp, conv, nil, gops, wantTypeData, wantErr,
					mocks, t)
			})

		t.Run("Should fail if receives not defined basic type", func(t *testing.T) {
			var (
				tp           = reflect.TypeFor[struct{}]()
				gops         = genops.New()
				wantTypeData = data.TypeData{}
				wantErr      = NewUnsupportedTypeError(tp)
			)
			testBuildDefinedTypeData(tp, nil, nil, gops, wantTypeData, wantErr, nil,
				t)
		})

	})

	t.Run("BuildStructData", func(t *testing.T) {

		testBuildStructData := func(tp reflect.Type, conv TypeNameConvertor,
			sops structops.Options,
			gops genops.Options,
			wantTypeData data.TypeData,
			wantErr error,
			mocks []*mok.Mock,
			t *testing.T,
		) {
			t.Helper()
			b := NewTypeDataBuilder(conv, gops)
			d, err := b.BuildStructData(tp, sops)
			asserterror.EqualError(err, wantErr, t)
			asserterror.EqualDeep(d, wantTypeData, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		}

		t.Run("Should work", func(t *testing.T) {
			var (
				tp           = reflect.TypeFor[struct_testdata.MyStruct]()
				sops         = structops.New()
				gops         = genops.New()
				wantTypeData = data.TypeData{
					FullName: "testdata.MyStruct",

					Fields: []data.FieldData{
						{FullName: "int", FieldName: "Int"},
						{FullName: "string", FieldName: "Str"},
					},
					Sops: sops,
					Gops: gops,
				}
				wantErr error = nil
				conv          = mock.NewTypeNameConvertor().RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "github.com/mus-format/musgen-go/testdata/struct/testdata.MyStruct", t)
						return wantTypeData.FullName, nil
					}).RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "int", t)
						return wantTypeData.Fields[0].FullName, nil
					}).RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "string", t)
						return wantTypeData.Fields[1].FullName, nil
					})
				mocks = []*mok.Mock{conv.Mock}
			)
			testBuildStructData(tp, conv, sops, gops, wantTypeData, wantErr, mocks, t)
		})

		t.Run("If Converter.ConvertToFullName fails with an error, BuildStructData should return it",
			func(t *testing.T) {
				var (
					tp           = reflect.TypeFor[struct_testdata.MyStruct]()
					sops         = structops.New()
					gops         = genops.New()
					wantTypeData = data.TypeData{}
					wantErr      = errors.New("Converter.ConvertToFullName error")
					conv         = mock.NewTypeNameConvertor().RegisterConvertToFullName(
						func(cname typename.CompleteName) (fname typename.FullName, err error) {
							err = wantErr
							return
						})
					mocks = []*mok.Mock{conv.Mock}
				)
				testBuildStructData(tp, conv, sops, gops, wantTypeData, wantErr, mocks,
					t)
			})

		t.Run("If second Converter.ConvertToFullName fails with an error, BuildStructData should return it",
			func(t *testing.T) {
				var (
					tp           = reflect.TypeFor[struct_testdata.MyStruct]()
					sops         = structops.New()
					gops         = genops.New()
					wantTypeData = data.TypeData{}
					wantErr      = errors.New("Converter.ConvertToFullName error")

					conv = mock.NewTypeNameConvertor().RegisterConvertToFullName(
						func(cname typename.CompleteName) (fname typename.FullName, err error) {
							return "testdata.MyStruct", nil
						}).RegisterConvertToFullName(
						func(cname typename.CompleteName) (fname typename.FullName, err error) {
							err = wantErr
							return
						})
					mocks = []*mok.Mock{conv.Mock}
				)
				testBuildStructData(tp, conv, sops, gops, wantTypeData, wantErr, mocks,
					t)
			})

		t.Run("Should fail if receives not a struct type", func(t *testing.T) {
			var (
				tp   = reflect.TypeFor[prim_testdata.MyInt]()
				sops = structops.New()
				gops = genops.New()

				wantTypeData = data.TypeData{}
				wantErr      = NewUnexpectedDefinedTypeError(tp)
			)
			testBuildStructData(tp, nil, sops, gops, wantTypeData, wantErr, nil, t)
		})

		t.Run("Should fail if receives not a defined struct type", func(t *testing.T) {
			var (
				tp   = reflect.TypeFor[struct{}]()
				sops = structops.New()
				gops = genops.New()

				wantTypeData = data.TypeData{}
				wantErr      = NewUnsupportedTypeError(tp)
			)
			testBuildStructData(tp, nil, sops, gops, wantTypeData, wantErr, nil, t)
		})

		t.Run("Should fail if len(sops.Fields) != number of struct fields",
			func(t *testing.T) {
				var (
					tp   = reflect.TypeFor[struct_testdata.MyStruct]()
					sops = structops.New()
					gops = genops.New()

					wantTypeData = data.TypeData{}
					wantErr      = NewWrongFieldsCountError(2)
				)
				structops.Apply([]structops.SetOption{structops.WithField(nil)}, &sops)
				testBuildStructData(tp, nil, sops, gops, wantTypeData, wantErr, nil, t)
			})

	})

	t.Run("BuildInterfaceData", func(t *testing.T) {

		testBuildInterfaceData := func(tp reflect.Type, conv TypeNameConvertor,
			iops introps.Options,
			gops genops.Options,
			wantTypeData data.TypeData,
			wantErr error,
			mocks []*mok.Mock,
			t *testing.T,
		) {
			t.Helper()
			b := NewTypeDataBuilder(conv, gops)
			d, err := b.BuildInterfaceData(tp, iops)
			asserterror.EqualError(err, wantErr, t)
			asserterror.EqualDeep(d, wantTypeData, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		}

		t.Run("Should work", func(t *testing.T) {
			var (
				tp   = reflect.TypeFor[intr_testdata.MyInterface]()
				iops = introps.New()
				gops = genops.New()
			)
			introps.Apply([]introps.SetOption{
				introps.WithImpl(reflect.TypeFor[intr_testdata.Impl1]()),
				introps.WithImpl(reflect.TypeFor[intr_testdata.Impl2]()),
			}, &iops)

			var (
				wantTypeData = data.TypeData{
					FullName: "testdata.MyInterface",
					Impls:    []typename.FullName{"testdata.Impl1", "testdata.Impl2"},
					Iops:     iops,
					Gops:     gops,
				}
				wantErr error = nil

				conv = mock.NewTypeNameConvertor().RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "github.com/mus-format/musgen-go/testdata/interface/testdata.MyInterface", t)
						return wantTypeData.FullName, nil
					}).RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "github.com/mus-format/musgen-go/testdata/interface/testdata.Impl1", t)
						return wantTypeData.Impls[0], nil
					}).RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "github.com/mus-format/musgen-go/testdata/interface/testdata.Impl2", t)
						return wantTypeData.Impls[1], nil
					})
				mocks = []*mok.Mock{conv.Mock}
			)
			testBuildInterfaceData(tp, conv, iops, gops, wantTypeData, wantErr, mocks,
				t)
		})

		t.Run("If Converter.ConvertToFullName fails with an error, BuildInterfaceData should return it",
			func(t *testing.T) {
				var (
					tp   = reflect.TypeFor[intr_testdata.MyInterface]()
					iops = introps.New()
					gops = genops.New()
				)
				introps.Apply([]introps.SetOption{
					introps.WithImpl(reflect.TypeFor[intr_testdata.Impl1]()),
					introps.WithImpl(reflect.TypeFor[intr_testdata.Impl2]()),
				}, &iops)
				var (
					wantTypeData = data.TypeData{}
					wantErr      = errors.New("Converter.ConvertToFullName error")
					conv         = mock.NewTypeNameConvertor().RegisterConvertToFullName(
						func(cname typename.CompleteName) (fname typename.FullName, err error) {
							err = wantErr
							return
						})
					mocks = []*mok.Mock{conv.Mock}
				)

				testBuildInterfaceData(tp, conv, iops, gops, wantTypeData, wantErr, mocks,
					t)
			})

		t.Run("If Converter.ConvertToFullName fails with an error, BuildInterfaceData should return it",
			func(t *testing.T) {
				var (
					tp   = reflect.TypeFor[intr_testdata.MyInterface]()
					iops = introps.New()
					gops = genops.New()
				)
				introps.Apply([]introps.SetOption{
					introps.WithImpl(reflect.TypeFor[intr_testdata.Impl1]()),
					introps.WithImpl(reflect.TypeFor[intr_testdata.Impl2]()),
				}, &iops)
				var (
					wantTypeData = data.TypeData{}
					wantErr      = errors.New("Converter.ConvertToFullName error")
					conv         = mock.NewTypeNameConvertor().RegisterConvertToFullName(
						func(cname typename.CompleteName) (fname typename.FullName, err error) {
							fname = "testdata.MyInterface"
							return
						}).RegisterConvertToFullName(
						func(cname typename.CompleteName) (fname typename.FullName, err error) {
							err = wantErr
							return
						})
					mocks = []*mok.Mock{conv.Mock}
				)

				testBuildInterfaceData(tp, conv, iops, gops, wantTypeData, wantErr, mocks,
					t)
			})

		t.Run("Should fail if receives not an interface type", func(t *testing.T) {
			var (
				tp   = reflect.TypeFor[prim_testdata.MyInt]()
				iops = introps.New()
				gops = genops.New()

				wantTypeData = data.TypeData{}
				wantErr      = NewUnexpectedDefinedTypeError(tp)
			)
			testBuildInterfaceData(tp, nil, iops, gops, wantTypeData, wantErr, nil, t)
		})

		t.Run("Should fail if receives not a not defined interface type", func(t *testing.T) {
			var (
				tp   = reflect.TypeFor[any]()
				iops = introps.New()
				gops = genops.New()

				wantTypeData = data.TypeData{}
				wantErr      = NewUnsupportedTypeError(tp)
			)
			testBuildInterfaceData(tp, nil, iops, gops, wantTypeData, wantErr, nil, t)
		})

	})

	t.Run("BuildDTSDesc", func(t *testing.T) {

		t.Run("Should work", func(t *testing.T) {
			var (
				gops         = genops.New()
				wantTypeData = data.TypeData{
					FullName: "testdata.MyInt",
					Gops:     gops,
				}
				wantErr error = nil

				conv = mock.NewTypeNameConvertor().RegisterConvertToFullName(
					func(cname typename.CompleteName) (fname typename.FullName, err error) {
						asserterror.Equal(cname, "github.com/mus-format/musgen-go/testdata/primitive/testdata.MyInt", t)
						return wantTypeData.FullName, nil
					})

				b = NewTypeDataBuilder(conv, gops)

				mocks = []*mok.Mock{conv.Mock}
			)
			d, err := b.BuildDTSData(reflect.TypeFor[prim_testdata.MyInt]())
			asserterror.EqualError(err, wantErr, t)
			asserterror.EqualDeep(d, wantTypeData, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("Should fail if receives an unsupported type", func(t *testing.T) {
			var (
				tp           = reflect.TypeFor[struct{}]()
				gops         = genops.New()
				wantTypeData = data.TypeData{}
				wantErr      = NewUnsupportedTypeError(tp)
				b            = NewTypeDataBuilder(nil, gops)
			)
			d, err := b.BuildDTSData(tp)
			asserterror.EqualError(err, wantErr, t)
			asserterror.EqualDeep(d, wantTypeData, t)
		})
	})

	t.Run("BuildTimeData", func(t *testing.T) {
		var (
			gops         = genops.New()
			wantTypeData = data.TypeData{
				FullName:       "testdata.MyTime",
				SourceFullName: "time.Time",
				Fields: []data.FieldData{
					{FullName: "time.Time"},
				},
				Gops: gops,
			}
			wantErr error = nil

			conv = mock.NewTypeNameConvertor().RegisterConvertToFullName(
				func(cname typename.CompleteName) (fname typename.FullName, err error) {
					asserterror.Equal(cname, "github.com/mus-format/musgen-go/testdata/time/testdata.MyTime", t)
					return wantTypeData.FullName, nil
				})

			b = NewTypeDataBuilder(conv, gops)

			mocks = []*mok.Mock{conv.Mock}
		)
		d, err := b.BuildTimeData(reflect.TypeFor[time_testdata.MyTime](), nil)
		asserterror.EqualError(err, wantErr, t)
		asserterror.EqualDeep(d, wantTypeData, t)

		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

}
