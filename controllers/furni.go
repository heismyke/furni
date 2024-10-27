package controllers

import (
	"net/http"
	"text/template"

	"github.com/heismyke/furni/data"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from here "))
}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.tmpl")
    if err != nil {
        http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, data.GetAll())
    if err != nil {
        http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
        return
    }
}
