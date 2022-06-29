package main

import "fmt"

type persona struct {
	dni    string
	nombre string
	edad   int
	sexo   string
	peso   float32
	altura float32
}

// Constructor por defecto
func NewPersona1(dni string) *persona {
	return &persona{
		dni:  dni,
		sexo: "H",
	}
}

// Constructor 2
func NewPersona2(dni string, nombre string, edad int, sexo string) *persona {
	return &persona{
		dni:    dni,
		nombre: nombre,
		edad:   edad,
		sexo:   sexo,
	}
}

// Constructor 3
func NewPersona3(dni string, nombre string, edad int, sexo string, peso float32, altura float32) *persona {
	return &persona{
		dni:    dni,
		nombre: nombre,
		edad:   edad,
		sexo:   sexo,
		peso:   peso,
		altura: altura,
	}
}

func main() {
	pers1 := NewPersona1("1561")
	pers2 := NewPersona2("74062106", "Christian", 24, "F")
	fmt.Println(*pers2, *pers1)
}
