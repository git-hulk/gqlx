package builder

import "strings"

type Directive struct {
	name string
	args Arguments
}

func (d *Directive) Kind() string {
	return "directive"
}

func (d *Directive) String() string {
	var b strings.Builder
	b.WriteString("@" + d.name)
	if len(d.args) == 0 {
		return b.String()
	}

	b.WriteString(d.args.String())
	return b.String()
}

type Directives []*Directive

func (d *Directives) Add(directives ...*Directive) {
	*d = append(*d, directives...)
}

func (d Directives) String() string {
	if len(d) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString(" ")
	for i, directive := range d {
		b.WriteString(directive.String())
		if i < len(d)-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}
