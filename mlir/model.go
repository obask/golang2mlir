package mlir

import "fmt"

type Operator struct {
	Name         string
	Dialect      string
	Operands     []ValueId
	OperandTypes []SimpleType
	Regions      []Region
	ReturnNames  []ValueId
	ReturnTypes  []SimpleType
	Attributes   AttributesMap
}

func (op *Operator) SetReturnName() {
	sprintf := fmt.Sprintf("%p", op)
	if op.ReturnNames == nil {
		op.ReturnNames = []ValueId{ValueId("%" + sprintf[len(sprintf)-3:])}
	}
}

type AttributesMap map[string]Attribute

type Region struct {
	Label *Label
	Items []Operator
}

type Label struct {
	Name        BlockId
	ParamValues []ValueId
	ParamTypes  []SimpleType
}

type BlockId string

type ValueId string

type Attribute interface {
	attributeImpl()
}

type SimpleType string

type FunctionTypeAttr struct {
	Params     []ValueId
	Types      []SimpleType
	ReturnType SimpleType
}

func (s FunctionTypeAttr) attributeImpl() {}

type NumberAttr int

func (s NumberAttr) attributeImpl() {}

type StringAttr string

func (s StringAttr) attributeImpl() {}
