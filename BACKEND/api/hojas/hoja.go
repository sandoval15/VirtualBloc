package hojas

type Hoja struct {
	ID     int    `json:"id"`
	Libro  int    `json:"libro"`
	Numero int    `json:"numero"`
	Texto  string `json:"texto"`
}
