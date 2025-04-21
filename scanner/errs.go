package scanner

import (
	"fmt"
)

func NewUnsupportedQualifiedNameError[T QualifiedName](name T) error {
	return fmt.Errorf("unsupported '%v' QualifiedTypeName", name)
}
