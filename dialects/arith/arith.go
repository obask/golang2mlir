package arith

import (
	"awesomeProject/ir"
)

func MakeConstant() *ir.Operator {
	return &ir.Operator{
		Dialect:     "arith",
		Name:        "constant",
		Arguments:   nil,
		Blocks:      nil,
		ReturnNames: nil,
		Attributes:  map[string]ir.Attribute{},
	}
}
