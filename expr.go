package dfa

import (
	"fmt"
	"unicode/utf8"
)

func Str(s string) *M {
	bs := []byte(s)
	ss := make(States, 0, len(bs)+1)
	for i, b := range bs {
		ss = append(ss, stateTo(b, i+1))
	}
	ss = append(ss, finalState())
	return &M{0, ss}
}

func Between(lo, hi rune) *M {
	if lo > hi {
		lo, hi = hi, lo
	}
	if lo < 0 || hi > 0x10ffff {
		panic(fmt.Errorf("invalid range for unicode point: 0x%x, 0x%x", lo, hi))
	}
	u := &u8s{}
	u.between(lo, hi)
	m, err := u.m()
	if err != nil {
		panic(err)
	}
	return m.Minimize()
}

func BetweenByte(s, e byte) *M {
	return &M{0, States{
		stateBetween(s, e, 1),
		finalState(),
	}}
}

func Char(s string) (m *M) {
	ms := make([]*M, 0, 64)
	for _, r := range s {
		if r == utf8.RuneError {
			panic("invalid rune")
		}
		ms = append(ms, Between(r, r))
	}
	m, err := orMany(ms)
	if err != nil {
		panic(err)
	}
	return m
}

func CharClass(name string) *M {
	m, err := charClass(name)
	if err != nil {
		panic(err)
	}
	return m
}

/*
1. If m2's start state is final, then f1's final should be kept in con(m1, m2).
*/

func Con(ms ...interface{}) *M {
	m, err := conMany(toMs(ms))
	if err != nil {
		panic(err)
	}
	return m.Minimize()
}
func conMany(ms []*M) (*M, error) {
	return opMany((*M).con, ms)
}
func (m1 *M) con(m2 *M) (*M, error) {
	m := m1.clone()
	if m2 == nil {
		return m, nil
	}
	m2 = m2.clone()
	m2.shiftID(len(m.States) - 1)
	for i := range m.States {
		if f := &m.States[i]; f.final() {
			if err := f.connect(m2.startState()); err != nil {
				return nil, err
			}
			if !m2.startState().final() {
				f.Label = notFinal
			}
		}
	}
	m.States = append(m.States, m2.States[1:]...)
	return m, nil
}

func Or(ms ...interface{}) *M {
	m, err := orMany(toMs(ms))
	if err != nil {
		panic(err)
	}
	return m
}
func orMany(ms []*M) (*M, error) {
	return opMany((*M).or, ms)
}
func (m1 *M) or(m2 *M) (*M, error) {
	if m1 == nil {
		return m2.clone(), nil
	}
	if m2 == nil {
		return m1.clone(), nil
	}
	m, err := newMerger(m1, m2, union{}).merge()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func And(ms ...interface{}) *M {
	m, err := andMany(toMs(ms))
	if err != nil {
		panic(err)
	}
	return m
}
func andMany(ms []*M) (*M, error) {
	return opMany((*M).and, ms)
}
func (m1 *M) and(m2 *M) (*M, error) {
	if m1 == nil {
		return m2.clone(), nil
	}
	if m2 == nil {
		return m1.clone(), nil
	}
	m, err := newMerger(m1, m2, intersection{}).merge()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *M) zeroOrMore() *M {
	m = m.loop()
	m.startState().Label = defaultFinal
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
	m, err := conMany(ms)
	if err != nil {
		panic(err)
	}
	return m
}

func (m *M) AtMost(n int) *M {
	var err error
	ms := make([]*M, n)
	ms[0] = m
	for i := 1; i < len(ms); i++ {
		ms[i], err = ms[i-1].con(m)
		if err != nil {
			panic(err)
		}
	}
	m, err = orMany(ms)
	if err != nil {
		panic(err)
	}
	return m.Optional()
}

func (m *M) loop() *M {
	m = m.clone()
	m.eachFinal(func(f *S) {
		if err := f.connect(m.startState()); err != nil {
			panic(err)
		}
	})
	return m
}

func (m *M) Loop(filters ...func(b byte) bool) *M {
	m = m.clone()
	m.eachFinal(func(f *S) {
		if err := f.filterConnect(m.startState(), filters); err != nil {
			panic(err)
		}
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
	m.States[0].Label = defaultFinal
	return m
}

func Optional(ms ...*M) *M {
	m, err := conMany(ms)
	if err != nil {
		panic(err)
	}
	return m.Optional()
}

func (m *M) Complement() *M {
	m = m.clone()
	for i := range m.States {
		f := &m.States[i]
		if f.final() {
			f.Label = notFinal
		} else {
			f.Label = defaultFinal
		}
	}
	return m.deleteUnreachable()
}

func (m *M) Exclude(ms ...interface{}) *M {
	m, err := newMerger(m, Or(ms...), difference{}).merge()
	if err != nil {
		panic(err)
	}
	return m.deleteUnreachable()
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
		m, err := conMany(ms)
		if err != nil {
			panic(err)
		}
		return m
	case 2:
		lo, hi := limit[0], limit[1]
		if lo > hi {
			lo, hi = hi, lo
		}
		ms := make([]*M, 0, hi-lo+1)
		for n := lo; n <= hi; n++ {
			ms = append(ms, m.Repeat(n))
		}
		m, err := orMany(ms)
		if err != nil {
			panic(err)
		}
		return m
	}
	panic("repeat should have zero to two arguments")
}

func opMany(op func(_, _ *M) (*M, error), ms []*M) (*M, error) {
	switch len(ms) {
	case 0:
		return nil, nil
	case 1:
		return ms[0].clone(), nil
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
			if m, err := op(ms[i], ms[i+1]); err == nil {
				ms[cur] = m
			} else {
				return nil, err
			}
			cur++
		}
		ms = ms[:cur]
	}
	return ms[0], nil
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

func (m *M) Minimize() *M {
	m, err := m.minimize()
	if err != nil {
		panic(err)
	}
	return m
}
