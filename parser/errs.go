package parser

import (
	"fmt"
	"reflect"
)

// NewUnsupportedType creates an error that happens when one of the Parse
// functions receives an unsupported type.
func NewUnsupportedType(t reflect.Type) error {
	return NewUnsupportedTypeStr(t.String())
}

func NewUnsupportedTypeStr(t string) error {
	return fmt.Errorf("unsupported '%v' type", t)
}
