package structs

// Getters
func (l *Libro) GetISBN() string {
	return l.ISBN
}
func (l *Libro) GetTitulo() string {
	return l.Titulo
}
func (l *Libro) GetAutor() string {
	return l.Autor
}
func (l *Libro) GetNPaginas() int {
	return l.NPag
}

// Setters
func (l *Libro) SetISBN(i string) {
	l.ISBN = i
}
func (l *Libro) SetTitulo(t string) {
	l.Titulo = t
}
func (l *Libro) SetAutor(a string) {
	l.Autor = a
}
func (l *Libro) SetNPaginas(n int) {
	l.NPag = n
}
