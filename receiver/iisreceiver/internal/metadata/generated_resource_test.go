// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetIisApplicationPool("iis.application_pool-val")
			rb.SetIisSite("iis.site-val")

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return 0

			switch test {
			case "default":
				assert.Equal(t, 2, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 2, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("iis.application_pool")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "iis.application_pool-val", val.Str())
			}
			val, ok = res.Attributes().Get("iis.site")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "iis.site-val", val.Str())
			}
		})
	}
}
