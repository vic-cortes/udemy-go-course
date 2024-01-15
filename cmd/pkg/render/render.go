package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/vmcortesf/udemy-course/cmd/pkg/config"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template

	if app.UseCache {
		// Create Template cache
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not render template")
	}

	buffer := new(bytes.Buffer)
	_ = t.Execute(buffer, nil)

	_, err := buffer.WriteTo(w)

	if err != nil {
		log.Println("Error writing template to browser :", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all the files named *.html in the templates folder
	pages, err := filepath.Glob("./templates/*.html")

	if err != nil {
		return myCache, err
	}

	// Loop through the pages one by one
	for _, page := range pages {
		name := filepath.Base(page)

		// ts = Template Set
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
