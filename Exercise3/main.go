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

// Functions

func esFuerte(c string) bool {
	may := []string{"A", "B", "C", "D"}
	//min := []string{"a","b","c","d"}
	//num := []string{"1","2","3","4"}
	cMay := 0
	for j := 0; j < len(c)-1; j++ {
		for i := 0; i < len(may); i++ {
			if string(c[j]) == may[i] {
				cMay += 1
			}
		}
	}
	if cMay >= 2 {
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
	fmt.Println(esFuerte("aa123"))
}
