package views

import (
	"html/template"
	"sync"
)

type Renderer struct {
	cache      map[string]*template.Template
	cacheMutex *sync.RWMutex
}

func Provider() *Renderer {
	var templatesCount = 0
	templatesDir, err := templatesFS.ReadDir("templates")
	if err == nil {
		templatesCount = len(templatesDir)
	}

	return &Renderer{
		cache:      make(map[string]*template.Template, templatesCount),
		cacheMutex: &sync.RWMutex{},
	}
}
