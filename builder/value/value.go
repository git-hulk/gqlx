package value

import (
	"fmt"
	"strings"
)

type ValueKind int

const (
	ValueKindNull ValueKind = iota + 1
	ValueKindString
	ValueKindInt
	ValueKindFloat
	ValueKindBoolean
	ValueKindEnum
	ValueKindList
	ValueKindObject
	ValueKindVariable
)

type Value struct {
	kind ValueKind
	val  any
}

func (v *Value) Kind() ValueKind {
	return v.kind
}

func (v *Value) String() string {
	switch v.kind {
	case ValueKindNull:
		return "null"
	case ValueKindString:
		return "\"" + v.val.(string) + "\""
	case ValueKindInt, ValueKindFloat:
		return fmt.Sprintf("%v", v.val)
	case ValueKindBoolean:
		if v.val.(bool) {
			return "true"
		}
		return "false"
	case ValueKindEnum:
		return v.val.(string)
	case ValueKindList:
		list := v.val.([]*Value)
		var b strings.Builder
		b.WriteString("[")
		for i, item := range list {
			b.WriteString(item.String())
			if i < len(list)-1 {
				b.WriteString(", ")
			}
		}
		b.WriteString("]")
		return b.String()
	case ValueKindObject:
		obj := v.val.(map[string]*Value)
		var b strings.Builder
		b.WriteString("{")
		tmp := make([]string, 0, len(obj))
		for key, value := range obj {
			tmp = append(tmp, fmt.Sprintf("%s: %s", key, value.String()))
		}
		b.WriteString(strings.Join(tmp, ", "))
		b.WriteString("}")
		return b.String()
	case ValueKindVariable:
		return "$" + v.val.(string)
	default:
		return ""
	}
}

func Null() *Value {
	return &Value{kind: ValueKindNull, val: nil}
}

func String(val string) *Value {
	return &Value{kind: ValueKindString, val: val}
}

func Int(val int) *Value {
	return &Value{kind: ValueKindInt, val: val}
}

func Float(val float64) *Value {
	return &Value{kind: ValueKindFloat, val: val}
}

func Boolean(val bool) *Value {
	return &Value{kind: ValueKindBoolean, val: val}
}

func Enum(val string) *Value {
	return &Value{kind: ValueKindEnum, val: val}
}

func List(values ...*Value) *Value {
	return &Value{kind: ValueKindList, val: values}
}

func Object(values map[string]*Value) *Value {
	return &Value{kind: ValueKindObject, val: values}
}

func Variable(name string) *Value {
	return &Value{kind: ValueKindVariable, val: name}
}
