package libros

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"virtualbloc/common"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	functions := map[string]common.Function{
		http.MethodGet:  GetLibros,
		http.MethodPost: SaveLibro,
	}

	common.Handler(w, r, functions)
}

func GetLibros(con *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := con.Query("call GetLibros()")
	if err != nil {
		http.Error(w, "Error al consultar", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	listaLibros := []Libro{}

	for rows.Next() {
		var l Libro
		if err := rows.Scan(&l.ID, &l.Nombre, &l.Color, &l.Icono, &l.Hojas); err != nil {
			http.Error(w, "Error al leer los datos", http.StatusInternalServerError)
			return
		}
		listaLibros = append(listaLibros, l)
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(listaLibros); err != nil {
		http.Error(w, "Error al generar JSON", http.StatusInternalServerError)
		return
	}
}

func SaveLibro(con *sql.DB, w http.ResponseWriter, r *http.Request) {
	var l Libro

	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if _, err := con.Exec("call SaveLibro(?, ?, ?, ?, ?)", l.ID, l.Nombre, l.Color, l.Icono, l.Hojas); err != nil {
		http.Error(w, "Error al guardar el libro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
