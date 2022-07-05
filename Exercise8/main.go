package main

import (
	st "e8/structs"
	"fmt"
)

func main() {
	// Crear Estudiantes
	est1 := st.NewEstudiante("Christian", 24, "M", 10)
	est2 := st.NewEstudiante("Eduardo", 25, "M", 8)
	est3 := st.NewEstudiante("Karolina", 28, "F", 4)
	est4 := st.NewEstudiante("Jesus", 20, "M", 10)

	// Crear Profesor
	prof1 := st.NewProfesor("Juan", 44, "M", "matematicas")

	// Crear Aula
	aula1 := st.NewAula(1, 5, "matematicas")

	// Crear Clase
	clase1 := st.NewClase(&prof1, &aula1)

	// Agregando Estudiantes
	clase1.AgregarEstudiante(&est1, &est2, &est3, &est4)

	if clase1.RevisarClase() {
		fmt.Printf("La clase de %s puede dictarse con normalidad\n", clase1.Materia)
	} else {
		fmt.Printf("La clase de %s no puede dictarse\n", clase1.Materia)
	}
	aprCla1 := clase1.EstudiantesAprobados()
	fmt.Println("Estudiantes aprobados:")
	for _, e := range aprCla1 {
		fmt.Printf("Estudiante: %s - Promedio: %v\n", e.Nombre, e.Calificacion)
	}
}
