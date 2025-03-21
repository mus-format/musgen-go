package parser

import (
	"reflect"
	"strconv"
)

func parseType(t reflect.Type, currPkg string) (str string, err error) {
	stars, t := parsePtrType(t)
	if supportedDefinedType(t) || definedStructType(t) || interfaceType(t) {
		if currPkg == t.PkgPath() {
			str = stars + t.Name()
			return
		}
		str = stars + t.String()
		return
	}
	if basicType(t) {
		str = stars + t.String()
		return
	}
	if t.Kind() == reflect.Array {
		return parseArrayType(stars, t, currPkg)
	}
	if t.Kind() == reflect.Slice {
		return parseSliceType(stars, t, currPkg)
	}
	if t.Kind() == reflect.Map {
		return parseMapType(stars, t, currPkg)
	}
	err = NewUnsupportedType(t)
	return
}

func parsePtrType(t reflect.Type) (stars string, at reflect.Type) {
	at = t
	for {
		if !ptrType(at) {
			return
		}
		at = at.Elem()
		stars += "*"
		continue
	}
}

func parseArrayType(stars string, t reflect.Type, pkgPath string) (str string,
	err error) {
	elemStr, err := parseType(t.Elem(), pkgPath)
	if err != nil {
		return "", err
	}
	str = stars + "[" + strconv.Itoa(t.Len()) + "]" + elemStr
	return
}

func parseSliceType(stars string, t reflect.Type, pkgPath string) (str string,
	err error) {
	elemStr, err := parseType(t.Elem(), pkgPath)
	if err != nil {
		return "", err
	}
	str = stars + "[]" + elemStr
	return
}

func parseMapType(stars string, t reflect.Type, pkgPath string) (str string,
	err error) {
	keyStr, err := parseType(t.Key(), pkgPath)
	if err != nil {
		return "", err
	}
	elemStr, err := parseType(t.Elem(), pkgPath)
	if err != nil {
		return "", err
	}
	str = stars + "map[" + keyStr + "]" + elemStr
	return
}
