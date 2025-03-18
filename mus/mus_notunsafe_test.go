package musgen

import (
	"os"
	"reflect"
	"testing"

	genops "github.com/mus-format/musgen-go/options/generate"
	"github.com/mus-format/musgen-go/testdata/pkg1"
)

func TestNotUnsafeFileGenerator(t *testing.T) {

	t.Run("Test FileGenerator.Generate pkg1", func(t *testing.T) {
		g := NewFileGenerator(genops.WithPackage("pkg1"),
			genops.WithImports([]string{
				"github.com/mus-format/musgen-go/testdata/pkg2",
			}),
			genops.WithNotUnsafe(),
		)

		err := g.AddStruct(reflect.TypeFor[pkg1.StructNotUnsafe]())
		if err != nil {
			t.Fatal(err)
		}

		bs, err := g.Generate()
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile("../testdata/pkg1/mus-format_notunsafe.gen.go", bs, 0755)
		if err != nil {
			t.Fatal(err)
		}

	})

}
