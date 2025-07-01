package value

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValue_String(t *testing.T) {
	require.Equal(t, "null", Null().String())
	require.Equal(t, "\"test\"", String("test").String())
	require.Equal(t, "123", Int(123).String())
	require.Equal(t, "123.456", Float(123.456).String())
	require.Equal(t, "true", Boolean(true).String())
	require.Equal(t, "false", Boolean(false).String())
	require.Equal(t, "ENUM_VALUE", Enum("ENUM_VALUE").String())
	require.Equal(t, "$variable", Variable("variable").String())
	require.Equal(t, "[1, 2, 3]", List(Int(1), Int(2), Int(3)).String())
	require.Equal(t, "{key: \"value\"}", Object(map[string]*Value{"key": String("value")}).String())

	// Nested Lists and Objects
	require.Equal(t, "[[1, 2], [3, 4]]", List(List(Int(1), Int(2)), List(Int(3), Int(4))).String())
	require.Equal(t, "{key: {nested_key: \"nested_value\"}}", Object(map[string]*Value{
		"key": Object(map[string]*Value{
			"nested_key": String("nested_value"),
		})}).String())
}
