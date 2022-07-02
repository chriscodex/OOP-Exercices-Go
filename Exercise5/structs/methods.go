package structs

import "fmt"

// Getters
func (s *Serie) GetTitulo() string {
	return s.Titulo
}

func (s *Serie) GetNTemporadas() int {
	return s.NTemporadas
}

func (s *Serie) GetEntregado() bool {
	return s.Entregado
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

func (s *Serie) SetEntregado(e bool) {
	s.Entregado = e
}

func (s *Serie) SetGenero(g string) {
	s.Genero = g
}

func (s *Serie) SetCreador(c string) {
	s.Creador = c
}

func (s Serie) String() string {
	if s.GetEntregado() {
		return fmt.Sprintf("Titulo de la serie: %s\nNumero de temporadas: %d\nEstado de entrega: Entregado\nGenero: %s\nCreador: %s",
			s.GetTitulo(), s.GetNTemporadas(), s.GetGenero(), s.GetCreador())
	} else {
		return fmt.Sprintf("Titulo de la serie: %s\nNumero de temporadas: %d\nEstado de entrega: No Entregado\nGenero: %s\nCreador: %s",
			s.GetTitulo(), s.GetNTemporadas(), s.GetGenero(), s.GetCreador())
	}
}
