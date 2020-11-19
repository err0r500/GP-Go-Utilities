package vongo

import (
	"encoding/json"
	"reflect"
)

// ValueOf return the reflect Value of d. In case of slice or map
// it reduces to a new primitive type.
func ValueOf(d interface{}) reflect.Value {
	v := reflect.ValueOf(d)

	if v.Type().Kind() == reflect.Slice || v.Type().Kind() == reflect.Map {
		inner := v.Type().Elem()
		switch inner.Kind() {
		case reflect.Ptr:
			v = reflect.New(inner.Elem()).Elem()
		default:
			v = reflect.New(inner).Elem()
		}
	} else if v.Type().Kind() == reflect.Ptr {
		return ValueOf(reflect.Indirect(v).Interface())
	}

	return v
}

func isSlice(s interface{}) bool {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		return false
	}

	return true
}

func isMap(s interface{}) bool {
	if reflect.TypeOf(s).Kind() != reflect.Map {
		return false
	}

	return true
}

func printOptional(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func interfaceName(i interface{}) string {
	var n string

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		inner := v.Type().Elem()
		switch inner.Kind() {
		case reflect.Ptr:
			n = v.Type().Elem().Elem().Name()
		default:
			n = v.Type().Elem().Name()
		}
	case reflect.Ptr:
		return interfaceName(reflect.Indirect(v).Interface())
	default:
		n = v.Type().Name()
	}

	return n
}
