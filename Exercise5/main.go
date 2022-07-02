package main

import (
	st "e5/structs"
	"fmt"
)

func main() {
	se := st.NewSerie3("The boys", 3, true, "action", "some1")
	fmt.Println(se.GetEntregado())
}
