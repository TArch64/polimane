package helpers

import "html/template"

func Get() template.FuncMap {
	return template.FuncMap{
		"beadsGrid": beadsGrid,
		"isPresent": isPresent,
	}
}
