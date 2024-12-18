package basegen

type TypeOptionsBuilder interface {
	BuildTypeOptions() *Options
}

type BoolOptions struct {
	Validator string
}

func (m BoolOptions) BuildTypeOptions() *Options {
	return &Options{}
}

type NumOptions struct {
	Encoding  NumEncoding
	Validator string
}

func (m NumOptions) BuildTypeOptions() *Options {
	return &Options{Encoding: m.Encoding, Validator: m.Validator}
}

type StringOptions struct {
	LenEncoding  NumEncoding
	LenValidator string
	Validator    string
}

func (m StringOptions) BuildTypeOptions() *Options {
	return &Options{
		LenEncoding:  m.LenEncoding,
		LenValidator: m.LenValidator,
		Validator:    m.Validator,
	}
}

type SliceOptions struct {
	LenEncoding  NumEncoding
	LenValidator string
	Elem         TypeOptionsBuilder
	Validator    string
}

func (o SliceOptions) BuildTypeOptions() *Options {
	opts := &Options{
		LenEncoding:  o.LenEncoding,
		LenValidator: o.LenValidator,
		Validator:    o.Validator,
	}
	if o.Elem != nil {
		opts.Elem = o.Elem.BuildTypeOptions()
	}
	return opts
}

type ArrayOptions struct {
	LenEncoding NumEncoding
	Elem        TypeOptionsBuilder
	Validator   string
}

func (o ArrayOptions) BuildTypeOptions() *Options {
	opts := &Options{
		LenEncoding: o.LenEncoding,
		Validator:   o.Validator,
	}
	if o.Elem != nil {
		opts.Elem = o.Elem.BuildTypeOptions()
	}
	return opts
}

type MapOptions struct {
	LenEncoding  NumEncoding
	LenValidator string
	Key          TypeOptionsBuilder
	Elem         TypeOptionsBuilder
	Validator    string
}

func (m MapOptions) BuildTypeOptions() *Options {
	tm := &Options{
		LenEncoding:  m.LenEncoding,
		LenValidator: m.LenValidator,
		Validator:    m.Validator,
	}
	if m.Key != nil {
		tm.Key = m.Key.BuildTypeOptions()
	}
	if m.Elem != nil {
		tm.Elem = m.Elem.BuildTypeOptions()
	}
	return tm
}

type PtrOptions struct {
	Validator string
}

func (m PtrOptions) BuildTypeOptions() *Options {
	return &Options{
		Validator: m.Validator,
	}
}
