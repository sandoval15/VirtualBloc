package main

import (
	"fmt"
	"net/http"
	"os"
	"virtualbloc/api/hojas"
	"virtualbloc/api/libros"
)

func direccion() string {
	if puerto := os.Getenv("PORT"); puerto != "" {
		return ":" + puerto
	}

	return "localhost:8080"
}

func main() {
	http.HandleFunc("/libros", libros.Handler)
	http.HandleFunc("/hojas", hojas.Handler)

	dir := direccion()
	fmt.Printf("Servidor corriendo en http://%s\n", dir)

	err := http.ListenAndServe(dir, nil)
	if err != nil {
		fmt.Printf("Error al arrancar el servidor: %v\n", err)
	}
}
