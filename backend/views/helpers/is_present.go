package helpers

import "reflect"

func isPresent(value any) bool {
	if value == nil {
		return false
	}
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		return !v.IsNil()
	}
	return true
}
