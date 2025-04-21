package typeops

import (
	"testing"

	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestOptions(t *testing.T) {
	var (
		o                = Options{}
		wantIgnore       = true
		wantNumEncoding  = Raw
		wantSourceType   = Time
		wantTimeUint     = Milli
		wantValidator    = "ValidateType"
		wantLenEncoding  = Varint
		wantLenValidator = "ValidateLength"
		wantKey          = &Options{Ignore: true}
		wantElem         = &Options{Ignore: true}
	)
	Apply([]SetOption{
		WithIgnore(),
		WithNumEncoding(Raw),
		WithSourceType(Time),
		WithTimeUnit(Milli),
		WithValidator("ValidateType"),
		WithLenEncoding(Varint),
		WithLenValidator("ValidateLength"),
		WithKey(WithIgnore()),
		WithElem(WithIgnore()),
	}, &o)
	asserterror.EqualDeep(o.Ignore, wantIgnore, t)
	asserterror.Equal(o.NumEncoding, wantNumEncoding, t)
	asserterror.Equal(o.SourceType, wantSourceType, t)
	asserterror.Equal(o.TimeUnit, wantTimeUint, t)
	asserterror.Equal(o.Validator, wantValidator, t)
	asserterror.Equal(o.LenEncoding, wantLenEncoding, t)
	asserterror.Equal(o.LenValidator, wantLenValidator, t)
	asserterror.Equal(*o.Key, *wantKey, t)
	asserterror.Equal(*o.Elem, *wantElem, t)

	t.Run("Hash", func(t *testing.T) {
		o1 := Options{}
		Apply([]SetOption{
			WithIgnore(),
			WithNumEncoding(Raw),
			WithSourceType(Time),
			WithTimeUnit(Milli),
			WithValidator("ValidateType"),
			WithLenEncoding(Varint),
			WithLenValidator("ValidateLength"),
			WithKey(WithIgnore()),
			WithElem(WithIgnore()),
		}, &o1)
		o2 := Options{}
		Apply([]SetOption{
			WithIgnore(),
			WithNumEncoding(Raw),
			WithSourceType(Time),
			WithTimeUnit(Milli),
			WithValidator("ValidateType"),
			WithLenEncoding(Varint),
			WithLenValidator("ValidateLength"),
			WithKey(WithIgnore()),
			WithElem(WithIgnore()),
		}, &o2)
		asserterror.Equal(o1.Hash(), o2.Hash(), t)
	})
}
