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
	fm := &FastM{make([]FastS, len(m.States))}
	for i := range m.States {
		s, fs := &m.States[i], &fm.States[i]
		fs.Label = int(s.Label - labeledFinalStart)
		for j := range s.Table {
			trans := &s.Table[j]
			for b := int(trans.Lo); b <= int(trans.Hi); b++ {
				fs.Trans[b] = &fm.States[trans.Next]
			}
		}
	}
	return fm
}

func (m *FastM) Count() int {
	return len(m.States)
}

func (m *FastM) Size() int {
	return int(reflect.TypeOf(FastS{}).Size()) * len(m.States)
}
