package views

import "html/template"

type Renderer interface {
	Render(options *RenderOptions) (string, error)
}

type RendererImpl struct {
	cache map[string]*template.Template
}

func Provider() Renderer {
	var templatesCount = 0
	templatesDir, err := templatesFS.ReadDir("templates")
	if err == nil {
		templatesCount = len(templatesDir)
	}

	return &RendererImpl{
		cache: make(map[string]*template.Template, templatesCount),
	}
}
