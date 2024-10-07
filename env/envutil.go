package env

import (
	"fmt"
	"reflect"
	"time"
)

func ToRyeValue(val any) Object {
	switch v := val.(type) {
	case float64:
		return *NewDecimal(v)
	case int:
		return *NewInteger(int64(v))
	case int64:
		return *NewInteger(v)
	case string:
		return *NewString(v)
	case rune:
		return *NewString(string(v))
	case map[string]any:
		return *NewDict(v)
	case []any:
		return *NewList(v)
	case *List:
		return *v
	case *Block:
		return *v
	case *Object:
		return *v
	case Object:
		return v
	case nil:
		return nil
	case time.Time:
		return *NewString(v.Format(time.RFC3339))
	default:
		fmt.Println(val)
		// TODO-FIXME
		return Void{}
	}
}

func IsPointer(x any) bool {
	return reflect.TypeOf(x).Kind() == reflect.Pointer
}

func IsPointer2(x reflect.Value) bool {
	return x.Kind() == reflect.Ptr
}

func DereferenceAny(x any) any {
	if reflect.TypeOf(x).Kind() == reflect.Ptr {
		return reflect.ValueOf(x).Elem().Interface()
	}
	return x
}

func ReferenceAny(x any) any {
	return &x
}
