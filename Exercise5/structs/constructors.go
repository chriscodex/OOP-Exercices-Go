package structs

func NewSerie() Serie {
	return Serie{
		NTemporadas: 3,
	}
}

func NewSerie2(t string, c string) Serie {
	s1 := NewSerie()
	return Serie{
		Titulo:      t,
		NTemporadas: s1.GetNTemporadas(),
		Entregado:   s1.GetEntregado(),
		Genero:      s1.GetGenero(),
		Creador:     c,
	}
}

func NewSerie3(t string, nt int, e bool, g string, c string) Serie {
	return Serie{
		Titulo:      t,
		NTemporadas: nt,
		Entregado:   e,
		Genero:      g,
		Creador:     c,
	}
}

func (v *Videojuego) NewVideojuego() Videojuego {
	return Videojuego{
		HorasEstimadas: 10,
	}
}

func (v *Videojuego) NewVideoJuego2(t string, he float32) Videojuego {
	vj := v.NewVideojuego()
	return Videojuego{
		Titulo:         t,
		HorasEstimadas: he,
		Entregado:      vj.GetEntregado(),
		Genero:         vj.GetGenero(),
		Compania:       vj.GetCompania(),
	}
}

func (v *Videojuego) NewVideoJuego3(t string, he float32, e bool, g string, c string) Videojuego {
	return Videojuego{
		Titulo:         t,
		HorasEstimadas: he,
		Entregado:      e,
		Genero:         g,
		Compania:       c,
	}
}
