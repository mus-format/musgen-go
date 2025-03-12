package tdesc

import "errors"

// ErrWrongOptionsAmount indicates that the amount of Options differs
// from the number structure fields.
var ErrWrongOptionsAmount = errors.New("wrong Options amount")

// ErrNoOneImplTypeSpecified indicates that len(Options.ImplTypes) == 0.
var ErrNoOneImplTypeSpecified = errors.New("no one ImplType specified")
