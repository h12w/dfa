package dfa

import "fmt"

type S struct {
	Label StateLabel
	Table TransTable
}
type TransTable []Trans
type Trans struct {
	Lo, Hi byte
	Next   int
}
type transArray [256]int

func stateBetween(lo, hi byte, next int) S {
	a := newTransArray()
	a.setBetween(lo, hi, next)
	return S{Table: a.toTransTable()}
}

func stateTo(b byte, next int) S {
	a := newTransArray()
	a.set(b, next)
	return S{Table: a.toTransTable()}
}

func finalState() S {
	return S{Label: defaultFinal}
}

func (s *S) final() bool {
	return s.Label.final()
}

func (s *S) connect(o *S) {
	a := s.Table.toTransArray()
	o.each(func(t *Trans) {
		a.setBetween(t.Lo, t.Hi, t.Next)
	})
	s.Table = a.toTransTable()
}

func (s *S) filterConnect(o *S, filters []func(byte) bool) {
	a := s.Table.toTransArray()
	o.each(func(t *Trans) {
		b := t.Lo
		for {
			connect := true
			for _, filter := range filters {
				if !filter(b) {
					connect = false
					break
				}
			}
			if connect {
				a.set(b, t.Next)
			}
			if b == t.Hi {
				break
			}
			b++
		}
	})
	s.Table = a.toTransTable()
}

func (s *S) clone() S {
	return S{s.Label, s.Table.clone()}
}

func (s *S) iter() TableIter {
	if s == nil {
		return TableIter{
			n: invalidID,
		}
	}
	return s.Table.iter()
}

type TableIter struct {
	t *TransTable
	i int
	b byte
	n int
}

func (it *TableIter) next() (rb byte, rnext int) {
	if it.n == invalidID {
		return 0, invalidID
	}
	rb, rnext = it.b, it.n
	if it.b == (*it.t)[it.i].Hi {
		it.i++
		if it.i < len(*it.t) {
			it.b = (*it.t)[it.i].Lo
			it.n = (*it.t)[it.i].Next
		} else {
			it.n = invalidID
		}
	} else {
		it.b++
	}
	return
}

func (t *TransTable) iter() TableIter {
	if len(*t) == 0 {
		return TableIter{
			n: invalidID,
		}
	}
	return TableIter{
		t: t,
		b: (*t)[0].Lo,
		n: (*t)[0].Next,
	}
}

func (s *S) each(visit func(*Trans)) {
	s.Table.each(visit)
}

func (s *S) seqNext(b byte) (sid int) {
	for i := range s.Table {
		if s.Table[i].Lo <= b && b <= s.Table[i].Hi {
			return s.Table[i].Next
		}
	}
	return invalidID
}

func (s *S) next(b byte) (sid int) {
	min, max := 0, len(s.Table)-1
	for min <= max {
		mid := (min + max) / 2
		if b < s.Table[mid].Lo {
			max = mid - 1
		} else if s.Table[mid].Lo <= b && b <= s.Table[mid].Hi {
			return s.Table[mid].Next
		} else {
			min = mid + 1
		}
	}
	return invalidID
}

func (Table *TransTable) each(visit func(*Trans)) {
	for i := range *Table {
		visit(&(*Table)[i])
	}
}

func (Table *TransTable) clone() TransTable {
	return append(TransTable(nil), *Table...)
}

func (Table *TransTable) toTransArray() transArray {
	a := newTransArray()
	Table.each(func(t *Trans) {
		for i := int(t.Lo); i <= int(t.Hi); i++ {
			a[i] = t.Next
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
		panic(fmt.Errorf("byte %s has already been assigned a different S %d", quote(b), next))
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

func (a *transArray) toTransTable() (Table TransTable) {
	i := 0
	for ; i < len(a); i++ {
		if a[i] >= 0 {
			break
		}
	}
	if i == 256 {
		return
	}
	Table = append(Table, Trans{byte(i), byte(i), a[i]})
	i++
	for ; i < len(a); i++ {
		if a[i] >= 0 {
			b := byte(i)
			last := Table[len(Table)-1]
			if b == last.Hi+1 && a[i] == last.Next {
				Table[len(Table)-1].Hi = b
			} else {
				Table = append(Table, Trans{b, b, a[i]})
			}
		}
	}
	return
}

func (a *transArray) toState() S {
	return S{Table: a.toTransTable()}
}
