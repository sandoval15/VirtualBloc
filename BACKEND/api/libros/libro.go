package libros

type Libro struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Color  string `json:"color"`
	Icono  string `json:"icono"`
	Hojas  int    `json:"hojas"`
}
