package basegen

type Generator interface {
	GenerateFnName(fnType FnType) (name string)
	GenerateMarshalCall(vname string) (call string)
	GenerateUnmarshalCall() (call string)
	GenerateSizeCall(vname string) (call string)
	GenerateSkipCall() (call string)
}

type GeneratorFactory interface {
	NewNumGenerator(conf Conf, tp string, opts *Options) Generator
	NewStringGenerator(conf Conf, opts *Options) Generator
	NewBoolGenerator(conf Conf, tp string, opts *Options) Generator
	NewSliceGenerator(conf Conf, tp, prefix string, opts *Options) Generator
	NewMapGenerator(conf Conf, tp, prefix string, opts *Options) Generator
	NewPtrGenerator(conf Conf, tp, prefix string, opts *Options) Generator
	NewSAIGenerator(conf Conf, tp, prefix string, opts *Options) Generator
}
