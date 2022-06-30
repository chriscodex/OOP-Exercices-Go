package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

type password struct {
	longitud   int
	contrasena string
}

// Constructor
func newPassword1(long int) *password {
	l := math.Pow(10, float64(long))
	c := strconv.Itoa(rand.Intn(int(l)))
	return &password{
		longitud:   long,
		contrasena: c,
	}
}

func newPassword2(longitud int, contrasena string) *password {
	return &password{
		longitud:   longitud,
		contrasena: contrasena,
	}
}

// Methods
func (p *password) generarPassword() {
	l := math.Pow(10, float64(p.longitud))
	p.contrasena = strconv.Itoa(rand.Intn(int(l)))
}

// Getters
func (p *password) getLongitud() int {
	return p.longitud
}

func (p *password) getContrasena() string {
	return p.contrasena
}

// Setters
func (p *password) setLongitud(l int) {
	p.longitud = l
}

func main() {
	var inp int
	fmt.Println("Indique el tamano del array")
	fmt.Scanln(&inp)
	t := []password{}
	for i := 0; i < inp; i++ {
		var s int
		fmt.Scanln(&s)
		t = append(t, password{
			longitud: s,
		})
		fmt.Println(t[i])
	}
}
