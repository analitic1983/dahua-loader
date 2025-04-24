package json

import "reflect"

func MarshalMap[T any](input any, serializer func(T) map[string]interface{}) interface{} {
	v := reflect.ValueOf(input)

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		length := v.Len()
		result := make([]map[string]interface{}, 0, length)

		for i := 0; i < length; i++ {
			item := v.Index(i).Interface().(T)
			result = append(result, serializer(item))
		}

		return result

	default:
		return serializer(input.(T))
	}
}
