package ir

import (
	"fmt"
	"io"
	"strings"
)

func MapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

// RenderArgsOperators
//
//	protected fun renderOperandsOperators(op: Operator, indent: String): MutableList<String> {
//	       val args = mutableListOf<String>()
//	       for (argument in op.arguments) {
//	           if (argument.name == "load") {
//	               args += (argument.attributes["value"] as StringAttr).value
//	           } else {
//	               args += renderToSB(argument, indent, isOperand = true)
//	               sb.appendLine()
//	           }
//	       }
//	       return args
//	   }
//fn (op *Operator) RenderArgsOperators(w io.Writer, indent string) []string {
//	var result []string
//	for _, arg := range op.Arguments {
//		if arg.Name == "load" {
//			result = append(result, string(arg.Attributes["value"].(StringAttr)))
//		} else {
//			result = append(result, arg.RenderTo(w, indent))
//			//_, _ = fmt.Fprint(w, '\n')
//		}
//	}
//	return result
//}

func (op *Operator) RenderTo(w io.Writer, indent string) {
	op.renderGenericLhs(w, indent)
	op.renderDefaultRhs(w, indent)
}

func (op *Operator) renderDefaultRhs(w io.Writer, indent string) {
	switch op.T {
	case FuncCall:
		operands := MapSlice(op.Arguments, func(t ValueId) string { return "%" + string(t.ReturnName()) })
		joined := strings.Join(operands, ", ")
		_, _ = fmt.Fprintf(w, "func.call NAME(%s) ", joined)
		op.Attributes.RenderTo(w, indent)
		_, _ = fmt.Fprintf(w, ": () -> ()")
		break
	case FuncFunc:
		operands := MapSlice(op.Arguments, func(t ValueId) string { return "%" + string(t.ReturnName()) })
		joined := strings.Join(operands, ", ")
		_, _ = fmt.Fprintf(w, "func.func @%s(%s) ", op.Attributes["sym_name"], joined)
		op.Blocks[0].RenderTo(w, indent)
		break
	case ScfYield:
		_, _ = fmt.Fprintf(w, "scf.yield")
		break
	case ScfIf:
		operands := MapSlice(op.Arguments, func(t ValueId) string { return "%" + string(t.ReturnName()) })
		joined := strings.Join(operands, ", ")
		_, _ = fmt.Fprintf(w, "scf.if %s ", joined)
		op.Blocks[0].RenderTo(w, indent)
		if len(op.Blocks) > 1 {
			_, _ = fmt.Fprintf(w, " else ")
			op.Blocks[0].RenderTo(w, indent)
		}
		break
	default:
		_, _ = fmt.Fprintf(w, "\"%s.%s\"", op.Dialect, op.Name)
		operands := MapSlice(op.Arguments, func(t ValueId) string { return "%" + string(t.ReturnName()) })
		joined := strings.Join(operands, ", ")
		_, _ = fmt.Fprintf(w, "(%s) ", joined)
		RenderRegions(op.Blocks, w, indent)
		op.Attributes.RenderTo(w, indent)
		_, _ = fmt.Fprintf(w, ": () -> ()")
	}
}

func (op *Operator) renderGenericLhs(w io.Writer, indent string) {
	_, _ = fmt.Fprint(w, indent)
	if op.ReturnNames != nil {
		tmp := MapSlice(op.ReturnNames, func(t SimpleName) string { return "%" + string(t) })
		_, _ = fmt.Fprintf(w, "%s = ", strings.Join(tmp, ", "))
	}
}

func RenderRegions(regions []Block, w io.Writer, indent string) {
	if regions == nil {
		return
	}
	_, _ = fmt.Fprint(w, "(")
	for i, r := range regions {
		if i > 0 {
			_, _ = fmt.Fprint(w, ", ")
		}
		r.RenderTo(w, indent)
	}
	_, _ = fmt.Fprint(w, ") ")
}

func (l Label) RenderTo(w io.Writer, indent string) {
	_, _ = fmt.Fprintf(w, "%s%s(", indent, l.Name)
	var results []string
	for i := range l.ParamTypes {
		results = append(results, fmt.Sprintf("%s: %s", l.ParamValues[i], l.ParamTypes[i]))
	}
	joined := strings.Join(results, ", ")
	_, _ = fmt.Fprintf(w, "%s):\n", joined)
}

func (region Block) RenderTo(w io.Writer, indent string) {
	_, _ = fmt.Fprint(w, "{\n")
	if region.Label != nil {
		region.Label.RenderTo(w, indent)
	}
	for _, o := range region.Items {
		o.RenderTo(w, indent+"    ")
		_, _ = fmt.Fprintf(w, "\n")
	}
	_, _ = fmt.Fprintf(w, "%s}", indent)
}

func (attributes AttributesMap) RenderTo(w io.Writer, indent string) {
	if len(attributes) == 0 {
		return
	}
	_, _ = fmt.Fprint(w, "{")
	isFirst := true
	for k, v := range attributes {
		if isFirst {
			isFirst = false
		} else {
			_, _ = fmt.Fprint(w, ", ")
		}
		switch expr := v.(type) {
		case StringAttr:
			_, _ = fmt.Fprintf(w, "%s: \"%v\"", k, expr)
		default:
			_, _ = fmt.Fprintf(w, "%s: %v", k, expr)
		}
	}
	_, _ = fmt.Fprint(w, "} ")
}
