package mlir

import (
	"fmt"
	"io"
	"strings"
)

//type printer struct {
//	output io.Writer
//	indent int         // current indentation level
//}
//
//func Fprint(w io.Writer, fset *token.FileSet, x any) error {
//	p := printer{
//		output: w,
//		indent: 0,
//	}
//
//
//	p.print("das")
//	return nil
//}

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
	_, _ = fmt.Fprintf(w, "%s.%s", o.Dialect, o.Name)
	operands := mapSlice(o.Operands, func(t ValueId) string { return string(t) })
	joined := strings.Join(operands, ", ")
	_, _ = fmt.Fprintf(w, "(%s) ", joined)
	renderRegions(o.Regions, w, indent)
	//	TODO: render attributes, types...
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
	_, _ = fmt.Fprint(w, ")")

}

func (block BasicBlock) RenderTo(w io.Writer, indent string) {
	_, _ = fmt.Fprintf(w, "%s%v\n", indent, block.Label)
	for _, o := range block.Items {
		o.RenderTo(w, indent)
	}

}

func (region Region) RenderTo(w io.Writer, indent string) {
	_, _ = fmt.Fprint(w, "{\n")
	for i, block := range region.Blocks {
		if i > 0 {
			_, _ = fmt.Fprint(w, "\n")
		}
		block.RenderTo(w, indent)
	}
	_, _ = fmt.Fprint(w, "}\n")

}

//
//
