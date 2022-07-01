package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	comp := lettersAndNumbers()
	pass := ""
	for i := 0; i < p.longitud; i++ {
		first := rand.Intn(3)
		rMay := rand.Intn(25)
		rMin := rand.Intn(25)
		rNum := rand.Intn(9)
		switch first {
		case 0:
			pass += string(comp[0][rMay])
		case 1:
			pass += string(comp[1][rMin])
		case 2:
			pass += string(comp[2][rNum])
		}
	}
	p.contrasena = pass
}

func (p *password) esFuerte() bool {
	comp := lettersAndNumbers()
	cMay := 0
	cMin := 0
	cNum := 0
	for j := 0; j < len(p.contrasena); j++ {
		for i := 0; i < len(comp[0]); i++ {
			if p.contrasena[j] == comp[0][i] {
				cMay += 1
			}
		}
		for i := 0; i < len(comp[1]); i++ {
			if p.contrasena[j] == comp[1][i] {
				cMin += 1
			}
		}
		for i := 0; i < len(comp[2]); i++ {
			if p.contrasena[j] == comp[2][i] {
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

func crearArrayPasswords(inp int) []password {
	t := []password{}
	for i := 0; i < inp; i++ {
		var s int
		fmt.Printf("Indique la longitud del password %d\n", i+1)
		fmt.Scanln(&s)
		t = append(t, password{
			longitud: s,
		})
		t[i].generarPassword()
	}
	return t
}

func lettersAndNumbers() []string {
	may := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	min := strings.ToLower(may)
	num := "123456789"
	return []string{may, min, num}
}

func 

func main() {
	var inp int
	fmt.Println("Indique el tamano del array")
	fmt.Scanln(&inp)
	t := crearArrayPasswords(inp)
	fmt.Println(t)
	b := []bool{}
	for _, i := range t {
		b = append(b, i.esFuerte())
	}
	fmt.Println(b)
	// p := newPassword1(8)
	// p.generarPassword()
	// p.contrasena = "AAB12345a"
	// fmt.Println(p.esFuerte())
}
