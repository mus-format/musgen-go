package musgen

import "fmt"

func NewFileGeneratorError(cause error) error {
	return fmt.Errorf("ensure that all musgen.FileGenerator options are set correctly, cause: %w", cause)
}
