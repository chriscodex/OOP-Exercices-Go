package structs

import (
	"fmt"
	"math/rand"
	"time"
)

func (p *Profesor) RevisarDisponibilidad() {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(5)
	if a != 0 {
		p.Disponibilidad = true
	} else {
		p.Disponibilidad = false
	}
}

func (c *Clase) AgregarEstudiante(e ...*Estudiante) {
	pr := Profesor{}
	au := Aula{}
	if len(e)+len(c.Estudiantes) <= c.Aula.MaxEst && c.Aula != au && c.Profesor != pr {
		c.Estudiantes = append(c.Estudiantes, e...)
	} else {
		fmt.Println("Error, estudiantes exceden el limite del salon")
		return
	}
}

func (c *Clase) MostrarEstudiantes() []Estudiante {
	est := []Estudiante{}
	for _, e := range c.Estudiantes {
		est = append(est, *e)
	}
	return est
}

func (c *Clase) RevisarClase() bool {
	return len(c.Estudiantes) > c.Aula.MaxEst/2
}

func (c *Clase) MostrarAprobados() []Estudiante {
	apr := []Estudiante{}
	for _, e := range c.Estudiantes {
		if e.Calificacion > 5 {
			apr = append(apr, *e)
		}
	}
	return apr
}
