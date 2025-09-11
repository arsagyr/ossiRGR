package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"rgr/model"
)

func Output(w http.ResponseWriter, r *http.Request) {

	vars := map[string]interface{}{
		"x":  F.Arguement,
		"pi": 3.14159265359,
		"E":  2.71828182845,
	}

	result, err := model.EvaluateExpression(F.Formula, vars)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	s := strconv.FormatFloat(result, 'f', 6, 32)
	data := ViewData{
		Title: "Ответ ",
		Text:  s,
	}

	tmpl, _ := template.ParseFiles("static/templates/index.html")
	tmpl.Execute(w, data)
}
