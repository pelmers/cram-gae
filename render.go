package cram

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// render a static file in the templates/ directory
func render(w http.ResponseWriter, tmpl string) {
	// prepend templates/ to tmpl string
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("error parsing template: ", err)
	}
	if err := t.Execute(w, ""); err != nil {
		log.Print("error writing template: ", err)
	}
}
