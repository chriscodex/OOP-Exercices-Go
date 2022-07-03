package main

import (
	st "ex5/structs"
	"fmt"
)

func main() {
	arraySeries := [5]st.Serie{
		st.NewSerie3("The boys", 12, true, "action", "unknown"),
		st.NewSerie3("Breaking Bad", 12, true, "action", "unknown"),
		st.NewSerie(),
		st.NewSerie(),
		st.NewSerie(),
	}
	arrayVideojuegos := [5]st.Videojuego{
		st.NewVideojuego(),
		st.NewVideoJuego2("Mario Bros", 36),
		st.NewVideoJuego2("WC3", 36),
		st.NewVideojuego(),
		st.NewVideojuego(),
	}
	arraySeries[0].Entregar()
	arrayVideojuegos[1].Entregar()
	fmt.Println("Series Entregadas")
	var smay int
	var vmay float32
	cs := 0
	cv := 0
	for i, e := range arraySeries {
		if e.NTemporadas > smay {
			smay = e.NTemporadas
		}
		if e.Entregado {
			fmt.Println(i + 1)
			fmt.Println(e)
		}
	}
	fmt.Println("\nSeries con más temporadas")
	for _, e := range arraySeries {
		if e.NTemporadas == smay {
			cs += 1
			fmt.Println(cs)
			fmt.Println(e)
		}
	}
	fmt.Println("\nVideojuegos Entregados")
	for i, e := range arrayVideojuegos {
		if e.HorasEstimadas > vmay {
			vmay = e.HorasEstimadas
		}
		if e.Entregado {
			fmt.Println(i + 1)
			fmt.Println(e)
		}
	}
	fmt.Println("Videojuegos con mayor número de horas estimadas")
	for _, e := range arrayVideojuegos {
		if e.HorasEstimadas == vmay {
			cv += 1
			fmt.Println(cv)
			fmt.Println(e)
		}
	}
}
