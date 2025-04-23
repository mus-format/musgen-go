package typename

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestStrToPkg(t *testing.T) {
	testCases := []struct {
		str     string
		wantPkg Package
		wantErr error
	}{
		{
			str:     "pkg",
			wantPkg: Package("pkg"),
		},
		{
			str:     "+++",
			wantErr: NewInvalidPackageError("+++"),
		},
	}
	for _, c := range testCases {
		pkg, err := StrToPackage(c.str)
		asserterror.EqualError(err, c.wantErr, t)
		asserterror.Equal(pkg, c.wantPkg, t)
	}
}
