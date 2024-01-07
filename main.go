package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"


// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Parse the template file
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	// Execute the template
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
} 

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}