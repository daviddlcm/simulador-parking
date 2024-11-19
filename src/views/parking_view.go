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
	Overlay      *fyne.Container // Contenedor para manejar los elementos visuales
	Label        *widget.Label
	parkingSpots []fyne.Position
	slotRects    map[int]*canvas.Rectangle // Mapa para rastrear los rectángulos por slot
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
		Overlay:      overlay,
		Label:        label,
		parkingSpots: parkingSpots,
		slotRects:    make(map[int]*canvas.Rectangle), // Inicializar el mapa de slots
	}
}

// DrawRedRectangle dibuja un rectángulo rojo en un slot específico
func (v *ParkingView) DrawRedRectangle(x, y float32, slotID int) {
	// Si ya existe un rectángulo en este slot, no hacer nada
	if _, exists := v.slotRects[slotID]; exists {
		return
	}

	// Crear un rectángulo rojo
	rect := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255}) // Rojo
	rect.Resize(fyne.NewSize(30, 30))                                   // Tamaño del rectángulo
	rect.Move(fyne.NewPos(x, y))                                        // Posición en la pantalla

	// Añadir el rectángulo al contenedor overlay
	v.Overlay.Add(rect)
	v.slotRects[slotID] = rect // Guardar el rectángulo en el mapa
	v.Overlay.Refresh()        // Refrescar la interfaz para mostrar el cambio
}

// RemoveRedRectangle elimina un rectángulo rojo en un slot específico
func (v *ParkingView) RemoveRedRectangle(slotID int) {
	// Validar que el slotID sea válido
	if rect, exists := v.slotRects[slotID]; exists {
		// Remover el rectángulo del contenedor overlay
		v.Overlay.Remove(rect)
		delete(v.slotRects, slotID) // Eliminar el rectángulo del mapa
		v.Overlay.Refresh()         // Refrescar la interfaz para reflejar el cambio
		fmt.Printf("Rectángulo eliminado del slot %d\n", slotID)
	} else {
		fmt.Printf("No se encontró un rectángulo en el slot %d\n", slotID)
	}
}

// UpdateState actualiza el estado mostrado en la vista
func (v *ParkingView) UpdateState(espaciosDisponibles, capacidad int, id, cajon int, accion string) {
	// Actualizar el texto del label
	v.Label.SetText("Espacios disponibles: " + strconv.Itoa(espaciosDisponibles) + "/" + strconv.Itoa(capacidad))
	// También puedes imprimir en consola si lo necesitas
	fmt.Println("Evento desde interfaz: Carro", id, accion, ". Cajón:", cajon, "Espacios disponibles:", espaciosDisponibles, "/", capacidad)

	// Dibujar o eliminar el rectángulo rojo en el slot correspondiente
	if accion == "entra" {
		v.DrawRedRectangle(v.parkingSpots[cajon-1].X, v.parkingSpots[cajon-1].Y, cajon)
	} else if accion == "sale" {
		v.RemoveRedRectangle(cajon)
	}
}
