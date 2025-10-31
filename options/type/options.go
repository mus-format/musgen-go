// Package typeops provides options for registering defined types in the
// code generator.
package typeops

import (
	"crypto/md5"
	"strconv"
)

// Options configures code generation behavior for a type.
type Options struct {
	Ignore bool

	NumEncoding NumEncoding

	SourceType SourceType
	TimeUnit   TimeUnit

	Validator    string
	LenEncoding  NumEncoding
	LenValidator string

	Key  *Options
	Elem *Options
}

func (o Options) Hash() [16]byte {
	bs := []byte{}
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
	return md5.Sum(bs)
}

func (o Options) LenSer() string {
	return o.LenEncoding.LenSer()
}

type SetOption func(o *Options)

// WithIgnore returns a SetOption that marks the field to be excluded from code
// generation. When true:
//   - The field will be omitted from serialization/deserialization logic.
//   - No validation will be generated for this field.
//
// This is typically used for:
//   - Private fields.
//   - Fields marked as deprecated
//
// Applies to: struct fields only.
func WithIgnore() SetOption {
	return func(o *Options) { o.Ignore = true }
}

// WithNumEncoding returns a SetOption that sets the binary encoding scheme
// (Varint or Raw) for numeric values.
//
// Applies to: all integer and float types (int, uint, float32, float64, etc.)
func WithNumEncoding(enc NumEncoding) SetOption {
	return func(o *Options) { o.NumEncoding = enc }
}

// WithTimeUnit returns a SetOption that specifies the source of the defined
// type. When set, activates type-specific generation rules and may influence
// the available options.
func WithSourceType(t SourceType) SetOption {
	return func(o *Options) { o.SourceType = t }
}

// WithTimeUnit returns a SetOption that specifies the time unit precision.
//
// Applies to: time.Time or types marked with SourceType = Time.
func WithTimeUnit(tu TimeUnit) SetOption {
	return func(o *Options) { o.TimeUnit = tu }
}

// WithValidator returns a SetOption that registers a validation function, which
// should have the following signature:
//
//	func(value T) error
//
// Example:
//   - ValidateEmail       // Local package function.
//   - pkg.CheckRange      // External package function.
//
// Applies to: all types.
func WithValidator(validator string) SetOption {
	return func(o *Options) { o.Validator = validator }
}

// WithLenEncoding returns a SetOption that sets the binary encoding scheme
// (Varint or Raw) used for serializing length of variable-length data types.
//
// Applies to: string, array, slice, map types.
func WithLenEncoding(enc NumEncoding) SetOption {
	return func(o *Options) { o.LenEncoding = enc }
}

// WithLenValidator returns a SetOption that registers a name of the length
// validator function, which should have the following signature:
//
//	func(length int) error
//
// Example:
//   - ValidateLength       	 // Local package function.
//   - pkg.ValidateLength      // External package function.
//
// Applies to: string, array, slice, map types.
func WithLenValidator(validator string) SetOption {
	return func(o *Options) { o.LenValidator = validator }
}

// WithKey returns a SetOption that configures type-specific serialization and
// validation options for map keys.
//
// Applies to: map type.
func WithKey(ops ...SetOption) SetOption {
	return func(o *Options) {
		ko := &Options{}
		Apply(ops, ko)
		o.Key = ko
	}
}

// WithElem returns a SetOption that configures type-specific serialization and
// validation options for collection elements.
//
// Applies to: array, slice, map types.
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
