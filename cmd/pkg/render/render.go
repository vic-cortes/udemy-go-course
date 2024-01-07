package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	// Parse the template file
	log.Println("Parsing template :", tmpl)
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")

	// Execute the template
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

// template cache
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check if the template is already in the cache
	_, inMap := tc[t]

	if !inMap {
		// Need to create the template
		log.Println("Creating template cache for :", t)
		err = createTemplateCache(t)

		if err != nil {
			log.Println("Error creating template cache :", err)
			return
		}

	} else {
		log.Println("Using cached Template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println("Error creating template cache :", err)
		return
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// add the template to the cache
	tc[t] = tmpl

	return nil
}
