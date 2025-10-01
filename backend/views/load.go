package views

import (
	"embed"
	"fmt"
	"html/template"

	"polimane/backend/views/helpers"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

func (r *RendererImpl) load(view string) (*template.Template, error) {
	r.cacheMutex.RLock()
	if tmpl, ok := r.cache[view]; ok {
		r.cacheMutex.RUnlock()
		return tmpl, nil
	}
	r.cacheMutex.RUnlock()

	r.cacheMutex.Lock()
	defer r.cacheMutex.Unlock()

	name := fmt.Sprintf("templates/%s.tmpl", view)

	source, err := templatesFS.ReadFile(name)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.
		New(name).
		Funcs(helpers.Get()).
		Parse(string(source))

	if err != nil {
		return nil, err
	}

	r.cache[view] = tmpl
	return tmpl, nil
}
