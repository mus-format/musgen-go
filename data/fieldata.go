package data

import "github.com/mus-format/musgen-go/typename"

type FieldData struct {
	FullName  typename.FullName // field type, contains pkg, generics
	FieldName string            // field name
}
