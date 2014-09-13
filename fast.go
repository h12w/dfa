package dfa

import "reflect"

type FastM struct {
	States []FastS
}
type FastS struct {
	Label int
	Trans [256]*FastS
}
type Action struct {
}

func (m *M) ToFast() *FastM {
	fm := &FastM{make([]FastS, len(m.states))}
	for i := range m.states {
		fm.States[i] = m.states[i].toFast(fm)
	}
	return fm
}

func (m *FastM) Count() int {
	return len(m.States)
}

func (m *FastM) Size() int {
	return int(reflect.TypeOf(FastS{}).Size()) * len(m.States)
}

func (s *state) toFast(fm *FastM) (fs FastS) {
	fs.Label = int(s.label - labeledFinalStart)
	for _, trans := range s.table.a {
		b := trans.s
		for {
			fs.Trans[b] = &fm.States[trans.next]
			if b == trans.e {
				break
			}
			b++
		}
	}
	return
}
