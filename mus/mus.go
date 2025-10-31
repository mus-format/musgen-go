// Package musgen provides a code generator for the MUS binary serialization
// format.
//
//go:generate go run gen/main.go
package musgen

import (
	"bytes"
	"go/parser"
	"go/token"
	"reflect"
	"text/template"

	"github.com/mus-format/musgen-go/data"
	"github.com/mus-format/musgen-go/data/builders"
	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
	"golang.org/x/tools/imports"
)

// NewCodeGenerator creates a new CodeGenerator.
func NewCodeGenerator(ops ...genops.SetOption) (g *CodeGenerator, err error) {
	gops := genops.New()
	if err = genops.Apply(ops, &gops); err != nil {
		return
	}
	var (
		crossgenTypes = map[typename.FullName]struct{}{}
		typeBuilder   = builders.NewTypeDataBuilder(builders.NewConverter(gops),
			gops)
		anonBuilder = builders.NewAnonDataBuilder(typeBuilder, gops)
	)
	baseTmpl := template.New("base")
	registerFuncs(typeBuilder, anonBuilder, crossgenTypes, baseTmpl, gops)
	err = registerTemplates(baseTmpl)
	if err != nil {
		panic(err)
	}
	g = &CodeGenerator{
		baseTmpl:      baseTmpl,
		typeBuilder:   typeBuilder,
		anonBuilder:   anonBuilder,
		crossgenTypes: crossgenTypes,
		anonMap:       map[data.AnonSerName]data.AnonData{},
		genSl:         []fileData{},
		bs:            []byte{},
		gops:          gops,
	}
	return
}

// CodeGenerator is responsible for generating MUS serialization code.
type CodeGenerator struct {
	baseTmpl      *template.Template
	typeBuilder   TypeDataBuilder
	anonBuilder   AnonSerDataBuilder
	crossgenTypes map[typename.FullName]struct{}
	anonMap       map[data.AnonSerName]data.AnonData
	genSl         []fileData
	dtmTypes      [][]string
	bs            []byte
	gops          genops.Options
}

// AddDefinedType adds the specified type to the CodeGenerator to produce a
// serializer for it. This method supports types defined with the following
// source types: number, string, array, slice, map, pointer.
func (g *CodeGenerator) AddDefinedType(t reflect.Type, ops ...typeops.SetOption) (
	err error,
) {
	var tops *typeops.Options
	if len(ops) > 0 {
		tops = &typeops.Options{}
		typeops.Apply(ops, tops)
	}
	typeData, err := g.typeBuilder.BuildDefinedTypeData(t, tops)
	if err != nil {
		return
	}
	g.fillCrossgen(t, typeData.FullName)
	g.anonBuilder.Collect(typeData.Fields[0].FullName, g.anonMap, tops)
	g.genSl = append(g.genSl, fileData{definedTypeSerTmpl, typeData})
	return
}

// AddStruct adds the specified type to the CodeGenerator to produce a
// serializer for it. This method supports types definined with the struct
// source type.
func (g *CodeGenerator) AddStruct(t reflect.Type, ops ...structops.SetOption) (
	err error,
) {
	sops := structops.New()
	if len(ops) > 0 {
		structops.Apply(ops, &sops)
	}
	var typeData data.TypeData
	if sops.Tops != nil && sops.Tops.SourceType == typeops.Time {
		typeData, err = g.typeBuilder.BuildTimeData(t, sops.Tops)
		if err != nil {
			return
		}
		g.fillCrossgen(t, typeData.FullName)
		g.genSl = append(g.genSl, fileData{definedTypeSerTmpl, typeData})
		return
	}
	typeData, err = g.typeBuilder.BuildStructData(t, sops)
	if err != nil {
		return
	}
	for i, fd := range typeData.SerializedFields() {
		var tops *typeops.Options
		if len(sops.Fields) > 0 {
			tops = sops.Fields[i]
		}
		g.anonBuilder.Collect(fd.FullName, g.anonMap, tops)
	}
	g.fillCrossgen(t, typeData.FullName)
	g.genSl = append(g.genSl, fileData{structSerTmpl, typeData})
	return
}

// AddDTS adds the specified type to the CodeGenerator to produce a DTS
// definition for it. This method supports all types acceptable by the
// AddDefinedType, AddStruct, and AddInterface methods.
func (g *CodeGenerator) AddDTS(t reflect.Type, ops ...typeops.SetOption) (
	err error,
) {
	tops := typeops.Options{}
	if len(ops) > 0 {
		typeops.Apply(ops, &tops)
	}
	typeData, err := g.typeBuilder.BuildDTSData(t)
	if err != nil {
		return
	}
	g.genSl = append(g.genSl, fileData{dtsTmpl, typeData})
	return
}

// AddInterface adds the specified type to the CodeGenerator to produce a
// serializer for it. This method supports types definined with the interface
// source type.
func (g *CodeGenerator) AddInterface(t reflect.Type, ops ...introps.SetOption) (
	err error,
) {
	iops := introps.Options{}
	if ops != nil {
		introps.Apply(ops, &iops)
	}
	typeData, err := g.typeBuilder.BuildInterfaceData(t, iops)
	if err != nil {
		return
	}
	g.fillCrossgen(t, typeData.FullName)
	g.genSl = append(g.genSl, fileData{interfaceSerTmpl, typeData})
	return
}

