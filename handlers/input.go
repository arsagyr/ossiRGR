package handlers

import (
	"fmt"
	"log"
	"net/http"
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
		arguement := r.FormValue("arguement")
		fmt.Printf("X is %s \n", arguement)

		F.Formula = formula
		F.Arguement, err = strconv.ParseFloat(arguement, 64)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/output", 301)
	} else {
		http.ServeFile(w, r, "static/templates/input.html")
	}
}
