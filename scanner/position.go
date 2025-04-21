package scanner

type Position int

const (
	UndefinedPosition Position = iota
	Key
	Elem
	Param
	// FirstParam
	// Param
	// LastParam
	// SingleParam
)
