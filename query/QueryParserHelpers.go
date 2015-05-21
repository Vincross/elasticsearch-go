package query

import (
	"encoding/json"
	"reflect"
)

func toJson(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func wrapper(name string, v interface{}) map[string]interface{} {
	query := map[string]interface{}{
		name: v,
	}
	return query
}

func convert(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.Slice:
		var res []interface{}
		for i := 0; i < value.Len(); i++ {
			res = append(res, convert(value.Index(i)))
		}
		return res
	case reflect.Ptr:
		if !value.IsNil() {
			return convert(value.Elem())
		}
	case reflect.String:
		return value.String()
	case reflect.Float64:
		return value.Float()
	case reflect.Int:
		return value.Int()
	case reflect.Bool:
		return value.Bool()
	}
	return value.Interface()
}

func convertable(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Slice:
		return value.Len() > 0
	case reflect.Ptr:
		return !value.IsNil()
	}
	return true
}

func convertStruct(x interface{}) interface{} {
	query := make(map[string]interface{})

	t := reflect.ValueOf(x).Type()
	v := reflect.ValueOf(x)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.FieldByName(field.Name)
		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}
		if convertable(value) {
			query[tag] = convert(value)
		}
	}

	return query

}
