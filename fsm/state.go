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
	Name    StateName
	Variant Variant
	Transit func() StateName
}
