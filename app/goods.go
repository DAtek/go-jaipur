package app

import (
	"fmt"
	"jaipur/core"
	"sort"
	"strings"
)

const separator = ", "

type goodAbbreviationMap map[string]core.GoodType

var goodRepresentations = map[core.GoodType]string{
	core.GoodDiamond: "(D)iamond",
	core.GoodGold:    "(G)old",
	core.GoodSilver:  "(Si)lver",
	core.GoodCloth:   "(Cl)oth",
	core.GoodSpice:   "(Sp)ice",
	core.GoodLeather: "(L)eather",
	core.GoodCamel:   "(Ca)mel",
}

var goodAbbreviations = goodAbbreviationMap{
	"D":  core.GoodDiamond,
	"G":  core.GoodGold,
	"Si": core.GoodSilver,
	"Cl": core.GoodCloth,
	"Sp": core.GoodSpice,
	"L":  core.GoodLeather,
	"Ca": core.GoodCamel,
}

var inverseGoodRepresentation = map[string]core.GoodType{}

func formatGoodMap(goods core.GoodMap) string {
	if len(inverseGoodRepresentation) == 0 {
		for key, value := range goodRepresentations {
			inverseGoodRepresentation[value] = key
		}
	}

	parts := []string{}

	keys := []core.GoodType{}
	for key := range goods {
		keys = append(keys, key)
	}

	goodNames := []string{}
	for _, key := range keys {
		goodNames = append(goodNames, goodRepresentations[key])
	}

	sort.Strings(goodNames)

	for _, key := range goodNames {
		goodType := inverseGoodRepresentation[key]
		value := goods[goodType]
		if value == 0 {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s: %d", key, value))
	}

	return strings.Join(parts, separator)
}

func (m goodAbbreviationMap) find(abbreviation string) (core.GoodType, bool) {
	abbreviation = strings.ToLower(abbreviation)
	for key, value := range goodAbbreviations {
		if strings.ToLower(key) == abbreviation {
			return value, true
		}
	}

	return 0, false
}
