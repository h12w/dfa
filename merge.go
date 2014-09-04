package dfa

import (
	"container/list"
	"fmt"
)

type merger struct {
	m1, m2, m *M
	l         *list.List
	idm       map[[2]int]int
	mergeMethod
}
type mergeMethod interface {
	mergeLabel(s1, s2 *state) stateLabel
	eachEdge(s1, s2 *state, visit func(b byte, id1, id2 int))
}

func newMerger(m1, m2 *M, method mergeMethod) *merger {
	return &merger{
		m1:          m1,
		m2:          m2,
		m:           &M{},
		l:           list.New(),
		idm:         make(map[[2]int]int),
		mergeMethod: method}
}

func (q *merger) merge() *M {
	if q.m1 == nil {
		return q.m2.clone()
	} else if q.m2 == nil {
		return q.m1.clone()
	}
	q.getID(q.m1.start, q.m2.start)
	for q.l.Len() > 0 {
		id, id1, id2 := q.get()
		q.m.states[id] = q.mergeState(q.m1.state(id1), q.m2.state(id2))
	}
	return q.m
}

func (q *merger) mergeState(s1, s2 *state) state {
	a := newTransArray()
	q.eachEdge(s1, s2, func(b byte, id1, id2 int) {
		a[b] = q.getID(id1, id2)
	})
	return state{a.toTransTable(), q.mergeLabel(s1, s2)}
}

// getKey merges trivial final states into a unique merged state(-2,-2).
func (q *merger) getKey(id1, id2 int) [2]int {
	const trivialFinalID = -2
	if id1 >= 0 && q.m1.states[id1].trivialFinal() {
		id1 = trivialFinalID
		if id2 < 0 {
			id2 = trivialFinalID
		}
	}
	if id2 >= 0 && q.m2.states[id2].trivialFinal() {
		id2 = trivialFinalID
		if id1 < 0 {
			id1 = trivialFinalID
		}
	}
	return [2]int{id1, id2}
}
func (s *state) trivialFinal() bool {
	return s.label == defaultFinal && len(s.table) == 0
}

func (q *merger) getID(id1, id2 int) int {
	key := q.getKey(id1, id2)
	if id, ok := q.idm[key]; ok {
		return id
	}
	id := len(q.m.states)
	q.idm[key] = id
	q.m.states = append(q.m.states, state{})
	q.put(id, id1, id2)
	return id
}

func (q *merger) put(id, id1, id2 int) {
	q.l.PushFront([3]int{id, id1, id2})
}

func (q *merger) get() (id, id1, id2 int) {
	v := q.l.Remove(q.l.Back()).([3]int)
	return v[0], v[1], v[2]
}

type intersection struct{}

func (intersection) mergeLabel(s1, s2 *state) stateLabel {
	l1, l2 := s1.label, s2.label
	if l1 == l2 {
		return l1
	}
	return notFinal
}

func (intersection) eachEdge(s1, s2 *state, visit func(b byte, id1, id2 int)) {
	it1, it2 := s1.iter(), s2.iter()
	b1, id1 := it1()
	b2, id2 := it2()
	for id1 >= 0 && id2 >= 0 {
		if b1 == b2 {
			visit(b1, id1, id2)
			b1, id1 = it1()
			b2, id2 = it2()
		} else if b1 < b2 {
			b1, id1 = it1()
		} else {
			b2, id2 = it2()
		}
	}
}

type union struct{}

func (union) mergeLabel(s1, s2 *state) stateLabel {
	if s1 == nil {
		return s2.label
	}
	if s2 == nil {
		return s1.label
	}
	f1, f2 := s1.label, s2.label
	if f1 > defaultFinal && f2 > defaultFinal && f1 != f2 {
		panic(fmt.Errorf("conflict label: %d & %d", f1.toExternal(), f2.toExternal()))
	}
	return finalMax(f1, f2)
}
func finalMax(a, b stateLabel) stateLabel {
	if a > b {
		return a
	}
	return b
}

func (union) eachEdge(s1, s2 *state, visit func(b byte, id1, id2 int)) {
	it1, it2 := s1.iter(), s2.iter()
	b1, next1 := it1()
	b2, next2 := it2()
	for {
		b := b1
		id1, id2 := next1, next2
		if id1 < 0 && id2 < 0 {
			break
		} else if id1 < 0 {
			b = b2
			b2, next2 = it2()
		} else if id2 < 0 {
			b = b1
			b1, next1 = it1()
		} else {
			if b1 == b2 {
				b1, next1 = it1()
				b2, next2 = it2()
			} else if b1 < b2 {
				id2 = invalidID
				b1, next1 = it1()
			} else {
				b = b2
				id1 = invalidID
				b2, next2 = it2()
			}
		}
		visit(b, id1, id2)
	}
}

type difference struct{}

func (difference) mergeLabel(s1, s2 *state) stateLabel {
	if s1 == nil {
		return notFinal
	}
	if s2 == nil {
		return s1.label
	}
	f1, f2 := s1.label, s2.label
	if f2.final() {
		return notFinal
	}
	return f1
}

func (difference) eachEdge(s1, s2 *state, visit func(b byte, id1, id2 int)) {
	it1, it2 := s1.iter(), s2.iter()
	b1, next1 := it1()
	b2, next2 := it2()
	for {
		b := b1
		id1, id2 := next1, next2
		if id1 < 0 && id2 < 0 {
			break
		} else if id1 < 0 {
			b = b2
			b2, next2 = it2()
		} else if id2 < 0 {
			b = b1
			b1, next1 = it1()
		} else {
			if b1 == b2 {
				b1, next1 = it1()
				b2, next2 = it2()
			} else if b1 < b2 {
				id2 = invalidID
				b1, next1 = it1()
			} else {
				b = b2
				id1 = invalidID
				b2, next2 = it2()
			}
		}
		visit(b, id1, id2)
	}
}
