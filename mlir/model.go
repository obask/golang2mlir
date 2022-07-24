package mlir

import "fmt"

type Operator struct {
	Name       string
	Dialect    string
	Operands   []ValueId
	Regions    []Region
	ReturnName ValueId
	Attributes AttributesMap
}

func (operator *Operator) SetReturnName() {
	sprintf := fmt.Sprintf("%p", operator)
	operator.ReturnName = ValueId("%" + sprintf[len(sprintf)-3:])
}

type AttributesMap map[string]Attribute

type Region []BasicBlock

type BasicBlock struct {
	Label *BlockLabel
	Items []Operator
}

type BlockLabel struct {
	Name        BlockId
	ParamValues []ValueId
	ParamTypes  []KotlinType
}

type BlockId string

type ValueId string

type Attribute interface {
	attributeImpl()
}

type KotlinType string

type FunctionTypeAttr struct {
	Params     []ValueId
	Types      []KotlinType
	ReturnType KotlinType
}

func (s FunctionTypeAttr) attributeImpl() {}

type NumberAttr int

func (s NumberAttr) attributeImpl() {}

type StringAttr string

func (s StringAttr) attributeImpl() {}
