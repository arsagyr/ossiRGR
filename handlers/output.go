package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"rgr/model"
)

func Output(w http.ResponseWriter, r *http.Request) {
	var s string = model.IntrgrateTrapzoidString(F.Formula, F.BorderA, F.BorderB)
	// model.IntrgrateString(F.Formula, F.BorderA, F.BorderB)
	// var s0 string = model.IntrgrateTrapzoidString(F.Formula, F.BorderA, F.BorderB)
	// s = s + " | " + s0
	fmt.Printf("Answer is %s \n", s)

	data := ViewData{
		Title: "Ответ ",
		Text:  s,
	}

	tmpl, _ := template.ParseFiles("static/templates/index.html")
	tmpl.Execute(w, data)
}
