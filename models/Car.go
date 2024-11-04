// Car.go
package models

import (
	"time"
)

func SimularCarro(id int, estacionamiento *Estacionamiento) {
	estacionamiento.VehiculoEntra(id)
	time.Sleep(1 * time.Second) // Simulación de tiempo dentro del estacionamiento
	estacionamiento.VehiculoSale(id)
}
