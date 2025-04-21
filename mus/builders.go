package musgen

import (
	"reflect"

	"github.com/mus-format/musgen-go/data"
	introps "github.com/mus-format/musgen-go/options/interface"
	structops "github.com/mus-format/musgen-go/options/struct"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/typename"
)

type TypeDataBuilder interface {
	BuildDefinedTypeData(t reflect.Type, tops *typeops.Options) (d data.TypeData,
		err error)
	BuildStructData(t reflect.Type, sops structops.Options) (d data.TypeData,
		err error)
	BuildInterfaceData(t reflect.Type, iops introps.Options) (d data.TypeData,
		err error)
	BuildDTSData(t reflect.Type) (d data.TypeData, err error)
	BuildTimeData(t reflect.Type, tops *typeops.Options) (d data.TypeData,
		err error)
	ToFullName(t reflect.Type) typename.FullName
	ToRelName(fname typename.FullName) typename.RelName
}

type AnonSerDataBuilder interface {
	Build(name typename.FullName, tops *typeops.Options) (d data.AnonData,
		ok bool, err error)
	Collect(name typename.FullName, m map[data.AnonSerName]data.AnonData,
		tops *typeops.Options) (err error)
}
