package arith

import (
	"awesomeProject/hlir"
)

func MakeConstant() *hlir.Operator {
	return &hlir.Operator{
		Dialect:     "arith",
		Name:        "constant",
		Arguments:   nil,
		Blocks:      nil,
		ReturnNames: nil,
		Attributes:  map[string]hlir.Attribute{},
	}
}
