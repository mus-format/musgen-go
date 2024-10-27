package testdata

func ValidateZeroValue[T comparable](t T) (err error) {
	if t == *new(T) {
		err = ErrZeroValue
	}
	return
}

func ValidateLength(l int) (err error) {
	if l > 0 {
		err = ErrTooLong
	}
	return
}
