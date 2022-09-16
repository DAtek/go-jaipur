package core

type (
	Amount      uint8
	GoodType    uint8
	Score       uint8
	goodMap     map[GoodType]Amount
	JaipurError string
	Name        string
)

func (e JaipurError) Error() string {
	return string(e)
}
