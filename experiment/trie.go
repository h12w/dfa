package dfa

import (
	"bytes"
	"fmt"
	"io"
)

type transTrie struct {
	a []trieNode
}
type trieNode struct {
	next  nextCode
	mask  byte
	value byte
}
type nextCode int

func (t *transTable) toTransTrie() *transTrie {
	a := t.toTransArray()
	return a.toTransTrie()
}
func (a *transArray) toTransTrie() *transTrie {
	var trie transTrie
	trie.add((*a)[:], 0, 255, 0x80)
	return &trie
}

func (t *transTrie) add(a []int, lo, hi int, mask byte) (int, bool) {
	if hi-lo > 0 {
		mid := lo + (hi-lo+1)/2 - 1
		zero, zeroAdded := t.add(a, lo, mid, mask>>1)
		one, oneAdded := t.add(a, mid+1, hi, mask>>1)
		if zeroAdded && oneAdded {
			t.a = append(t.a, trieNode{encodeNext(zero, one), mask, 0})
			return len(t.a) - 1, true
		} else if zeroAdded && !oneAdded {
			return zero, true
		} else if oneAdded && !zeroAdded {
			return one, true
		}
	} else if hi == lo && a[lo] >= 0 {
		t.a = append(t.a, trieNode{nextCode(a[lo]), 0xFF, byte(lo)})
		return len(t.a) - 1, true
	}
	return -1, false
}

func encodeNext(zero, one int) nextCode {
	return nextCode((zero & 0xFFFF) | ((one & 0xFFFF) << 16))
}

func (c nextCode) decodeNext() (zero, one int) {
	return int(c & 0xFFFF), int((c >> 16) & 0xFFFF)
}

func (n *trieNode) isLeaf() bool {
	return n.mask == 0xFF
}

func (t *transTrie) writeGo(w io.Writer, sid int) {
	if len(t.a) == 0 {
		fmt.Fprintln(w, "goto sEnd")
		return
	}
	fmt.Fprintf(w, "goto s%d_t%d\n", sid, len(t.a)-1)
	for i, n := range t.a {
		fmt.Fprintf(w, "s%d_t%d:\n", sid, i)
		if n.isLeaf() {
			fmt.Fprintf(w, "if src[p] == 0x%.2x {\n", n.value)
			fmt.Fprintf(w, "goto s%d", n.next)
			fmt.Fprintln(w, "}")
			fmt.Fprintln(w, "goto sEnd")
		} else {
			zero, one := n.next.decodeNext()
			fmt.Fprintf(w, "if src[p] & 0x%.2x != 0 {\n", n.mask)
			fmt.Fprintf(w, "goto s%d_t%d", sid, one)
			fmt.Fprintln(w, "} else {")
			fmt.Fprintf(w, "goto s%d_t%d", sid, zero)
			fmt.Fprintln(w, "}")
		}
	}
}

func (t *transTrie) dump() string {
	var w bytes.Buffer
	for i, n := range t.a {
		fmt.Fprintf(&w, "n%d\n", i)
		w.WriteString(n.dump())
	}
	return w.String()
}

func (n *trieNode) dump() string {
	var w bytes.Buffer
	if n.isLeaf() {
		fmt.Fprintf(&w, "\t%-7ss%d\n", quote(n.value), n.next)
	} else {
		zero, one := n.next.decodeNext()
		fmt.Fprintf(&w, "\t&%.2x=0  n%d\n", n.mask, zero)
		fmt.Fprintf(&w, "\t&%.2x=1  n%d\n", n.mask, one)
	}
	return w.String()
}
