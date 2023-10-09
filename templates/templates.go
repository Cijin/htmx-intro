package templates

import (
	"html/template"
	"io"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func New(path string) *Template {
	return &Template{
		Templates: template.Must(template.ParseGlob(path)),
	}
}
