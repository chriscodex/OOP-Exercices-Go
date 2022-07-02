package structs

func NewSerie() Serie {
	s := Serie{
		NTemporadas: 3,
		Entregado:   false,
	}
	return s
}

func NewSerie2(t string, c string) Serie {
	s1 := NewSerie()
	s := Serie{
		Titulo:      t,
		NTemporadas: s1.NTemporadas,
		Entregado:   s1.Entregado,
		Genero:      s1.Genero,
		Creador:     c,
	}
	return s
}

func NewSerie3(t string, nt int, e bool, g string, c string) Serie {
	s := Serie{
		Titulo:      t,
		NTemporadas: nt,
		Entregado:   e,
		Genero:      g,
		Creador:     c,
	}
	return s
}
