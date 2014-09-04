package dfa

import "reflect"

type FastM struct {
	States []FastS
}
type FastS struct {
	Label int
	Trans [256]*FastS
}

func (m *M) ToFast() *FastM {
	fm := &FastM{make([]FastS, len(m.states))}
	for i := range m.states {
		fm.States[i] = m.states[i].toFast(fm)
	}
	return fm
}

func (m *FastM) Size() int {
	return int(reflect.TypeOf(FastS{}).Size()) * len(m.States)
}

func (s *state) toFast(fm *FastM) (fs FastS) {
	fs.Label = int(s.label - labeledFinalStart)
	for _, trans := range s.table {
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

// Match greedily matches the DFA against src.
func (m *FastM) Match(src []byte) (size, label int, matched bool) {
	var (
		s   = &m.States[0]
		pos int
	)
	for {
		if s.Label >= 0 {
			size = pos
			label = s.Label
			matched = true
		}
		if pos == len(src) {
			break
		}
		s = s.Trans[src[pos]]
		if s == nil {
			break
		}
		pos++
	}
	return
}
