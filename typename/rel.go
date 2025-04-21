package typename

import "strings"

// RelName examples: "TypeName", "pkg.TypeName".
type RelName FullName

func (n RelName) WithoutSquares() (str string) {
	str = string(n)
	open := strings.Index(str, "[")
	if open == -1 {
		return str // no type parameters
	}
	return str[:open]
}
