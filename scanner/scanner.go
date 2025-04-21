package scanner

import (
	"path/filepath"

	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
)

type QualifiedName interface {
	typename.CompleteName | typename.FullName
}

func Scan[T QualifiedName](name T, op Op[T], tops *typeops.Options,
	ops ...SetOption) (err error) {
	o := Options{}
	Apply(ops, &o)
	return scan(name, op, tops, o, 0)
}

func scan[T QualifiedName](name T, op Op[T], tops *typeops.Options, o Options,
	position Position) (err error) {
	// defined type
	if t, ok := ParseDefinedType(name); ok {
		t.Kind = Defined
		t.Position = position

		switch t.Position {
		// case SingleParam, FirstParam, Param, LastParam:
		case Param:
			t = fixParamPkgPath(t)
		}

		if err = op.ProcessType(t, tops); err != nil {
			return
		}
		pOp := op
		if o.WithoutParams {
			pOp = ignoreOp[T]{}
		}
		for i := range t.Params {
			if i == 0 {
				pOp.ProcessLeftSquare()
			}
			if err = scan(t.Params[i], pOp, nil, o, Param); err != nil {
				return
			}
			if i < len(t.Params)-1 {
				pOp.ProcessComma()
			}
			if i == len(t.Params)-1 {
				pOp.ProcessRightSquare()
			}
		}
		return
	}
	// container type
	if t, keyType, elemType, kind, ok := parseContainerType(name); ok {
		var (
			keyTops  *typeops.Options = nil
			elemTops *typeops.Options = nil
		)
		t.Kind = kind
		t.Position = position
		if err = op.ProcessType(t, tops); err != nil {
			return
		}
		if keyType != "" {
			if tops != nil {
				keyTops = tops.Key
			}
			op.ProcessLeftSquare()
			if err = scan(keyType, op, keyTops, o, Key); err != nil {
				return
			}
			op.ProcessRightSquare()
		}
		if tops != nil {
			elemTops = tops.Elem
		}
		if err = scan(elemType, op, elemTops, o, Elem); err != nil {
			return
		}
		return
	}
	// primitive type
	if t, ok := parsePrimitiveType(name); ok {
		t.Kind = Prim
		t.Position = position
		return op.ProcessType(t, tops)
	}
	return NewUnsupportedQualifiedNameError(name)
}

func fixParamPkgPath[T QualifiedName](t Type[T]) Type[T] {
	t.PkgPath = typename.PkgPath(filepath.Join(string(t.PkgPath), string(t.Pkg)))
	return t
}
