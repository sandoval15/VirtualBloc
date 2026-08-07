package main

import (
	"fmt"
	"net/http"
	"virtualbloc/api/hojas"
	"virtualbloc/api/libros"
)

func main() {
	http.HandleFunc("/libros", libros.Handler)
	http.HandleFunc("/hojas", hojas.Handler)

	puerto := ":8080"
	fmt.Printf("Servidor corriendo en http://localhost%s\n", puerto)

	err := http.ListenAndServe("localhost"+puerto, nil)
	if err != nil {
		fmt.Printf("Error al arrancar el servidor: %v\n", err)
	}
}
