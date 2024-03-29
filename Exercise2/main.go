package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// Exercise 2 from https://www.discoduroderoer.es/ejercicios-propuestos-y-resueltos-programacion-orientado-a-objetos-java/

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

// Methods
func (p *persona) calcularIMC() int {
	valor := (p.peso / float32(math.Pow(float64(p.altura), 2)))
	switch {
	case valor < 20:
		return -1
	case valor >= 20 && valor <= 25:
		return 0
	default:
		return 1
	}
}

func (p *persona) esMayorDeEdad() bool {
	if p.edad >= 18 {
		return true
	}
	return false
}

func (p *persona) generarDNI() {
	p.dni = strconv.Itoa(rand.Intn(100000000))
}

// Setters

func main() {
	pers3 := NewPersona3("1231244", "Eduardo", 24, "M", 26, 1)
	fmt.Println(pers3.calcularIMC())
	fmt.Println(pers3.esMayorDeEdad())
	pers3.generarDNI()
	fmt.Println(pers3)
}
