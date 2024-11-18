package views

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ParkingView representa la vista del estacionamiento
type ParkingView struct {
	Container    *fyne.Container
	Label        *widget.Label
	parkingSpots []fyne.Position
}

// NewParkingView crea una vista del estacionamiento
func NewParkingView() *ParkingView {
	// Crear el label para los espacios disponibles
	label := widget.NewLabel("Espacios disponibles: 20/20")

	// Crear la imagen de fondo
	backgroundImage := canvas.NewImageFromFile("src/assets/parking-image.png")
	backgroundImage.FillMode = canvas.ImageFillStretch

	// Contenedor para el estacionamiento
	overlay := container.NewWithoutLayout()
	overlay.Resize(fyne.NewSize(800, 600)) // Ajusta el tamaño del overlay
	overlay.Move(fyne.NewPos(0, 0))        // Posición inicial

	// Ajustar las posiciones para que estén alineadas correctamente
	parkingSpots := []fyne.Position{}
	xStart, yTop, yBottom := float32(185), float32(120), float32(200) // Valores ajustados
	spotSpacing := float32(55)                                        // Separación de los cajones

	// Fila superior
	for i := 0; i < 10; i++ {
		parkingSpots = append(parkingSpots, fyne.Position{X: xStart + float32(i)*spotSpacing, Y: yTop})
	}

	// Fila inferior
	for i := 0; i < 10; i++ {
		parkingSpots = append(parkingSpots, fyne.Position{X: xStart + float32(i)*spotSpacing, Y: yBottom})
	}

	for _, spot := range parkingSpots {
		rect := canvas.NewRectangle(color.Black) // Cicla los colores
		rect.Resize(fyne.NewSize(30, 30))        // Tamaño del cajón
		rect.Move(spot)                          // Posición del cajón
		overlay.Add(rect)
	}

	// Crear el contenedor principal
	container := container.NewMax(
		backgroundImage, // Imagen de fondo
		overlay,         // Contenedor para los cajones
		label,           // Label encima de todo
	)

	return &ParkingView{
		Container:    container,
		Label:        label,
		parkingSpots: parkingSpots,
	}
}

// UpdateState actualiza el estado mostrado en la vista
func (v *ParkingView) UpdateState(espaciosDisponibles, capacidad int, id, cajon int, accion string) {
	// Actualizar el texto del label
	v.Label.SetText("Espacios disponibles: " + strconv.Itoa(espaciosDisponibles) + "/" + strconv.Itoa(capacidad))
	// También puedes imprimir en consola si lo necesitas
	fmt.Println("Evento desde interfaz: Carro", id, "entró. Cajón:", cajon, "Espacios disponibles:", espaciosDisponibles, "/", capacidad)
}
