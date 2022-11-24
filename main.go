package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Users struct {
	Name   string
	Skills string
	Age    int
}

func Index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")

	user := Users{"Jose Macias", "Estudiante de ASIR", 24}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
	}

}

func main() {

	http.HandleFunc("/", Index)

	//Creación del servidor
	fmt.Println("El servidor esta corriendo en puerto 8081")
	fmt.Println("Run server: http://localhost:8081")
	http.ListenAndServe("localhost:8081", nil)

}
