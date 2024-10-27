package mus

import (
	"os"
	"reflect"
	"testing"

	"github.com/mus-format/musgen-go/basegen"
	"github.com/mus-format/musgen-go/testdata/pkg1"
)

const UnsafeStreamPrefix = "UnsafeStream"

func TestUnsafeStreamFileGenerator(t *testing.T) {
	g, err := NewFileGenerator(basegen.Conf{Package: "pkg1", Unsafe: true, Stream: true})
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddAliasWith(reflect.TypeFor[pkg1.BoolAlias](), UnsafeStreamPrefix, nil)
	if err != nil {
		t.Fatal(err)
	}
	bs, err := g.Generate()
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("../testdata/pkg1/mus-format-unsafe-stream-alias.gen.go", bs, 0755)
	if err != nil {
		t.Fatal(err)
	}
}
