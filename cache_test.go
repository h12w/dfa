package dfa

import (
	"fmt"
	"os"
	"testing"
)

func SkipTestGenCache(t *testing.T) {
	w, _ := os.Create("cache.go")
	defer w.Close()

	fmt.Fprintln(w, "package dfa")
	fmt.Fprint(w, "var classL = ")
	L, _ := charClass("L")
	L.WriteGo(w, "dfa")
}