// RegisterInterface registers an interface type and all of its implementations
// with the code generator.
//
// DTM values are generated automatically, so there is no need to assign them
// manually.
//
// This helper method is equivalent to calling, in order:
//
//	AddStruct/AddDefinedType → AddDTS → AddInterface
func (g *CodeGenerator) RegisterInterface(t reflect.Type,
	ops ...introps.SetRegisterOption,
) (err error) {
	rops := introps.RegisterOptions{}
	if ops != nil {
		introps.ApplyRegister(ops, &rops)
	}
	var (
		l     = len(rops.StructImpls) + len(rops.DefinedTypeImpls)
		iops  = make([]introps.SetOption, 0, l)
		types = make([]string, 0, l)
	)
	for _, impl := range rops.StructImpls {
		err = g.AddStruct(impl.Type, impl.Ops...)
		if err != nil {
			return
		}
		err = g.AddDTS(impl.Type)
		if err != nil {
			return
		}
		iops = append(iops, introps.WithImpl(impl.Type))
		types = append(types, impl.Type.Name())
	}
	for _, impl := range rops.DefinedTypeImpls {
		err = g.AddDefinedType(impl.Type, impl.Ops...)
		if err != nil {
			return
		}
		err = g.AddDTS(impl.Type)
		if err != nil {
			return
		}
		iops = append(iops, introps.WithImpl(impl.Type))
		types = append(types, impl.Type.Name())
	}
	g.addDTM(types)
	return g.AddInterface(t, iops...)
}

func (g *CodeGenerator) addDTM(types []string) {
	g.dtmTypes = append(g.dtmTypes, types)
}

// Generate produces the serialization code. The output is intended to be saved
// to a file.
func (g *CodeGenerator) Generate() (bs []byte, err error) {
	tmp := g.generatePackage()
	tmp = append(tmp, g.generateImports()...)
	tmp = append(tmp, g.generateDTMs()...)
	tmp = append(tmp, g.generateAnonymosDefinitions()...)
	tmp = append(tmp, g.generateSerializers()...)
	bs, err = imports.Process("", tmp, nil)
	if err != nil {
		err = ErrCodeGenFailed
		return tmp, err
	}
	err = g.checkSyntax(bs)
	if err != nil {
		err = ErrCodeGenFailed
		return tmp, err
	}
	return
}

func (g *CodeGenerator) generatePackage() (bs []byte) {
	return g.generatePart(packageTmpl, g.gops)
}

func (g *CodeGenerator) generateImports() (bs []byte) {
	return g.generatePart(importsTmpl, g.gops)
}

func (g *CodeGenerator) generateAnonymosDefinitions() (bs []byte) {
	return g.generatePart(anonDefinitionsTmpl, struct {
		Map map[data.AnonSerName]data.AnonData
		Ops genops.Options
	}{
		Map: g.anonMap,
		Ops: g.gops,
	})
}

func (g *CodeGenerator) generateSerializers() (bs []byte) {
	for _, item := range g.genSl {
		bs = append(bs, g.generatePart(item.fileName, item.typeData)...)
	}
	return
}

func (g *CodeGenerator) generatePart(tmplName string, a any) (bs []byte) {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := g.baseTmpl.ExecuteTemplate(buf, tmplName, a)
	if err != nil {
		panic(err)
	}
	bs = buf.Bytes()
	return
}

func (g *CodeGenerator) generateDTMs() (bs []byte) {
	buf := bytes.Buffer{}
	for _, types := range g.dtmTypes {
		err := g.baseTmpl.ExecuteTemplate(&buf, dtmsDefinitionTmpl, types)
		if err != nil {
			panic(err)
		}
	}
	bs = buf.Bytes()
	return
}

func (g *CodeGenerator) fillCrossgen(t reflect.Type, fullName typename.FullName) {
	if g.crossGeneration(t) {
		g.crossgenTypes[fullName] = struct{}{}
	}
}

func (g *CodeGenerator) crossGeneration(t reflect.Type) bool {
	return t.PkgPath() != string(g.gops.PkgPath)
}

func (g *CodeGenerator) checkSyntax(bs []byte) (err error) {
	var (
		fs  = token.NewFileSet()
		src = string(bs)
	)
	_, err = parser.ParseFile(fs, "", src, parser.AllErrors)
	return
}

type fileData struct {
	fileName string
	typeData data.TypeData
}

func registerFuncs(typeBuilder TypeDataBuilder, anonBuilder AnonSerDataBuilder,
	crossgenTypes map[typename.FullName]struct{},
	tmpl *template.Template,
	gops genops.Options,
) {
	var (
		keywordFns = NewKeywordFns(typeBuilder, crossgenTypes, gops)
		optionFns  = NewOptionFns(typeBuilder)
		utilFns    = NewUtilFns(keywordFns)
		serFns     = NewSerFns(keywordFns, utilFns, anonBuilder, typeBuilder)
	)
	m := map[string]any{}
	keywordFns.RegisterItself(m)
	optionFns.RegisterItself(m)
	utilFns.RegisterItself(m, tmpl)
	serFns.RegisterItself(m)
	tmpl.Funcs(m)
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
