package service

import (
	"reflect"
	"strings"
)

func CheckSafeId(obj any) bool {
	completeSafeId(obj)
	return false
}

func completeSafeId(obj any) {

	v := reflect.ValueOf(obj)
	safeId := v.Elem().FieldByName("SafeId").String()
	if safeId == "" {
		namespace := v.Elem().FieldByName("Namespace").String()
		safeId, _, _ = strings.Cut(namespace, "-")
	}
	safeId = strings.ToLower(safeId)
	reflect.Indirect(v).FieldByName("SafeId").SetString(safeId)

}
