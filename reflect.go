package tool

import "reflect"

func GetFullName(typ reflect.Type) string {
	return typ.PkgPath() + "." + typ.Name()
}
