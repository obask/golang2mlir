package ir

import "fmt"

type OpType string

const (
	ArithConstant OpType = "arith.constant"
	BuiltinCast          = "func.call"
	FuncCall             = "func.call"
	FuncFunc             = "func.func"
	FuncConstant         = "func.constant"
	ScfIf                = "scf.if"
	ScfWhile             = "scf.while"
	ScfYield             = "scf.yield"
	GenericUast          = "uast.*"
)

type Operator struct {
	T             OpType
	ReturnNames   []SimpleName
	Dialect       string
	Name          string
	Blocks        []Block
	Arguments     []ValueId
	Attributes    AttributesMap
	ArgumentTypes []SimpleType
	ReturnTypes   []SimpleType
}

func (op *Operator) SetReturnName() {
	sprintf := fmt.Sprintf("%p", op)
	if op.ReturnNames == nil {
		op.ReturnNames = append(op.ReturnNames, SimpleName(sprintf[len(sprintf)-3:]))
	}
}

func (op *Operator) Attr(field string, value string) {
	op.Attributes[field] = StringAttr(value)
}

type AttributesMap map[string]Attribute

type Block struct {
	Label *Label
	Items []Operator
}

type Label struct {
	Name        BlockId
	ParamValues []ValueId
	ParamTypes  []SimpleType
}

type BlockId string

type Attribute interface {
	attributeImpl()
}

type SimpleName string

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

type ReferenceAttr string

func (s ReferenceAttr) attributeImpl() {}
