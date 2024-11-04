package main

import (
	"fmt"
	"math/rand"
	"parking/models"
	"sync"
	"time"
)

func main() {
	numCarros := 10
	estacionamiento := models.NewEstacionamiento(5)
	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= numCarros; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			models.SimularCarro(id, estacionamiento)
		}(i)
	}

	wg.Wait()
	fmt.Println("Simulación completada")
}
