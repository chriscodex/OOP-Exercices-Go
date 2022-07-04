package structs

func NewLibro(i string, t string, a string, n int) Libro {
	return Libro{
		ISBN:   i,
		Titulo: t,
		Autor:  a,
		NPag:   n,
	}
}
