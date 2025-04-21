package typename

import (
	"go/token"
)

type Pkg string

func StrToPkg(str string) (pkg Pkg, err error) {
	if !token.IsIdentifier(str) {
		err = NewInvalidPkgError(str)
		return
	}
	pkg = Pkg(str)
	return
}
