package typename

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrTypeMismatch = errors.New("type mismatch")

func NewNotDefinedTypeError(t reflect.Type, kind string) error {
	return fmt.Errorf("%v is not a defined %v type", t, kind)
}

func NewInvalidPkgPathError(str string) error {
	return fmt.Errorf("invalid '%v' pkg path format", str)
}

func NewInvalidPackageError(str string) error {
	return fmt.Errorf("invalid '%v' package format", str)
}

func NewMultiPointerError(t reflect.Type) error {
	return fmt.Errorf("do not support multi-pointer types, like %v", t)
}

func NewUnsupportedTypeError(t reflect.Type) error {
	return fmt.Errorf("can't get a name for the %v type", t)
}
