package main

import (
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

func main() {
	// Crear la aplicación Fyne
	myApp := app.New()
	mainWindow := myApp.NewWindow("Simulador de Parking")

	// Configuración inicial del estacionamiento y la vista
	estacionamiento := models.NewEstacionamiento(20)
	parkingView := views.NewParkingView()

	// Crear la escena principal
	mainScene := scenes.NewMainScene(estacionamiento, parkingView)

	// Configurar el contenido de la ventana con la vista
	content := container.NewMax(parkingView.Container)
	mainWindow.SetContent(content)

	// Ejecutar la simulación de carros en segundo plano
	numCarros := 100
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	go func() {
		for i := 1; i <= numCarros; i++ {
			wg.Add(1)
			go funcionCarro(i, estacionamiento, &wg)
			time.Sleep(time.Millisecond * 500) // Retraso entre entradas de carros
		}
		wg.Wait()
	}()

	// Mostrar la interfaz gráfica
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
	_ = mainScene
}
