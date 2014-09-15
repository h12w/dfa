package dfa

const (
	leadx = 0x80
	maskx = 0x3f
)

var (
	runeRanges = []rune{
		0x7f,
		0x7ff,
		0xffff,
		0x10ffff,
	}
	mask = []rune{
		0x3f,
		0xfff,
		0x3ffff,
	}
	lead = []byte{
		0x00,
		0xc0,
		0xe0,
		0xf0,
	}
)

// u8s represents the union of multiple u8's and is used to construct an UTF8
// DFA by collecting u8's.
type u8s struct {
	a []u8
}

func (u *u8s) between(lo, hi rune) {
	for _, l := range runeRanges {
		if lo <= l {
			if hi > l {
				u.betweenRune(lo, l)
				lo = l + 1
			} else {
				u.betweenRune(lo, hi)
				break
			}
		}
	}
}

func (u *u8s) betweenRune(lo, hi rune) {
	for _, m := range mask {
		if lo&^m != hi&^m {
			if lo&m != 0 {
				u.add(lo, lo|m)
				lo = (lo | m) + 1
			}
			if hi&m != m {
				u.add(hi&^m, hi)
				hi = (hi &^ m) - 1
			}
		} else {
			break
		}
	}
	if lo <= hi {
		u.add(lo, hi)
	}
}

func (s *u8s) add(lo, hi rune) {
	s.a = append(s.a, u8{lo, hi})
}

func (s *u8s) m() (*M, error) {
	ms := make([]*M, len(s.a))
	for i := range s.a {
		ms[i] = s.a[i].m()
	}
	return orMany(ms)
}

type u8 struct {
	lo, hi rune
}

func (u *u8) m() *M {
	lo, hi := u.lo, u.hi
	n := calcN(hi)
	m := &M{0, make(States, n+1)}
	m.States[n] = finalState()
	for i := n - 1; i >= 1; i-- {
		l, h := byte(lo&maskx), byte(hi&maskx)
		m.States[i] = stateBetween(l|leadx, h|leadx, i+1)
		lo >>= 6
		hi >>= 6
	}
	m.States[0] = stateBetween(byte(lo)|lead[n-1], byte(hi)|lead[n-1], 1)
	return m
}
func calcN(hi rune) int {
	n := 0
	for i, l := range runeRanges {
		if hi <= l {
			n = i + 1
			break
		}
	}
	return n
}
