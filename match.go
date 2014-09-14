package dfa

// Match greedily matches the DFA against src.
func (m *M) Match(src []byte) (size, Label int, matched bool) {
	var (
		s, matchedState *S
		sid             = 0
		pos, matchedPos int
	)
	for sid >= 0 {
		s = &m.States[sid]
		if s.Label >= defaultFinal {
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
		return matchedPos, matchedState.Label.toExternal(), true
	}
	return 0, -1, false
}

// Match greedily matches the DFA against src.
func (m *FastM) Match(src []byte, p int) (size, Label int, matched bool) {
	cur := &m.States[0]
	pos := p
	matchedPos := pos
	for {
		if cur.Label >= 0 {
			matchedPos = pos
			Label = cur.Label
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
