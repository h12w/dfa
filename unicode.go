package dfa

import (
	"fmt"
	"unicode"
)

func charClass(name string) (m *M, err error) {
	switch name {
	case "L":
		return classL, nil
	}
	ms := []*M{}
	err = eachRangeInCharClass(name, func(lo, hi rune) {
		ms = append(ms, Between(lo, hi))
	})
	if err != nil {
		return nil, err
	}
	m, err = orMany(ms)
	if err != nil {
		return nil, err
	}
	return m.minimize()
}

func eachRangeInCharClass(name string, visit func(lo, hi rune)) error {
	table := unicode.Categories[name]
	if table == nil {
		table = unicode.Scripts[name]
		if table == nil {
			return fmt.Errorf("character class %s not exists", name)
		}
	}

	for _, xr := range table.R16 {
		eachRangeWithStride(rune(xr.Lo), rune(xr.Hi), rune(xr.Stride), visit)
	}
	for _, xr := range table.R32 {
		eachRangeWithStride(rune(xr.Lo), rune(xr.Hi), rune(xr.Stride), visit)
	}
	return nil
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
