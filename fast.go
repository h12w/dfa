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
	fm := &FastM{make([]FastS, len(m.States))}
	for i := range m.States {
		fm.States[i] = m.States[i].toFast(fm)
	}
	return fm
}

func (m *FastM) Count() int {
	return len(m.States)
}

func (m *FastM) Size() int {
	return int(reflect.TypeOf(FastS{}).Size()) * len(m.States)
}

func (s *S) toFast(fm *FastM) (fs FastS) {
	fs.Label = int(s.Label - labeledFinalStart)
	for _, trans := range s.Table {
		b := trans.Lo
		for {
			fs.Trans[b] = &fm.States[trans.Next]
			if b == trans.Hi {
				break
			}
			b++
		}
	}
	return
}
