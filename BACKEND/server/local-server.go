package main

import (
	"fmt"
	"net/http"
	"virtualbloc/api/hojas"
	"virtualbloc/api/libros"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	html := `
	<!DOCTYPE html>
	<html lang="es">
	<head>
		<meta charset="UTF-8">
		<title>Prueba Local Go</title>
		<style>
			body { font-family: sans-serif; text-align: center; margin-top: 50px; background-color: #f0f2f5; }
			h1 { color: #0070f3; }
		</style>
	</head>
	<body>
		<h1>VALERIA ABSOL ES TONTO</h1>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

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
