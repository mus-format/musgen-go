package basegen

type Generator interface {
	GenerateFnName(fnType FnType) (name string)
	GenerateMarshalCall(vname string) (call string)
	GenerateUnmarshalCall() (call string)
	GenerateSizeCall(vname string) (call string)
	GenerateSkipCall() (call string)
}

type GeneratorFactory interface {
	NewNumGenerator(conf Conf, tp string, meta *Metadata) Generator
	NewStringGenerator(conf Conf, meta *Metadata) Generator
	NewBoolGenerator(conf Conf, tp string, meta *Metadata) Generator
	NewSliceGenerator(conf Conf, tp, prefix string, meta *Metadata) Generator
	NewMapGenerator(conf Conf, tp, prefix string, meta *Metadata) Generator
	NewPtrGenerator(conf Conf, tp, prefix string, meta *Metadata) Generator
	NewSAIGenerator(conf Conf, tp, prefix string, meta *Metadata) Generator
}
