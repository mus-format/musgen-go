package genops

import (
	"crypto/md5"
	"go/token"
	"reflect"
	"strconv"

	"github.com/mus-format/musgen-go/typename"
	"golang.org/x/mod/module"
)

type ImportPath string

type Alias string

// New creates a new Options.
func New() Options {
	return Options{
		SerNames:      make(map[reflect.Type]string),
		importAliases: make(map[ImportPath]Alias),
	}
}

// Options specifies configuration options for code generation.
type Options struct {
	PkgPath       typename.PkgPath
	Unsafe        bool
	NotUnsafe     bool
	Stream        bool
	Imports       []string
	SerNames      map[reflect.Type]string
	pkg           typename.Pkg
	importAliases map[ImportPath]Alias
}

func (o Options) MarshalSignatureLastParam() string {
	if o.Stream {
		return msigLastParamStream
	}
	return msigLastParam
}

func (o Options) MarshalLastParam(first bool) string {
	if o.Stream {
		return mLastParamStream
	}
	if first {
		return mLastParamFirst
	}
	return mLastParam
}

func (o Options) UnmarshalSignatureLastParam() string {
	if o.Stream {
		return usigLastParamStream
	}
	return usigLastParam
}

func (o Options) UnmarshalLastParam(first bool) string {
	if o.Stream {
		return uLastParamStream
	}
	if first {
		return uLastParamFirst
	}
	return uLastParam
}

func (o Options) SkipSignatureLastParam() string {
	if o.Stream {
		return sksiqLastParamStream
	}
	return sksiqLastParam
}

func (o Options) SkipLastParam() string {
	if o.Stream {
		return skLastParamStream
	}
	return skLastParam
}

func (o Options) Package() typename.Pkg {
	if o.pkg != "" {
		return o.pkg
	}
	return o.PkgPath.Package()
}

func (o Options) ModImportName() string {
	if o.Stream {
		return modImportNameStream
	}
	return modImportName
}

func (o Options) ExtPackageName() string {
	if o.Stream {
		return extPackageNameStream
	}
	return extPackageName
}

func (o Options) ImportAliases() map[ImportPath]Alias {
	return o.importAliases
}

func (o Options) Hash() [16]byte {
	var bs = []byte{}
	bs = append(bs, []byte(o.PkgPath)...)
	bs = append(bs, []byte(strconv.FormatBool(o.Stream))...)
	bs = append(bs, []byte(strconv.FormatBool(o.Unsafe))...)
	for i := range o.Imports {
		bs = append(bs, []byte(o.Imports[i])...)
	}
	return md5.Sum(bs)
}

// -----------------------------------------------------------------------------

type SetOption func(o *Options) error

// WithPkgPath returns a SetOption that configures the package path for the
// generated file. The path must match the standard Go package path format
// (e.g., github.com/user/project/pkg) and could be obtained like:
//
//	pkgPath := reflect.TypeFor[YourType]().PkgPath()
func WithPkgPath(str string) SetOption {
	return func(o *Options) (err error) {
		o.PkgPath, err = typename.StrToPkgPath(str)
		return
	}
}

// WithPackage returns a SetOption that sets the package name for the generated
// file.
func WithPackage(str string) SetOption {
	return func(o *Options) (err error) {
		o.pkg, err = typename.StrToPkg(str)
		return
	}
}

// WithUnsafe returns a SetOption that enables unsafe code generation.
func WithUnsafe() SetOption {
	return func(o *Options) (err error) {
		o.Unsafe = true
		return
	}
}

// WithNotUnsafe returns a SetOption that enables unsafe code generation without
// side effects. When applied, the generator will avoid unsafe operations for
// the string type.
func WithNotUnsafe() SetOption {
	return func(o *Options) (err error) {
		o.NotUnsafe = true
		return
	}
}

// WithStream returns a SetOption that enables streaming code generation.
// When applied, the generator will produce code that can process data
// incrementally rather than requiring complete in-memory representations.
func WithStream() SetOption {
	return func(o *Options) (err error) {
		o.Stream = true
		return
	}
}

// WithImport returns a SetOption that adds the given import path to the list
// of imports.
func WithImport(importPath string) SetOption {
	return func(o *Options) (err error) {
		if err = module.CheckImportPath(importPath); err != nil {
			err = NewInvalidImportPathError(importPath)
			return
		}
		o.Imports = append(o.Imports, `"`+importPath+`"`)
		return
	}
}

// WithImportAlias returns a SetOption that adds the given import path and alias
// to the list of imports.
func WithImportAlias(importPath, alias string) SetOption {
	return func(o *Options) (err error) {
		if err = module.CheckImportPath(importPath); err != nil {
			err = NewInvalidImportPathError(importPath)
			return
		}
		if !token.IsIdentifier(alias) {
			err = NewInvalidAliasError(alias)
			return
		}
		var (
			p = ImportPath(importPath)
			a = Alias(alias)
		)
		if err := uniqImportAlias(p, a, *o); err != nil {
			return err
		}
		o.Imports = append(o.Imports, alias+" "+`"`+importPath+`"`)
		o.importAliases[p] = a
		return
	}
}

// WithSerName returns a SetOption that registers a custom serializer name for a
// specific type.
func WithSerName(t reflect.Type, serName string) SetOption {
	return func(o *Options) (err error) {
		o.SerNames[t] = serName
		return
	}
}

func Apply(ops []SetOption, o *Options) (err error) {
	for i := range ops {
		if ops[i] != nil {
			err = ops[i](o)
			if err != nil {
				return
			}
		}
	}
	if o.PkgPath == "" {
		return ErrEmptyPkgPath
	}
	// Should check the pkg format here, because WithPackage() is optional.
	if _, err = typename.StrToPkg(string(o.Package())); err != nil {
		return
	}
	if err = addImportAlias(*o); err != nil {
		return
	}
	return
}

func uniqImportAlias(importPath ImportPath, alias Alias, o Options) (
	err error) {
	if _, pst := o.importAliases[importPath]; pst {
		err = NewDuplicateImportPath(importPath)
		return
	}
	for _, existAlias := range o.importAliases {
		if existAlias == alias {
			err = NewDuplicateImportAlias(alias)
			return
		}
	}
	return
}

func addImportAlias(o Options) (err error) {
	var (
		p = ImportPath(o.PkgPath)
		a = Alias(o.Package())
	)
	if err = uniqImportAlias(p, a, o); err != nil {
		return
	}
	o.importAliases[p] = a
	return
}
