package repository

import (
	"reflect"

	"gorm.io/gorm"
)

type AfterScan interface {
	AfterScan() error
}

func DoAfterScan(stmt *gorm.Statement) {
	_ = stmt.
		Callback().
		Query().
		After("gorm:after_query").
		Register("hooks:after_scan", func(db *gorm.DB) {
			if db.Error != nil || db.Statement.Dest == nil {
				return
			}

			val := reflect.ValueOf(db.Statement.Dest)
			if !val.IsValid() || (val.Kind() == reflect.Ptr && val.IsNil()) {
				return
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
						if db.Error = scanner.AfterScan(); db.Error != nil {
							return
						}
						break
					}
				}
			}
		})
}
