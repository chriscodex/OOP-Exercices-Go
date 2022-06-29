package main

import "fmt"

// Exercise 1 from https://www.discoduroderoer.es/ejercicios-propuestos-y-resueltos-programacion-orientado-a-objetos-java/

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

// Special methods
func (c *Cuenta) ingresar(cant float64) {
	if cant > 0 {
		c.cantidad += cant
	}
}

func (c *Cuenta) retirar(cant float64) {
	if c.cantidad-cant < 0 {
		fmt.Println("Error, Está intentando hacer un retiro mayor a su saldo")
	} else {
		c.cantidad -= cant
		fmt.Println("Saldo retirado con éxito")
	}
}

func main() {
	miCuenta := NewCuenta("Christian", 50.0)
	fmt.Println(*miCuenta)
	fmt.Println(miCuenta.getTitular())
	miCuenta.setTitular("Peter")
	miCuenta.setCantidad(1.25)
	fmt.Println(miCuenta.toString())
	miCuenta.ingresar(30)
	fmt.Println(miCuenta.toString())
	miCuenta.retirar(18)
	fmt.Println(miCuenta.getCantidad())
}
