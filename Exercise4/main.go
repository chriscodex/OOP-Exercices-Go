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

type precFinal interface {
	precioFinal() float64
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

func newLavadora3(c float64) lavadora {
	e1 := newElectrodometico()
	l := lavadora{}
	l.precioBase = e1.precioBase
	l.color = e1.color
	l.consumoEnergetico = e1.consumoEnergetico
	l.peso = e1.peso
	l.carga = c
	return l
}

// methods of Lavadora

func (l *lavadora) getCarga() float64 {
	return l.carga
}

func (l *lavadora) precioFinal() float64 {
	e1 := electrodomestico{
		precioBase:        l.precioBase,
		color:             l.color,
		consumoEnergetico: l.consumoEnergetico,
		peso:              l.peso,
	}
	pfl := e1.precioFinal()
	if l.carga >= 30 {
		pfl += 50
		return pfl
	}
	return pfl
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
	pfl := func() float64 {
		pfl := float64(e.getPrecioBase())
		switch e.getConsumoEnergetico() {
		case "A":
			pfl += 100
			return float64(pfl)
		case "B":
			pfl += 80
			return float64(pfl)
		case "C":
			pfl += 60
			return float64(pfl)
		case "D":
			pfl += 50
			return float64(pfl)
		case "E":
			pfl += 30
			return float64(pfl)
		case "F":
			pfl += 10
			return float64(pfl)
		default:
			return float64(pfl)
		}
	}()
	p := e.getPeso()
	switch {
	case p >= 0 && p <= 19:
		pfl += 10
		return pfl
	case p >= 20 && p <= 49:
		pfl += 50
		return pfl
	case p >= 50 && p <= 79:
		pfl += 80
		return pfl
	case p >= 80 && p <= 100:
		pfl += 100
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
	l1 := newLavadora3(32.0)

	fmt.Println(l1.precioFinal())
}
