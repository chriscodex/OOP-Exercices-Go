package structs

// Methods of electrodomestico
// Getters
func (e *Electrodomestico) GetPrecioBase() int {
	return e.PrecioBase
}

func (e *Electrodomestico) GetColor() string {
	return e.Color
}

func (e *Electrodomestico) GetConsumoEnergetico() string {
	return e.ConsumoEnergetico
}

func (e *Electrodomestico) GetPeso() float64 {
	return e.Peso
}

// Method itself
func (e *Electrodomestico) ComprobarConsumoEnergetico(let string) {
	if let != e.GetConsumoEnergetico() {
		e.ConsumoEnergetico = "F"
	}
}

func (e *Electrodomestico) ComprobarColor(color string) {
	if color != e.GetColor() {
		e.Color = "Blanco"
	}
}

func (e *Electrodomestico) PrecioFinal() float64 {
	pf := func() float64 {
		pfl := float64(e.GetPrecioBase())
		switch e.GetConsumoEnergetico() {
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
	p := e.GetPeso()
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

// Lavadora
// Getters
func (l *Lavadora) GetCarga() float64 {
	return l.Carga
}

func (l *Lavadora) PrecioFinal() float64 {
	e1 := Electrodomestico{
		PrecioBase:        l.PrecioBase,
		Color:             l.Color,
		ConsumoEnergetico: l.ConsumoEnergetico,
		Peso:              l.Peso,
	}
	pf := e1.PrecioFinal()
	if l.Carga >= 30 {
		pf += 50
		return pf
	}
	return pf
}

// Televisor
// Getters
func (t *Televisor) getResolucion() float64 {
	return t.Resolucion
}
func (t *Televisor) getSintonizador() bool {
	return t.Sintonizador
}

func (t *Televisor) precioFinal() float64 {
	e := NewElectrodomestico3(t.PrecioBase, t.Color, t.ConsumoEnergetico, t.Peso)
	pf := e.PrecioFinal()
	pf1 := func() float64 {
		if t.Resolucion >= 40 {
			return pf * 1.3
		} else {
			return pf
		}
	}()
	if t.Sintonizador {
		return pf1 + float64(50)
	} else {
		return pf1
	}
}
