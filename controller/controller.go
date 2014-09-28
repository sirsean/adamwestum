package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/sirsean/adamwestum/config"
)

var indexTemplate = buildTemplate("index.html")

func Index(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, nil)
}

func buildTemplate(file string) *template.Template {
	return template.Must(template.ParseFiles(templatePath(file)))
}

func templatePath(file string) string {
	return fmt.Sprintf("%s/template/%s", config.Get().Host.Path, file)
}
