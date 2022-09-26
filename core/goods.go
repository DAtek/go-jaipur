package core

const (
	GoodDiamond GoodType = iota
	GoodGold
	GoodSilver
	GoodCloth
	GoodSpice
	GoodLeather
	GoodCamel
)

var expensiveGoods = []GoodType{GoodDiamond, GoodGold, GoodSilver}
