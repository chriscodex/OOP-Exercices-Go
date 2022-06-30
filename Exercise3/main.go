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

func main() {
	p := newPassword1(5)
	fmt.Println(*p)
}
