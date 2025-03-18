package musgen

import (
	"testing"
	"time"

	"github.com/mus-format/musgen-go/testdata/pkg1"
)

func TestNotUnsafeGeneratedCode(t *testing.T) {

	t.Run("StructNotUnsafe should be serializable", func(t *testing.T) {
		var (
			v = pkg1.StructNotUnsafe{
				String: "abc",
				Int:    3,
				Time:   time.Unix(time.Now().Unix(), 0),
			}
		)
		testSerializability(v, pkg1.StructNotUnsafeMUS, nil, t)
	})

}
