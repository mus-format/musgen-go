package parser

import (
	"regexp"
	"strconv"
)

var TypeName = typeName{}

type typeName struct{}

func (n typeName) ParseArray(t string) (elemType string, length int, ok bool) {
	re := regexp.MustCompile(`^\[(\d+)\](.+$)`)
	match := re.FindStringSubmatch(t)
	if len(match) == 3 {
		length, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		return match[2], length, true
	}
	return
}

func (n typeName) ParseSlice(t string) (elemType string, ok bool) {
	re := regexp.MustCompile(`^\[\](.+$)`)
	match := re.FindStringSubmatch(t)
	if len(match) == 2 {
		return match[1], true
	}
	return
}

func (n typeName) ParseMap(t string) (keyType, elemType string, ok bool) {
	re := regexp.MustCompile(`^map\[(.+?)\](.+$)`)
	match := re.FindStringSubmatch(t)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return
}

func (n typeName) ParsePtr(t string) (stars, elemType string, ok bool) {
	re := regexp.MustCompile(`(^\*)(.+$)`)
	match := re.FindStringSubmatch(t)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return
}
