package structs

import "fmt"

type Electrodomestico struct {
	PrecioBase        int
	Color             string
	ConsumoEnergetico string
	Peso              float64
}

type Lavadora struct {
	Electrodomestico
	Carga float64
}

type Televisor struct {
	Electrodomestico
	Resolucion   float64
	Sintonizador bool
}

func Hello() {
	fmt.Println("hello")
}
