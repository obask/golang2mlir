package mlir

type Operator struct {
	Name       string
	Dialect    string
	Operands   []ValueId
	Regions    []Region
	ReturnName ValueId
	Attributes map[string]Attribute
}

type Region struct {
	Blocks []BasicBlock
}

type BasicBlock struct {
	Label BlockLabel
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

func (s *FunctionTypeAttr) attributeImpl() {}

type NumberAttr string

func (s *NumberAttr) attributeImpl() {}

type StringAttr string

func (s *StringAttr) attributeImpl() {}

//
//
//