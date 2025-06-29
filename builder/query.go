package builder

// QueryBuilder represents a GraphQL query builder that allows adding selections and directives.
type QueryBuilder struct {
	name       string
	args       Arguments
	selections Selections
	directives Directives
	fragments  []*FragmentDef
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		selections: make(Selections, 0),
		directives: make(Directives, 0),
	}
}

func (qb *QueryBuilder) Name(name string) *QueryBuilder {
	qb.name = name
	return qb
}

func (qb *QueryBuilder) AddArguments(args ...*Argument) *QueryBuilder {
	qb.args.Add(args...)
	return qb
}

func (qb *QueryBuilder) AddDirectives(directives ...*Directive) *QueryBuilder {
	qb.directives.Add(directives...)
	return qb
}

func (qb *QueryBuilder) AddSelections(selections ...Selection) *QueryBuilder {
	qb.selections.Add(selections...)
	return qb
}

func (qb *QueryBuilder) DeclareFragment(name, on string, selections Selections) *QueryBuilder {
	qb.fragments = append(qb.fragments, &FragmentDef{
		name:       name,
		on:         on,
		selections: selections,
	})
	return qb
}

func (qb *QueryBuilder) Validate() error {
	// TODO: named fragments must be declared before using
	// TODO: declare variables only allowed at the top level
	return nil
}

func (qb *QueryBuilder) String() string {
	f := NewFormatter()
	for _, fragment := range qb.fragments {
		fragment.Format(f)
	}

	f.WriteString("query")
	if qb.name != "" {
		f.WriteString(" " + qb.name)
	}
	f.WriteString(qb.args.String())
	if len(qb.directives) > 0 {
		f.WriteString(" ").WriteString(qb.directives.String())
	}
	f.WriteString(" {").NewLine()
	f.IncreaseLevel()
	for _, selection := range qb.selections {
		selection.Format(f)
	}
	f.DecreaseLevel()
	f.WriteString("}").NewLine()
	return f.String()
}
