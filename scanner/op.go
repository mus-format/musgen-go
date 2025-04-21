package scanner

import typeops "github.com/mus-format/musgen-go/options/type"

type Op[T QualifiedName] interface {
	ProcessType(t Type[T], tops *typeops.Options) error
	ProcessLeftSquare()
	ProcessComma()
	ProcessRightSquare()
}

type ignoreOp[T QualifiedName] struct{}

func (o ignoreOp[T]) ProcessType(t Type[T], tops *typeops.Options) (err error) {
	return
}
func (o ignoreOp[T]) ProcessLeftSquare()  {}
func (o ignoreOp[T]) ProcessComma()       {}
func (o ignoreOp[T]) ProcessRightSquare() {}
