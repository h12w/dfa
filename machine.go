package dfa

type M struct {
	Start int
	States
}
type States []S

func (m *M) String() string {
	return m.dump()
}

func (m *M) clone() *M {
	if m == nil {
		return nil
	}
	return &M{m.Start, m.States.clone()}
}

func (m *M) startState() *S {
	return &m.States[m.Start]
}

func (m *M) As(Label int) *M {
	m.States.eachFinal(func(f *S) {
		f.Label = StateLabel(Label).toInternal()
	})
	return m
}

func (ss States) eachFinal(visit func(*S)) {
	for i := range ss {
		if ss[i].final() {
			visit(&ss[i])
		}
	}
}

func (ss States) clone() States {
	ss = append(States(nil), ss...)
	for i := range ss {
		ss[i] = ss[i].clone()
	}
	return ss
}

func (ss States) state(id int) *S {
	if id < 0 {
		return nil
	}
	return &ss[id]
}

func (ss States) shiftID(offset int) {
	for i := range ss {
		for j := range ss[i].Table {
			ss[i].Table[j].Next += offset
		}
	}
}
