package views

import (
	"regexp"
	"strings"
)

var whitespaceRegex = regexp.MustCompile(`\s+`)

func (r *Renderer) minify(html string) string {
	html = whitespaceRegex.ReplaceAllString(html, " ")
	html = strings.ReplaceAll(html, "> <", "><")
	return strings.TrimSpace(html)
}
