package helper

import "reflect"

func GetJSONFieldNames(v interface{}) map[string]string {
	fieldNames := make(map[string]string)

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag != "" {
			fieldNames[field.Name] = jsonTag
		}
	}

	return fieldNames
}
