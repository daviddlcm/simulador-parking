// package main

// import (
// 	"parking/src/scenes"
// )

// func main() {
// 	mainScene := scenes.NewMainScene()
// 	mainScene.Show()
// }

package main

import (
	"fmt"
	"math/rand"
	"parking/src/models"
	"sync"
	"time"
)

func funcionCarro(id int, estacionamiento *models.Estacionamiento, wg *sync.WaitGroup) {
	defer wg.Done()
	models.SimularCarro(id, estacionamiento)
}

func main() {
	numCarros := 10
	estacionamiento := models.NewEstacionamiento(5)
	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= numCarros; i++ {
		wg.Add(1)
		go funcionCarro(i, estacionamiento, &wg)
	}

	wg.Wait()
	fmt.Println("Simulación completada")
}
