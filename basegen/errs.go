package basegen

import (
	"errors"
	"fmt"
)

// ErrNotAlias happens on building alias TypeDesc for not an alias type.
var ErrNotAlias = errors.New("not an alias")

// ErrNotAlias happens on building struct TypeDesc for not a struct type.
var ErrNotStruct = errors.New("not a struct")

// ErrNotAlias happens on building interface TypeDesc for not an interface type.
var ErrNotInterface = errors.New("not an interface")

// ErrWrongOptionsAmount indicates than the amount of Options differs
// from the number structure fields.
var ErrWrongOptionsAmount = errors.New("wrong Options amount")

var ErrEmptyOneof = errors.New("Options.Oneof should contain at least one item")

var ErrUnsupportedArrayType = errors.New("do not support array type")

func NewUnexpectedFnType(fnType FnType) error {
	return fmt.Errorf("unexpected %v function type", fnType)
}
