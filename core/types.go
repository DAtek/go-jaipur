package core

type (
	Amount      uint8
	GoodType    uint8
	Score       uint8
	GoodMap     map[GoodType]Amount
	JaipurError string
	Name        string
)

func (e JaipurError) Error() string {
	return string(e)
}

func (goodMap GoodMap) Copy() GoodMap {
	newMap := GoodMap{}
	for k, v := range goodMap {
		newMap[k] = v
	}
	return newMap
}
