package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
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

// Functions

func esFuerte(c string) bool {
	may := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	min := strings.ToLower(may)
	num := "123456789"
	cMay := 0
	cMin := 0
	cNum := 0
	for j := 0; j < len(c); j++ {
		for i := 0; i < len(may); i++ {
			if c[j] == may[i] {
				cMay += 1
			}
		}
		for i := 0; i < len(min); i++ {
			if c[j] == min[i] {
				cMin += 1
			}
		}
		for i := 0; i < len(num); i++ {
			if c[j] == num[i] {
				cNum += 1
			}
		}
	}
	if cMay >= 2 && cMin >= 1 && cNum >= 5 {
		return true
	} else {
		return false
	}
}

func crearArray(inp int) []password {
	t := []password{}
	for i := 0; i < inp; i++ {
		var s int
		fmt.Printf("Indique la longitud del password %d\n", i+1)
		fmt.Scanln(&s)
		t = append(t, password{
			longitud: s,
		})
		t[i].generarPassword()
		fmt.Println(t[i])
	}
	return t
}

func main() {
	// var inp int
	// fmt.Println("Indique el tamano del array")
	// fmt.Scanln(&inp)
	// t := crearArray(inp)
	// fmt.Println(t)
	fmt.Println(esFuerte("AZu8675443"))
}
