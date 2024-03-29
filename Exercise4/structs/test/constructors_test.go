package structs

import (
	st "e4/structs"
	"testing"
)

func TestNewElectrodomestico(t *testing.T) {
	tables := []struct {
		n st.Electrodomestico
	}{
		{
			st.Electrodomestico{
				PrecioBase:        100,
				Color:             "Blanco",
				ConsumoEnergetico: "F",
				Peso:              5,
			},
		},
	}
	for _, item := range tables {
		ne := st.NewElectrodomestico()
		if ne != item.n {
			t.Errorf("New electrodomestico created with incorrect values as default")
		}
	}
}
