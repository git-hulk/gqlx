package builder

import (
	"testing"

	"github.com/sebdah/goldie/v2"
)

func TestBuilder_Query(t *testing.T) {
	b := New()
	testCases := []struct {
		Query    *QueryBuilder
		Snapshot string
	}{
		{
			Query: b.Query().Name("user").
				AddSelections(
					NewField("id").Alias("user_id"),
					NewField("name").AddArguments(
						ValueArgument("age", Int(30)),
						ValueArgument("var", Variable("var")),
						TypedArgument("sex", "Sex", nil),
						TypedArgument("status", "UserStatus", String("active")),
					),
				),
			Snapshot: "basic",
		},
		{
			Query: b.Query().Name("foo").AddSelections(
				NewField("id"),
				NamedFragment("named_fragment"),
				InlineFragment("inline_fragment").AddSelections(
					NewField("bar"),
				),
			),
			Snapshot: "fragment",
		},
		{
			Query: b.Query().Name("foo").
				DeclareFragment("Hello", "Character", Selections{
					NewField("id"),
					NewField("name").AddArguments(ValueArgument("age", Int(30))),
				}).
				AddSelections(
					NewField("status"),
					NamedFragment("Hello"),
				),
			Snapshot: "fragment_with_definition",
		},
	}

	// Load the string from the snapshot file and compare the result.
	for _, testCase := range testCases {
		t.Run(testCase.Snapshot, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithNameSuffix(".golden.graphql"),
				goldie.WithDiffEngine(goldie.ColoredDiff),
				goldie.WithFixtureDir("testdata/query"))
			g.Assert(t, testCase.Snapshot, []byte(testCase.Query.String()))
		})
	}
}
