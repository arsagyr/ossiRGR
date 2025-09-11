package handlers

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"

	"rgr/model"
)

func Output(w http.ResponseWriter, r *http.Request) {
	var s string
	vars1 := map[string]interface{}{
		"x":  F.BorderA,
		"pi": 3.14159265359,
		"E":  2.71828182845,
	}
	vars2 := map[string]interface{}{
		"x":  F.BorderB,
		"pi": 3.14159265359,
		"E":  2.71828182845,
	}

	fxa, err := model.EvaluateExpression(F.Formula, vars1)
	if err != nil {
		s = err.Error()
	} else {
		fmt.Printf("func(a) = %f \n", fxa)
		fxb, err := model.EvaluateExpression(F.Formula, vars2)
		if err != nil {
			s = err.Error()
		} else {
			fmt.Printf("func(b) = %f \n", fxb)
			s = strconv.FormatFloat(math.Abs(F.BorderB-F.BorderA)*(math.Abs(fxb)+math.Abs(fxa)/2), 'f', 6, 32)
		}
	}

	fmt.Println(s)

	data := ViewData{
		Title: "Ответ ",
		Text:  s,
	}

	tmpl, _ := template.ParseFiles("static/templates/index.html")
	tmpl.Execute(w, data)
}
