package handlers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"rgr/model"
	"strconv"
)

func Input(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		formula := r.FormValue("formula")
		fmt.Printf("Formula is : %s \n", formula)
		borderA := r.FormValue("borderA")
		fmt.Printf("a is %s \n", borderA)
		borderB := r.FormValue("borderB")
		fmt.Printf("b is %s \n", borderB)
		points := r.FormValue("points")
		fmt.Printf("Number of points is %s \n", points)
		F.Formula = formula
		F.BorderA, err = strconv.ParseFloat(borderA, 64)
		if err != nil {
			log.Println(err)
		}
		F.BorderB, err = strconv.ParseFloat(borderB, 64)
		if err != nil {
			log.Println(err)
		}
		m, err := strconv.Atoi(points)
		if err != nil {
			log.Println(err)
		}
		n := uint(math.Abs(float64(m)))
		model.PointNumbers = &n
		http.Redirect(w, r, "/output", 301)
	} else {
		http.ServeFile(w, r, "static/templates/input.html")
	}
}
