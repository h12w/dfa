package dfa

func (m *M) deleteUnreachable() *M {
	reachable := m.reachable()
	for i := range reachable {
		if !reachable[i] {
			m.States.each(func(s *S) {
				a := s.Table.toTransArray()
				for b := range a {
					if a.get(byte(b)) == i {
						a.clear(byte(b))
					}
				}
				s.Table = a.toTransTable()
			})
		}
	}
	return m
}
func (m *M) reachable() []bool {
	reachFinal := make([]bool, m.States.count())
	for i := range m.States {
		if m.States[i].final() {
			reachFinal[i] = true
		}
	}
	more := true
	for more {
		more = false
		for i := range reachFinal {
			if !reachFinal[i] {
				for j := range m.States[i].Table {
					next := m.States[i].Table[j].Next
					if next >= 0 && reachFinal[next] {
						reachFinal[i] = true
						more = true
						break
					}
				}
			}
		}
	}
	return reachFinal
}

func (m *M) minimize() (*M, error) {
	if m == nil {
		return nil, nil
	}
	n := m.States.count()
	diff := newDiff(n)
	diff.eachFalse(func(i, j int) {
		s, t := m.States[i], m.States[j]
		if s.Label != t.Label || !s.Table.positionEqual(&t.Table) {
			diff.set(i, j)
		}
	})
	for diff.hasNewDiff {
		diff.hasNewDiff = false
		diff.eachFalse(func(i, j int) {
			s, t := m.States[i].Table, m.States[j].Table
			for k := range s {
				sNext, tNext := s[k].Next, t[k].Next
				if sNext != tNext && diff.get(sNext, tNext) {
					diff.set(i, j)
					break
				}
			}
		})
	}
	idm := make(map[int]int)
	diff.eachFalse(func(i, j int) {
		idm[j] = i
	})
	if len(idm) > 0 {
		m.each(func(s *S) {
			s.each(func(t *Trans) {
				if small, ok := idm[t.Next]; ok {
					t.Next = small
				}
			})
		})
	}
	return m.and(m) // m.and(m) is also a way to remove unreachable nodes
}

func (t *TransTable) positionEqual(o *TransTable) bool {
	if len(*t) != len(*o) {
		return false
	}
	for i := range *t {
		tt, oo := (*t)[i], (*o)[i]
		if tt.Lo != oo.Lo || tt.Hi != oo.Hi {
			return false
		}
	}
	return true
}

// 0: 1, 2, ..., n-1
// 1:    2, ..., n-1
// ...
// n-2:          n-1
type boolPairs struct {
	n          int
	a          []bool
	hasNewDiff bool
}

func newDiff(n int) *boolPairs {
	return &boolPairs{n, make([]bool, n*(n-1)/2), false}
}

func (d *boolPairs) set(i, j int) {
	d.hasNewDiff = true
	d.a[d.index(i, j)] = true
}

func (d *boolPairs) get(i, j int) bool {
	return d.a[d.index(i, j)]
}

func (d *boolPairs) index(i, j int) int {
	if i == j {
		panic("i should never be equal to j")
	} else if i > j {
		i, j = j, i
	}
	return (2*d.n-i-1)*i/2 + (j - i - 1)
}

func (d *boolPairs) eachFalse(visit func(int, int)) {
	for i := d.n - 2; i >= 0; i-- { // reverse order so the smaller comes later
		for j := i + 1; j <= d.n-1; j++ {
			if !d.get(i, j) {
				visit(i, j)
			}
		}
	}
}
