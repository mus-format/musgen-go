//go:generate go run gen/main.go
package musgen

import (
	"bytes"
	"reflect"
	"text/template"

	"github.com/mus-format/musgen-go/adesc"
	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/tdesc"
	"golang.org/x/tools/imports"
)

// NewFileGenerator creates a new FileGenerator.
func NewFileGenerator(ops ...genops.SetOption) *FileGenerator {
	gops := genops.Options{}
	genops.Apply(ops, &gops)
	baseTmpl := template.New("base")

	registerFuncs(baseTmpl)
	err := registerTemplates(baseTmpl)
	if err != nil {
		panic(err)
	}
	return &FileGenerator{
		gops:     gops,
		baseTmpl: baseTmpl,
		m:        map[adesc.AnonymousName]adesc.AnonymousDesc{},
		bs:       []byte{},
	}
}

// FileGenerator is responsible for generating MUS serialization code. Add one
// or more types to it and call Generate.
type FileGenerator struct {
	gops     genops.Options
	baseTmpl *template.Template
	m        map[adesc.AnonymousName]adesc.AnonymousDesc
	bs       []byte
}

// AddDefinedType adds the specified type to the FileGenerator to produce a
// serializer for it. This method supports types defined with the following
// source types: number, string, array, slice, map, pointer.
func (g *FileGenerator) AddDefinedType(t reflect.Type, ops ...typeops.SetOption) (
	err error) {
	var tops *typeops.Options
	if len(ops) > 0 {
		tops = &typeops.Options{}
		if len(ops) > 0 {
			typeops.Apply(ops, tops)
		}
	}
	td, err := tdesc.MakeDefinedTypeDesc(t, tops, g.gops)
	if err != nil {
		return
	}
	adesc.Collect(td.Fields[0].Type, tops, g.gops, g.m)
	g.bs = append(g.bs, g.generatePart("defined_type_ser.tmpl", td)...)
	return
}

// AddStruct adds the specified type to the FileGenerator to produce a
// serializer for it. This method supports types definined with the struct
// source type.
func (g *FileGenerator) AddStruct(t reflect.Type, ops ...structops.SetOption) (
	err error) {
	var sops structops.Options
	if len(ops) > 0 {
		sops = structops.Apply(ops, structops.NewOptions())
	}

	if sops.Type != nil && sops.Type.SourceType == structops.Time {
		td := tdesc.MakeTimeDesc(t, sops.Type.Ops, g.gops)
		g.bs = append(g.bs, g.generatePart("defined_type_ser.tmpl", td)...)
		return
	}

	td, err := tdesc.MakeStructDesc(t, sops, g.gops)
	if err != nil {
		return
	}
	for i := range td.Fields {
		adesc.Collect(td.Fields[i].Type, td.Fields[i].Tops, g.gops, g.m)
	}
	g.bs = append(g.bs, g.generatePart("struct_ser.tmpl", td)...)
	return
}

// AddDTS adds the specified type to the FileGenerator to produce a DTS
// definition for it. This method supports all types acceptable by the
// AddDefinedType, AddStruct, and AddInterface methods.
func (g *FileGenerator) AddDTS(t reflect.Type) (err error) {
	td, err := tdesc.MakeDTSDesc(t, g.gops)
	if err != nil {
		return
	}
	g.bs = append(g.bs, g.generatePart("dts.tmpl", td)...)
	return
}

// AddInterface adds the specified type to the FileGenerator to produce a
// serializer for it. This method supports types definined with the interface
// source type.
func (g *FileGenerator) AddInterface(t reflect.Type, ops ...introps.SetOption) (
	err error) {
	iops := introps.NewOptions(t)
	if ops != nil {
		introps.Apply(ops, &iops)
	}
	td, err := tdesc.MakeInterfaceDesc(t, iops, g.gops)
	if err != nil {
		return
	}
	g.bs = append(g.bs, g.generatePart("interface_ser.tmpl", td)...)
	return
}

// Generate produces the serialization code. The result should then be saved to
// a file.
func (g *FileGenerator) Generate() (bs []byte, err error) {
	bs = g.generatePackage()
	bs = append(bs, g.generateImports()...)
	bs = append(bs, g.generateAnonymosDefinitions()...)
	bs = append(bs, g.bs...)

	bs, err = imports.Process("", bs, nil)
	return
}

func (g *FileGenerator) generatePackage() (bs []byte) {
	return g.generatePart("package.tmpl", g.gops)
}

func (g *FileGenerator) generateImports() (bs []byte) {
	return g.generatePart("imports.tmpl", g.gops)
}

func (g *FileGenerator) generateAnonymosDefinitions() (bs []byte) {
	return g.generatePart("anonymous_definitions.tmpl", struct {
		Map map[adesc.AnonymousName]adesc.AnonymousDesc
		Ops genops.Options
	}{
		Map: g.m,
		Ops: g.gops,
	})
}

func (g *FileGenerator) generatePart(tmplName string, a interface{}) (bs []byte) {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := g.baseTmpl.ExecuteTemplate(buf, tmplName, a)
	if err != nil {
		panic(err)
	}
	bs = buf.Bytes()
	return
}

func registerFuncs(tmpl *template.Template) {
	tmpl.Funcs(map[string]any{
		"SerializerOf":      SerializerOf,
		"KeySerializer":     KeySerializer,
		"ElemSerializer":    ElemSerializer,
		"SerType":           SerType,
		"SerVar":            SerVar,
		"FieldsLen":         FieldsLen,
		"ArrayType":         ArrayType,
		"PtrType":           PtrType,
		"Fields":            Fields,
		"VarName":           VarName,
		"TmpVarName":        TmpVarName,
		"CurrentTypeOf":     CurrentTypeOf,
		"MakeFieldTmplPipe": MakeFieldTmplPipe,
		"add":               MakeAddFunc(),
		"minus":             MakeMinusFunc(),
		"include":           MakeIncludeFunc(tmpl),
		"MakeStringOps":     MakeStringOps,
		"MakeArrayOps":      MakeArrayOps,
		"MakeByteSliceOps":  MakeByteSliceOps,
		"MakeSliceOps":      MakeSliceOps,
		"MakeMapOps":        MakeMapOps,
	})
}

func registerTemplates(tmpl *template.Template) (err error) {
	for name, template := range templates {
		_, err = tmpl.New(name).Parse(template)
		if err != nil {
			return
		}
	}
	return nil
}
