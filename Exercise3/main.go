package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Exercise 3 from https://www.discoduroderoer.es/ejercicios-propuestos-y-resueltos-programacion-orientado-a-objetos-java/

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
	maps := lettersAndNumbers()
	pass := ""
	for i := 0; i < p.longitud; i++ {
		first := rand.Intn(3)
		rMay := rand.Intn(25)
		rMin := rand.Intn(25)
		rNum := rand.Intn(9)
		switch first {
		case 0:
			pass += string(maps["may"][rMay])
		case 1:
			pass += string(maps["min"][rMin])
		case 2:
			pass += string(maps["num"][rNum])
		}
	}
	p.contrasena = pass
}

func (p *password) esFuerte() bool {
	maps := lettersAndNumbers()
	cMay := 0
	cMin := 0
	cNum := 0
	for j := 0; j < len(p.contrasena); j++ {
		for i := 0; i < len(maps["may"]); i++ {
			if p.contrasena[j] == maps["may"][i] {
				cMay += 1
			}
		}
		for i := 0; i < len(maps["min"]); i++ {
			if p.contrasena[j] == maps["min"][i] {
				cMin += 1
			}
		}
		for i := 0; i < len(maps["num"]); i++ {
			if p.contrasena[j] == maps["num"][i] {
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

func lettersAndNumbers() map[string]string {
	maps := make(map[string]string)
	maps["may"] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	maps["min"] = strings.ToLower(maps["may"])
	maps["num"] = "123456789"
	return maps
}

// func lettersAndNumbers() []string {
// 	maps := make(map[string]string)
// 	maps["may"] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	maps["min"] = strings.ToLower(maps["may"])
// 	maps["num"] = "123456789"
// 	may := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	min := strings.ToLower(may)
// 	num := "123456789"
// 	return []string{may, min, num}
// }

func crearArrayBool(t []password) []bool {
	b := []bool{}
	for _, i := range t {
		b = append(b, i.esFuerte())
	}
	return b
}

func main() {
	var inp int
	fmt.Println("Indique el tamano del array")
	fmt.Scanln(&inp)
	t := crearArrayPasswords(inp)
	fmt.Println(t)
	b := crearArrayBool(t)
	fmt.Println(b)
	for i := 0; i < len(t); i++ {
		if b[i] {
			fmt.Printf("La contrasena es: %v y es fuerte\n", t[i].getContrasena())
		} else {
			fmt.Printf("La contrasena es: %v y es debil\n", t[i].getContrasena())
		}
	}
}
