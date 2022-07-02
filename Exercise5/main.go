package main

import (
	st "ex5/structs"
	"fmt"
)

func main() {
	se := st.NewSerie3("The boys", 3, true, "action", "some1")
	vj := st.NewVideojuego()
	vj.Devolver()
	se.Devolver()
	se.IsEntregado()
	fmt.Println(vj.Entregado)
	fmt.Println(se.Entregado)
}
