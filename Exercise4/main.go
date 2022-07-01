package main

type electrodomestico struct {
	precioBase        int
	color             string
	consumoEnergetico string
	peso              float64
}

// constructor
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

func main() {
	e1 := newElectrodometico()
	e1.peso = 100
	//p := e1.precioFinal()
	//fmt.Println(p)
}
