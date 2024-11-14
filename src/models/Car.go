package models

import (
	"math/rand"
	"time"
)

func SimularCarro(id int, estacionamiento *Estacionamiento) {
	estacionamiento.VehiculoEntra(id) // Tiempo aleatorio entre 3 y 5 segundos dentro del estacionamiento
	tiempoEstacionado := time.Duration(3+rand.Intn(3)) * time.Second
	//fmt.Printf("Carro %d estacionado por %v segundos\n", id, tiempoEstacionado.Seconds())
	time.Sleep(tiempoEstacionado)
	estacionamiento.VehiculoSale(id)
}
