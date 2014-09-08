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

// Match greedily matches the DFA against src.
func (m *FastM) Match(src []byte, p int) (size, label int, matched bool) {
	cur := &m.States[0]
	pos := p
	matchedPos := pos
	for {
		if cur.Label >= 0 {
			matchedPos = pos
			label = cur.Label
			matched = true
		}
		if pos == len(src) {
			break
		}
		if cur = cur.Trans[src[pos]]; cur == nil {
			break
		}
		pos++
	}
	size = matchedPos - p
	return
}
