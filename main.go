package main

import (
	"fmt"
	"net/http"

	mux "github.com/gorilla/mux"

	handlers "rgr/handlers"
	model "rgr/model"
)

func runServer() {
	f := model.Function{
		BorderA: 0,
		BorderB: 1,
		Formula: "x",
	}
	pointnumber := 1
	model.PointNumbers = &pointnumber

	handlers.F = &f

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.MainPage)
	router.HandleFunc("/input", handlers.Input)
	router.HandleFunc("/output", handlers.Output)

	//не работают стили - проверить надобнА!
	rmux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static/"))
	rmux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//

	http.Handle("/", router)
	fmt.Println("Server is listening...")

	http.ListenAndServe("localhost:8181", nil)
}

func main() {
	runServer()
}
