package handlers

import (
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"

	config "github.com/sidramon/AxoloGo/Config"
	"github.com/sidramon/AxoloGo/models"
)

func Home(w http.ResponseWriter,  r *http.Request) {
	user := &models.TemplateData{
		Username: "Testeur123",
		Email: "testeur@hotmail.com",

	}

	renderTemplate(w, "home", user)
}

func Contact(w http.ResponseWriter,  r *http.Request) {
	renderTemplate(w, "contact", &models.TemplateData{

	})
}

var appConfig *config.Config

func CreateTemplates(app *config.Config) {
	appConfig = app
}

func renderTemplate(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
	templateCache := appConfig.TemplateCache

	tmpl, ok := templateCache[tmplName + ".page.tmpl"]

	if !ok {
		http.Error(w, "Le template n'existe pas !", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, td)
	buffer.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}

		cache[name] = tmpl
	}

	return cache, nil
}