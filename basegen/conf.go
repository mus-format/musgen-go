package basegen

// Conf configures Basegen.
type Conf struct {
	Package string   // Defines the package of the generated file.
	Unsafe  bool     // If true, an unsafe code will be generated.
	Stream  bool     // If true, a streaming code will be generated.
	Imports []string // These strings will be added to the imports of the generated file.
}

func (c Conf) ModImportName() string {
	if c.Stream {
		return "muss"
	}
	return "mus"
}

func (c Conf) MarshalParamSignature() string {
	if c.Stream {
		return "w muss.Writer"
	}
	return "bs []byte"
}

func (c Conf) MarshalParam() string {
	if c.Stream {
		return "w"
	}
	return "bs[n:]"
}

func (c Conf) MarshalReturnValues() string {
	if c.Stream {
		return "(n int, err error)"
	}
	return "(n int)"
}

func (c Conf) UnmarshalParamSignature() string {
	if c.Stream {
		return "r muss.Reader"
	}
	return "bs []byte"
}

func (c Conf) UnmarshalParam() string {
	if c.Stream {
		return "r"
	}
	return "bs[n:]"
}

func (c Conf) SkipParamSignature() string {
	if c.Stream {
		return "r muss.Reader"
	}
	return "bs []byte"
}

func (c Conf) SkipParam() string {
	if c.Stream {
		return "r"
	}
	return "bs[n:]"
}
