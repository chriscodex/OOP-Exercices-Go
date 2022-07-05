package main

import (
	st "e8/structs"
	"fmt"
)

func main() {
	//totalAlumnos := 4

	est1 := st.NewEstudiante("Christian", 24, "M", 10)
	est2 := st.NewEstudiante("Eduardo", 25, "M", 8)
	est3 := st.NewEstudiante("Karolina", 28, "F", 8)
	est4 := st.NewEstudiante("Jesus", 20, "M", 10)

	prof1 := st.NewProfesor("Juan", 44, "M", "matematicas")
	prof2 := st.NewProfesor("Pedro", 45, "'M", "matematicas")

	aula1 := st.NewAula(1, 5, "matematicas")

	fmt.Println("Clase:")
	clase1 := st.NewClase(&prof1, &aula1)
	clase1.AgregarEstudiante(&est1, &est2, &est3, &est4)
	clase2 := st.NewClase(&prof2, &aula1)
	// fmt.Println(clase1.RevisarClase())
	fmt.Println(clase2)
	fmt.Println(clase1.MostrarEstudiantes())
	if clase1.RevisarClase() {
		fmt.Println("La clase 1 puede dictarse con normalidad")
	} else {
		fmt.Println("La clase 1 no puede dictarse")
	}
}
