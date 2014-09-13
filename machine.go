package dfa

type M struct {
	states
	start int
}
type states []state

func (m *M) String() string {
	return m.dump()
}

func (m *M) clone() *M {
	if m == nil {
		return nil
	}
	return &M{m.states.clone(), m.start}
}

func (m *M) Count() int {
	return m.states.count()
}

func (m *M) startState() *state {
	return &m.states[m.start]
}

func (m *M) As(label int) *M {
	m.states.eachFinal(func(f *state) {
		f.label = stateLabel(label).toInternal()
	})
	return m
}

func (ss states) each(visit func(*state)) {
	for i := range ss {
		visit(&ss[i])
	}
}

func (ss states) eachFinal(visit func(*state)) {
	for i := range ss {
		if ss[i].final() {
			visit(&ss[i])
		}
	}
}

func (ss states) count() int {
	return len(ss)
}

func (ss states) clone() states {
	ss = append(states(nil), ss...)
	for i := range ss {
		ss[i] = ss[i].clone()
	}
	return ss
}

func (ss states) state(id int) *state {
	if id < 0 {
		return nil
	}
	return &ss[id]
}

func (ss states) shiftID(offset int) {
	ss.each(func(s *state) {
		s.each(func(t *trans) {
			t.next += offset
		})
	})
}
