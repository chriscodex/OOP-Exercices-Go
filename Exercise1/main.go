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

// Setters
func (c *Cuenta) setTitular(t string) {
	c.titular = t
}

func (c *Cuenta) setCantidad(cant float64) {
	c.cantidad = cant
}

// Function toString
func (c *Cuenta) toString() string {
	return fmt.Sprintf("La cuenta de %s tiene %v euros", c.titular, c.cantidad)
}

func main() {
	miCuenta := NewCuenta("Christian", 50.0)
	fmt.Println(*miCuenta)
	fmt.Println(miCuenta.getTitular())
	miCuenta.setTitular("Peter")
	miCuenta.setCantidad(1.25)
	fmt.Println(miCuenta.toString())
}
