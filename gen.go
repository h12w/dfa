package dfa

import (
	"fmt"
	"io"
	"os"
)

func (m *M) SaveGo(file, pac string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprint(f, `package `)
	fmt.Fprint(f, pac)
	fmt.Fprintln(f, "")
	fmt.Fprintln(f, `func match(src []byte, pos int) (size, label int, matched bool) {`)
	fmt.Fprintln(f, `start := pos`)
	fmt.Fprintln(f, `p := pos`)
	fmt.Fprintln(f, `matchedPos := pos`)
	fmt.Fprintln(f, `goto s0`)
	for i, s := range m.states {
		s.writeGo(f, i)
	}
	fmt.Fprintln(f, `sEnd:`)
	fmt.Fprintln(f, `size = matchedPos - start`)
	fmt.Fprintln(f, `return`)
	fmt.Fprintln(f, `}`)
	return nil
}

func (s *state) writeGo(w io.Writer, id int) {
	fmt.Fprintf(w, "s%d:\n", id)
	if s.label >= labeledFinalStart {
		fmt.Fprintln(w, `matched = true`)
		fmt.Fprintln(w, `matchedPos = pos`)
		fmt.Fprintf(w, "label = %d\n", s.label.toExternal())
	}
	fmt.Fprintln(w, `if pos == len(src) {`)
	fmt.Fprintln(w, `goto sEnd`)
	fmt.Fprintln(w, `}`)
	fmt.Fprintln(w, `p = pos`)
	fmt.Fprintln(w, `pos++`)
	s.table.writeGo(w)
	fmt.Fprintln(w, `goto sEnd`)
}

func (t *transTable) writeGo(w io.Writer) {
	m := make(map[int][]byte)
	for _, trans := range t.a {
		b := trans.s
		for {
			m[trans.next] = append(m[trans.next], b)
			if b == trans.e {
				break
			}
			b++
		}
	}
	fmt.Fprintln(w, `switch src[p] {`)
	for next, a := range m {
		fmt.Fprintf(w, "case ")
		for i, b := range a {
			if i == 0 {
				fmt.Fprintf(w, "0x%.2x", b)
			} else {
				fmt.Fprintf(w, ", 0x%.2x", b)
			}
		}
		fmt.Fprintln(w, ":")
		fmt.Fprintf(w, "goto s%d\n", next)
	}
	fmt.Fprintln(w, `}`)

	//	fmt.Fprintln(w, `switch src[p] {`)
	//	for _, trans := range t.a {
	//		trans.writeGo(w)
	//	}
	//	fmt.Fprintln(w, `}`)
}

func (t *trans) writeGo(w io.Writer) {
	b := t.s
	fmt.Fprintf(w, "case ")
	for {
		if b == t.e {
			fmt.Fprintf(w, "0x%.2x", b)
			break
		} else {
			fmt.Fprintf(w, "0x%.2x, ", b)
		}
		b++
	}
	fmt.Fprintln(w, ":")
	fmt.Fprintf(w, "goto s%d\n", t.next)
}
