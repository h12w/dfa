package dfa

import (
	"container/list"
	"fmt"
)

type merger struct {
	m1, m2, m *M
	l         *list.List
	idm       []int
	idmCount  int
	mergeMethod
}
type mergeMethod interface {
	mergeLabel(s1, s2 *S) (StateLabel, error)
	eachEdge(s1, s2 *S, visit func(b byte, id1, id2 int))
}

func newMerger(m1, m2 *M, method mergeMethod) *merger {
	n1, n2 := 0, 0
	if m1 != nil {
		n1 = len(m1.States) + 1
	}
	if m2 != nil {
		n2 = len(m2.States) + 1
	}
	return &merger{
		m1:          m1,
		m2:          m2,
		m:           &M{},
		l:           list.New(),
		idm:         make([]int, n1*n2),
		mergeMethod: method}
}

func (q *merger) merge() (m *M, err error) {
	if q.m1 == nil {
		return q.m2.clone(), nil
	} else if q.m2 == nil {
		return q.m1.clone(), nil
	}
	q.getID(q.m1.Start, q.m2.Start)
	for q.l.Len() > 0 {
		id, id1, id2 := q.get()
		q.m.States[id], err = q.mergeState(q.m1.S(id1), q.m2.S(id2))
		if err != nil {
			return nil, err
		}
	}
	return q.m, nil
}

func (q *merger) mergeState(s1, s2 *S) (S, error) {
	var a transArray
	q.eachEdge(s1, s2, func(b byte, id1, id2 int) {
		a.set(b, q.getID(id1, id2))
	})
	label, err := q.mergeLabel(s1, s2)
	return S{label, a.toTransTable()}, err
}

func (s *S) trivialFinal() bool {
	return s.Label == defaultFinal && len(s.Table) == 0
}

func (q *merger) getID(id1, id2 int) int {
	key := (id1+1)*(len(q.m2.States)+1) + (id2 + 1)
	if id := q.idm[key] - 1; id >= 0 {
		return id
	}
	id := q.idmCount
	q.idmCount++
	q.idm[key] = id + 1
	q.m.States = append(q.m.States, S{})
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

func (intersection) mergeLabel(s1, s2 *S) (StateLabel, error) {
	l1, l2 := s1.Label, s2.Label
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

type union struct{}

func (union) mergeLabel(s1, s2 *S) (StateLabel, error) {
	if s1 == nil {
		return s2.Label, nil
	}
	if s2 == nil {
		return s1.Label, nil
	}
	f1, f2 := s1.Label, s2.Label
	if f1 > defaultFinal && f2 > defaultFinal && f1 != f2 {
		return -1, fmt.Errorf("conflict Label: %d & %d", f1.toExternal(), f2.toExternal())
	}
	return finalMax(f1, f2), nil // TODO reove finalMax
}
func finalMax(a, b StateLabel) StateLabel {
	if a > b {
		return a
	}
	return b
}

func (union) eachEdge(s1, s2 *S, visit func(b byte, id1, id2 int)) {
	it1, it2 := s1.iter(), s2.iter()
	b1, next1 := it1.next()
	b2, next2 := it2.next()
	for {
		b := b1
		id1, id2 := next1, next2
		if id1 < 0 && id2 < 0 {
			break
		} else if id1 < 0 {
			b = b2
			b2, next2 = it2.next()
		} else if id2 < 0 {
			b = b1
			b1, next1 = it1.next()
		} else {
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
		}
		visit(b, id1, id2)
	}
}

type difference struct{}

func (difference) mergeLabel(s1, s2 *S) (StateLabel, error) {
	if s1 == nil {
		return notFinal, nil
	}
	if s2 == nil {
		return s1.Label, nil
	}
	f1, f2 := s1.Label, s2.Label
	if f2.final() {
		return notFinal, nil
	}
	return f1, nil
}

func (difference) eachEdge(s1, s2 *S, visit func(b byte, id1, id2 int)) {
	it1, it2 := s1.iter(), s2.iter()
	b1, next1 := it1.next()
	b2, next2 := it2.next()
	for {
		b := b1
		id1, id2 := next1, next2
		if id1 < 0 && id2 < 0 {
			break
		} else if id1 < 0 {
			b = b2
			b2, next2 = it2.next()
		} else if id2 < 0 {
			b = b1
			b1, next1 = it1.next()
		} else {
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
		}
		visit(b, id1, id2)
	}
}
