package builders

import (
	"fmt"
	"reflect"

	"github.com/mus-format/musgen-go/classifier"
	"github.com/mus-format/musgen-go/data"
	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
)

// TODO TypeDataBuilder should use Parser interface:
//
// type Parser interface {
// 	ParseTypeName(t reflect.Type) (typename.TypeName, error)
// 	ParseSourceTypeName(t reflect.Type) (typename.TypeName, error)
// }

func NewTypeDataBuilder(conv TypeNameConvertor, gops genops.Options) TypeDataBuilder {
	return TypeDataBuilder{conv, gops}
}

type TypeDataBuilder struct {
	conv TypeNameConvertor
	gops genops.Options
}

func (b TypeDataBuilder) BuildDefinedTypeData(t reflect.Type,
	tops *typeops.Options) (d data.TypeData, err error) {
	if !classifier.DefinedBasicType(t) {
		err = b.notDefinedTypeError(t)
		return
	}
	var (
		typeName   typename.FullName
		sourceType typename.FullName
	)
	if typeName, err = b.parseTypeName(t); err != nil {
		return
	}
	if sourceType, err = b.parseSourceTypeName(t); err != nil {
		return
	}
	d.FullName = typeName
	d.SourceFullName = sourceType
	d.Fields = []data.FieldData{{FullName: d.SourceFullName}}
	d.Tops = tops
	d.Gops = b.gops
	return
}

func (b TypeDataBuilder) BuildStructData(t reflect.Type, sops structops.Options) (
	d data.TypeData, err error) {
	if !classifier.DefinedStruct(t) {
		err = b.notStructError(t)
		return
	}
	if err = b.checkFields(t, sops); err != nil {
		return
	}
	var (
		typeName typename.FullName
		fields   []typename.FullName
	)
	if typeName, err = b.parseTypeName(t); err != nil {
		return
	}
	if fields, err = b.parseFields(t); err != nil {
		return
	}
	d.FullName = typeName
	d.Fields = b.fieldsToData(t, fields)
	d.Sops = sops
	d.Gops = b.gops
	return
}

func (b TypeDataBuilder) BuildInterfaceData(t reflect.Type, iops introps.Options) (
	d data.TypeData, err error) {
	if !classifier.DefinedInterface(t) {
		err = b.notInterfaceError(t)
		return
	}
	var (
		typeName typename.FullName
		impls    []typename.FullName
	)
	if typeName, err = b.parseTypeName(t); err != nil {
		return
	}
	if impls, err = b.parseImpls(iops); err != nil {
		return
	}
	d.FullName = typeName
	d.Impls = impls
	d.Iops = iops
	d.Gops = b.gops
	return
}

func (b TypeDataBuilder) BuildDTSData(t reflect.Type) (
	d data.TypeData, err error) {
	if !(classifier.DefinedBasicType(t) || classifier.DefinedStruct(t) ||
		classifier.DefinedInterface(t)) {
		err = NewUnsupportedTypeError(t)
		return
	}
	d.FullName, err = b.parseTypeName(t)
	d.Gops = b.gops
	return
}

func (b TypeDataBuilder) BuildTimeData(t reflect.Type, tops *typeops.Options) (
	d data.TypeData, err error) {
	if !classifier.DefinedStruct(t) {
		err = NewNotStructError(t)
		return
	}
	typeName, err := b.parseTypeName(t)
	if err != nil {
		return
	}
	d.FullName = typeName
	d.SourceFullName = "time.Time"
	d.Fields = []data.FieldData{{FullName: d.SourceFullName}}
	d.Tops = tops
	d.Gops = b.gops
	return
}

func (b TypeDataBuilder) ToFullName(t reflect.Type) (fname typename.FullName) {
	cname, err := typename.TypeCompleteName(t)
	if err != nil {
		return
	}
	fname, err = b.conv.ConvertToFullName(cname)
	if err != nil {
		panic(fmt.Sprintf("can't convert %v to FullName, cause: %v", t, err))
	}
	return
}

func (b TypeDataBuilder) ToRelName(fname typename.FullName) (
	rname typename.RelName) {
	rname, err := b.conv.ConvertToRelName(fname)
	if err != nil {
		panic(fmt.Sprintf("can't convert %v to RelName, cause: %v", fname, err))
	}
	return
}

func (b TypeDataBuilder) checkFields(t reflect.Type, sops structops.Options) (
	err error) {
	if len(sops.Fields) == 0 {
		return
	}
	var (
		want   = t.NumField()
		actual = len(sops.Fields)
	)
	if actual != want {
		err = NewWrongFieldsCountError(want)
	}
	return
}

func (b TypeDataBuilder) fieldsToData(t reflect.Type,
	fields []typename.FullName) (sl []data.FieldData) {
	sl = make([]data.FieldData, len(fields))
	for i := range fields {
		sl[i] = data.FieldData{
			FullName:  fields[i],
			FieldName: t.Field(i).Name,
		}
	}
	return
}

func (b TypeDataBuilder) parseTypeName(t reflect.Type) (
	typeName typename.FullName, err error) {
	cname, err := typename.TypeCompleteName(t)
	if err != nil {
		return
	}
	return b.conv.ConvertToFullName(cname)
}

func (b TypeDataBuilder) parseSourceTypeName(t reflect.Type) (
	typeName typename.FullName, err error) {
	cname, err := typename.SourceTypeCompleteName(t)
	if err != nil {
		return
	}
	return b.conv.ConvertToFullName(cname)
}

func (b TypeDataBuilder) parseFields(t reflect.Type) (
	fields []typename.FullName, err error) {
	fields = make([]typename.FullName, t.NumField())
	for i := range t.NumField() {
		ft := t.Field(i).Type
		if fields[i], err = b.parseTypeName(ft); err != nil {
			return
		}
	}
	return
}

func (b TypeDataBuilder) parseImpls(iops introps.Options) (
	impls []typename.FullName, err error) {
	impls = make([]typename.FullName, len(iops.Impls))
	for i, impl := range iops.Impls {
		if impls[i], err = b.parseTypeName(impl); err != nil {
			return
		}
	}
	return
}

func (b TypeDataBuilder) notDefinedTypeError(t reflect.Type) error {
	switch {
	case classifier.DefinedStruct(t):
		return NewUnexpectedStructTypeError(t)
	case classifier.DefinedInterface(t):
		return NewUnexpectedInterfaceTypeError(t)
	default:
		return NewUnsupportedTypeError(t)
	}
}

func (b TypeDataBuilder) notStructError(t reflect.Type) error {
	switch {
	case classifier.DefinedBasicType(t):
		return NewUnexpectedDefinedTypeError(t)
	case classifier.DefinedInterface(t):
		return NewUnexpectedInterfaceTypeError(t)
	default:
		return NewUnsupportedTypeError(t)
	}
}

func (b TypeDataBuilder) notInterfaceError(t reflect.Type) error {
	switch {
	case classifier.DefinedBasicType(t):
		return NewUnexpectedDefinedTypeError(t)
	case classifier.DefinedStruct(t):
		return NewUnexpectedStructTypeError(t)
	default:
		return NewUnsupportedTypeError(t)
	}
}
