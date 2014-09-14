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
	return m
}

func Con(ms ...interface{}) *M {
	return conMany(toMs(ms))
}
func conMany(ms []*M) *M {
	return opMany((*M).con, ms).Minimize()
}
func (m1 *M) con(m2 *M) *M {
	m := m1.clone()
	if m2 == nil {
		return m
	}
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

func Or(ms ...interface{}) *M {
	return orMany(toMs(ms))
}
func orMany(ms []*M) *M {
	return opMany((*M).or, ms)
}
func (m1 *M) or(m2 *M) *M {
	return newMerger(m1, m2, union{}).merge()
}

func And(ms ...interface{}) *M {
	return andMany(toMs(ms))
}
func andMany(ms []*M) *M {
	return opMany((*M).and, ms)
}
func (m1 *M) and(m2 *M) *M {
	return newMerger(m1, m2, intersection{}).merge()
}

func (m *M) zeroOrMore() *M {
	m = m.loop()
	m.startState().label = defaultFinal
	return m
}

func (m *M) AtLeast(n int) *M {
	switch n {
	case 0:
		return m.zeroOrMore()
	case 1:
		return m.oneOrMore()
	}
	ms := make([]*M, n)
	for i := range ms {
		ms[i] = m
	}
	ms[n-1] = m.oneOrMore()
	return conMany(ms)
}

func (m *M) AtMost(n int) *M {
	ms := make([]*M, n)
	ms[0] = m
	for i := 1; i < len(ms); i++ {
		ms[i] = ms[i-1].con(m)
	}
	return orMany(ms).Optional()
}

func (m *M) loop() *M {
	m = m.clone()
	m.eachFinal(func(f *state) {
		f.connect(m.startState())
	})
	return m
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

func (m *M) oneOrMore() *M {
	return m.loop()
}

func (m *M) Optional() *M {
	m = m.clone()
	m.states[0].label = defaultFinal
	return m
}

func Optional(ms ...*M) *M {
	return conMany(ms).Optional()
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
	return m.deleteUnreachable()
}

func (m *M) Exclude(ms ...interface{}) *M {
	return newMerger(m, Or(ms...), difference{}).merge().deleteUnreachable()
}

func (m *M) Repeat(limit ...int) *M {
	switch len(limit) {
	case 0:
		return m.zeroOrMore()
	case 1:
		n := limit[0]
		ms := make([]*M, n)
		for i := range ms {
			ms[i] = m
		}
		return conMany(ms)
	case 2:
		lo, hi := limit[0], limit[1]
		if lo > hi {
			lo, hi = hi, lo
		}
		ms := make([]*M, 0, hi-lo+1)
		for n := lo; n <= hi; n++ {
			ms = append(ms, m.Repeat(n))
		}
		return orMany(ms)
	}
	panic("repeat should have zero to two arguments")
}

func opMany(op func(_, _ *M) *M, ms []*M) *M {
	switch len(ms) {
	case 0:
		return nil
	case 1:
		return ms[0].clone()
	}
	for len(ms) > 1 {
		if len(ms)%2 != 0 {
			if ms[len(ms)-1] == nil {
				ms = ms[:len(ms)-1]
			} else {
				ms = append(ms, nil)
			}
		}
		cur := 0
		for i := 0; i < len(ms)-1; i += 2 {
			ms[cur] = op(ms[i], ms[i+1])
			cur++
		}
		ms = ms[:cur]
	}
	return ms[0]
}

func toMs(a []interface{}) []*M {
	ms := make([]*M, len(a))
	for i := range a {
		switch o := a[i].(type) {
		case *M:
			ms[i] = o
		case string:
			ms[i] = Str(o)
		default:
			panic("type of argument should be either string or *M")
		}
	}
	return ms
}
