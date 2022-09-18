package app

import (
	"jaipur/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatGoodMap(t *testing.T) {
	scenarios := []struct {
		goods  core.GoodMap
		wanted string
	}{
		{
			goods:  core.GoodMap{core.GoodDiamond: 2},
			wanted: "Diamond: 2",
		},
		{
			goods:  core.GoodMap{core.GoodCamel: 3, core.GoodCloth: 1},
			wanted: "Camel: 3, Cloth: 1",
		},
	}

	for _, s := range scenarios {
		t.Run("Test good map representation", func(t *testing.T) {
			result := formatGoodMap(s.goods)

			assert.Equal(t, s.wanted, result)
		})
	}
}
