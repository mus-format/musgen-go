package scanner

import "github.com/mus-format/musgen-go/typename"

type Type[T QualifiedName] struct {
	PkgPath   typename.PkgPath
	Stars     string
	Pkg       typename.Pkg
	Name      typename.TypeName
	Params    []T
	ArrLength string

	KeyType  T
	ElemType T

	Kind     Kind
	Position Position
}
