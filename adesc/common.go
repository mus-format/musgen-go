package adesc

import (
	"crypto/md5"
	"encoding/base64"
	"strings"

	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
)

const (
	PlusSignReplacement  = "Σ"
	SlashSignReplacement = "Δ"
	EqualSignReplacement = "Ξ"
)

func subMake(t string, tops *typeops.Options, gops genops.Options,
	m map[AnonymousName]AnonymousDesc) {
	dd, ok := Collect(t, tops, gops, m)
	if ok {
		if _, pst := m[dd.Name]; !pst {
			m[dd.Name] = dd
		}
	}
	return
}

func anonymousName(kind, t string, gops genops.Options, tops *typeops.Options) AnonymousName {
	bs := []byte(t)
	gh := gops.Hash()
	bs = append(bs, gh[:]...)
	if tops != nil {
		th := tops.Hash()
		bs = append(bs, th[:]...)
	}
	h := md5.Sum(bs)
	str := base64.StdEncoding.EncodeToString(h[:])
	str = strings.ReplaceAll(str, "+", PlusSignReplacement)
	str = strings.ReplaceAll(str, "/", SlashSignReplacement)
	str = strings.ReplaceAll(str, "=", EqualSignReplacement)
	return AnonymousName(kind + str)
}

func validatorStr(tp, vl string) string {
	return "com.ValidatorFn[" + tp + "](" + vl + ")"
}
