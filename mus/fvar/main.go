package main

import (
	"os"

	"github.com/ymz-ncnk/fvar"
)

func main() {
	bs, err := fvar.Fvar{}.Generate(fvar.Conf{
		Folder:  "templates",
		Package: "mus",
		VarName: "templates",
	})
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("templates.gen.go", bs, 0755)
	if err != nil {
		panic(err)
	}
}
