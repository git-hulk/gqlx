package builder

const (
	KindQuery        = "query"
	KindMutation     = "mutation"
	KindSubscription = "subscription"
)

type Builder struct {
	// one of "query", "mutation", "subscription"
	kind       string
	name       string
	args       Arguments
	selections Selections
	directives Directives
	fragments  []*FragmentDef
}

func New() *Builder {
	return &Builder{}
}

func Query() *Builder {
	return &Builder{
		kind: KindQuery,
	}
}

func Mutation() *Builder {
	return &Builder{
		kind: KindMutation,
	}
}

func Subscription() *Builder {
	return &Builder{
		kind: KindSubscription,
	}
}

func (builder *Builder) Name(name string) *Builder {
	builder.name = name
	return builder
}

func (builder *Builder) AddArguments(args ...*typedArgument) *Builder {
	for _, arg := range args {
		builder.args.Add(arg)
	}
	return builder
}

func (builder *Builder) AddDirectives(directives ...*Directive) *Builder {
	builder.directives.Add(directives...)
	return builder
}

func (builder *Builder) AddSelections(selections ...Selection) *Builder {
	builder.selections.Add(selections...)
	return builder
}

func (builder *Builder) DeclareFragment(name, on string, selections Selections) *Builder {
	builder.fragments = append(builder.fragments, &FragmentDef{
		name:       name,
		on:         on,
		selections: selections,
	})
	return builder
}

func (builder *Builder) Validate() error {
	// TODO: named fragments must be declared before using
	// TODO: declare variables only allowed at the top level
	return nil
}

func (builder *Builder) String() string {
	f := NewFormatter()
	for _, fragment := range builder.fragments {
		fragment.Format(f)
	}

	f.WriteString(builder.kind)
	if builder.name != "" {
		f.WriteString(" " + builder.name)
	}
	f.WriteString(builder.args.String())
	if len(builder.directives) > 0 {
		f.WriteString(" ").WriteString(builder.directives.String())
	}
	f.WriteString(" {").NewLine()
	f.IncreaseLevel()
	for _, selection := range builder.selections {
		selection.Format(f)
	}
	f.DecreaseLevel()
	f.WriteString("}").NewLine()
	return f.String()
}
