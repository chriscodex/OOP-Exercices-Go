package structs

import "fmt"

// Serie
// Getters
func (s *Serie) GetTitulo() string {
	return s.Titulo
}
func (s *Serie) GetNTemporadas() int {
	return s.NTemporadas
}
func (s *Serie) GetGenero() string {
	return s.Genero
}
func (s *Serie) GetCreador() string {
	return s.Creador
}

// Setters
func (s *Serie) SetTitulo(t string) {
	s.Titulo = t
}
func (s *Serie) SetNTemporadas(nt int) {
	s.NTemporadas = nt
}
func (s *Serie) SetGenero(g string) {
	s.Genero = g
}
func (s *Serie) SetCreador(c string) {
	s.Creador = c
}

func (s Serie) String() string {
	if s.Entregado {
		return fmt.Sprintf("Titulo de la serie: %s\nNumero de temporadas: %d\nEstado de entrega: Entregado\nGenero: %s\nCreador: %s",
			s.GetTitulo(), s.GetNTemporadas(), s.GetGenero(), s.GetCreador())
	} else {
		return fmt.Sprintf("Titulo de la serie: %s\nNumero de temporadas: %d\nEstado de entrega: No Entregado\nGenero: %s\nCreador: %s",
			s.GetTitulo(), s.GetNTemporadas(), s.GetGenero(), s.GetCreador())
	}
}

// Videojuego
// Getters
func (v *Videojuego) GetTitulo() string {
	return v.Titulo
}
func (v *Videojuego) GetHorasEstimadas() float32 {
	return v.HorasEstimadas
}
func (v *Videojuego) GetGenero() string {
	return v.Genero
}
func (v *Videojuego) GetCompania() string {
	return v.Compania
}

// Setters
func (v *Videojuego) SetTitulo(t string) {
	v.Titulo = t
}
func (v *Videojuego) SetHEstimadas(he float32) {
	v.HorasEstimadas = he
}
func (v *Videojuego) SetGenero(g string) {
	v.Genero = g
}
func (v *Videojuego) SetCompania(c string) {
	v.Compania = c
}

func (v *Videojuego) String() string {
	if v.Entregado {
		return fmt.Sprintf("Titulo: %s\nHoras Estimadas: %v\nEstado de entrega: Entregado\nGenero: %s\nCompania: %s",
			v.GetTitulo(), v.GetHorasEstimadas(), v.GetGenero(), v.GetCompania())
	}
	return fmt.Sprintf("Titulo: %s\nHoras Estimadas: %v\nEstado de entrega: No Entregado\nGenero: %s\nCompania: %s",
		v.GetTitulo(), v.GetHorasEstimadas(), v.GetGenero(), v.GetCompania())
}
