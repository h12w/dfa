package dfa

import "unicode/utf8"

func Str(s string) *M {
	bs := []byte(s)
	ss := make(states, 0, len(bs)+1)
	for i, b := range bs {
		ss = append(ss, stateTo(b, i+1))
	}
	ss = append(ss, finalState())
	return &M{ss, 0}
}

func Between(lo, hi rune) *M {
	if lo > hi {
		lo, hi = hi, lo
	}
	if lo < 0 || hi > 0x10ffff {
		panic("invalid range for unicode point")
	}
	u := &u8s{}
	u.between(lo, hi)
	return u.m()
}

func BetweenByte(s, e byte) *M {
	return &M{states{
		stateBetween(s, e, 1),
		finalState(),
	}, 0}
}

func Char(s string) (m *M) {
	for _, r := range s {
		if r == utf8.RuneError {
			panic("invalid rune")
		}
		m = or2(m, Between(r, r))
	}
	return m
}

func opMany(op func(_, _ *M) *M, ms []*M) *M {
	switch len(ms) {
	case 0:
		return nil
	case 1:
		return ms[0].clone()
	}
	m := ms[0].clone()
	for i := 1; i < len(ms); i++ {
		m = op(m, ms[i])
	}
	return m
}

func Con(ms ...*M) *M {
	return opMany(con2, ms)
}
func con2(m1, m2 *M) *M {
	m := m1.clone()
	m2 = m2.clone()
	m2.shiftID(m.states.count() - 1)
	m.eachFinal(func(f *state) {
		f.connect(m2.startState())
		if !m2.startState().final() {
			f.label = notFinal
		}
	})
	m.states = append(m.states, m2.states[1:]...)
	return m
}

func Or(ms ...*M) *M {
	return opMany(or2, ms)
}
func or2(m1, m2 *M) *M {
	return newMerger(m1, m2, union{}).merge()
}

func And(ms ...*M) *M {
	return opMany(and2, ms)
}
func and2(m1, m2 *M) *M {
	return newMerger(m1, m2, intersection{}).merge()
}

func (m *M) ZeroOrMore() *M {
	m = m.OneOrMore()
	if len(m.states) == 2 {
		m.states = m.states[1:]
		m.shiftID(-1)
	}
	m.startState().label = defaultFinal
	return m
}

func ZeroOrMore(ms ...*M) *M {
	return Con(ms...).ZeroOrMore()
}

func (m *M) Loop(filter func(b byte) bool) *M {
	m = m.clone()
	m.eachFinal(func(f *state) {
		f.filterConnect(m.startState(), filter)
	})
	return m
}

func (m *M) OneOrMore() *M {
	m = m.clone()
	m.eachFinal(func(f *state) {
		f.connect(m.startState())
	})
	return m
}

func OneOrMore(ms ...*M) *M {
	return Con(ms...).OneOrMore()
}

func (m *M) ZeroOrOne() *M {
	m = m.clone()
	m.states[0].label = defaultFinal
	return m
}

func ZeroOrOne(ms ...*M) *M {
	return Con(ms...).ZeroOrOne()
}

func (m *M) Complement() *M {
	m = m.clone()
	m.each(func(f *state) {
		if f.final() {
			f.label = notFinal
		} else {
			f.label = defaultFinal
		}
	})
	return m
}

func (m *M) Exclude(ms ...*M) *M {
	ex := Or(ms...)
	ex.eachFinal(func(s *state) {
		s.label = 9999 // TODO remove this hack
	})
	return newMerger(m, ex, difference{}).merge().deleteUnreachable()
}

func (m *M) Repeat(n int) *M {
	mm := m.clone()
	for i := 0; i < n-1; i++ {
		mm = con2(mm, m)
	}
	return mm
}
