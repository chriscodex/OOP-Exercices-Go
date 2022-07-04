package main

import (
	st "e6/structs"
	"fmt"
)

func main() {
	l1 := st.NewLibro("9780590353403", "Harry Potter y la piedra filosofal", "J. K. Rowling", 320)
	l2 := st.NewLibro("9788869185182", "Harry Potter y la cámara secreta", "J. K. Rowling", 251)
	fmt.Println(l1, "\n", l2)
	if l1.GetNPaginas() > l2.GetNPaginas() {
		fmt.Printf("El libro con mayor número de paginas es %s\n", l1.GetTitulo())
	} else {
		fmt.Printf("El libro con mayor número de paginas es %s\n", l2.GetTitulo())
	}
}
