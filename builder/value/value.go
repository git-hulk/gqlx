package value

import (
	"fmt"
	"strings"
)

type ValueKind int

const (
	kindNull ValueKind = iota + 1
	kindString
	kindInt
	kindFloat
	kindBoolean
	kindEnum
	kindList
	kindObject
	kindVariable
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
	case kindNull:
		return "null"
	case kindString:
		return "\"" + v.val.(string) + "\""
	case kindInt, kindFloat:
		return fmt.Sprintf("%v", v.val)
	case kindBoolean:
		if v.val.(bool) {
			return "true"
		}
		return "false"
	case kindEnum:
		return v.val.(string)
	case kindList:
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
	case kindObject:
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
	case kindVariable:
		return "$" + v.val.(string)
	default:
		return ""
	}
}

func Null() *Value {
	return &Value{kind: kindNull, val: nil}
}

func String(val string) *Value {
	return &Value{kind: kindString, val: val}
}

func Int(val int) *Value {
	return &Value{kind: kindInt, val: val}
}

func Float(val float64) *Value {
	return &Value{kind: kindFloat, val: val}
}

func Boolean(val bool) *Value {
	return &Value{kind: kindBoolean, val: val}
}

func Enum(val string) *Value {
	return &Value{kind: kindEnum, val: val}
}

func List(values ...*Value) *Value {
	return &Value{kind: kindList, val: values}
}

func Object(values map[string]*Value) *Value {
	return &Value{kind: kindObject, val: values}
}

func Variable(name string) *Value {
	return &Value{kind: kindVariable, val: name}
}
