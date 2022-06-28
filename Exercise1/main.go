package main

import "fmt"

type Cuenta struct {
	titular  string
	cantidad float64
}

// Constructor 1
func NewCuenta(titular string, cantidad ...float64) *Cuenta {
	if len(cantidad) == 0 {
		return &Cuenta{
			titular:  titular,
			cantidad: 0.0,
		}
	}
	return &Cuenta{
		titular:  titular,
		cantidad: cantidad[0],
	}
}

// Getters
func (c *Cuenta) getTitular() string {
	return c.titular
}

func (c *Cuenta) getCantidad() float64 {
	return c.cantidad
}

func main() {
	miCuenta := NewCuenta("Christian", 50.0)
	fmt.Println(*miCuenta)
	fmt.Println(miCuenta.getTitular())
}
