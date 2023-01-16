package hlir

import (
	"fmt"
	"io"
	"strings"
)

func mapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func (op *Operator) RenderTo(w io.Writer, indent string) {
	_, _ = fmt.Fprint(w, indent)
	if op.ReturnNames != nil {
		tmp := mapSlice(op.ReturnNames, func(t ValueId) string { return string(t) })
		_, _ = fmt.Fprintf(w, "%s = ", strings.Join(tmp, ", "))
	}
	_, _ = fmt.Fprintf(w, "\"%s.%s\"", op.Dialect, op.Name)
	operands := mapSlice(op.Arguments, func(t ValueId) string { return string(t) })
	joined := strings.Join(operands, ", ")
	_, _ = fmt.Fprintf(w, "(%s) ", joined)
	renderRegions(op.Blocks, w, indent)
	op.Attributes.RenderTo(w, indent)
	_, _ = fmt.Fprintf(w, ": () -> ()\n")
}

func renderRegions(regions []Block, w io.Writer, indent string) {
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
		o.RenderTo(w, indent+"  ")
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
