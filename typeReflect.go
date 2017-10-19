package tool

import (
	"log"
	"reflect"
	"time"
)

func IsZero(v reflect.Value) (isValid bool) {

	if !v.CanInterface() {
		return false
	}
	value := v.Interface()
	switch v.Kind() {
	case reflect.String:
		if value.(string) == "" {
			isValid = true
		}
	case reflect.Bool:
		if value.(bool) {
			isValid = true
		}
	case reflect.Int:
		if value.(int) == 0 {
			isValid = true
		}
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		if v.Len() != 0 {
			isValid = true
		}
	case reflect.Struct:
		if _v, ok := value.(time.Time); ok {
			if _v.IsZero() {
				isValid = true
			}
		}
	default:
		isValid = false
	}
	log.Print(isValid)
	return isValid
}
