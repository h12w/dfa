package dfa

import "fmt"

type state struct {
	table transTable
	label stateLabel
}
type transTable struct {
	a []trans
}
type trans struct {
	s, e byte
	next int
}
type transArray [256]int

func stateBetween(lo, hi byte, next int) state {
	a := newTransArray()
	a.setBetween(lo, hi, next)
	return state{table: a.toTransTable()}
}

func stateTo(b byte, next int) state {
	a := newTransArray()
	a.set(b, next)
	return state{table: a.toTransTable()}
}

func finalState() state {
	return state{label: defaultFinal}
}

func (s *state) final() bool {
	return s.label.final()
}

func (s *state) connect(o *state) {
	a := s.table.toTransArray()
	o.each(func(t *trans) {
		a.setBetween(t.s, t.e, t.next)
	})
	s.table = a.toTransTable()
}

func (s *state) filterConnect(o *state, filters []func(byte) bool) {
	a := s.table.toTransArray()
	o.each(func(t *trans) {
		b := t.s
		for {
			connect := true
			for _, filter := range filters {
				if !filter(b) {
					connect = false
					break
				}
			}
			if connect {
				a.set(b, t.next)
			}
			if b == t.e {
				break
			}
			b++
		}
	})
	s.table = a.toTransTable()
}

func (s *state) clone() state {
	return state{s.table.clone(), s.label}
}

func (s *state) iter() tableIter {
	if s == nil {
		return tableIter{
			n: invalidID,
		}
	}
	return s.table.iter()
}

type tableIter struct {
	t *transTable
	i int
	b byte
	n int
}

func (it *tableIter) next() (rb byte, rnext int) {
	if it.n == invalidID {
		return 0, invalidID
	}
	rb, rnext = it.b, it.n
	if it.b == it.t.a[it.i].e {
		it.i++
		if it.i < len(it.t.a) {
			it.b = it.t.a[it.i].s
			it.n = it.t.a[it.i].next
		} else {
			it.n = invalidID
		}
	} else {
		it.b++
	}
	return
}

func (t *transTable) iter() tableIter {
	if len(t.a) == 0 {
		return tableIter{
			n: invalidID,
		}
	}
	return tableIter{
		t: t,
		b: t.a[0].s,
		n: t.a[0].next,
	}
}

func (s *state) each(visit func(*trans)) {
	s.table.each(visit)
}

func (s *state) seqNext(b byte) (sid int) {
	for i := range s.table.a {
		if s.table.a[i].s <= b && b <= s.table.a[i].e {
			return s.table.a[i].next
		}
	}
	return invalidID
}

func (s *state) next(b byte) (sid int) {
	min, max := 0, len(s.table.a)-1
	for min <= max {
		mid := (min + max) / 2
		if b < s.table.a[mid].s {
			max = mid - 1
		} else if s.table.a[mid].s <= b && b <= s.table.a[mid].e {
			return s.table.a[mid].next
		} else {
			min = mid + 1
		}
	}
	return invalidID
}

func (table *transTable) each(visit func(*trans)) {
	for i := range table.a {
		visit(&table.a[i])
	}
}

func (table *transTable) clone() transTable {
	return transTable{append([]trans(nil), table.a...)}
}

func (table *transTable) toTransArray() transArray {
	a := newTransArray()
	table.each(func(t *trans) {
		for i := int(t.s); i <= int(t.e); i++ {
			a[i] = t.next
		}
	})
	return a
}

func newTransArray() (a transArray) {
	for i := range a {
		a[i] = invalidID
	}
	return a
}

func (a *transArray) set(b byte, next int) *transArray {
	if a[b] < 0 {
		a[b] = next
		return a
	}
	if a[b] != next {
		panic(fmt.Errorf("byte %s has already been assigned a different state %d", quote(b), next))
	}
	return a
}

func (a *transArray) setBetween(lo, hi byte, next int) *transArray {
	b := lo
	for {
		a.set(b, next)
		if b == hi {
			break
		} else {
			b++
		}
	}
	return a
}

func (a *transArray) toTransTable() (table transTable) {
	i := 0
	for ; i < len(a); i++ {
		if a[i] >= 0 {
			break
		}
	}
	if i == 256 {
		return
	}
	table.a = append(table.a, trans{byte(i), byte(i), a[i]})
	i++
	for ; i < len(a); i++ {
		if a[i] >= 0 {
			b := byte(i)
			last := table.a[len(table.a)-1]
			if b == last.e+1 && a[i] == last.next {
				table.a[len(table.a)-1].e = b
			} else {
				table.a = append(table.a, trans{b, b, a[i]})
			}
		}
	}
	return
}

func (a *transArray) toState() state {
	return state{table: a.toTransTable()}
}
