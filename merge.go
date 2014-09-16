package dfa

import (
	"container/list"
	"fmt"
)

type merger struct {
	m1, m2, m *M
	l         *list.List
	idMap
	mergeMethod
	conMode bool
}
type idMap struct {
	a     []int
	n2    int
	count int
}
type mergeMethod interface {
	mergeLabel(l1, l2 StateLabel) (StateLabel, error)
	eachEdge(s1, s2 *S, visit func(b byte, id1, id2 int))
}

func newMerger(m1, m2 *M, method mergeMethod) *merger {
	n1 := len(m1.States) + 1
	n2 := len(m2.States) + 1
	return &merger{
		m1:          m1,
		m2:          m2,
		m:           &M{},
		l:           list.New(),
		idMap:       idMap{a: make([]int, n1*n2), n2: n2},
		mergeMethod: method}
}

func (m *merger) concatMode() *merger {
	m.conMode = true
	return m
}

func (m *merger) merge() (*M, error) {
	if m.conMode {
		m.add(m.m1.Start, -1)
	} else {
		m.add(m.m1.Start, m.m2.Start)
	}
	for m.l.Len() > 0 {
		id, id1, id2 := m.get()
		var err error
		m.m.States[id], err = m.mergeState(m.m1.state(id1), m.m2.state(id2))
		if err != nil {
			return nil, err
		}
	}
	return m.m, nil
}

func (m *merger) mergeState(s1, s2 *S) (S, error) {
	var a transArray
	m.eachEdge(s1, s2, func(b byte, id1, id2 int) {
		a.set(b, m.add(id1, id2))
	})
	label, err := m.mergeLabel(s1.label(), s2.label())
	return S{label, a.toTransTable()}, err
}

func (m *merger) add(id1, id2 int) int {
	if m.conMode {
		if id1 >= 0 && m.m1.States[id1].final() {
			id2 = m.m2.Start
		}
	}
	id, isNew := m.getID(id1, id2)
	if isNew {
		m.m.States = append(m.m.States, S{})
		m.l.PushFront([3]int{id, id1, id2})
	}
	return id
}
func (m *idMap) getID(id1, id2 int) (id int, isNew bool) {
	id1++
	id2++
	key := id1*m.n2 + id2
	if id = m.a[key] - 1; id >= 0 {
		return id, false
	}
	id = m.count
	m.count++
	m.a[key] = id + 1
	return id, true
}

func (m *merger) get() (id, id1, id2 int) {
	v := m.l.Remove(m.l.Back()).([3]int)
	return v[0], v[1], v[2]
}

type intersection struct{}

func (intersection) mergeLabel(l1, l2 StateLabel) (StateLabel, error) {
	if l1 == l2 {
		return l1, nil
	}
	return notFinal, nil
}

func (intersection) eachEdge(s1, s2 *S, visit func(b byte, id1, id2 int)) {
	it1, it2 := s1.iter(), s2.iter()
	b1, id1 := it1.next()
	b2, id2 := it2.next()
	for id1 >= 0 && id2 >= 0 {
		if b1 == b2 {
			visit(b1, id1, id2)
			b1, id1 = it1.next()
			b2, id2 = it2.next()
		} else if b1 < b2 {
			b1, id1 = it1.next()
		} else {
			b2, id2 = it2.next()
		}
	}
}

type union struct {
	unionEdge
}

func (union) mergeLabel(l1, l2 StateLabel) (StateLabel, error) {
	if l1 > defaultFinal && l2 > defaultFinal && l1 != l2 {
		return -1, fmt.Errorf("conflict Label: %d & %d", l1.toExternal(), l2.toExternal())
	}
	if l1 > l2 {
		// same as:
		// !l2.final() && l1.final() ||
		// l2.final() && l1.labeled()
		return l1, nil
	}
	// same as:
	// l1 == l2 ||
	// !l1.final() && l2.final() ||
	// l1.final() && l2.labeled()
	return l2, nil
}

type unionEdge struct{}

func (unionEdge) eachEdge(s1, s2 *S, visit func(b byte, id1, id2 int)) {
	it1, it2 := s1.iter(), s2.iter()
	b1, next1 := it1.next()
	b2, next2 := it2.next()
	for {
		b := b1
		id1, id2 := next1, next2
		if id1 >= 0 && id2 >= 0 {
			if b1 == b2 {
				b1, next1 = it1.next()
				b2, next2 = it2.next()
			} else if b1 < b2 {
				id2 = invalidID
				b1, next1 = it1.next()
			} else {
				b = b2
				id1 = invalidID
				b2, next2 = it2.next()
			}
		} else if id1 >= 0 {
			b = b1
			b1, next1 = it1.next()
		} else if id2 >= 0 {
			b = b2
			b2, next2 = it2.next()
		} else {
			break
		}

		visit(b, id1, id2)
	}
}

type difference struct{}

func (difference) mergeLabel(l1, l2 StateLabel) (StateLabel, error) {
	if l2.final() {
		return notFinal, nil
	}
	return l1, nil
}

func (difference) eachEdge(s1, s2 *S, visit func(b byte, id1, id2 int)) {
	it1, it2 := s1.iter(), s2.iter()
	b1, next1 := it1.next()
	b2, next2 := it2.next()
	for {
		b := b1
		id1, id2 := next1, next2
		if id1 >= 0 && id2 >= 0 {
			if b1 == b2 {
				b1, next1 = it1.next()
				b2, next2 = it2.next()
			} else if b1 < b2 {
				id2 = invalidID
				b1, next1 = it1.next()
			} else {
				b = b2
				id1 = invalidID
				b2, next2 = it2.next()
			}
		} else if id1 >= 0 {
			b = b1
			b1, next1 = it1.next()
		} else {
			break
		}

		visit(b, id1, id2)
	}
}

type concat struct {
	acceptFirst bool
	unionEdge
}

func (c concat) mergeLabel(l1, l2 StateLabel) (StateLabel, error) {
	if l2.final() {
		return l2, nil
	}
	if c.acceptFirst && l1.final() {
		return l1, nil
	}
	return l2, nil
}
