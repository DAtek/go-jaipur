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
			wanted: "(D)iamond: 2",
		},
		{
			goods:  core.GoodMap{core.GoodCloth: 1, core.GoodCamel: 3},
			wanted: "(Ca)mel: 3, (Cl)oth: 1",
		},
		{
			goods:  core.GoodMap{core.GoodCloth: 1, core.GoodCamel: 0},
			wanted: "(Cl)oth: 1",
		},
	}

	for _, s := range scenarios {
		t.Run("Test good map representation", func(t *testing.T) {
			result := formatGoodMap(s.goods)

			assert.Equal(t, s.wanted, result)
		})
	}
}

func TestGoodAbbreviations(t *testing.T) {
	scenarios := []struct {
		abbreviation   string
		wantedGoodType core.GoodType
	}{
		{"SP", core.GoodSpice},
		{"Sp", core.GoodSpice},
		{"sp", core.GoodSpice},
		{"sP", core.GoodSpice},
		{"d", core.GoodDiamond},
	}

	for _, s := range scenarios {
		t.Run("Test good type found", func(t *testing.T) {
			found, ok := goodAbbreviations.find(s.abbreviation)

			assert.True(t, ok)
			assert.Equal(t, s.wantedGoodType, found)
		})
	}

	t.Run("Test good type not found", func(t *testing.T) {
		_, ok := goodAbbreviations.find("non existent def")

		assert.False(t, ok)
	})
}
