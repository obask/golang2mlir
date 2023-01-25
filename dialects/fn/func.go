package fn

import (
	"awesomeProject/ir"
)

func MakeCallOp() *ir.Operator {
	return &ir.Operator{
		T:          ir.FuncCall,
		Attributes: map[string]ir.Attribute{},
	}
}
