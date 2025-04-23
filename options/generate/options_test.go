package genops

import (
	"reflect"
	"testing"

	prim_testdata "github.com/mus-format/musgen-go/testdata/primitive"
	"github.com/mus-format/musgen-go/typename"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestOptions(t *testing.T) {

	t.Run("Apply should work", func(t *testing.T) {
		var (
			o             = New()
			wantPkgPath   = "github.com/user/project"
			wantUnsafe    = true
			wantNotUnsafe = true
			wantStream    = true
			wantImports   = []string{"\"import1\""}
			wantSerNames  = map[reflect.Type]string{
				reflect.TypeFor[prim_testdata.MyInt](): "MyAwesomeInt",
			}
		)
		Apply([]SetOption{
			WithPkgPath(wantPkgPath),
			WithUnsafe(),
			WithNotUnsafe(),
			WithStream(),
			WithImport("import1"),
			WithSerName(reflect.TypeFor[prim_testdata.MyInt](), "MyAwesomeInt"),
		}, &o)
		asserterror.Equal(o.PkgPath, typename.PkgPath(wantPkgPath), t)
		asserterror.Equal(o.Unsafe, wantUnsafe, t)
		asserterror.Equal(o.NotUnsafe, wantNotUnsafe, t)
		asserterror.Equal(o.Stream, wantStream, t)
		asserterror.EqualDeep(o.Imports, wantImports, t)
		asserterror.EqualDeep(o.SerNames, wantSerNames, t)
	})

	t.Run("Hash", func(t *testing.T) {
		o1 := New()
		Apply([]SetOption{
			WithPkgPath("github.com/user/project"),
			WithUnsafe(),
			WithNotUnsafe(),
			WithStream(),
			WithImport("import1"),
			WithSerName(reflect.TypeFor[prim_testdata.MyInt](), "MyAwesomeInt"),
		}, &o1)
		o2 := New()
		Apply([]SetOption{
			WithPkgPath("github.com/user/project"),
			WithUnsafe(),
			WithNotUnsafe(),
			WithStream(),
			WithImport("import1"),
			WithSerName(reflect.TypeFor[prim_testdata.MyInt](), "MyAwesomeInt"),
		}, &o2)
		asserterror.Equal(o1.Hash(), o2.Hash(), t)
	})

	t.Run("WithImport", func(t *testing.T) {

		t.Run("Should fail if it receives an invalid ImportPath", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImport(""),
			}, &o)
			asserterror.EqualError(err, NewInvalidImportPathError(""), t)
		})

	})

	t.Run("WithImportAlias", func(t *testing.T) {

		t.Run("Should fail if receives two items with the same alias", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImportAlias("github.com/user/project1", "alias"),
				WithImportAlias("github.com/user/project2", "alias"),
			}, &o)
			asserterror.EqualError(err, NewDuplicateImportAlias("alias"), t)
		})

		t.Run("Should fail if receives two similar pkgPath", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImportAlias("github.com/user/project", "alias1"),
				WithImportAlias("github.com/user/project", "alias2"),
			}, &o)
			asserterror.EqualError(err,
				NewDuplicateImportPath("github.com/user/project"), t)
		})

		t.Run("Should fail if it receives an invalid ImportPath", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImportAlias("", "alias"),
			}, &o)
			asserterror.EqualError(err, NewInvalidImportPathError(""), t)
		})

		t.Run("Should fail if it receives an invalid alias", func(t *testing.T) {
			o := New()
			err := Apply([]SetOption{
				WithImportAlias("github.com/user/project", "123"),
			}, &o)
			asserterror.EqualError(err, NewInvalidAliasError("123"), t)
		})

	})

	t.Run("MarshalSignatureLastParam", func(t *testing.T) {
		var (
			o                             = New()
			wantMarshalSignatureLastParam = msigLastParam
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(o.MarshalSignatureLastParam(),
			wantMarshalSignatureLastParam, t)
	})

	t.Run("MarshalSignatureLastParam with Stream option", func(t *testing.T) {
		var (
			o                             = New()
			wantMarshalSignatureLastParam = msigLastParamStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(o.MarshalSignatureLastParam(),
			wantMarshalSignatureLastParam, t)
	})

	t.Run("MarshalLastParam", func(t *testing.T) {
		var (
			o                         = New()
			wantMarshalLastParam      = mLastParam
			wantMarshalLastParamFirst = mLastParamFirst
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(o.MarshalLastParam(false), wantMarshalLastParam, t)
		asserterror.Equal(o.MarshalLastParam(true), wantMarshalLastParamFirst, t)
	})

	t.Run("MarshalLastParam with Stream option", func(t *testing.T) {
		var (
			o                    = New()
			wantMarshalLastParam = mLastParamStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(o.MarshalLastParam(false), wantMarshalLastParam, t)
		asserterror.Equal(o.MarshalLastParam(true), wantMarshalLastParam, t)
	})

	t.Run("UnmarshalLastParam", func(t *testing.T) {
		var (
			o                           = New()
			wantUnmarshalLastParam      = uLastParam
			wantUnmarshalLastParamFirst = uLastParamFirst
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(o.UnmarshalLastParam(false), wantUnmarshalLastParam, t)
		asserterror.Equal(o.UnmarshalLastParam(true), wantUnmarshalLastParamFirst, t)
	})

	t.Run("UnmarshalLastParam with Stream option", func(t *testing.T) {
		var (
			o                      = New()
			wantUnmarshalLastParam = uLastParamStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(o.UnmarshalLastParam(false), wantUnmarshalLastParam, t)
		asserterror.Equal(o.UnmarshalLastParam(true), wantUnmarshalLastParam, t)
	})

	t.Run("SkipLastParam", func(t *testing.T) {
		var (
			o                 = New()
			wantSkipLastParam = skLastParam
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(o.SkipLastParam(), wantSkipLastParam, t)
	})

	t.Run("SkipLastParam with Stream option", func(t *testing.T) {
		var (
			o                 = New()
			wantSkipLastParam = skLastParamStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(o.SkipLastParam(), wantSkipLastParam, t)
	})

	t.Run("ModImportName", func(t *testing.T) {
		var (
			o                 = New()
			wantModImportName = modImportName
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(o.ModImportName(), wantModImportName, t)
	})

	t.Run("ModImportName with Stream option", func(t *testing.T) {
		var (
			o                 = New()
			wantModImportName = modImportNameStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(o.ModImportName(), wantModImportName, t)
	})

	t.Run("ExtPackageName", func(t *testing.T) {
		var (
			o                  = New()
			wantExtPackageName = extPackageName
		)
		Apply([]SetOption{}, &o)
		asserterror.Equal(o.ExtPackageName(), wantExtPackageName, t)
	})

	t.Run("ExtPackageName with Stream option", func(t *testing.T) {
		var (
			o                  = New()
			wantExtPackageName = extPackageNameStream
		)
		Apply([]SetOption{WithStream()}, &o)
		asserterror.Equal(o.ExtPackageName(), wantExtPackageName, t)
	})

	t.Run("Package", func(t *testing.T) {
		var (
			o           = New()
			wantPackage = "exts"
		)
		Apply([]SetOption{WithPackage(wantPackage)}, &o)
		asserterror.Equal(o.Package, typename.Package(wantPackage), t)
	})

}
