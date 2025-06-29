package builder

import "strings"

const DefaultIndent = "  "

type Formatter struct {
	indent      string
	indentLevel int

	buffer strings.Builder
}

func NewFormatter() *Formatter {
	return &Formatter{
		indent:      DefaultIndent,
		indentLevel: 0,
	}
}

func (f *Formatter) IncreaseLevel() {
	f.indentLevel++
}

func (f *Formatter) DecreaseLevel() {
	f.indentLevel--
	if f.indentLevel < 0 {
		f.indentLevel = 0
	}
}

func (f *Formatter) Indent() string {
	return strings.Repeat(f.indent, f.indentLevel)
}

func (f *Formatter) WriteIndent() *Formatter {
	f.buffer.WriteString(f.Indent())
	return f
}

func (f *Formatter) WriteString(s string) *Formatter {
	f.buffer.WriteString(s)
	return f
}

func (f *Formatter) NewLine() *Formatter {
	f.buffer.WriteString("\n")
	return f
}

func (f *Formatter) String() string {
	return f.buffer.String()
}
