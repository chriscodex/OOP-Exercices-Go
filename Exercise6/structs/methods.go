package structs

// Getters
func (l *Libro) getISBN() string {
	return l.ISBN
}
func (l *Libro) getTitulo() string {
	return l.Titulo
}
func (l *Libro) getAutor() string {
	return l.Autor
}
func (l *Libro) getNPaginas() int {
	return l.NPag
}
