package typename

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestFullName(t *testing.T) {

	t.Run("Pkg", func(t *testing.T) {
		testCases := []struct {
			name    FullName
			wantPkg Pkg
		}{
			{name: "pkg.Type", wantPkg: "pkg"},
			{name: "Type", wantPkg: ""},
			{name: "pkg123.Type", wantPkg: "pkg123"},
			{name: "pkg.Type[pkg.Type]", wantPkg: "pkg"},
			{name: "", wantPkg: ""},
		}
		for _, c := range testCases {
			asserterror.Equal(c.name.Pkg(), c.wantPkg, t)
		}
	})

	t.Run("TypeName", func(t *testing.T) {
		testCases := []struct {
			name         FullName
			wantTypeName TypeName
		}{
			{name: "pkg.Type", wantTypeName: "Type"},
			{name: "Type", wantTypeName: "Type"},
			{name: "pkg123.Type", wantTypeName: "Type"},
			{name: "pkg.Type[pkg.Type]", wantTypeName: "Type[pkg.Type]"},
			{name: "", wantTypeName: ""},
		}
		for _, c := range testCases {
			asserterror.Equal(c.name.TypeName(), c.wantTypeName, t)
		}
	})

}
