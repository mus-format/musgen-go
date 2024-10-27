//go:generate go run fvar/main.go
package basegen

import (
	"bytes"
	"reflect"

	"text/template"
)

// New creates a new Basegen.
func New(conf Conf, fn GenerateFnCall) (g *Basegen, err error) {
	baseTmpl := template.New("base")
	registerFuncs(baseTmpl, fn)
	err = registerTemplates(baseTmpl)
	if err != nil {
		return nil, err
	}
	return &Basegen{conf, baseTmpl}, err
}

type Basegen struct {
	conf     Conf
	baseTmpl *template.Template
}

func (g *Basegen) GenerateInterface(tp reflect.Type, prefix string,
	meta Metadata) (bs []byte, err error) {
	td, err := BuildInterfaceTypeDesc(g.conf, tp, prefix, meta)
	if err != nil {
		return
	}
	tmplFile := "main_interface.tmpl"
	return g.generate(td, tmplFile)
}

func (g *Basegen) GenerateAlias(tp reflect.Type, prefix string,
	meta *Metadata) (bs []byte, err error) {
	return g.generateAlias(tp, prefix, meta, "main_alias_struct.tmpl")
}

func (g *Basegen) GenerateStruct(tp reflect.Type, prefix string,
	meta []*Metadata) (bs []byte, err error) {
	return g.generateStruct(tp, prefix, meta, "main_alias_struct.tmpl")
}

func (g *Basegen) GenerateAliasDTS(tp reflect.Type, prefix string,
	meta *Metadata) (bs []byte, err error) {
	return g.generateAlias(tp, prefix, meta, "main_dts.tmpl")
}

func (g *Basegen) GenerateStructDTS(tp reflect.Type, prefix string,
	meta []*Metadata) (bs []byte, err error) {
	return g.generateStruct(tp, prefix, meta, "main_dts.tmpl")
}

// // TODO tps should be ordered by DTMs.
// func (g *Basegen) GenerateDVS(tps []reflect.Type) (bs []byte, err error) {
// 	sl := make([]string, len(tps))
// 	for i := 0; i < len(tps); i++ {
// 		sl = append(sl, tps[i].Name())
// 	}
// 	return g.generate(sl, "main_dvs.tmpl")
// }

func (g Basegen) generateAlias(tp reflect.Type, prefix string,
	meta *Metadata, tmplFile string) (bs []byte, err error) {
	td, err := BuildAliasTypeDesc(g.conf, tp, prefix, meta)
	if err != nil {
		return
	}
	return g.generate(td, tmplFile)
}

func (g Basegen) generateStruct(tp reflect.Type, prefix string,
	meta []*Metadata, tmplFile string) (bs []byte, err error) {
	td, err := BuildStructTypeDesc(g.conf, tp, prefix, meta)
	if err != nil {
		return
	}
	return g.generate(td, tmplFile)
}

func (g Basegen) generate(td any, tmplFile string) (bs []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	err = g.baseTmpl.ExecuteTemplate(buf, tmplFile, td)
	if err != nil {
		return
	}
	bs = buf.Bytes()
	return
}

func registerFuncs(tmpl *template.Template, fn GenerateFnCall) {
	tmpl.Funcs(map[string]any{
		"GenerateFnCall": fn,
		"FieldsLen":      FieldsLen,
		"Fields":         Fields,
		"Receiver":       Receiver,
		"CurrentTypeOf":  CurrentTypeOf,
		"minus":          MakeMinusFunc(),
		"include":        MakeIncludeFunc(tmpl),
	})
}

func registerTemplates(tmpl *template.Template) (err error) {
	var childTmpl *template.Template
	for name, template := range templates {
		childTmpl = tmpl.New(name)
		_, err = childTmpl.Parse(template)
		if err != nil {
			return
		}
	}
	return nil
}
