package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create Template cache
	tc, err := createTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	// Get the template from the cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	err = t.Execute(buffer, nil)

	if err != nil {
		log.Println("Error executing template :", err)
	}

	_, err = buffer.WriteTo(w)

	if err != nil {
		log.Println("Error writing template to browser :", err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
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
