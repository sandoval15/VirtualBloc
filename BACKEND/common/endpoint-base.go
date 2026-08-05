package common

import (
	"database/sql"
	"net/http"
	"virtualbloc/db"
)

type Function func(*sql.DB, http.ResponseWriter, *http.Request)

func Handler(w http.ResponseWriter, r *http.Request, f map[string]Function) {
	w.Header().Set("Content-Type", "application/json")
	con, err := db.GetConnection()
	if err != nil {
		http.Error(w, `"Error de conexión"`, http.StatusInternalServerError)
		return
	}

	if fn, exists := f[r.Method]; exists {
		fn(con, w, r)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
