package structs

type Serie struct {
	Titulo      string
	NTemporadas int
	Entregado   bool
	Genero      string
	Creador     string
}

type Videojuego struct {
	Titulo         string
	HorasEstimadas float32
	Entregado      bool
	Genero         string
	Compania       string
}
