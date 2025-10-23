package model

import (
	"fmt"
	"sync"

	"gorm.io/gorm/schema"
)

func GetColumns(model interface{}) []string {
	s, _ := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
	var columns []string

	for _, field := range s.Fields {
		if field.DBName != "" && field.Creatable {
			columns = append(columns, fmt.Sprintf("%s.%s", s.Table, field.DBName))
		}
	}

	return columns
}
