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
	fmt.Fprintln(f, `func match(src []byte, pos int) (size, Label int, matched bool) {`)
	fmt.Fprintln(f, `start := pos`)
	fmt.Fprintln(f, `p := pos`)
	fmt.Fprintln(f, `matchedPos := pos`)
	fmt.Fprintln(f, `goto s0`)
	for i, s := range m.States {
		s.writeGo(f, i)
	}
	fmt.Fprintln(f, `sEnd:`)
	fmt.Fprintln(f, `size = matchedPos - start`)
	fmt.Fprintln(f, `return`)
	fmt.Fprintln(f, `}`)
	return nil
}

func (s *S) writeGo(w io.Writer, id int) {
	fmt.Fprintf(w, "s%d:\n", id)
	if s.Label >= labeledFinalStart {
		fmt.Fprintln(w, `matched = true`)
		fmt.Fprintln(w, `matchedPos = pos`)
		fmt.Fprintf(w, "Label = %d\n", s.Label.toExternal())
	}
	fmt.Fprintln(w, `if pos == len(src) {`)
	fmt.Fprintln(w, `goto sEnd`)
	fmt.Fprintln(w, `}`)
	fmt.Fprintln(w, `p = pos`)
	fmt.Fprintln(w, `pos++`)
	s.Table.writeGo(w)
	fmt.Fprintln(w, `goto sEnd`)
}

func (t *TransTable) writeGo(w io.Writer) {
	m := make(map[int][]byte)
	for _, Trans := range *t {
		b := Trans.Lo
		for {
			m[Trans.Next] = append(m[Trans.Next], b)
			if b == Trans.Hi {
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
	//	for _, Trans := range *t {
	//		Trans.writeGo(w)
	//	}
	//	fmt.Fprintln(w, `}`)
}

func (t *Trans) writeGo(w io.Writer) {
	b := t.Lo
	fmt.Fprintf(w, "case ")
	for {
		if b == t.Hi {
			fmt.Fprintf(w, "0x%.2x", b)
			break
		} else {
			fmt.Fprintf(w, "0x%.2x, ", b)
		}
		b++
	}
	fmt.Fprintln(w, ":")
	fmt.Fprintf(w, "goto s%d\n", t.Next)
}
