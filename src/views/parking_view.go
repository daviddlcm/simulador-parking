package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// ParkingView representa la vista del estacionamiento con solo la imagen de fondo
type ParkingView struct {
	Container *fyne.Container
}

// NewParkingView crea una vista de estacionamiento con solo la imagen de fondo
func NewParkingView() *ParkingView {
	pv := &ParkingView{
		Container: container.NewWithoutLayout(),
	}

	// Cargar la imagen de fondo desde el archivo
	imagePath := "src/assets/parking-image.png"
	backgroundImage := canvas.NewImageFromFile(imagePath)
	backgroundImage.FillMode = canvas.ImageFillStretch
	backgroundImage.Resize(fyne.NewSize(800, 600)) // Tamaño de la imagen como fondo

	// Agregar la imagen de fondo al contenedor
	pv.Container.Add(backgroundImage)

	return pv
}
