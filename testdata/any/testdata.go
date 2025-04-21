package testdata

type MyAny any

type MyAnySlice []any

type MyAnyStruct struct {
	Any any
}

type MyAnyGenericSlice[T any] []T
