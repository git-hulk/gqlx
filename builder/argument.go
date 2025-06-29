package builder

import "strings"

type DataType string

type Argument struct {
	name       string
	val        *Value
	valType    DataType
	defaultVal *Value // Optional default value for the argument
}

func ValueArgument(name string, val *Value) *Argument {
	return &Argument{
		name: name,
		val:  val,
	}
}

func TypedArgument(name string, valType DataType, val *Value) *Argument {
	return &Argument{
		name:       name,
		valType:    valType,
		defaultVal: val,
	}
}

func declaredVarArgument(name string, valType DataType, val *Value) *Argument {
	return &Argument{
		name:       "$" + name,
		valType:    valType,
		defaultVal: val, // No default value for declared variables
	}
}

func (a *Argument) String() string {
	if a.val != nil {
		return a.name + ": " + a.val.String()
	}

	result := a.name + ": " + string(a.valType)
	if a.defaultVal != nil {
		result += " = " + a.defaultVal.String()
	}
	return result
}

type Arguments []*Argument

func (a *Arguments) Add(args ...*Argument) {
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
