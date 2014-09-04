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
	return u.m().Minimize()
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
		m = m.or(Between(r, r))
	}
	return m.Minimize()
}

func Con(ms ...*M) *M {
	return opMany((*M).con, ms).Minimize()
}
func (m1 *M) con(m2 *M) *M {
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
	return opMany((*M).or, ms).Minimize()
}
func (m1 *M) or(m2 *M) *M {
	return newMerger(m1, m2, union{}).merge()
}

func And(ms ...*M) *M {
	return opMany((*M).and, ms)
}
func (m1 *M) and(m2 *M) *M {
	return newMerger(m1, m2, intersection{}).merge()
}

func (m *M) ZeroOrMore() *M {
	m = m.loop()
	m.startState().label = defaultFinal
	return m
}

func ZeroOrMore(ms ...*M) *M {
	return Con(ms...).ZeroOrMore()
}

func (m *M) loop() *M {
	m = m.clone()
	m.eachFinal(func(f *state) {
		f.connect(m.startState())
	})
	return m.Minimize()
}

func (m *M) Loop(filters ...func(b byte) bool) *M {
	m = m.clone()
	m.eachFinal(func(f *state) {
		f.filterConnect(m.startState(), filters)
	})
	return m.Minimize()
}
func IfNot(bs ...byte) func(byte) bool {
	return func(input byte) bool {
		for _, b := range bs {
			if b == input {
				return false
			}
		}
		return true
	}
}

func (m *M) OneOrMore() *M {
	return m.loop()
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
	return newMerger(m, Or(ms...), difference{}).merge().deleteUnreachable()
}

func (m *M) Repeat(n int) *M {
	mm := m.clone()
	for i := 0; i < n-1; i++ {
		mm = mm.con(m)
	}
	return mm
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
