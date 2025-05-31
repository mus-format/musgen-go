package builders

import (
	"crypto/md5"
	"strings"

	"github.com/mus-format/musgen-go/data"
	genops "github.com/mus-format/musgen-go/options/generate"
	typeops "github.com/mus-format/musgen-go/options/type"
	"github.com/mus-format/musgen-go/scanner"
	"github.com/mus-format/musgen-go/typename"
)

func NewCollectAnonOp(m map[data.AnonSerName]data.AnonData,
	typeBuilder TypeDataBuilder, gops genops.Options) *CollectAnonOp {
	return &CollectAnonOp{m: m, typeBuilder: typeBuilder, gops: gops}
}

// CollectAnonOp
type CollectAnonOp struct {
	m             map[data.AnonSerName]data.AnonData
	firstAnonData *data.AnonData
	typeBuilder   TypeDataBuilder
	gops          genops.Options
}

func (o *CollectAnonOp) ProcessType(t scanner.Type[typename.FullName],
	tops *typeops.Options) (err error) {
	var (
		d  data.AnonData
		ok bool
	)
	if t.Stars != "" {
		return o.processPtrType(t, tops)
	}
	switch t.Kind {
	case scanner.Prim:
		d, ok = o.makePrimData(t, tops)
	case scanner.Array:
		d, ok = o.makeArrayData(t, tops)
	case scanner.Slice:
		if t.ElemType == "byte" || t.ElemType == "uint8" {
			d, ok = o.makeByteSliceData(t, tops)
		} else {
			d, ok = o.makeSliceData(t, tops)
		}
	case scanner.Map:
		d, ok = o.makeMapDesc(t, tops)
	}
	if ok {
		o.putToMap(d)
	}
	return
}

func (o *CollectAnonOp) ProcessLeftSquare() {}

func (o *CollectAnonOp) ProcessComma() {}

func (o *CollectAnonOp) ProcessRightSquare() {}

func (o *CollectAnonOp) FirstAnonData() *data.AnonData {
	return o.firstAnonData
}

func (o *CollectAnonOp) processPtrType(t scanner.Type[typename.FullName],
	tops *typeops.Options) (err error) {
	d, _ := o.makePtrData(t, tops)
	o.putToMap(d)

	t.Stars = trimOneStar(t.Stars)
	var elemTops *typeops.Options
	if tops != nil {
		elemTops = tops.Elem
	}

	return o.ProcessType(t, elemTops)
}

func (o *CollectAnonOp) makePtrData(t scanner.Type[typename.FullName],
	tops *typeops.Options) (d data.AnonData, ok bool) {
	str := t.Stars + string(typename.MakeFullName(t.Package, t.Name))
	return data.AnonData{
		AnonSerName: anonSerName(data.AnonPtr, typename.TypeName(str), tops, o.gops),
		ElemType:    typename.FullName(trimOneStar(str)),
		Kind:        data.AnonPtr,
		Tops:        tops,
	}, true
}

func (o *CollectAnonOp) makePrimData(t scanner.Type[typename.FullName],
	tops *typeops.Options) (d data.AnonData, ok bool) {
	if string(t.Name) == "string" && tops != nil {
		lenSer := "nil"
		if tops.LenEncoding != typeops.UndefinedNumEncoding &&
			tops.LenEncoding != typeops.VarintPositive {
			lenSer = tops.LenSer()
			ok = true
		}
		lenVl := "nil"
		if tops.LenValidator != "" {
			lenVl = o.validatorStr("int", tops.LenValidator)
			ok = true
		}
		if ok {
			return data.AnonData{
				AnonSerName: anonSerName(data.AnonString, t.Name, tops, o.gops),
				Kind:        data.AnonString,
				LenSer:      lenSer,
				LenVl:       lenVl,
				Tops:        tops,
			}, ok
		}
	}
	return
}

func (o *CollectAnonOp) makeArrayData(t scanner.Type[typename.FullName],
	tops *typeops.Options) (d data.AnonData, ok bool) {
	var (
		lenSer = "nil"
		elemVl = "nil"
	)
	if tops != nil {
		if tops.LenEncoding != typeops.UndefinedNumEncoding {
			lenSer = tops.LenSer()
		}
		if tops.Elem != nil {
			if tops.Elem.Validator != "" {
				elemVl = o.validatorStr(t.ElemType, tops.Elem.Validator)
			}
		}
	}
	return data.AnonData{
		AnonSerName: anonSerName(data.AnonArray, t.Name, tops, o.gops),
		Kind:        data.AnonArray,
		ArrType:     typename.MakeFullName(t.Package, t.Name),
		ArrLength:   t.ArrLength,
		LenSer:      lenSer,
		ElemType:    t.ElemType,
		ElemVl:      elemVl,
		Tops:        tops,
	}, true
}

