package builder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelection_Field(t *testing.T) {
	testCases := []struct {
		Selection Selection
		GQL       string
	}{
		{
			// only with field name
			Selection: &Field{name: "id"},
			GQL:       "id\n",
		},
		{
			// with field name and arguments
			Selection: &Field{
				name: "user",
				args: []Argument{
					FromValue("id", String("123")),
				},
			},
			GQL: "user(id: \"123\")\n",
		},
		{
			// with field name, arguments, and directives
			Selection: &Field{
				name: "age",
				directives: []*Directive{
					{name: "include", args: []Argument{
						FromValue("if", Boolean(true)),
					}},
				},
				args: []Argument{
					FromValue("id", String("123")),
				},
			},
			GQL: "age(id: \"123\") @include(if: true)\n",
		},
	}
	for _, testCase := range testCases {
		f := NewFormatter()
		testCase.Selection.Format(f)
		require.Equal(t, testCase.GQL, f.String(), "String output should match expected for selection")
	}
}
