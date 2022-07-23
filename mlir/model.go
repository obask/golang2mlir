package mlir

type Operator struct {
	Name       string
	Dialect    string
	Operands   []ValueId
	Regions    []Region
	ReturnName ValueId
	Attributes AttributesMap
}

type AttributesMap map[string]Attribute

type Region []BasicBlock

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

func (s FunctionTypeAttr) attributeImpl() {}

type NumberAttr int

func (s NumberAttr) attributeImpl() {}

type StringAttr string

func (s StringAttr) attributeImpl() {}
