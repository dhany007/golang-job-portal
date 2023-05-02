package utils

import (
	"encoding/json"
	"reflect"
	"time"
)

func IsNil(value interface{}) (res bool) {
	condition1 := value == nil

	isPtr := reflect.TypeOf(value).Kind() == reflect.Ptr
	condition2 := isPtr && reflect.ValueOf(value).IsNil()

	return condition1 || condition2
}

func InterfaceToString(value interface{}) (res string) {
	if !IsNil(value) {
		val := reflect.ValueOf(value)
		switch val.Kind() {
		case reflect.String:
			res = val.String()
		case reflect.Ptr:
			res = InterfaceToString(reflect.Indirect(val))
		default:
			switch valx := value.(type) {
			case []byte:
				res = string(valx)
			case time.Time:
				res = valx.Format(time.RFC3339Nano)
			default:
				byt, err := json.Marshal(value)
				if err == nil {
					res = string(byt)
				}
			}
		}
	}

	return
}
