package typename

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestStrToPkgPath(t *testing.T) {
	testCases := []struct {
		str         string
		wantPkgPath PkgPath
		wantErr     error
	}{
		{
			str:         "github.com/user/project",
			wantPkgPath: PkgPath("github.com/user/project"),
		},
		{
			str:     "+++",
			wantErr: NewInvalidPkgPathError("+++"),
		},
	}
	for _, c := range testCases {
		pkgPath, err := StrToPkgPath(c.str)
		asserterror.EqualError(err, c.wantErr, t)
		asserterror.Equal(pkgPath, c.wantPkgPath, t)
	}
}

func TestPkgPath(t *testing.T) {
	testCases := []struct {
		pkgPath PkgPath
		wantPkg Package
	}{
		{
			pkgPath: PkgPath("github.com/user/project"),
			wantPkg: Package("project"),
		},
		{
			pkgPath: "+++",
			wantPkg: Package("+++"),
		},
	}
	for _, c := range testCases {
		pkg := Package(c.pkgPath.Base())
		asserterror.Equal(pkg, c.wantPkg, t)
	}
}
