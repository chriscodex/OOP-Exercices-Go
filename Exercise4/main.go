package main

import "fmt"

type electrodomestico struct {
	precioBase        int
	color             string
	consumoEnergetico string
	peso              float64
}

type lavadora struct {
	electrodomestico
	carga float64
}

// constructor electrodomestico
func newElectrodometico() electrodomestico {
	return electrodomestico{
		precioBase:        100,
		color:             "Blanco",
		consumoEnergetico: "F",
		peso:              5,
	}
}

func newElectrodometico2(precio int, peso float64) electrodomestico {
	return electrodomestico{
		precioBase:        precio,
		color:             "Blanco",
		consumoEnergetico: "F",
		peso:              peso,
	}
}

func newElectrodometico3(precio int, color string, consumo string, peso float64) electrodomestico {
	return electrodomestico{
		precioBase:        precio,
		color:             color,
		consumoEnergetico: consumo,
		peso:              peso,
	}
}

// Constructor lavadora

func newLavadora() lavadora {
	l := lavadora{}
	l.precioBase = 100
	l.color = "Blanco"
	l.consumoEnergetico = "F"
	l.peso = 5
	l.carga = 5
	return l
}

func newLavadora2(pr int, pe float64) lavadora {
	l := lavadora{}
	l.precioBase = pr
	l.color = "Blanco"
	l.consumoEnergetico = "F"
	l.peso = pe
	l.carga = 5
	return l
}

// getters

func (e *electrodomestico) getPrecioBase() int {
	return e.precioBase
}

func (e *electrodomestico) getColor() string {
	return e.color
}

func (e *electrodomestico) getConsumoEnergetico() string {
	return e.consumoEnergetico
}

func (e *electrodomestico) getPeso() float64 {
	return e.peso
}

// Methods

func (e *electrodomestico) comprobarConsumoEnergetico(let string) {
	if let != e.getConsumoEnergetico() {
		e.consumoEnergetico = "F"
	}
}

func (e *electrodomestico) comprobarColor(color string) {
	if color != e.getColor() {
		e.color = "Blanco"
	}
}

func (e *electrodomestico) precioFinal() float64 {
	pfl := float64(e.getPrecioBase())
	c1 := make(chan float64)
	go func() {
		switch e.getConsumoEnergetico() {
		case "A":
			pfl += 100
			c1 <- float64(pfl)
		case "B":
			pfl += 80
			c1 <- float64(pfl)
		case "C":
			pfl += 60
			c1 <- float64(pfl)
		case "D":
			pfl += 50
			c1 <- float64(pfl)
		case "E":
			pfl += 30
			c1 <- float64(pfl)
		case "F":
			pfl += 10
			c1 <- float64(pfl)
		default:
			c1 <- float64(pfl)
		}
	}()
	p := e.getPeso()
	switch {
	case p >= 0 && p <= 19:
		pfl = <-c1 + 10
		return pfl
	case p >= 20 && p <= 49:
		pfl = <-c1 + 50
		return pfl
	case p >= 50 && p <= 79:
		pfl = <-c1 + 80
		return pfl
	case p >= 80 && p <= 100:
		pfl = <-c1 + 100
		return pfl
	default:
		return pfl
	}
}

func main() {
	e1 := newElectrodometico()
	e1.peso = 100
	p := e1.precioFinal()
	fmt.Println(p)
}
