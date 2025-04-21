package musgen

// NewTmplEngineError creates a new TmplEngineError.
func NewTmplEngineError(bs []byte, cause error) error {
	return &TmplEngineError{bs, cause}
}

// TmplEngineError represents a template processing failure, preserving both
// the original template bytes and the underlying error cause.
type TmplEngineError struct {
	bs    []byte
	cause error
}

func (e *TmplEngineError) ByteSlice() []byte {
	return e.bs
}

func (e *TmplEngineError) Error() string {
	return "ensure that all FileGenerator options are set correctly"
}

func (e *TmplEngineError) Unwrap() error {
	return e.cause
}
