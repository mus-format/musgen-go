package genops

import (
	"errors"
	"fmt"
)

var ErrEmptyPkgPath = errors.New("option PkgPath is not set")

func NewInvalidImportPathError(val string) error {
	return fmt.Errorf("invalid '%v' import path format", val)
}

func NewInvalidAliasError(val string) error {
	return fmt.Errorf("invalid '%v' alias format", val)
}

func NewDuplicateImportPath(path ImportPath) error {
	return fmt.Errorf("duplicate '%v' import path in ImportAlias option", path)
}

func NewDuplicateImportAlias(alias Alias) error {
	return fmt.Errorf("duplicate '%v' alias in ImportAlias option", alias)
}
