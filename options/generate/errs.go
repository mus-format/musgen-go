package genops

import (
	"errors"
	"fmt"
)

var ErrEmptyPkgPath = errors.New("option PkgPath is not set")

func NewInvalidImportPathError(val string) error {
	return fmt.Errorf("invalid '%v' import path", val)
}

func NewInvalidAliasError(val string) error {
	return fmt.Errorf("invalid '%v' package alias", val)
}

func NewDuplicateImportPath(path ImportPath) error {
	return fmt.Errorf("duplicate '%v' import path in musgen.FileGenerator options", path)
}

func NewDuplicateImportAlias(alias Alias) error {
	return fmt.Errorf("duplicate '%v' package alias in musgen.FileGenerator options",
		alias)
}
