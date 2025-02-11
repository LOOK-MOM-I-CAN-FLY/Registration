package handlers

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/templates/index.html")
	if err != nil {
		http.Error(w, "Error download page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
