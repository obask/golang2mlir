package ir

type ValueId interface {
	ReturnName() SimpleName
	ReturnType() SimpleType
}

type ConstValue struct {
	Name SimpleName
	Type SimpleType
}

func (v ConstValue) ReturnName() SimpleName {
	return v.Name
}

func (v ConstValue) ReturnType() SimpleType {
	return v.Type
}

type SsaValue struct {
	Ref *Operator
}

func (v SsaValue) ReturnName() SimpleName {
	return v.Ref.ReturnNames[0]
}

func (v SsaValue) ReturnType() SimpleType {
	return v.Ref.ReturnTypes[0]
}
