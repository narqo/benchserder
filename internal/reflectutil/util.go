package reflectutil

import (
	"fmt"
	"reflect"
)

func IsNilInterface(value interface{}) bool {
	if value == nil {
		return true
	}
	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func HasNoNilFields(value interface{}) error {
	v, err := getStructValue(reflect.ValueOf(value))
	if err != nil {
		return err
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		switch fv.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			if fv.IsNil() {
				return fmt.Errorf("%s.%s is nil", t.Name(), t.Field(i).Name)
			}
		default:
		}
	}
	return nil
}

func getStructValue(v reflect.Value) (reflect.Value, error) {
	switch v.Kind() {
	case reflect.Struct:
		return v, nil
	case reflect.Ptr, reflect.Interface:
		return getStructValue(v.Elem())
	default:
		return v, fmt.Errorf("Bad argument, need struct or pointer to struct")
	}
}
