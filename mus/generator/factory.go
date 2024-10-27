package generator

import "github.com/mus-format/musgen-go/basegen"

type GeneratorFactory struct{}

func (f GeneratorFactory) NewNumGenerator(conf basegen.Conf, tp string,
	meta *basegen.Metadata) basegen.Generator {
	return NewNumGenerator(conf, tp, meta)
}

func (f GeneratorFactory) NewStringGenerator(conf basegen.Conf,
	meta *basegen.Metadata) basegen.Generator {
	return NewStringGenerator(conf, meta)
}

func (f GeneratorFactory) NewBoolGenerator(conf basegen.Conf, tp string,
	meta *basegen.Metadata) basegen.Generator {
	return NewBoolGenerator(conf, tp, meta)
}

func (f GeneratorFactory) NewSliceGenerator(conf basegen.Conf, tp, prefix string,
	meta *basegen.Metadata) basegen.Generator {
	return NewSliceGenerator(conf, tp, prefix, meta)
}

func (f GeneratorFactory) NewMapGenerator(conf basegen.Conf, tp, prefix string,
	meta *basegen.Metadata) basegen.Generator {
	return NewMapGenerator(conf, tp, prefix, meta)
}

func (f GeneratorFactory) NewPtrGenerator(conf basegen.Conf, tp, prefix string,
	meta *basegen.Metadata) basegen.Generator {
	return NewPtrGenerator(conf, tp, prefix, meta)
}

func (f GeneratorFactory) NewSAIGenerator(conf basegen.Conf, tp, prefix string,
	meta *basegen.Metadata) basegen.Generator {
	return NewSAIGenerator(conf, tp, prefix, meta)
}

var GenerateSubFn = basegen.BuildGenerateSubFn(GeneratorFactory{})
