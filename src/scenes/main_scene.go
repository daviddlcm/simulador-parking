package scenes

import (
	"parking/src/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// MainScene representa la escena principal de la aplicación
type MainScene struct {
	Window      fyne.Window
	ParkingView *views.ParkingView
}

// NewMainScene crea y configura la escena principal con la vista del estacionamiento
func NewMainScene() *MainScene {
	a := app.New() // Creamos la aplicación
	win := a.NewWindow("Simulador de Parking")
	parkingView := views.NewParkingView() // Creamos la vista con 20 espacios

	// Contenedor principal
	container := container.NewVBox(parkingView.Container)
	win.SetContent(container)

	return &MainScene{
		Window:      win,
		ParkingView: parkingView,
	}
}

// Show muestra la escena principal
func (s *MainScene) Show() {
	s.Window.Resize(fyne.NewSize(800, 600))
	s.Window.ShowAndRun()
}
