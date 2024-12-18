package generator

import "github.com/mus-format/musgen-go/basegen"

type GeneratorFactory struct{}

func (f GeneratorFactory) NewNumGenerator(conf basegen.Conf, tp string,
	opts *basegen.Options) basegen.Generator {
	return NewNumGenerator(conf, tp, opts)
}

func (f GeneratorFactory) NewStringGenerator(conf basegen.Conf,
	opts *basegen.Options) basegen.Generator {
	return NewStringGenerator(conf, opts)
}

func (f GeneratorFactory) NewBoolGenerator(conf basegen.Conf, tp string,
	opts *basegen.Options) basegen.Generator {
	return NewBoolGenerator(conf, tp, opts)
}

func (f GeneratorFactory) NewSliceGenerator(conf basegen.Conf, tp, prefix string,
	opts *basegen.Options) basegen.Generator {
	return NewSliceGenerator(conf, tp, prefix, opts)
}

func (f GeneratorFactory) NewMapGenerator(conf basegen.Conf, tp, prefix string,
	opts *basegen.Options) basegen.Generator {
	return NewMapGenerator(conf, tp, prefix, opts)
}

func (f GeneratorFactory) NewPtrGenerator(conf basegen.Conf, tp, prefix string,
	opts *basegen.Options) basegen.Generator {
	return NewPtrGenerator(conf, tp, prefix, opts)
}

func (f GeneratorFactory) NewSAIGenerator(conf basegen.Conf, tp, prefix string,
	opts *basegen.Options) basegen.Generator {
	return NewSAIGenerator(conf, tp, prefix, opts)
}

var GenerateSubFn = basegen.BuildGenerateSubFn(GeneratorFactory{})
