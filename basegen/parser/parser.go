package parser

import (
	"reflect"
	"strconv"
)

// Parse can accept interface, alias or sturct type, for other types returns
// ErrUnsupportedType.
// For interface type returns intr == true.
// For alias type creates string representation of the underlying type and
// returns it as the aliasOf value.
// For struct type (alias to struct is also a sturct) - for each struct
// field creates a string representation of its type, puts all this in
// the fieldsTypes value.
func Parse(tp reflect.Type) (intr bool, aliasOf string, fieldsTypes []string,
	err error) {
	if tp == nil {
		err = ErrUnsupportedType
		return
	}
	if interfaceType(tp) {
		intr = true
		return
	}
	if aliasType(tp) {
		aliasOf, err = parseAlias(tp)
		return
	}
	if definedStrucType(tp) {
		fieldsTypes, err = parseStruct(tp)
		return
	}
	err = ErrUnsupportedType
	return
}

// parseAlias tries to parse an alias type(not an alias to a struct).
func parseAlias(tp reflect.Type) (aliasOf string, err error) {
	if primitiveType(tp) {
		aliasOf = tp.Kind().String()
	} else {
		switch tp.Kind() {
		case reflect.Array:
			aliasOf, err = parseArrayType("", tp, tp.PkgPath())
		case reflect.Slice:
			aliasOf, err = parseSliceType("", tp, tp.PkgPath())
		case reflect.Map:
			aliasOf, err = parseMapType("", tp, tp.PkgPath())
		default:
			err = ErrUnsupportedType
		}
	}
	return
}

// parseStruct tries to parse a struct type or an alias to struct type.
func parseStruct(tp reflect.Type) (fieldsTypes []string, err error) {
	var (
		field     reflect.StructField
		fieldType string
	)
	fieldsTypes = []string{}
	for i := 0; i < tp.NumField(); i++ {
		field = tp.Field(i)
		fieldType, err = parseType(field.Type, tp.PkgPath())
		if err != nil {
			return
		}
		fieldsTypes = append(fieldsTypes, fieldType)
	}
	return
}

// parseType parses a type into the string representation. Complex types are
// parsed recursively.
// Translates a custom type to its name or package + name(if it's from
// another package).
func parseType(tp reflect.Type, currPkg string) (tpStr string, err error) {
	stars, tp := ParsePtrType(tp)
	if aliasType(tp) || definedStrucType(tp) || interfaceType(tp) {
		if currPkg == tp.PkgPath() {
			tpStr = stars + tp.Name()
			return
		}
		tpStr = stars + tp.String()
		return
	}
	if primitiveType(tp) {
		tpStr = stars + tp.String()
		return
	}
	if tp.Kind() == reflect.Array {
		return parseArrayType(stars, tp, currPkg)
	}
	if tp.Kind() == reflect.Slice {
		return parseSliceType(stars, tp, currPkg)
	}
	if tp.Kind() == reflect.Map {
		return parseMapType(stars, tp, currPkg)
	}
	err = ErrUnsupportedType
	return
}

// ParsePtrType returns pointer signs and an underlying type. If the underlying
// type is an alias to the pointer type returns an error.
func ParsePtrType(tp reflect.Type) (stars string, atp reflect.Type) {
	atp = tp
	for {
		if !ptrType(atp) {
			return
		}
		atp = atp.Elem()
		stars += "*"
		continue
	}
}

// parseArrayType returns a string representation of the array type.
func parseArrayType(stars string, tp reflect.Type, pkgPath string) (tpStr string,
	err error) {
	elemTpStr, err := parseType(tp.Elem(), pkgPath)
	if err != nil {
		return "", err
	}
	tpStr = stars + "[" + strconv.Itoa(tp.Len()) + "]" + elemTpStr
	return
}

// parseSliceType returns a string representation of the slice type.
func parseSliceType(stars string, tp reflect.Type, pkgPath string) (tpStr string,
	err error) {
	elemTpStr, err := parseType(tp.Elem(), pkgPath)
	if err != nil {
		return "", err
	}
	tpStr = stars + "[]" + elemTpStr
	return
}

// parseMapType returns a string representation of the map type.
func parseMapType(stars string, tp reflect.Type, pkgPath string) (tpStr string,
	err error) {
	keyTpStr, err := parseType(tp.Key(), pkgPath)
	if err != nil {
		return "", err
	}
	elemTpStr, err := parseType(tp.Elem(), pkgPath)
	if err != nil {
		return "", err
	}
	tpStr = stars + "map[" + keyTpStr + "]" + elemTpStr
	return
}

func hasPkg(tp reflect.Type) bool {
	return tp.PkgPath() != ""
}

func definedStrucType(tp reflect.Type) bool {
	return hasPkg(tp) && strucType(tp)
}

// struct or alias to struct, also returns true for struct{}{}
func strucType(tp reflect.Type) bool {
	return tp.Kind() == reflect.Struct
}

func interfaceType(tp reflect.Type) bool {
	return tp.Kind() == reflect.Interface
}

// alias to primitive type, array, slice or map
func aliasType(tp reflect.Type) bool {
	return hasPkg(tp) && !strucType(tp) && !interfaceType(tp)
}

func primitiveType(tp reflect.Type) bool {
	kind := tp.Kind()
	return kind == reflect.Bool ||
		kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64 ||
		kind == reflect.Float32 ||
		kind == reflect.Float64 ||
		// kind == reflect.Complex64 ||
		// kind == reflect.Complex128 ||
		kind == reflect.String
}

func ptrType(tp reflect.Type) bool {
	return tp.Kind() == reflect.Pointer
}
