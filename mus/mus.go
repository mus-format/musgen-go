//go:generate go run fvar/main.go
package mus

import (
	"bytes"
	"html/template"
	"reflect"

	"github.com/mus-format/musgen-go/basegen"
	"github.com/mus-format/musgen-go/mus/generator"
	"golang.org/x/tools/imports"
)

// NewFileGenerator creates a new FileGenerator.
func NewFileGenerator(conf basegen.Conf) (g FileGenerator, err error) {
	bg, err := basegen.New(conf,
		basegen.BuildGenerateFnCall(generator.GeneratorFactory{}))
	if err != nil {
		return
	}
	bs, err := makeFileHeader(conf)
	if err != nil {
		return
	}
	g = FileGenerator{
		bg: bg,
		bs: bs,
	}
	return
}

// FileGenerator generates Marshal/Unmarshal/Size/Skip functions that
// implement the MUS serialization format for alias, interface, struct or DTS
// types.
// More info about DTS could be found at https://github.com/mus-format/mus-dts-go.
type FileGenerator struct {
	bg *basegen.Basegen
	bs []byte
}

// AddAlias adds an alias type to the generator.
//
// Returns basegen.ErrNotAlias if type is not an alias.
func (b *FileGenerator) AddAlias(tp reflect.Type) (err error) {
	return b.AddAliasWith(tp, "", nil)
}

// AddAliasWith adds an alias type to the generator.
//
// Prefix allows to have several Marshal/Unmarshal/Size/Skip functions for one
// alias type. For example, at the same time we can have MarshalAlias and
// MarshaUnsafeAlias, where Unsafe is a prefix. With opts, we can configure the
// serialization of an alias type, specify encoding, validator, etc.
//
// Returns basegen.ErrNotAlias if type is not an alias.
func (b *FileGenerator) AddAliasWith(tp reflect.Type, prefix string,
	opts basegen.TypeOptionsBuilder) (err error) {
	var m *basegen.Options
	if opts != nil {
		m = opts.BuildTypeOptions()
	}
	bs, err := b.bg.GenerateAlias(tp, prefix, m)
	if err != nil {
		return
	}
	b.bs = append(b.bs, bs...)
	b.bs = append(b.bs, []byte("\n")...)
	return
}

// AddStruct adds a struct type to the generator.
//
// Returns besegen.ErrNotStruct if type is not a struct.
func (b *FileGenerator) AddStruct(tp reflect.Type) (err error) {
	return b.AddStructWith(tp, "", nil)
}

// AddStruct adds a struct type to the generator.
//
// Prefix allows to have several Marshal/Unmarshal/Size/Skip functions for one
// struct type. For example, at the same time we can have MarshalStruct and
// MarshaUnsafeStruct, where Unsafe is a prefix. With opts, we can configure the
// serialization of struct fields, specify encoding, validator, skip a field, etc.
//
// Returns basegen.ErrNotStruct if type is not a struct.
func (b *FileGenerator) AddStructWith(tp reflect.Type, prefix string,
	opts basegen.StructOptionsBuilder) (err error) {
	var m []*basegen.Options
	if opts != nil {
		m = opts.BuildStructOptions()
	}
	bs, err := b.bg.GenerateStruct(tp, prefix, m)
	if err != nil {
		return
	}
	b.bs = append(b.bs, bs...)
	b.bs = append(b.bs, []byte("\n")...)
	return
}

// AddStruct adds an interface type to the generator.
//
// Returns besegen.ErrNotInterface if type is not an interface.
func (b *FileGenerator) AddInterface(tp reflect.Type,
	opts basegen.InterfaceOptionsBuilder) (err error) {
	return b.AddInterfaceWith(tp, "", opts)
}

