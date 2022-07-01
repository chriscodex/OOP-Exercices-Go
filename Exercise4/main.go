package main

import "fmt"

// Class electrodoméstico
type electrodomestico struct {
	precioBase        int
	color             string
	consumoEnergetico string
	peso              float64
}

// Class lavadora
type lavadora struct {
	electrodomestico
	carga float64
}

// Class television
type televisor struct {
	electrodomestico
	resolucion   float64
	sintonizador bool
}

type precFinal interface {
	precioFinal() float64
}

// constructor electrodomestico
func newElectrodomestico() electrodomestico {
	return electrodomestico{
		precioBase:        100,
		color:             "Blanco",
		consumoEnergetico: "F",
		peso:              5,
	}
}

func newElectrodomestico2(precio int, peso float64) electrodomestico {
	l := newElectrodomestico()
	l.precioBase = precio
	l.peso = peso
	return l
}

func newElectrodomestico3(precio int, color string, consumo string, peso float64) electrodomestico {
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
	l := newLavadora()
	l.precioBase = pr
	l.peso = pe
	return l
}

func newLavadora3(c float64) lavadora {
	l := newLavadora()
	l.carga = c
	return l
}

// Constructor television

func newTelevisor() televisor {
	e := newElectrodomestico()
	t := televisor{}
	t.precioBase = e.precioBase
	t.color = e.color
	t.consumoEnergetico = e.consumoEnergetico
	t.peso = e.peso
	t.resolucion = 20
	t.sintonizador = false
	return t
}

func newTelevisor2(pr int, pe float64) televisor {
	t := newTelevisor()
	t.precioBase = pr
	t.peso = pe
	return t
}

func newTelevisor3(re float64, sint bool) televisor {
	t := newTelevisor()
	t.resolucion = re
	t.sintonizador = sint
	return t
}

// methods of televisor
// getters
func (t *televisor) getResolucion() float64 {
	return t.resolucion
}
func (t *televisor) getSintonizador() bool {
	return t.sintonizador
}

func (t *televisor) precioFinal() float64 {
	e := newElectrodomestico3(t.precioBase, t.color, t.consumoEnergetico, t.peso)
	pf := e.precioFinal()
	pf1 := func() float64 {
		if t.resolucion >= 40 {
			return pf * 1.3
		} else {
			return pf
		}
	}()
	if t.sintonizador {
		return pf1 + float64(50)
	} else {
		return pf1
	}
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
	pf := e1.precioFinal()
	if l.carga >= 30 {
		pf += 50
		return pf
	}
	return pf
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
	pf := func() float64 {
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
		pf += 10
		return pf
	case p >= 20 && p <= 49:
		pf += 50
		return pf
	case p >= 50 && p <= 79:
		pf += 80
		return pf
	case p >= 80 && p <= 100:
		pf += 100
		return pf
	default:
		return pf
	}
}

func main() {
	e1 := newElectrodomestico()
	e1.peso = 100
	p := e1.precioFinal()
	fmt.Println(p)
	l1 := newLavadora3(32.0)
	fmt.Println(l1.precioFinal())
	t1 := newTelevisor3(50, true)
	fmt.Println(t1.precioFinal())
	//
	//tam = 10
	arrayElectrodomesticos := []electrodomestico{}
	arrayElectrodomesticos = append(arrayElectrodomesticos, l1.electrodomestico)
	fmt.Println(arrayElectrodomesticos)

}
