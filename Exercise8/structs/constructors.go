package structs

import "fmt"

func NewEstudiante(n string, e int, s string, c float32) Estudiante {
	est := Estudiante{}
	est.Nombre = n
	est.Edad = e
	est.Sexo = s
	est.Calificacion = c
	return est
}

func validarMateria(m string) bool {
	materias := []string{"matematicas", "filosofia", "fisica"}
	flag := false
	for _, e := range materias {
		if m == e {
			flag = true
		}
	}
	return flag
}

func NewProfesor(n string, e int, s string, m string) Profesor {
	fl := validarMateria(m)
	if fl {
		pro := Profesor{}
		pro.Nombre = n
		pro.Edad = e
		pro.Sexo = s
		pro.MateriaD = m
		pro.Disponibilidad = true
		return pro
	} else {
		fmt.Println("Error, materia no valida")
		return Profesor{}
	}
}

func NewAula(id int, me int, c string) Aula {
	fl := validarMateria(c)
	if fl {
		a := Aula{
			Id:             id,
			MaxEst:         me,
			CursoD:         c,
			Disponibilidad: true,
		}
		return a
	} else {
		fmt.Println("Error, materia no valida")
		return Aula{}
	}
}

func NewClase(p *Profesor, a *Aula) Clase {
	p.RevisarDisponibilidad()
	if p.MateriaD == a.CursoD {
		if p.Disponibilidad {
			if a.Disponibilidad {
				p.Disponibilidad = false
				a.Disponibilidad = false
				return Clase{
					Profesor: *p,
					Aula:     *a,
				}
			} else {
				fmt.Println("Error, el aula se encuentra ocupada")
				return Clase{}
			}
		} else {
			fmt.Println("Error, el profesor no se encuentra disponible")
			return Clase{}
		}
	} else {
		fmt.Println("Error, la materia del docente no coincide con la del aula")
		return Clase{}
	}
}
