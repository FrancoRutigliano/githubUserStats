package utils

import (
	"html/template"
	"net/http"
)

const (
	templateDir  = "./internal/web/template/"
	templateBase = templateDir + "base.html"
)

func Render(w http.ResponseWriter, page string, data any) error {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
		return err
	}
	return nil
}
