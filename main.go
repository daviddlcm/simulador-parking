package main

import (
	"math"
	"math/rand"
	"parking/src/models"
	"parking/src/scenes"
	"parking/src/views"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func funcionCarro(id int, estacionamiento *models.Estacionamiento, wg *sync.WaitGroup) {
	defer wg.Done()
	models.SimularCarro(id, estacionamiento)
}

func tiempoEntreLlegadasPoisson(lambda float64) time.Duration {
	u := rand.Float64()
	return time.Duration(-math.Log(1-u)/lambda) * time.Second
}

func generarVehiculos(numCarros int, lambda float64, estacionamiento *models.Estacionamiento, wg *sync.WaitGroup) {
	for i := 1; i <= numCarros; i++ {
		wg.Add(1)
		go funcionCarro(i, estacionamiento, wg)
		time.Sleep(tiempoEntreLlegadasPoisson(lambda))
	}
	wg.Wait()
}

func main() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("Simulador de Parking")

	estacionamiento := models.NewEstacionamiento(20)
	parkingView := views.NewParkingView()

	mainScene := scenes.NewMainScene(estacionamiento, parkingView)

	content := container.NewMax(parkingView.Container)
	mainWindow.SetContent(content)

	numCarros := 100
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	go generarVehiculos(numCarros, 1.0, estacionamiento, &wg)

	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
	_ = mainScene
}
