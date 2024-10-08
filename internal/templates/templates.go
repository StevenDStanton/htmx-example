package templates

import (
	"io"
	"text/template"
)

type Templates struct {
	templates *template.Template
}

func NewTemplates(path string) *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob(path)),
	}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
