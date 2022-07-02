package structs

// Constructor of electrodomestico
func NewElectrodomestico() Electrodomestico {
	return Electrodomestico{
		PrecioBase:        100,
		Color:             "Blanco",
		ConsumoEnergetico: "F",
		Peso:              5,
	}
}

func NewElectrodomestico2(precio int, peso float64) Electrodomestico {
	l := NewElectrodomestico()
	l.PrecioBase = precio
	l.Peso = peso
	return l
}

func NewElectrodomestico3(precio int, color string, consumo string, peso float64) Electrodomestico {
	return Electrodomestico{
		PrecioBase:        precio,
		Color:             color,
		ConsumoEnergetico: consumo,
		Peso:              peso,
	}
}

// Constructors of Lavadora

func NewLavadora() Lavadora {
	l := Lavadora{}
	l.PrecioBase = 100
	l.Color = "Blanco"
	l.ConsumoEnergetico = "F"
	l.Peso = 5
	l.Carga = 5
	return l
}

func NewLavadora2(pr int, pe float64) Lavadora {
	l := NewLavadora()
	l.PrecioBase = pr
	l.Peso = pe
	return l
}

func NewLavadora3(c float64) Lavadora {
	l := NewLavadora()
	l.Carga = c
	return l
}

// Constructor of televisor

func NewTelevisor() Televisor {
	e := NewElectrodomestico()
	t := Televisor{}
	t.PrecioBase = e.PrecioBase
	t.Color = e.Color
	t.ConsumoEnergetico = e.ConsumoEnergetico
	t.Peso = e.Peso
	t.Resolucion = 20
	t.Sintonizador = false
	return t
}

func NewTelevisor2(pr int, pe float64) Televisor {
	t := NewTelevisor()
	t.PrecioBase = pr
	t.Peso = pe
	return t
}

func NewTelevisor3(re float64, sint bool) Televisor {
	t := NewTelevisor()
	t.Resolucion = re
	t.Sintonizador = sint
	return t
}
