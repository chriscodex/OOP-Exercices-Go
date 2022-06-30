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

func main() {
	p := newPassword2(8, "12344")
	p.generarPassword()
	fmt.Println(*p)
}
