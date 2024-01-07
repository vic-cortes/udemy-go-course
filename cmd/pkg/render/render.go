package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Parse the template file
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.html")

	// Execute the template
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
} 