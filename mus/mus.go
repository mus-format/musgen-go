//go:generate go run gen/main.go
package musgen

import (
	"bytes"
	"go/parser"
	"go/token"
	"reflect"
	"text/template"

	"github.com/mus-format/musgen-go/data"
	"github.com/mus-format/musgen-go/databuild"
	genops "github.com/mus-format/musgen-go/options/generate"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
	"golang.org/x/tools/imports"
)

// NewFileGenerator creates a new FileGenerator.
func NewFileGenerator(ops ...genops.SetOption) (g *FileGenerator, err error) {
	gops := genops.New()
	if err = genops.Apply(ops, &gops); err != nil {
		return
	}
	var (
		crossgenTypes = map[typename.FullName]struct{}{}
		typeBuilder   = databuild.NewTypeDataBuilder(databuild.NewConverter(gops),
			gops)
		anonBuilder = databuild.NewAnonDataBuilder(typeBuilder, gops)
	)
	baseTmpl := template.New("base")
	registerFuncs(typeBuilder, anonBuilder, crossgenTypes, baseTmpl, gops)
	err = registerTemplates(baseTmpl)
	if err != nil {
		panic(err)
	}
	g = &FileGenerator{
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

// FileGenerator is responsible for generating MUS serialization code.
type FileGenerator struct {
	baseTmpl      *template.Template
	typeBuilder   TypeDataBuilder
	anonBuilder   AnonSerDataBuilder
	crossgenTypes map[typename.FullName]struct{}
	anonMap       map[data.AnonSerName]data.AnonData
	genSl         []fileData
	bs            []byte
	gops          genops.Options
}

// AddDefinedType adds the specified type to the FileGenerator to produce a
// serializer for it. This method supports types defined with the following
// source types: number, string, array, slice, map, pointer.
func (g *FileGenerator) AddDefinedType(t reflect.Type, ops ...typeops.SetOption) (
	err error) {
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

// AddStruct adds the specified type to the FileGenerator to produce a
// serializer for it. This method supports types definined with the struct
// source type.
func (g *FileGenerator) AddStruct(t reflect.Type, ops ...structops.SetOption) (
	err error) {
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

// AddDTS adds the specified type to the FileGenerator to produce a DTS
// definition for it. This method supports all types acceptable by the
// AddDefinedType, AddStruct, and AddInterface methods.
func (g *FileGenerator) AddDTS(t reflect.Type, ops ...typeops.SetOption) (
	err error) {
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

// AddInterface adds the specified type to the FileGenerator to produce a
// serializer for it. This method supports types definined with the interface
// source type.
func (g *FileGenerator) AddInterface(t reflect.Type, ops ...introps.SetOption) (
	err error) {
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

// Generate produces the serialization code. The output is intended to be saved
// to a file.
func (g *FileGenerator) Generate() (bs []byte, err error) {
	tmp := g.generatePackage()
	tmp = append(tmp, g.generateImports()...)
	tmp = append(tmp, g.generateAnonymosDefinitions()...)
	tmp = append(tmp, g.generateSerializers()...)
	bs, err = imports.Process("", tmp, nil)
	if err != nil {
		err = NewFileGeneratorError(err)
		return tmp, err
	}
	err = g.checkSyntax(bs)
	if err != nil {
		err = NewFileGeneratorError(err)
		return tmp, err
	}
	return
}

func (g *FileGenerator) generatePackage() (bs []byte) {
	return g.generatePart(packageTmpl, g.gops)
}

func (g *FileGenerator) generateImports() (bs []byte) {
	return g.generatePart(importsTmpl, g.gops)
}

func (g *FileGenerator) generateAnonymosDefinitions() (bs []byte) {
	return g.generatePart(anonDefinitionsTmpl, struct {
		Map map[data.AnonSerName]data.AnonData
		Ops genops.Options
	}{
		Map: g.anonMap,
		Ops: g.gops,
	})
}

func (g *FileGenerator) generateSerializers() (bs []byte) {
	for _, item := range g.genSl {
		bs = append(bs, g.generatePart(item.fileName, item.typeData)...)
	}
	return
}

func (g *FileGenerator) generatePart(tmplName string, a any) (bs []byte) {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := g.baseTmpl.ExecuteTemplate(buf, tmplName, a)
	if err != nil {
		panic(err)
	}
	bs = buf.Bytes()
	return
}

func (g *FileGenerator) fillCrossgen(t reflect.Type, fullName typename.FullName) {
	if g.crossGeneration(t) {
		g.crossgenTypes[fullName] = struct{}{}
	}
}

func (g *FileGenerator) crossGeneration(t reflect.Type) bool {
	return t.PkgPath() != string(g.gops.PkgPath)
}

func (g *FileGenerator) checkSyntax(bs []byte) (err error) {
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
