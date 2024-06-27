package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type User struct {
    Username string `validate:"required"`
    Email    string `validate:"required,email"`
    Phone    string `validate:"required"`
    Password string `validate:"required"`
}

func main() {
	app := fiber.New()
	// Sirve el archivo index.html desde la carpeta "public"
	app.Static("/", "./public")

	// Escucha en el puerto 3000
	app.Listen(":4000")
	fmt.Println("Servidor en el puerto 4000")
}	