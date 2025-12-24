package views

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

func (r *Renderer) load(view string) (*template.Template, error) {
	r.cacheMutex.RLock()
	if tmpl, ok := r.cache[view]; ok {
		r.cacheMutex.RUnlock()
		return tmpl, nil
	}
	r.cacheMutex.RUnlock()

	r.cacheMutex.Lock()
	defer r.cacheMutex.Unlock()

	if tmpl, ok := r.cache[view]; ok {
		return tmpl, nil
	}

	name := fmt.Sprintf("templates/%s.tmpl", view)

	source, err := templatesFS.ReadFile(name)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(name).Parse(string(source))
	if err != nil {
		return nil, err
	}

	r.cache[view] = tmpl
	return tmpl, nil
}
