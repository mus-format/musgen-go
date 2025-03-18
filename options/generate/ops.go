package genops

import (
	"crypto/md5"
	"strconv"
)

type Options struct {
	Package   string   // Defines the package of the generated file.
	Unsafe    bool     // If true, an unsafe code will be generated.
	NotUnsafe bool     // If true, a not unsafe code will be generated.
	Stream    bool     // If true, a streaming code will be generated.
	Imports   []string // These strings will be added to the imports of the generated file.
}

func (o Options) ModImportName() string {
	if o.Stream {
		return "muss"
	}
	return "mus"
}

func (o Options) MarshalSignatureLastParam() string {
	if o.Stream {
		return "w muss.Writer"
	}
	return "bs []byte"
}

func (o Options) MarshalLastParam(first bool) string {
	if o.Stream {
		return "w"
	}
	if first {
		return "bs"
	}
	return "bs[n:]"
}

func (o Options) UnmarshalSignatureLastParam() string {
	if o.Stream {
		return "r muss.Reader"
	}
	return "bs []byte"
}

func (o Options) UnmarshalLastParam(first bool) string {
	if o.Stream {
		return "r"
	}
	if first {
		return "bs"
	}
	return "bs[n:]"
}

func (o Options) SkipSignatureLastParam() string {
	if o.Stream {
		return "r muss.Reader"
	}
	return "bs []byte"
}

func (o Options) SkipLastParam() string {
	if o.Stream {
		return "r"
	}
	return "bs[n:]"
}

func (o Options) Hash() [16]byte {
	var bs = []byte{}
	bs = append(bs, []byte(o.Package)...)
	bs = append(bs, []byte(strconv.FormatBool(o.Stream))...)
	bs = append(bs, []byte(strconv.FormatBool(o.Unsafe))...)
	for i := range o.Imports {
		bs = append(bs, []byte(o.Imports[i])...)
	}
	return md5.Sum(bs)
}

type SetOption func(o *Options)

func WithPackage(p string) SetOption {
	return func(o *Options) {
		o.Package = p
	}
}

func WithUnsafe() SetOption {
	return func(o *Options) {
		o.Unsafe = true
	}
}

func WithNotUnsafe() SetOption {
	return func(o *Options) {
		o.NotUnsafe = true
	}
}

func WithStream() SetOption {
	return func(o *Options) {
		o.Stream = true
	}
}

func WithImports(imports []string) SetOption {
	return func(o *Options) {
		o.Imports = imports
	}
}

func Apply(ops []SetOption, o *Options) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
