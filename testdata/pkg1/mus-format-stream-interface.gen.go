// Code generated by musgen-go. DO NOT EDIT.

package pkg1

import (
	"fmt"

	dts "github.com/mus-format/mus-stream-dts-go"
	muss "github.com/mus-format/mus-stream-go"
)

func MarshalStreamInterfaceMUS(v Interface, w muss.Writer) (n int, err error) {
	switch tp := v.(type) {
	case InterfaceImpl1:
		return StreamInterfaceImpl1DTS.Marshal(tp, w)
	case InterfaceImpl2:
		return StreamInterfaceImpl2DTS.Marshal(tp, w)
	default:
		panic(fmt.Errorf("unexpected %v type", tp))
	}
}

func UnmarshalStreamInterfaceMUS(r muss.Reader) (v Interface, n int, err error) {
	dtm, n, err := dts.UnmarshalDTM(r)
	if err != nil {
		return
	}
	var n1 int
	switch dtm {
	case InterfaceImpl1DTM:
		v, n1, err = StreamInterfaceImpl1DTS.UnmarshalData(r)
	case InterfaceImpl2DTM:
		v, n1, err = StreamInterfaceImpl2DTS.UnmarshalData(r)
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
	}
	n += n1
	return
}

func SizeStreamInterfaceMUS(v Interface) (size int) {
	switch tp := v.(type) {
	case InterfaceImpl1:
		return StreamInterfaceImpl1DTS.Size(tp)
	case InterfaceImpl2:
		return StreamInterfaceImpl2DTS.Size(tp)
	default:
		panic(fmt.Errorf("unexpected %v type", tp))
	}
}

func SkipStreamInterfaceMUS(r muss.Reader) (n int, err error) {
	dtm, n, err := dts.UnmarshalDTM(r)
	if err != nil {
		return
	}
	var n1 int
	switch dtm {
	case InterfaceImpl1DTM:
		n1, err = StreamInterfaceImpl1DTS.SkipData(r)
	case InterfaceImpl2DTM:
		n1, err = StreamInterfaceImpl2DTS.SkipData(r)
	default:
		err = fmt.Errorf("unexpected %v DTM", dtm)
		return
	}
	n += n1
	return
}
