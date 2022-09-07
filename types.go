package jaipur

type (
	Amount      uint8
	ProductType uint8
	Score       uint8
	productMap  map[ProductType]Amount
	JaipurError string
	Name        string
)

func (e JaipurError) Error() string {
	return string(e)
}
