package repository

import (
	"reflect"
)

type AfterScan interface {
	AfterScan() error
}

func DoAfterScan(out interface{}) error {
	if out == nil {
		return nil
	}

	val := reflect.ValueOf(out)
	if !val.IsValid() || (val.Kind() == reflect.Ptr && val.IsNil()) {
		return nil
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if err := doElementAfterScan(val.Index(i)); err != nil {
				return err
			}
		}
	default:
		if err := doElementAfterScan(val); err != nil {
			return err
		}
	}

	return nil
}

func doElementAfterScan(v reflect.Value) error {
	if v.CanAddr() {
		if addr := v.Addr(); addr.CanInterface() {
			if scanner, ok := addr.Interface().(AfterScan); ok {
				return scanner.AfterScan()
			}
		}
	}
	if v.CanInterface() {
		if scanner, ok := v.Interface().(AfterScan); ok {
			return scanner.AfterScan()
		}
	}
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		if elem := v.Elem(); elem.CanInterface() {
			if scanner, ok := elem.Interface().(AfterScan); ok {
				return scanner.AfterScan()
			}
		}
	}
	return nil
}
