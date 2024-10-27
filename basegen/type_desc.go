package basegen

import (
	"reflect"
	"regexp"

	"github.com/mus-format/musgen-go/basegen/parser"
)

type FieldName string

type TypeDesc struct {
	Conf    Conf
	Package string
	Prefix  string
	Name    string
	AliasOf string
	OneOf   []string
	Fields  []FieldDesc
}

type FieldDesc struct {
	Name     string
	Type     string
	Metadata *Metadata
}

func BuildAliasTypeDesc(conf Conf, tp reflect.Type, prefix string,
	meta *Metadata) (td TypeDesc, err error) {
	_, aliasOf, _, err := parser.Parse(tp)
	if err != nil {
		return
	}
	if aliasOf == "" {
		err = ErrNotAlias
		return
	}
	td.Conf = conf
	td.Prefix = prefix
	td.Package = pkg(tp)
	td.Prefix = prefix
	td.Name = tp.Name()
	td.AliasOf = aliasOf
	td.Fields = []FieldDesc{
		{
			Type:     aliasOf,
			Metadata: (*Metadata)(meta),
		},
	}
	return
}

func BuildStructTypeDesc(conf Conf, tp reflect.Type, prefix string,
	meta []*Metadata) (td TypeDesc, err error) {
	_, _, fieldsTypes, err := parser.Parse(tp)
	if err != nil {
		return
	}
	if meta != nil && len(meta) != len(fieldsTypes) {
		err = ErrWrongMetadataAmount
		return
	}
	td.Conf = conf
	td.Prefix = prefix
	td.Package = pkg(tp)
	td.Name = tp.Name()
	td.Fields = make([]FieldDesc, len(fieldsTypes))
	for i := 0; i < len(td.Fields); i++ {
		td.Fields[i] = FieldDesc{
			Name: tp.Field(i).Name,
			Type: fieldsTypes[i],
		}
		if meta != nil {
			// TODO Check metadata type
			td.Fields[i].Metadata = (*Metadata)(meta[i])
		}
	}
	return
}

func BuildInterfaceTypeDesc(conf Conf, tp reflect.Type, prefix string,
	meta Metadata) (td TypeDesc, err error) {
	intr, _, _, err := parser.Parse(tp)
	if err != nil {
		return
	}
	if !intr {
		err = ErrNotInterface
	}
	if len(meta.OneOf) == 0 {
		err = ErrEmptyOneOf
	}
	for i := 0; i < len(meta.OneOf); i++ {
		_, _, _, err = parser.Parse(meta.OneOf[i])
		if err != nil {
			panic(err)
		}
	}
	td.Conf = conf
	td.Package = pkg(tp)
	td.Prefix = prefix
	td.Name = tp.Name()
	td.OneOf = make([]string, len(meta.OneOf))
	for i := 0; i < len(meta.OneOf); i++ {
		oneOf := meta.OneOf[i]
		if tp.PkgPath() != oneOf.PkgPath() {
			td.OneOf[i] = oneOf.String()
		} else {
			td.OneOf[i] = oneOf.Name()
		}
	}
	return
}

func pkg(tp reflect.Type) string {
	re := regexp.MustCompile(`^(.*)\.`)
	match := re.FindStringSubmatch(tp.String())
	if len(match) != 2 {
		return ""
	}
	return match[1]
}
