package main

import (
	st "e4/structs"
	"fmt"
)

// Exercise 4 from https://www.discoduroderoer.es/ejercicios-propuestos-y-resueltos-programacion-orientado-a-objetos-java/

func main() {
	arrayLavadora := []st.Lavadora{
		st.NewLavadora(),
		st.NewLavadora2(300, 20),
		st.NewLavadora2(200, 10),
	}
	arrayTelevisor := []st.Televisor{
		st.NewTelevisor(),
		st.NewTelevisor2(250, 35),
		st.NewTelevisor(),
	}
	c := make(chan float64)
	go func() {
		var suml float64
		for i, e := range arrayLavadora {
			fmt.Printf("El precio final de la lavadora %d es: %v\n", i+1, e.PrecioFinal())
			suml += e.PrecioFinal()
		}
		c <- suml
	}()
	go func() {
		var sumt float64
		for i, e := range arrayTelevisor {
			fmt.Printf("El precio final del televisor %d es: %v\n", i+1, e.PrecioFinal())
			sumt += e.PrecioFinal()
		}
		c <- sumt
	}()
	total := <-c
	total += <-c
	fmt.Printf("El precio de todos los electrodomésticos es: %v", total)
}