func (o *CollectAnonOp) makeByteSliceData(t scanner.Type[typename.FullName],
	tops *typeops.Options) (d data.AnonData, ok bool) {
	if tops != nil &&
		(tops.LenEncoding != typeops.UndefinedNumEncoding || tops.LenValidator != "") {
		var (
			lenSer = "nil"
			lenVl  = "nil"
		)
		if tops.LenEncoding != typeops.UndefinedNumEncoding {
			lenSer = tops.LenSer()
		}
		if tops.LenValidator != "" {
			lenVl = o.validatorStr("int", tops.LenValidator)
		}
		return data.AnonData{
			AnonSerName: anonSerName(data.AnonByteSlice, t.Name, tops, o.gops),
			Kind:        data.AnonByteSlice,
			LenSer:      lenSer,
			LenVl:       lenVl,
			Tops:        tops,
		}, true
	}
	return
}

func (o *CollectAnonOp) makeSliceData(t scanner.Type[typename.FullName],
	tops *typeops.Options) (d data.AnonData, ok bool) {
	var (
		lenSer = "nil"
		lenVl  = "nil"
		elemVl = "nil"
	)
	if tops != nil {
		if tops.LenEncoding != typeops.UndefinedNumEncoding {
			lenSer = tops.LenSer()
		}
		if tops.LenValidator != "" {
			lenVl = o.validatorStr("int", tops.LenValidator)
		}
		if tops.Elem != nil {
			if tops.Elem.Validator != "" {
				elemVl = o.validatorStr(t.ElemType, tops.Elem.Validator)
			}
		}
	}
	return data.AnonData{
		AnonSerName: anonSerName(data.AnonSlice, t.Name, tops, o.gops),
		Kind:        data.AnonSlice,
		LenSer:      lenSer,
		LenVl:       lenVl,
		ElemType:    t.ElemType,
		ElemVl:      elemVl,
		Tops:        tops,
	}, true
}

func (o *CollectAnonOp) makeMapDesc(t scanner.Type[typename.FullName],
	tops *typeops.Options) (d data.AnonData, ok bool) {
	var (
		lenSer = "nil"
		lenVl  = "nil"
		keyVl  = "nil"
		elemVl = "nil"
	)
	if tops != nil {
		if tops.LenEncoding != typeops.UndefinedNumEncoding {
			lenSer = tops.LenSer()
		}
		if tops.LenValidator != "" {
			lenVl = o.validatorStr("int", tops.LenValidator)
		}
		if tops.Key != nil {
			if tops.Key.Validator != "" {
				keyVl = o.validatorStr(t.KeyType, tops.Key.Validator)
			}
		}
		if tops.Elem != nil {
			if tops.Elem.Validator != "" {
				elemVl = o.validatorStr(t.ElemType, tops.Elem.Validator)
			}
		}
	}
	return data.AnonData{
		AnonSerName: anonSerName(data.AnonMap, t.Name, tops, o.gops),
		Kind:        data.AnonMap,
		LenSer:      lenSer,
		LenVl:       lenVl,
		KeyType:     t.KeyType,
		KeyVl:       keyVl,
		ElemType:    t.ElemType,
		ElemVl:      elemVl,
		Tops:        tops,
	}, true
}

func (o *CollectAnonOp) putToMap(d data.AnonData) {
	if o.firstAnonData == nil {
		o.firstAnonData = &d
	}
	o.m[d.AnonSerName] = d
}

func (o *CollectAnonOp) validatorStr(name typename.FullName, vl string) string {
	return "com.ValidatorFn[" + string(o.typeBuilder.ToRelName(name)) + "](" + vl + ")"
}

func anonSerName(kind data.AnonSerKind, typeName typename.TypeName,
	tops *typeops.Options, gops genops.Options) data.AnonSerName {
	bs := []byte(typeName)
	gh := gops.Hash()
	bs = append(bs, gh[:]...)
	if tops != nil {
		th := tops.Hash()
		bs = append(bs, th[:]...)
	}
	h := md5.Sum(bs)
	return data.AnonSerName(kind.String() +
		Base64KeywordEncoding.EncodeToString(h[:]))
}

func trimOneStar(s string) string {
	if strings.HasPrefix(s, "*") {
		return s[1:]
	}
	return s
}
