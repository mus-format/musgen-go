package parser

import "reflect"

func supportedTypeDef(t reflect.Type) (ok bool) {
	return t.PkgPath() != "" && (basicType(t) || containerType(t) || ptrType(t))
}

func basicType(t reflect.Type) (ok bool) {
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String:
		return true
	default:
		return
	}
}

func containerType(t reflect.Type) (ok bool) {
	switch t.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map:
		return true
	default:
		return
	}
}

func definedStructType(t reflect.Type) (ok bool) {
	return t.PkgPath() != "" && t.Kind() == reflect.Struct
}

func interfaceType(t reflect.Type) (ok bool) {
	return t.Kind() == reflect.Interface && t.NumMethod() > 0
}

func ptrType(t reflect.Type) bool {
	return t.Kind() == reflect.Pointer
}
