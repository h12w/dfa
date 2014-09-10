package dfa

// Match greedily matches the DFA against src.
func (m *M) Match(src []byte) (size, label int, matched bool) {
	var (
		s, matchedState *state
		sid             = 0
		pos, matchedPos int
	)
	for sid >= 0 {
		s = &m.states[sid]
		if s.label >= defaultFinal {
			matchedState = s
			matchedPos = pos
		}
		if pos < len(src) {
			sid = s.next(src[pos])
			if sid >= 0 {
				pos++
			}
		} else {
			break
		}
	}
	if matchedState != nil {
		return matchedPos, matchedState.label.toExternal(), true
	}
	return 0, -1, false
}

// Match greedily matches the DFA against src.
func (m *FastM) Match(src []byte, p int) (size, label int, matched bool) {
	cur := &m.States[0]
	pos := p
	matchedPos := pos
	for {
		if cur.Label >= 0 {
			matchedPos = pos
			label = cur.Label
			matched = true
		}
		if pos == len(src) {
			break
		}
		if cur = cur.Trans[src[pos]]; cur == nil {
			break
		}
		pos++
	}
	size = matchedPos - p
	return
}
