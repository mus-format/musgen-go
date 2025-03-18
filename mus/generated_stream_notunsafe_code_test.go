package musgen

import (
	"testing"
	"time"

	"github.com/mus-format/musgen-go/testdata/pkg1"
)

func TestStreamNotUnsafeGeneratedCode(t *testing.T) {

	t.Run("StructNotUnsafe should be serializable", func(t *testing.T) {
		var (
			v = pkg1.StructStreamNotUnsafe{
				String: "abc",
				Int:    3,
				Time:   time.Unix(time.Now().Unix(), 0),
			}
		)
		testStreamSerializability(v, pkg1.StructStreamNotUnsafeMUS, t)
	})

}
