package dfa

func (m *M) eachUnreachable(visit func(int)) {
	reachFinal := make([]bool, m.states.count())
	for i := range m.states {
		if m.states[i].final() {
			reachFinal[i] = true
		}
	}
	more := true
	for more {
		more = false
		for i := range reachFinal {
			if !reachFinal[i] {
				for j := range m.states[i].table {
					next := m.states[i].table[j].next
					if next >= 0 && reachFinal[next] {
						reachFinal[i] = true
						more = true
						break
					}
				}
			}
		}
	}
	for i, r := range reachFinal {
		if !r {
			visit(i)
		}
	}
}

func (m *M) deleteUnreachable() *M {
	m.eachUnreachable(func(i int) {
		m.states.each(func(s *state) {
			a := s.table.toTransArray()
			for b := range a {
				if a[b] == i {
					a[b] = invalidID //excludingID // TODO change this hack
				}
			}
			s.table = a.toTransTable()
		})
	})
	return m
}

func (m *M) Minimize() *M {
	n := m.states.count()
	diff := newDiff(n)
	diff.eachFalse(func(i, j int) {
		s, t := m.states[i], m.states[j]
		if s.label != t.label || s.table.toTransSet() != t.table.toTransSet() {
			diff.set(i, j)
		}
	})
	for diff.hasNewDiff {
		diff.hasNewDiff = false
		diff.eachFalse(func(i, j int) {
			s, t := m.states[i], m.states[j]
			si, ti := s.iter(), t.iter()
			_, sid := si()
			_, tid := ti()
			for sid != -1 && tid != -1 {
				if sid != tid && diff.get(sid, tid) {
					diff.set(i, j)
					break
				}
				_, sid = si()
				_, tid = ti()
			}
		})
	}
	idm := make(map[int]int)
	diff.eachFalse(func(i, j int) {
		idm[j] = i
	})
	if len(idm) > 0 {
		m.each(func(s *state) {
			s.each(func(t *trans) {
				if small, ok := idm[t.next]; ok {
					t.next = small
				}
			})
		})
	}
	return or2(m, m) // or2(m, m) is also a way to remove unreachable nodes
}

type transSet [256]bool

func (t *transTable) toTransSet() (s transSet) {
	t.each(func(t *trans) {
		for b := t.s; b <= t.e; b++ {
			s[b] = true
		}
	})
	return
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
		panic("i should not be equal to j")
	}
	if i > j {
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
