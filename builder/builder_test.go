package builder

import (
	"testing"

	"github.com/git-hulk/gqlx/builder/value"

	"github.com/sebdah/goldie/v2"
)

func TestBuilder_Query(t *testing.T) {
	testCases := []struct {
		Query    *Builder
		Snapshot string
	}{
		{
			Query: Query().Name("user").
				AddSelections(
					NewField("id").Alias("user_id"),
					NewField("name").AddArguments(
						FromValue("age", value.Int(30)),
						FromValue("var", value.Variable("var")),
						FromType("sex", "Sex", nil),
						FromType("status", "UserStatus", value.String("active")),
					),
				),
			Snapshot: "basic",
		},
		{
			Query: Query().Name("foo").AddSelections(
				NewField("id"),
				NamedFragment("named_fragment"),
				InlineFragment("inline_fragment").AddSelections(
					NewField("bar"),
				),
			),
			Snapshot: "fragment",
		},
		{
			Query: Query().Name("foo").
				DeclareFragment("Hello", "Character", Selections{
					NewField("id"),
					NewField("name").AddArguments(FromValue("age", value.Int(30))),
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

func TestBuilder_Mutation(t *testing.T) {
	testCases := []struct {
		Mutation *Builder
		Snapshot string
	}{
		{
			Mutation: Mutation().Name("rateFilm").AddArguments(
				FromType("episode", "Episode", nil),
				FromType("rating", "Rating", nil),
			).AddSelections(
				NewField("episode"),
				NewField("viewerRating"),
			),
			Snapshot: "basic",
		},

		// Multiple fields in mutation
		{
			Mutation: Mutation().Name("deleteStarship").AddArguments(
				FromType("id", "ID!", value.String("3001")),
			).AddSelections(
				NewField("firstShip").Alias("first_ship").AddArguments(
					FromValue("id", value.String("3001")),
				),
				NewField("secondShip").Alias("second_ship").AddArguments(
					FromValue("id", value.String("3002")),
				),
			),
			Snapshot: "multiple_fields",
		},
		// Variables in mutation
		{
			Mutation: Mutation().Name("CreateReviewForEpisode").AddArguments(
				FromType("ep", "Episode", nil),
				FromType("review", "ReviewInput", nil),
			).AddSelections(
				NewField("createReview").AddArguments(
					FromValue("episode", value.Variable("ep")),
					FromValue("review", value.Variable("review")),
				).AddSelections(
					NewField("stars"),
					NewField("commentary"),
				),
			),
			Snapshot: "variables",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Snapshot, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithNameSuffix(".golden.graphql"),
				goldie.WithDiffEngine(goldie.ColoredDiff),
				goldie.WithFixtureDir("testdata/mutation"))
			g.Assert(t, testCase.Snapshot, []byte(testCase.Mutation.String()))
		})
	}
}
