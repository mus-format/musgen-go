package scanner

type Options struct {
	WithoutParams bool
	Ignore        bool
}

type SetOption func(o *Options)

func WithoutParams() SetOption {
	return func(o *Options) { o.WithoutParams = true }
}

func Apply(ops []SetOption, o *Options) {
	for i := range ops {
		ops[i](o)
	}
}
