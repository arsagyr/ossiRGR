package handlers

import (
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	data := ViewData{
		Title: "Добро пожаловать! ",
		Text:  "Чтобы рассчитать интеграл, нажмите по кнопке",
	}

	tmpl, _ := template.ParseFiles("static/templates/index.html")
	tmpl.Execute(w, data)
}
