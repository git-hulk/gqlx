package builder

import (
	"strings"

	"github.com/git-hulk/gqlx/builder/value"
)

type DataType string

type Argument interface {
	IsValue() bool
	String() string
}

type valueArgument struct {
	name string
	val  *value.Value
}

type typedArgument struct {
	name       string
	valType    DataType
	defaultVal *value.Value // Optional default value for the argument
}

func FromValue(name string, val *value.Value) *valueArgument {
	return &valueArgument{
		name: name,
		val:  val,
	}
}

func (a *valueArgument) String() string {
	return a.name + ": " + a.val.String()
}

func (a *valueArgument) IsValue() bool {
	return true
}

func FromType(name string, valType DataType, val *value.Value) *typedArgument {
	return &typedArgument{
		name:       name,
		valType:    valType,
		defaultVal: val,
	}
}

func (a *typedArgument) String() string {
	result := a.name + ": " + string(a.valType)
	if a.defaultVal != nil {
		result += " = " + a.defaultVal.String()
	}
	return result
}

func (a *typedArgument) IsValue() bool {
	return false
}

type Arguments []Argument

func (a *Arguments) Add(args ...Argument) {
	*a = append(*a, args...)
}

func (a Arguments) String() string {
	if len(a) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("(")
	for i, arg := range a {
		b.WriteString(arg.String())
		if i < len(a)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(")")
	return b.String()
}
