package dfa

import (
	"testing"

	"h12.me/gspec"
)

func TestMatch(t *testing.T) {
	expect := gspec.Expect(t.FailNow)
	for _, testcase := range []struct {
		m     *M
		input string
		token string
		label int
	}{
		{threeToken(), "0x12A ", "0x12A", hexLabel},
		{threeToken(), "123 ", "123", decimalLabel},
		{threeToken(), "abc", "abc", identLabel},
		{s("ab").AtLeast(1).As(9), "aba", "ab", 9},
	} {
		size, label, ok := testcase.m.Match([]byte(testcase.input))
		expect("matched", ok).Equal(true)
		expect("matched label", label).Equal(testcase.label)
		expect("token", string(testcase.input[:size])).Equal(testcase.token)
	}
}
