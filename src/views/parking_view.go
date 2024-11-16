package views

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ParkingView representa la vista del estacionamiento
type ParkingView struct {
	Container *fyne.Container
	Label     *widget.Label
}

// NewParkingView crea una vista del estacionamiento
func NewParkingView() *ParkingView {
	// Crear el label para los espacios disponibles
	label := widget.NewLabel("Espacios disponibles: 20/20")
	// Crear la imagen de fondo
	backgroundImage := canvas.NewImageFromFile("src/assets/parking-image.png")
	backgroundImage.FillMode = canvas.ImageFillStretch

	// Usar layout para asegurar que la imagen ocupe todo el espacio
	// y colocar los elementos sobre ella (como el label)
	container := container.NewMax(backgroundImage) // Este layout asegura que la imagen ocupe toda la pantalla
	container.Add(label)                           // Agregar el label encima de la imagen

	// Crear y retornar el ParkingView
	return &ParkingView{
		Container: container,
		Label:     label,
	}
}

// UpdateState actualiza el estado mostrado en la vista
func (v *ParkingView) UpdateState(espaciosDisponibles, capacidad int) {
	// Actualizar el texto del label
	v.Label.SetText("Espacios disponibles: " + strconv.Itoa(espaciosDisponibles) + "/" + strconv.Itoa(capacidad))
	// También puedes imprimir en consola si lo necesitas
	fmt.Println("Espacios disponibles: " + strconv.Itoa(espaciosDisponibles) + "/" + strconv.Itoa(capacidad))
}
