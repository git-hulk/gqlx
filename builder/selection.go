package builder

// Selection allows multiple types of selections to be made, such as fields, fragments, or inline fragments.
type Selection interface {
	SelectionKind() string
	Format(f *Formatter)
}

type Selections []Selection

func (s *Selections) Add(selections ...Selection) {
	*s = append(*s, selections...)
}

type Field struct {
	name       string
	args       Arguments
	directives Directives
	alias      string // Optional alias for the field
}

func NewField(name string) *Field {
	return &Field{
		name:       name,
		args:       make(Arguments, 0),
		directives: make(Directives, 0),
	}
}

func (field *Field) Alias(alias string) *Field {
	field.alias = alias
	return field
}

func (field *Field) AddArguments(args ...*Argument) *Field {
	field.args.Add(args...)
	return field
}

func (field *Field) AddDirectives(directives ...*Directive) *Field {
	field.directives = append(field.directives, directives...)
	return field
}

func (field Field) SelectionKind() string {
	return "field"
}

func (field Field) Format(f *Formatter) {
	f.WriteIndent()
	if field.alias != "" {
		f.WriteString(field.alias).WriteString(": ")
	}
	f.WriteString(field.name).
		WriteString(field.args.String()).
		WriteString(field.directives.String()).
		NewLine()
}
