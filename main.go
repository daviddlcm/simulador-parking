package main

import (
	"fmt"
	"parking/models"
	"sync"
	"time"
)

func main() {
	numCarros := 10
	estacionamiento := models.NewEstacionamiento(5)
	var wg sync.WaitGroup

	for i := 1; i <= numCarros; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			estacionamiento.VehiculoEntra(id)
			time.Sleep(1 * time.Second)
			estacionamiento.VehiculoSale(id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Simulación completada")
}
