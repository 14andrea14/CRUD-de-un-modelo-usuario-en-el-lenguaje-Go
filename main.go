package main

import (
	//"fmt"      //formateo de lo que se esta manipulando
	"html/template"
	"log"      //monitoreo de terminal
	"net/http" //webb
	//separacion de la informacion en templates
)

var plantillas = template.Must(template.ParseGlob("x/*"))

func main() {
	http.HandleFunc("/", Inicio) //Acceder a la funcion inicio
	http.HandleFunc("/formulario", Formulario)

	log.Println("servidor corriendo...") //mensaje a la terminal

	err := http.ListenAndServe(":8080", nil) //inicializacion del servidor
	if err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %s\n", err)
	}
}

func Inicio(w http.ResponseWriter, r *http.Request) { //envio y rercibo de info w y r
	//fmt.Fprintf(w, "dios matenme") //mandando mensaje al navegador
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func Formulario(w http.ResponseWriter, r *http.Request) { //envio y rercibo de info w y r
	//fmt.Fprintf(w, "dios matenme") //mandando mensaje al navegador
	plantillas.ExecuteTemplate(w, "formulario", nil)
}
