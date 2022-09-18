package app

import (
	"fmt"
	"jaipur/core"
	"strings"
)

const separator = ", "

var goodsRepresentation = map[core.GoodType]string{
	core.GoodDiamond: "Diamond",
	core.GoodGold:    "Gold",
	core.GoodSilver:  "Silver",
	core.GoodCloth:   "Cloth",
	core.GoodSpice:   "Spice",
	core.GoodLeather: "Leather",
	core.GoodCamel:   "Camel",
}

func formatGoodMap(goods core.GoodMap) string {
	parts := []string{}

	for key, value := range goods {
		parts = append(parts, fmt.Sprintf("%s: %d", goodsRepresentation[key], value))
	}

	return strings.Join(parts, separator)
}
