package fsm

type (
	Variant   uint8
	StateName string
)

const (
	VariantStart Variant = iota
	VariantIntermediate
	VariantFinal
)

type State struct {
	name    StateName
	variant Variant
	transit func() StateName
}
