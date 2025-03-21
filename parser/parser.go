package parser

import (
	"reflect"
)

func ParseTypedef(t reflect.Type) (sourceType string, err error) {
	if !supportedDefinedType(t) {
		err = NewUnsupportedType(t)
		return
	}
	if basicType(t) {
		sourceType = t.Kind().String()
	} else if stars, elemType := parsePtrType(t); stars != "" {
		sourceType, err = parseType(elemType, elemType.PkgPath())
		if err == nil {
			sourceType = stars + sourceType
		}
	} else {
		switch t.Kind() {
		case reflect.Array:
			sourceType, err = parseArrayType("", t, t.PkgPath())
		case reflect.Slice:
			sourceType, err = parseSliceType("", t, t.PkgPath())
		case reflect.Map:
			sourceType, err = parseMapType("", t, t.PkgPath())
		default:
			panic(NewUnsupportedType(t))
		}
	}
	return
}

func ParseStructType(t reflect.Type) (fieldTypes []string, err error) {
	if !definedStructType(t) {
		err = NewUnsupportedType(t)
		return
	}
	var (
		field     reflect.StructField
		fieldType string
	)
	fieldTypes = []string{}
	for i := range t.NumField() {
		field = t.Field(i)
		fieldType, err = parseType(field.Type, t.PkgPath())
		if err != nil {
			return
		}
		fieldTypes = append(fieldTypes, fieldType)
	}
	return
}

func ParseInterfaceType(t reflect.Type) (err error) {
	if !interfaceType(t) {
		err = NewUnsupportedType(t)
	}
	return
}

func SupportedType(t reflect.Type) (err error) {
	_, err = ParseTypedef(t)
	if err == nil {
		return
	}
	_, err = ParseStructType(t)
	if err == nil {
		return
	}
	err = ParseInterfaceType(t)
	if err == nil {
		return
	}
	err = NewUnsupportedType(t)
	return
}
