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
	Oneof   []string
	Fields  []FieldDesc
}

type FieldDesc struct {
	Name    string
	Type    string
	Options *Options
}

func BuildAliasTypeDesc(conf Conf, tp reflect.Type, prefix string,
	opts *Options) (td TypeDesc, err error) {
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
			Type:    aliasOf,
			Options: (*Options)(opts),
		},
	}
	return
}

func BuildStructTypeDesc(conf Conf, tp reflect.Type, prefix string,
	opts []*Options) (td TypeDesc, err error) {
	_, _, fieldsTypes, err := parser.Parse(tp)
	if err != nil {
		return
	}
	if opts != nil && len(opts) != len(fieldsTypes) {
		err = ErrWrongOptionsAmount
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
		if opts != nil {
			// TODO Check Options type
			td.Fields[i].Options = (*Options)(opts[i])
		}
	}
	return
}

func BuildInterfaceTypeDesc(conf Conf, tp reflect.Type, prefix string,
	opts Options) (td TypeDesc, err error) {
	intr, _, _, err := parser.Parse(tp)
	if err != nil {
		return
	}
	if !intr {
		err = ErrNotInterface
	}
	if len(opts.Oneof) == 0 {
		err = ErrEmptyOneof
	}
	for i := 0; i < len(opts.Oneof); i++ {
		_, _, _, err = parser.Parse(opts.Oneof[i])
		if err != nil {
			panic(err)
		}
	}
	td.Conf = conf
	td.Package = pkg(tp)
	td.Prefix = prefix
	td.Name = tp.Name()
	td.Oneof = make([]string, len(opts.Oneof))
	for i := 0; i < len(opts.Oneof); i++ {
		oneOf := opts.Oneof[i]
		if tp.PkgPath() != oneOf.PkgPath() {
			td.Oneof[i] = oneOf.String()
		} else {
			td.Oneof[i] = oneOf.Name()
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
