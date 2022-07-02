package main

import (
	st "ex5/structs"
	"fmt"
)

func main() {
	arraySeries := [5]st.Serie{
		st.NewSerie3("The boys", 3, true, "action", "some1"),
		st.NewSerie(),
		st.NewSerie(),
		st.NewSerie(),
		st.NewSerie(),
	}
	arrayVideojuegos := [5]st.Videojuego{
		st.NewVideojuego(),
		st.NewVideoJuego2("Mario", 36),
		st.NewVideojuego(),
		st.NewVideojuego(),
		st.NewVideojuego(),
	}
	arraySeries[0].Entregar()
	arrayVideojuegos[1].Entregar()
	fmt.Println("Series")
	for i, e := range arraySeries {
		if arraySeries[i].Entregado {
			fmt.Println(e)
		}
	}
	fmt.Println("Videojuegos")
	for i, e := range arrayVideojuegos {
		if arrayVideojuegos[i].Entregado {
			fmt.Println(e)
		}
	}
}
