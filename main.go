package main

import (
	//"fmt"      //formateo de lo que se esta manipulando
	"database/sql"
	"html/template" //separacion de la informacion en templates
	"log"           //monitoreo de terminal
	"net/http"      //webb

	_ "github.com/go-sql-driver/mysql" //conexion con mysql a traves de git
)

func conexionBD() (conexion *sql.DB) { //configuracion BD
	Driver := "mysql"
	Usuario := "root"
	Contraseña := ""
	Nombre := "usuarios"

	conexion, err := sql.Open(Driver, Usuario+":"+Contraseña+"@tcp(127.0.0.1)/"+Nombre) //mensaje error para terminal en caso de fallo de inicio en la BD
	if err != nil {
		panic(err.Error())
	}
	return conexion

}

var plantillas = template.Must(template.ParseGlob("x/*")) //declaracion de las plantillas

func main() {
	http.HandleFunc("/", Inicio)               //Acceder a la funcion inicio
	http.HandleFunc("/formulario", Formulario) //Acceder a lectura del formulario
	http.HandleFunc("/insertar", Insertar)     //Acceder a la escritura del formulario

	log.Println("servidor corriendo...") //mensaje a la terminal

	err := http.ListenAndServe(":8080", nil) //inicializacion del servidor
	if err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %s\n", err) //mensaje de error para terminal en caso de que no se pueda iniciar el servidor
	}
}

func Inicio(w http.ResponseWriter, r *http.Request) { //envio y rercibo de info w y r

	conexionEstablecida := conexionBD()                                                                                                                                                                              //confirmacion conexion BD
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO usuarios(Nombre, Correo, Contraseña, Telefono, Imagen) VALUES ('andrea','alejandra.arismendi1@gmail.com','Andrea14*','04245003970','1010') ") //prueba ingreso de datos a la base de datos desde el codigo
	if err != nil {
		panic(err.Error())
	}
	insertarRegistros.Exec()

	//fmt.Fprintf(w, "dios matenme") //mandando mensaje al navegador //prueba en un principio local host
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func Formulario(w http.ResponseWriter, r *http.Request) { //recibo de informacion en el formulario
	//fmt.Fprintf(w, "dios matenme") //mandando mensaje al navegador
	plantillas.ExecuteTemplate(w, "formulario", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) { //escritura de informacion en el formulario
	if r.Method == "POST" {

		Nombre := r.FormValue("nombre")
		Correo := r.FormValue("correo")
		Contraseña := r.FormValue("contraseña")
		Telefono := r.FormValue("telefono")
		Imagen := r.FormValue("imagen")

		conexionEstablecida := conexionBD()                                                                                                             //confirmacion conexion BD
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO usuarios(Nombre, Correo, Contraseña, Telefono, Imagen) VALUES (?,?,?,?,?) ") //prueba ingreso de datos a la base de datos desde la interfaz
		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(Nombre, Correo, Contraseña, Telefono, Imagen)

		http.Redirect(w, r, "/", 301) //redireccionando a la misma pagina
	}

}
