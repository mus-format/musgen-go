package basegen

type StructOptionsBuilder interface {
	BuildStructOptions() []*Options
}

type StructOptions []FieldOptionsBuilder

func (m StructOptions) BuildStructOptions() (sl []*Options) {
	if m != nil {
		sl = make([]*Options, len(m))
		for i := 0; i < len(m); i++ {
			if m[i] != nil {
				fm := (FieldOptionsBuilder)(m[i]).BuildFieldOptions()
				sl[i] = fm
			}
		}
	}
	return
}

type FieldOptionsBuilder interface {
	BuildFieldOptions() *Options
}

type BoolFieldOptions struct {
	BoolOptions
	Ignore bool
}

func (m BoolFieldOptions) BuildFieldOptions() (opts *Options) {
	opts = m.BoolOptions.BuildTypeOptions()
	opts.Ignore = m.Ignore
	return
}

type NumFieldOptions struct {
	NumOptions
	Ignore bool
}

func (m NumFieldOptions) BuildFieldOptions() (opts *Options) {
	opts = m.NumOptions.BuildTypeOptions()
	opts.Ignore = m.Ignore
	return
}

type StringFieldOptions struct {
	StringOptions
	Ignore bool
}

func (m StringFieldOptions) BuildFieldOptions() (opts *Options) {
	opts = m.StringOptions.BuildTypeOptions()
	opts.Ignore = m.Ignore
	return
}

type SliceFieldOptions struct {
	SliceOptions
	Ignore bool
}

func (m SliceFieldOptions) BuildFieldOptions() (opts *Options) {
	opts = m.SliceOptions.BuildTypeOptions()
	opts.Ignore = m.Ignore
	return
}

type ArrayFieldOptions struct {
	ArrayOptions
	Ignore bool
}

func (m ArrayFieldOptions) BuildFieldOptions() (opts *Options) {
	opts = m.ArrayOptions.BuildTypeOptions()
	opts.Ignore = m.Ignore
	return
}

type MapFieldOptions struct {
	MapOptions
	Ignore bool
}

func (m MapFieldOptions) BuildFieldOptions() (opts *Options) {
	opts = m.MapOptions.BuildTypeOptions()
	opts.Ignore = m.Ignore
	return
}

type PtrFieldOptions struct {
	PtrOptions
	Ignore bool
}

func (m PtrFieldOptions) BuildTypeOptions() (opts *Options) {
	opts = m.PtrOptions.BuildTypeOptions()
	opts.Ignore = m.Ignore
	return
}

type CustomTypeFieldOptions struct {
	Prefix    string
	Ignore    bool
	Validator string
}

func (m CustomTypeFieldOptions) BuildFieldOptions() *Options {
	return &Options{
		Prefix:    m.Prefix,
		Ignore:    m.Ignore,
		Validator: m.Validator,
	}
}
