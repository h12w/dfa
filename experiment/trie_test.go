package dfa

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	a := newTransArray()
	trie := a.toTransTrie()
	fmt.Println(trie)
	a.set(0x20, 2)
	a.set(0x30, 3)
	a.set(0x10, 1)
	trie = a.toTransTrie()
	fmt.Println(trie.dump())
}
