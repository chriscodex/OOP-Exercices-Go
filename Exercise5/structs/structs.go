package structs

type Serie struct {
	Titulo        string
	NTemporadas   int
	Entregado     bool
	Genero        string
	Creador       string
	saveEntregado bool
}

type Videojuego struct {
	Titulo         string
	HorasEstimadas float32
	Entregado      bool
	Genero         string
	Compania       string
	saveEntregado  bool
}

type Entregable interface {
	Entregable()
	Devolver()
}