// AddInterfaceWith adds an interface type to the generator.
//
// Prefix allows to have several Marshal/Unmarshal/Size/Skip functions for one
// interface type. For example, at the same time we can have MarshalInterface
// and MarshaUnsafeInterface, where Unsafe is a prefix. Opts must contain one
// or more interface implementation types.
//
// Returns basegen.ErrNotInterface if type is not an interface or
// basegen.ErrEmptyOneof if opts's Oneof property is empty.
func (b *FileGenerator) AddInterfaceWith(tp reflect.Type, prefix string,
	opts basegen.InterfaceOptionsBuilder) (err error) {
	var m basegen.Options
	if opts != nil {
		m, err = opts.BuildInterfaceOptions()
		if err != nil {
			return
		}
	}
	bs, err := b.bg.GenerateInterface(tp, prefix, m)
	if err != nil {
		return
	}
	b.bs = append(b.bs, bs...)
	b.bs = append(b.bs, []byte("\n")...)
	return
}

// AddAliasDTS adds an alias DTS to the generator.
//
// Returns basegen.ErrNotAlias if type is not an alias.
func (b *FileGenerator) AddAliasDTS(tp reflect.Type) (err error) {
	return b.AddAliasDTSWith(tp, "", nil)
}

// AddAliasDTSWith adds an alias DTS to the generator.
//
// Prefix allows to have several Marshal/Unmarshal/Size/Skip functions for one
// alias type. For example, at the same time we can have MarshalAlias and
// MarshaUnsafeAlias, where Unsafe is a prefix. With opts, we can configure the
// serialization of an alias type, specify encoding, validator, etc.
//
// Returns basegen.ErrNotAlias if type is not an alias.
func (b *FileGenerator) AddAliasDTSWith(tp reflect.Type, prefix string,
	opts basegen.TypeOptionsBuilder) (err error) {
	var m *basegen.Options
	if opts != nil {
		m = opts.BuildTypeOptions()
	}
	bs, err := b.bg.GenerateAliasDTS(tp, prefix, m)
	if err != nil {
		return
	}
	b.bs = append(b.bs, bs...)
	b.bs = append(b.bs, []byte("\n")...)
	return
}

// AddStructDTS adds a struct DTS to the generator.
//
// Returns basegen.ErrNotStruct if type is not a struct.
func (b *FileGenerator) AddStructDTS(tp reflect.Type) (err error) {
	return b.AddStructDTSWith(tp, "", nil)
}

// AddStructDTS adds a struct DTS to the generator.
//
// Prefix allows to have several Marshal/Unmarshal/Size/Skip functions for one
// struct/DTS type. For example, at the same time we can have MarshalStruct and
// MarshaUnsafeStruct, where Unsafe is a prefix. With opts, we can configure the
// serialization of struct fields, specify encoding, validator, skip a field,
// etc.
//
// Returns basegen.ErrNotStruct if type is not a struct.
func (b *FileGenerator) AddStructDTSWith(tp reflect.Type, prefix string,
	opts basegen.StructOptionsBuilder) (err error) {
	var m []*basegen.Options
	if opts != nil {
		m = opts.BuildStructOptions()
	}
	bs, err := b.bg.GenerateStructDTS(tp, prefix, m)
	if err != nil {
		return
	}
	b.bs = append(b.bs, bs...)
	b.bs = append(b.bs, []byte("\n")...)
	return
}

// func (b *FileGenerator) AddDVS(tps []reflect.Type) (
// 	err error) {
// 	bs, err := b.bg.GenerateDVS(tps)
// 	if err != nil {
// 		return
// 	}
// 	b.bs = append(b.bs, bs...)
// 	b.bs = append(b.bs, []byte("\n")...)
// 	return
// }

// Generate generates the file content.
func (b *FileGenerator) Generate() (bs []byte, err error) {
	return imports.Process("", b.bs, nil)
}

func makeFileHeader(conf basegen.Conf) (bs []byte, err error) {
	var (
		fileHeaderTmpl = template.New("file_header.tmpl")
		buf            = bytes.NewBuffer(make([]byte, 0))
	)
	_, err = fileHeaderTmpl.Parse(templates["file_header.tmpl"])
	if err != nil {
		return
	}
	err = fileHeaderTmpl.ExecuteTemplate(buf, "file_header.tmpl", conf)
	if err != nil {
		return
	}
	bs = buf.Bytes()
	return
}
