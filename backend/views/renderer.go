package views

import (
	"html/template"
	"sync"
)

type Renderer interface {
	Render(options *RenderOptions) (string, error)
}

type RendererImpl struct {
	cache      map[string]*template.Template
	cacheMutex *sync.RWMutex
}

func Provider() Renderer {
	var templatesCount = 0
	templatesDir, err := templatesFS.ReadDir("templates")
	if err == nil {
		templatesCount = len(templatesDir)
	}

	return &RendererImpl{
		cache:      make(map[string]*template.Template, templatesCount),
		cacheMutex: &sync.RWMutex{},
	}
}
