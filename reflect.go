package tool

import "reflect"

func GetFullName(typ reflect.Type) string {
	return typ.PkgPath() + "." + typ.Name()
}

// get real reflect.Value from raw, across ptr
func GetRealReflectValue(raw reflect.Value) *reflect.Value {
	if raw.Kind() == reflect.Ptr || raw.Kind() == reflect.Interface {
		raw = raw.Elem()
	}
	return &raw
}
