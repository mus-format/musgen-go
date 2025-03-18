package typeops

import "fmt"

const (
	UndefinedTimeUnit TimeUnit = iota
	Sec
	Milli
	Micro
	Nano
	SecUTC
	MilliUTC
	MicroUTC
	NanoUTC
)

type TimeUnit int

func (u TimeUnit) Ser() string {
	switch u {
	case UndefinedTimeUnit, Sec:
		return "TimeUnix"
	case Milli:
		return "TimeUnixMilli"
	case Micro:
		return "TimeUnixMicro"
	case Nano:
		return "TimeUnixNano"
	case SecUTC:
		return "TimeUnixUTC"
	case MilliUTC:
		return "TimeUnixMilliUTC"
	case MicroUTC:
		return "TimeUnixMicroUTC"
	case NanoUTC:
		return "TimeUnixNanoUTC"
	default:
		panic(fmt.Sprintf("unexpected TimeUnit %v", u))
	}
}
