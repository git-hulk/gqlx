package builder

// FragmentDef represents a GraphQL fragment definition, which includes the name of the fragment,
// the type it applies to (the "on" type), and the selections within the fragment.
//
// We add a dedicated type for FragmentDef to distinguish between a fragment definition and a fragment usage
// to avoid mixing them in selections or queries, even though they share similar structures.
type FragmentDef struct {
	name       string
	on         string
	selections Selections
}

func DeclareFragment(name, on string) *FragmentDef {
	return &FragmentDef{
		name:       name,
		on:         on,
		selections: make(Selections, 0),
	}
}

func (fragmentDef *FragmentDef) AddSelections(selections ...Selection) *FragmentDef {
	fragmentDef.selections.Add(selections...)
	return fragmentDef
}

func (fragmentDef *FragmentDef) Format(f *Formatter) {
	f.WriteIndent().WriteString("fragment ").
		WriteString(fragmentDef.name).WriteString(" on ").
		WriteString(fragmentDef.on).WriteString(" {").NewLine()
	f.IncreaseLevel()
	for _, sel := range fragmentDef.selections {
		sel.Format(f)
	}
	f.DecreaseLevel()
	f.WriteIndent().WriteString("}").NewLine()
}

// Fragment represents a GraphQL fragment, which can be either a named fragment or an inline fragment,
// as well as a fragment definition.
type Fragment struct {
	// it's a fragment definition if name and on are both set
	// fragment Hello on User {
	//   foo
	//   bar
	// }

	// named fragments
	// {
	// 	 ... Hello
	//   foo
	// }
	name string

	// inline fragments
	// ... On Hello {
	//   bar
	// }
	on         string
	selections Selections
}

func NamedFragment(name string) *Fragment {
	return &Fragment{
		name: name,
	}
}

func InlineFragment(on string) *Fragment {
	return &Fragment{
		on:         on,
		selections: make(Selections, 0),
	}
}

func (fragment *Fragment) AddSelections(selections ...Selection) *Fragment {
	fragment.selections.Add(selections...)
	return fragment
}

func (fragment *Fragment) SelectionKind() string {
	return "fragment"
}

func (fragment *Fragment) Format(f *Formatter) {
	if fragment.name != "" {
		f.WriteIndent()
		f.WriteString("... ").WriteString(fragment.name).NewLine()
	} else {
		f.WriteIndent()
		f.WriteString("...").
			WriteString(" on ").WriteString(fragment.on).
			WriteString(" {").NewLine()
		f.IncreaseLevel()
		for _, sel := range fragment.selections {
			sel.Format(f)
		}
		f.DecreaseLevel()
		f.WriteIndent().WriteString("}").NewLine()
	}
}
