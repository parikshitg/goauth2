package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	data["title"] = "Dashboard"

	page, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}
