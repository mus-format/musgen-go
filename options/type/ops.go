package typeops

import (
	"crypto/md5"
	"strconv"
)

type Options struct {
	Ignore       bool
	NumEncoding  NumEncoding
	Validator    string
	LenEncoding  NumEncoding
	LenValidator string
	Key          *Options
	Elem         *Options
}

func (o Options) Hash() [16]byte {
	var bs = []byte{}
	bs = append(bs, []byte(strconv.FormatBool(o.Ignore))...)
	bs = append(bs, []byte(strconv.Itoa(int(o.NumEncoding)))...)
	bs = append(bs, []byte(o.Validator)...)
	bs = append(bs, []byte(strconv.Itoa(int(o.LenEncoding)))...)
	bs = append(bs, []byte(o.LenValidator)...)
	if o.Key != nil {
		kh := o.Key.Hash()
		bs = append(bs, []byte("key")...)
		bs = append(bs, kh[:]...)
	}
	if o.Elem != nil {
		eh := o.Elem.Hash()
		bs = append(bs, []byte("elem")...)
		bs = append(bs, eh[:]...)
	}
	// for i := range o.Oneof {
	// 	bs = append(bs, []byte(o.Oneof[i].Name())...)
	// }
	return md5.Sum(bs)
}

func (o Options) LenSer() string {
	return o.LenEncoding.LenSer()
}

type SetOption func(o *Options)

func WithIgnore() SetOption {
	return func(o *Options) { o.Ignore = true }
}

func WithNumEncoding(enc NumEncoding) SetOption {
	return func(o *Options) { o.NumEncoding = enc }
}

func WithValidator(validator string) SetOption {
	return func(o *Options) { o.Validator = validator }
}

func WithLenEncoding(enc NumEncoding) SetOption {
	return func(o *Options) { o.LenEncoding = enc }
}

func WithLenValidator(validator string) SetOption {
	return func(o *Options) { o.LenValidator = validator }
}

func WithKey(ops ...SetOption) SetOption {
	return func(o *Options) {
		ko := &Options{}
		Apply(ops, ko)
		o.Key = ko
	}
}

func WithElem(ops ...SetOption) SetOption {
	return func(o *Options) {
		eo := &Options{}
		Apply(ops, eo)
		o.Elem = eo
	}
}

func Apply(ops []SetOption, o *Options) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
