package dfa

import (
	"fmt"
	"io"
)

func (m *M) WriteGo(w io.Writer) {
	fmt.Fprintln(w, "&dfa.M{")
	fmt.Fprintln(w, "States: dfa.States{")
	for i := range m.States {
		m.States[i].writeGo(w)
	}
	fmt.Fprint(w, `}}`)
}

func (s *S) writeGo(w io.Writer) {
	fmt.Fprintln(w, "{")
	if s.Label != 0 {
		fmt.Fprintf(w, "Label: %d,\n", s.Label)
	}
	if s.Table != nil {
		fmt.Fprint(w, "Table: dfa.TransTable{")
		s.Table.writeGo(w)
		fmt.Fprint(w, "}")
	}
	fmt.Fprintln(w, "},")
}

func (t *TransTable) writeGo(w io.Writer) {
	if len(*t) > 1 {
		fmt.Fprintln(w, "")
	}
	for i := range *t {
		(*t)[i].writeGo(w)
	}
	if len(*t) > 1 {
		fmt.Fprintln(w, "")
	}
}

func (t *Trans) writeGo(w io.Writer) {
	fmt.Fprintf(w, "{0x%.2x, 0x%.2x, %d},\n", t.Lo, t.Hi, t.Next)
}
