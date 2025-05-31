package builders

import (
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	"github.com/mus-format/musgen-go/scanner"
	generic_testdata "github.com/mus-format/musgen-go/testdata/generic"
	prim_testdata "github.com/mus-format/musgen-go/testdata/primitive"
	"github.com/mus-format/musgen-go/typename"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestConverter(t *testing.T) {

	t.Run("ConvertToFullName", func(t *testing.T) {

		t.Run("Should work", func(t *testing.T) {
			gops := genops.New()
			genops.Apply([]genops.SetOption{
				genops.WithImportAlias("github.com/mus-format/musgen-go/testdata/primitive", "primitive"),
				genops.WithImportAlias("github.com/mus-format/musgen-go/testdata/generic", "generic"),
			}, &gops)

			testCases := []struct {
				cname        typename.CompleteName
				wantFullName typename.FullName
				wantErr      error
			}{
				{
					cname:        typename.MustTypeCompleteName(reflect.TypeFor[prim_testdata.MyInt]()),
					wantFullName: "primitive.MyInt",
				},
				{
					cname:        typename.MustTypeCompleteName(reflect.TypeFor[generic_testdata.MySlice[prim_testdata.MyInt]]()),
					wantFullName: "generic.MySlice[primitive.MyInt]",
				},
				{
					cname:        typename.MustTypeCompleteName(reflect.TypeFor[map[generic_testdata.MyArray[int]]map[prim_testdata.MyInt]string]()),
					wantFullName: "map[generic.MyArray[int]]map[primitive.MyInt]string",
				},
			}

			conv := NewConverter(gops)
			var (
				fname typename.FullName
				err   error
			)
			for _, c := range testCases {
				fname, err = conv.ConvertToFullName(c.cname)
				asserterror.EqualError(err, c.wantErr, t)
				asserterror.Equal(fname, c.wantFullName, t)
			}
		})

		t.Run("If scanner fails with an error ConvertToFullName should return it", func(t *testing.T) {
			var (
				cname typename.CompleteName = "++++++"
				conv                        = NewConverter(genops.New())

				wantFullName typename.FullName = ""
				wantErr      error             = scanner.NewUnsupportedQualifiedNameError(cname)
			)
			fname, err := conv.ConvertToFullName(cname)
			asserterror.Equal(fname, wantFullName, t)
			asserterror.EqualError(err, wantErr, t)
		})

		t.Run("Should fails if cname contains two pkgPath with the same alias",
			func(t *testing.T) {
				var (
					cname typename.CompleteName = "map[github.com/mus-format/musgen-go/testdata/primitive/testdata.MyInt]github.com/mus-format/musgen-go/testdata/container/testdata.MyArray"
					conv                        = NewConverter(genops.New())

					wantFullName typename.FullName = ""
					wantErr      error             = NewTwoPathsSameAliasError(
						"github.com/mus-format/musgen-go/testdata/primitive",
						"github.com/mus-format/musgen-go/testdata/container",
						"testdata")
				)
				fname, err := conv.ConvertToFullName(cname)
				asserterror.Equal(fname, wantFullName, t)
				asserterror.EqualError(err, wantErr, t)
			})

	})

	t.Run("ConvertToRelName", func(t *testing.T) {

		t.Run("Should work", func(t *testing.T) {
			gops := genops.New()
			genops.Apply([]genops.SetOption{genops.WithPackage("primitive")}, &gops)

			testCases := []struct {
				fname       typename.FullName
				wantRelName typename.RelName
				wantErr     error
			}{
				{
					fname:       "primitive.MyInt",
					wantRelName: "MyInt",
				},
				{
					fname:       "generic.MySlice[primitive.MyInt]",
					wantRelName: "generic.MySlice[MyInt]",
				},
				{
					fname:       "map[generic.MyArray[int]]map[primitive.MyInt]string",
					wantRelName: "map[generic.MyArray[int]]map[MyInt]string",
				},
			}

			conv := NewConverter(gops)
			var (
				rname typename.RelName
				err   error
			)
			for _, c := range testCases {
				rname, err = conv.ConvertToRelName(c.fname)
				asserterror.EqualError(err, c.wantErr, t)
				asserterror.Equal(rname, c.wantRelName, t)
			}
		})

		t.Run("If scanner fails with an error ConvertToRelName should return it", func(t *testing.T) {
			var (
				fname typename.FullName = "++++++"
				conv                    = NewConverter(genops.New())

				wantRelName typename.RelName = ""
				wantErr     error            = scanner.NewUnsupportedQualifiedNameError(fname)
			)
			rname, err := conv.ConvertToRelName(fname)
			asserterror.Equal(rname, wantRelName, t)
			asserterror.EqualError(err, wantErr, t)
		})

	})

}
