package structs

type Persona struct {
	Nombre string
	Edad   int
	Sexo   string
}

type Estudiante struct {
	Persona
	Calificacion float32
}

type Profesor struct {
	Persona
	MateriaD string
}

type Aula struct {
	Id     int
	MaxEst int
	CursoD string
}
