package classifier

import (
	"reflect"
	"testing"

	contr_testdata "github.com/mus-format/musgen-go/testdata/container"
	generic_testdata "github.com/mus-format/musgen-go/testdata/generic"
	intr_testdata "github.com/mus-format/musgen-go/testdata/interface"
	ptr_testdata "github.com/mus-format/musgen-go/testdata/pointer"
	prim_testdata "github.com/mus-format/musgen-go/testdata/primitive"
	struct_testdata "github.com/mus-format/musgen-go/testdata/struct"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestClassifier(t *testing.T) {
	var (
		undefinedStructType     = reflect.TypeFor[struct{}]()
		definedStructType       = reflect.TypeFor[struct_testdata.MyStruct]()
		doubleDefinedStructType = reflect.TypeFor[struct_testdata.DoubleDefinedMyStruct]()

		undefinedInterfaceType     = reflect.TypeFor[interface{ Print() }]()
		definedInterfaceType       = reflect.TypeFor[intr_testdata.MyInterface]()
		doubleDefinedInterfaceType = reflect.TypeFor[intr_testdata.DoubleDefinedMyInterface]()
		anyType                    = reflect.TypeFor[any]()

		undefinedPrimitiveType = reflect.TypeFor[int]()
		definedPrimitiveType   = reflect.TypeFor[prim_testdata.MyInt]()

		undefinedContainerType           = reflect.TypeFor[[]int]()
		definedContainerType             = reflect.TypeFor[contr_testdata.MySlice]()
		definedParametrizedContainerType = reflect.TypeFor[generic_testdata.MySlice[int]]()

		ptrPrimitiveType = reflect.TypeFor[*int]()
		ptrContainerType = reflect.TypeFor[*[]int]()
		ptrStructType    = reflect.TypeFor[*struct{}]()
		ptrInterfaceType = reflect.TypeFor[*interface{ Print() }]()

		ptrDefinedPrimitiveType = reflect.TypeFor[ptr_testdata.MyIntPtr]()
		ptrDefinedContainerType = reflect.TypeFor[ptr_testdata.MySlicePtr]()
		ptrDefinedStructType    = reflect.TypeFor[ptr_testdata.MyStructPtr]()
		ptrDefinedInterfaceType = reflect.TypeFor[ptr_testdata.MyInterfacePtr]()
	)

	t.Run("DefinedBasicType", func(t *testing.T) {
		testCases := []struct {
			tp   reflect.Type
			want bool
		}{
			{tp: undefinedStructType, want: false},
			{tp: definedStructType, want: false},
			{tp: doubleDefinedStructType, want: false},
			{tp: undefinedInterfaceType, want: false},
			{tp: definedInterfaceType, want: false},
			{tp: doubleDefinedInterfaceType, want: false},
			{tp: anyType, want: false},
			{tp: undefinedPrimitiveType, want: false},
			{tp: undefinedContainerType, want: false},
			{tp: ptrPrimitiveType, want: false},
			{tp: ptrContainerType, want: false},
			{tp: ptrStructType, want: false},
			{tp: ptrInterfaceType, want: false},

			{tp: definedPrimitiveType, want: true},
			{tp: definedContainerType, want: true},
			{tp: definedParametrizedContainerType, want: true},
			{tp: ptrDefinedPrimitiveType, want: true},
			{tp: ptrDefinedContainerType, want: true},
			{tp: ptrDefinedStructType, want: true},
			{tp: ptrDefinedInterfaceType, want: true},
		}
		for _, c := range testCases {
			asserterror.Equal(DefinedBasicType(c.tp), c.want, t)
		}
	})

	t.Run("DefinedStruct", func(t *testing.T) {
		testCases := []struct {
			tp   reflect.Type
			want bool
		}{
			{tp: undefinedStructType, want: false},
			{tp: undefinedInterfaceType, want: false},
			{tp: definedInterfaceType, want: false},
			{tp: doubleDefinedInterfaceType, want: false},
			{tp: anyType, want: false},
			{tp: undefinedPrimitiveType, want: false},
			{tp: definedPrimitiveType, want: false},
			{tp: undefinedContainerType, want: false},
			{tp: definedContainerType, want: false},
			{tp: definedParametrizedContainerType, want: false},
			{tp: ptrPrimitiveType, want: false},
			{tp: ptrContainerType, want: false},
			{tp: ptrStructType, want: false},
			{tp: ptrInterfaceType, want: false},
			{tp: ptrDefinedPrimitiveType, want: false},
			{tp: ptrDefinedContainerType, want: false},
			{tp: ptrDefinedStructType, want: false},
			{tp: ptrDefinedInterfaceType, want: false},

			{tp: definedStructType, want: true},
			{tp: doubleDefinedStructType, want: true},
		}
		for _, c := range testCases {
			asserterror.Equal(DefinedStruct(c.tp), c.want, t)
		}
	})

	t.Run("DefinedInterface", func(t *testing.T) {
		testCases := []struct {
			tp   reflect.Type
			want bool
		}{
			{tp: undefinedStructType, want: false},
			{tp: definedStructType, want: false},
			{tp: doubleDefinedStructType, want: false},
			{tp: undefinedInterfaceType, want: false},
			{tp: anyType, want: false},
			{tp: undefinedPrimitiveType, want: false},
			{tp: definedPrimitiveType, want: false},
			{tp: undefinedContainerType, want: false},
			{tp: definedContainerType, want: false},
			{tp: definedParametrizedContainerType, want: false},
			{tp: ptrPrimitiveType, want: false},
			{tp: ptrContainerType, want: false},
			{tp: ptrStructType, want: false},
			{tp: ptrInterfaceType, want: false},
			{tp: ptrDefinedPrimitiveType, want: false},
			{tp: ptrDefinedContainerType, want: false},
			{tp: ptrDefinedStructType, want: false},
			{tp: ptrDefinedInterfaceType, want: false},

			{tp: definedInterfaceType, want: true},
			{tp: doubleDefinedInterfaceType, want: true},
		}
		for _, c := range testCases {
			asserterror.Equal(DefinedNonEmptyInterface(c.tp), c.want, t)
		}
	})
}
