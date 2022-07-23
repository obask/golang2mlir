package mlir

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

func (o Operator) RenderTo(w io.Writer, indent string) {
	_, _ = fmt.Fprint(w, indent)
	if o.ReturnName != "" {
		_, _ = fmt.Fprintf(w, "%s = ", o.ReturnName)
	}
	_, _ = fmt.Fprintf(w, "\"%s.%s\"", o.Dialect, o.Name)
	operands := mapSlice(o.Operands, func(t ValueId) string { return string(t) })
	joined := strings.Join(operands, ", ")
	_, _ = fmt.Fprintf(w, "(%s) ", joined)
	renderRegions(o.Regions, w, indent)
	o.Attributes.RenderTo(w, indent)
	_, _ = fmt.Fprintf(w, ": () -> ()\n")
}

func renderRegions(regions []Region, w io.Writer, indent string) {
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

func (l BlockLabel) RenderTo(w io.Writer, indent string) {
	_, _ = fmt.Fprintf(w, "%s%s(", indent, l.Name)
	var results []string
	for i := range l.ParamTypes {
		results = append(results, fmt.Sprintf("%s: %s", l.ParamValues[i], l.ParamTypes[i]))
	}
	joined := strings.Join(results, ", ")
	_, _ = fmt.Fprintf(w, "%s):\n", joined)
}

func (block BasicBlock) RenderTo(w io.Writer, indent string) {
	block.Label.RenderTo(w, indent)
	for _, o := range block.Items {
		o.RenderTo(w, indent+"  ")
	}
}

func (region Region) RenderTo(w io.Writer, indent string) {
	_, _ = fmt.Fprint(w, "{\n")
	for i, block := range region {
		if i > 0 {
			_, _ = fmt.Fprint(w, "\n")
		}
		block.RenderTo(w, indent)
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
		_, _ = fmt.Fprintf(w, "%s: %v", k, v)
	}
	_, _ = fmt.Fprint(w, "} ")
}
