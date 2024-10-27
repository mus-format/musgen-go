package mus

import (
	"testing"

	"github.com/mus-format/musgen-go/testdata/pkg1"
)

func NewValidationError(cause error) *ValidationError {
	return &ValidationError{cause}
}

type ValidationError struct {
	cause error
}

func (e *ValidationError) Error() string {
	return e.cause.Error()
}

func TestGeneratedUnsafeStreamCode(t *testing.T) {

	t.Run("Test alias serializability", func(t *testing.T) {

		testStreamSerializability(pkg1.BoolAlias(true),
			pkg1.MarshalUnsafeStreamBoolAliasMUS,
			pkg1.UnmarshalUnsafeStreamBoolAliasMUS,
			pkg1.SizeUnsafeStreamBoolAliasMUS,
			pkg1.SkipUnsafeStreamBoolAliasMUS,
			t)

	})
}
