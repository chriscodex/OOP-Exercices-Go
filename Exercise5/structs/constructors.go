package structs

func NewSerie() Serie {
	return Serie{
		NTemporadas: 3,
	}
}

func NewSerie2(t string, c string) Serie {
	s1 := NewSerie()
	return Serie{
		Titulo:        t,
		NTemporadas:   s1.GetNTemporadas(),
		Entregado:     s1.Entregado,
		Genero:        s1.GetGenero(),
		Creador:       c,
		saveEntregado: s1.saveEntregado,
	}
}

func NewSerie3(t string, nt int, e bool, g string, c string) Serie {
	return Serie{
		Titulo:        t,
		NTemporadas:   nt,
		Entregado:     e,
		Genero:        g,
		Creador:       c,
		saveEntregado: e,
	}
}

func NewVideojuego() Videojuego {
	return Videojuego{
		HorasEstimadas: 10,
	}
}

func NewVideoJuego2(t string, he float32) Videojuego {
	vj := NewVideojuego()
	return Videojuego{
		Titulo:         t,
		HorasEstimadas: he,
		Entregado:      vj.Entregado,
		Genero:         vj.GetGenero(),
		Compania:       vj.GetCompania(),
		saveEntregado:  vj.saveEntregado,
	}
}

func NewVideoJuego3(t string, he float32, e bool, g string, c string) Videojuego {
	return Videojuego{
		Titulo:         t,
		HorasEstimadas: he,
		Entregado:      e,
		Genero:         g,
		Compania:       c,
		saveEntregado:  e,
	}
}
