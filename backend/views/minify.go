package views

import (
	"regexp"
	"strings"
)

var whitespaceRegex = regexp.MustCompile(`\s+`)

func (r *RendererImpl) minify(html string) string {
	html = whitespaceRegex.ReplaceAllString(html, " ")
	html = strings.ReplaceAll(html, "> <", "><")
	return strings.TrimSpace(html)
}
