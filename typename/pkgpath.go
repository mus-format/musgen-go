package typename

import (
	"path/filepath"

	"golang.org/x/mod/module"
)

type PkgPath string

func (p PkgPath) Package() Pkg {
	return Pkg(filepath.Base(string(p)))
}

func StrToPkgPath(str string) (path PkgPath, err error) {
	if err = module.CheckPath(str); err != nil {
		err = NewInvalidPkgPathError(str)
		return
	}
	path = PkgPath(str)
	return
}
