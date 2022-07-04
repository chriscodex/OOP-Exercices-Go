package main

import (
	st "e8/structs"
	"fmt"
)

func main() {
	// est1 := st.NewEstudiante("Christian",24,"M",10)
	// est2 := st.NewEstudiante("Eduardo",25,"M",8)
	// est3 := st.NewEstudiante("Karolina",28,"F",8)
	// est4 := st.NewEstudiante("Jesus",20,"M",10)
	prof1 := st.NewProfesor("Juan", 44, "M", "matematicas")
	fmt.Println(prof1)

}
