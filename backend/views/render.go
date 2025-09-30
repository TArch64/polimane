package views

import (
	"bytes"
)

type RenderOptions struct {
	View   string
	Data   interface{}
	Minify bool
}

func (r *RendererImpl) Render(options *RenderOptions) (string, error) {
	tmpl, err := r.load(options.View)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, options.Data); err != nil {
		return "", err
	}

	result := buf.String()

	if options.Minify {
		result = r.minify(result)
	}

	return result, nil
}
