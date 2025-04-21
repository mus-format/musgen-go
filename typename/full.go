package typename

import (
	"regexp"
	"strings"
)

const pkgTypeNamePattern = `^(\w+)\.(.+)$`

var pkgTypeNameRe = regexp.MustCompile(pkgTypeNamePattern)

func MakeFullName(pkg Pkg, name TypeName) FullName {
	if pkg == "" {
		return FullName(name)
	}
	return FullName(strings.Join([]string{string(pkg), string(name)}, "."))
}

// FullName examples: "*pkg.TypeName", "[]pkg.TypeName",
// "pkg.TypeName[pkg.Typename]".
type FullName string

func (n FullName) Pkg() Pkg {
	match := pkgTypeNameRe.FindStringSubmatch(string(n))
	if len(match) != 3 {
		return ""
	}
	return Pkg(match[1])
}

func (n FullName) TypeName() TypeName {
	match := pkgTypeNameRe.FindStringSubmatch(string(n))
	if len(match) != 3 {
		return TypeName(n)
	}
	return TypeName(match[2])
}
