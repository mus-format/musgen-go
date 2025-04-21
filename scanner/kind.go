package scanner

type Kind int

const (
	UndefinedKind Kind = iota
	Defined
	Array
	Slice
	Map
	Prim
)
