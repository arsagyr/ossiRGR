package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"rgr/model"
)

func Output(w http.ResponseWriter, r *http.Request) {
	var s string = model.TrapzoidMethodStringReturn(F.Formula, F.BorderA, F.BorderB)

	fmt.Println(s)

	data := ViewData{
		Title: "Ответ ",
		Text:  s,
	}

	tmpl, _ := template.ParseFiles("static/templates/index.html")
	tmpl.Execute(w, data)
}
