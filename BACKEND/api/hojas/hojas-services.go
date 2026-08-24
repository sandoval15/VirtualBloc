package hojas

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"virtualbloc/common"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	functions := map[string]common.Function{
		http.MethodGet:  GetHojas,
		http.MethodPost: SaveHojas,
	}

	common.Handler(w, r, functions)
}

func GetHojas(con *sql.DB, w http.ResponseWriter, r *http.Request) {
	libro := r.URL.Query().Get("libro")

	rows, err := con.Query("select * from hoja where libro = ?", libro)
	if err != nil {
		http.Error(w, "Error al consultar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	hojas := []Hoja{}

	for rows.Next() {
		var h Hoja
		if err := rows.Scan(&h.ID, &h.Libro, &h.Numero, &h.Texto); err != nil {
			http.Error(w, "Error al leer los datos: "+err.Error(), http.StatusInternalServerError)
			return
		}
		hojas = append(hojas, h)
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(hojas); err != nil {
		http.Error(w, "Error al generar JSON", http.StatusInternalServerError)
		return
	}
}

func SaveHojas(con *sql.DB, w http.ResponseWriter, r *http.Request) {
	var hojas []Hoja

	if err := json.NewDecoder(r.Body).Decode(&hojas); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	t, err := con.Begin()
	if err != nil {
		http.Error(w, "Error al iniciar la transacción", http.StatusInternalServerError)
		return
	}

	defer t.Rollback()

	e, err := t.Prepare("call SaveHoja(?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Error al preparar la consulta", http.StatusInternalServerError)
		return
	}

	defer e.Close()

	for _, h := range hojas {
		if _, err := e.Exec(h.ID, h.Libro, h.Numero, h.Texto); err != nil {
			http.Error(w, "Error al guardar las hojas", http.StatusInternalServerError)
			return
		}
	}

	if err = t.Commit(); err != nil {
		http.Error(w, "Error al confirmar la transacción", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteHojas(con *sql.DB, w http.ResponseWriter, r *http.Request) {
	libro := r.URL.Query().Get("libro")
	indice := r.URL.Query().Get("indice")

	if _, err := con.Exec("delete from hoja where libro = ? and numero >= ?", libro, indice); err != nil {
		http.Error(w, "Error al eliminar las hojas", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
