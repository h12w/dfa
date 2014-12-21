package dfa

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type GraphOption struct {
	FontName  string
	Timelabel bool
}

func (s *S) dump(ss []S, sid int) string {
	var w bytes.Buffer
	w.WriteByte('s')
	w.WriteString(strconv.Itoa(sid))
	if s.final() {
		w.WriteByte('$')
		if s.Label > defaultFinal {
			w.WriteString(strconv.Itoa(int(s.Label)))
		}
	}
	for _, Trans := range s.Table {
		w.WriteByte('\n')
		w.WriteString("\t")
		w.WriteString(Trans.dump())
	}
	return w.String()
}

func (t *Trans) dump() string {
	var w bytes.Buffer
	fmt.Fprintf(&w, "%-7s", t.rangeString())
	w.WriteString(" s")
	w.WriteString(strconv.Itoa(t.Next))
	return w.String()
}
func (t *Trans) rangeString() string {
	if t.Lo == t.Hi {
		return quote(t.Lo)
	}
	return quote(t.Lo) + "-" + quote(t.Hi)
}
func quote(b byte) string {
	switch b {
	case '\'', '"', '`':
		return string(rune(b))
	}
	if b < utf8.RuneSelf && strconv.IsPrint(rune(b)) {
		return strconv.QuoteRune(rune(b))
	}
	return fmt.Sprintf(`%.2x`, b)
}

func (m *M) dump() string {
	ss := make([]string, len(m.States))
	for i := range m.States {
		ss[i] = m.States[i].dump(m.States, i)
	}
	return strings.Join(ss, "\n")
}

func (m *M) SaveSVG(file string, opt ...*GraphOption) error {
	dotCmd := exec.Command("dot", "-Tsvg", "-o", file)
	w, err := dotCmd.StdinPipe()
	if err != nil {
		return err
	}
	ew, err := dotCmd.StderrPipe()
	if err != nil {
		return err
	}
	defer w.Close()
	go dotCmd.Run()
	go func() {
		buf, _ := ioutil.ReadAll(ew)
		fmt.Println(string(buf))
	}()
	if len(opt) == 0 {
		opt = []*GraphOption{{}}
	}
	return m.writeDotFormat(w, opt[0])
}

func (m *M) SaveDot(file string, opt ...*GraphOption) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	if len(opt) == 0 {
		opt = []*GraphOption{{}}
	}
	return m.writeDotFormat(f, opt[0])

}

func (m *M) writeDotFormat(writer io.Writer, opt *GraphOption) error {
	var w bytes.Buffer
	w.WriteString("digraph g {\n")
	if opt.Timelabel {
		fmt.Fprintf(&w, "graph [label=\"(%s)\", labeljust=right, fontsize=12];", time.Now().Format("2006-01-02 15:04:05"))
	}
	w.WriteString("\trankdir=LR;\n")
	if opt.FontName != "" {
		fmt.Fprintf(&w, "\tnode [fontname=\"%s\"];\n", opt.FontName)
		fmt.Fprintf(&w, "\tedge [fontname=\"%s\"];\n", opt.FontName)
	}
	w.WriteString("\tnode [fontsize=12, shape=circle, fixedsize=true, width=\".25\"];\n")
	w.WriteString("\tedge [fontsize=12];\n")
	w.WriteString("\tedge [arrowhead=lnormal];\n")
	w.WriteString("\tENTRY [shape=point, fixedsize=false, width=\".05\"];\n")
	w.WriteString("\tENTRY -> 0 [label=\"(input)\"];\n")
	if len(m.States) > 0 {
		for i := range m.States {
			s := &m.States[i]
			s.writeDotFormat(&w, i)
		}
	}
	w.WriteString("}")

	_, err := w.WriteTo(writer)
	return err
}
func (s *S) writeDotFormat(w io.Writer, sid int) {
	if s.final() {
		label := ""
		if s.Label > defaultFinal {
			label = "L" + strconv.Itoa(int(s.Label.toExternal()))
		}
		fmt.Fprintf(w, "\t%d [shape=doublecircle, width=\".18\", xlabel=\"%s\"];\n", sid, label)
	}
	m := make(map[int]bool)
	for _, Trans := range s.Table {
		if !m[Trans.Next] {
			fmt.Fprintf(w, "\t%d -> %d [label=\"%s\"];\n", sid, Trans.Next, dotEscape(s.Table.description(Trans.Next)))
			m[Trans.Next] = true
		}
	}
}
func (Table TransTable) description(sid int) (l string) {
	for _, Trans := range Table {
		if Trans.Next == sid {
			if l != "" {
				l += `\n`
			}
			l += Trans.rangeString()
		}
	}
	return
}
func dotEscape(s string) string {
	return strings.Replace(s, `"`, `\"`, -1)
}
