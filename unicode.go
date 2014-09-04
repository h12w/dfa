package dfa

import (
	"fmt"
	"unicode"
)

func CharClass(name string) (m *M) {
	eachRangeInCharClass(name, func(lo, hi rune) {
		m = m.or(Between(lo, hi))
	})
	return m.Minimize()
}

func eachRangeInCharClass(name string, visit func(lo, hi rune)) {
	table := unicode.Categories[name]
	if table == nil {
		table = unicode.Scripts[name]
	}
	if table == nil {
		panic(fmt.Errorf("character class %s not exists", name))
	}

	for _, xr := range table.R16 {
		eachRangeWithStride(rune(xr.Lo), rune(xr.Hi), rune(xr.Stride), visit)
	}
	for _, xr := range table.R32 {
		eachRangeWithStride(rune(xr.Lo), rune(xr.Hi), rune(xr.Stride), visit)
	}
}
func eachRangeWithStride(lo, hi, stride rune, visit func(lo, hi rune)) {
	if stride == 1 {
		visit(lo, hi)
	} else {
		for c := lo; c <= hi; c += stride {
			visit(c, c)
		}
	}
}
