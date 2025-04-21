package typename

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestStrToPkg(t *testing.T) {
	testCases := []struct {
		str     string
		wantPkg Pkg
		wantErr error
	}{
		{
			str:     "pkg",
			wantPkg: Pkg("pkg"),
		},
		{
			str:     "+++",
			wantErr: NewInvalidPkgError("+++"),
		},
	}
	for _, c := range testCases {
		pkg, err := StrToPkg(c.str)
		asserterror.EqualError(err, c.wantErr, t)
		asserterror.Equal(pkg, c.wantPkg, t)
	}
}
