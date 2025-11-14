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

	for i := 0; i < val.Len(); i++ {
		item := val.Index(i)

		for _, v := range []reflect.Value{item.Addr(), item} {
			if !v.CanInterface() {
				continue
			}
			if scanner, ok := v.Interface().(AfterScan); ok {
				if err := scanner.AfterScan(); err != nil {
					return err
				}
				break
			}
		}
	}

	return nil
}
