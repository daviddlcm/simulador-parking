package models

import (
	"math/rand"
	"time"
)

func SimularCarro(id int, estacionamiento *Estacionamiento) {
	estacionamiento.VehiculoEntra(id)
	tiempoEstacionado := time.Duration(3+rand.Intn(3)) * time.Second
	//fmt.Printf("Carro %d estacionado por %v segundos\n", id, tiempoEstacionado.Seconds())
	time.Sleep(tiempoEstacionado)
	estacionamiento.VehiculoSale(id)
}
