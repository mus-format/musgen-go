package typeops

import "fmt"

const (
	Undefined PtrType = iota
	Ordinary
	PM
)

type PtrType int

func (t PtrType) PackageName() string {
	switch t {
	case Undefined, Ordinary:
		return "ord"
	case PM:
		return "pm"
	default:
		panic(fmt.Errorf("unexpected %v encoding", int(t)))
	}
}
