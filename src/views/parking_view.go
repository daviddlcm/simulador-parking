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

type ParkingView struct {
	Container    *fyne.Container
	Overlay      *fyne.Container
	Label        *widget.Label
	parkingSpots []fyne.Position
	slotRects    map[int]*canvas.Rectangle
}

func NewParkingView() *ParkingView {
	label := widget.NewLabel("Espacios disponibles: 20/20")

	backgroundImage := canvas.NewImageFromFile("src/assets/parking-image.png")
	backgroundImage.FillMode = canvas.ImageFillStretch

	overlay := container.NewWithoutLayout()
	overlay.Resize(fyne.NewSize(800, 600))

	parkingSpots := []fyne.Position{}
	xStart, yTop, yBottom := float32(185), float32(120), float32(200)
	spotSpacing := float32(55)

	// Fila superior
	for i := 0; i < 10; i++ {
		parkingSpots = append(parkingSpots, fyne.Position{X: xStart + float32(i)*spotSpacing, Y: yTop})
	}

	// Fila inferior
	for i := 0; i < 10; i++ {
		parkingSpots = append(parkingSpots, fyne.Position{X: xStart + float32(i)*spotSpacing, Y: yBottom})
	}

	for _, spot := range parkingSpots {
		rect := canvas.NewRectangle(color.Black)
		rect.Resize(fyne.NewSize(30, 30))
		rect.Move(spot)
		overlay.Add(rect)
	}

	container := container.NewMax(
		backgroundImage,
		overlay,
		label,
	)

	return &ParkingView{
		Container:    container,
		Overlay:      overlay,
		Label:        label,
		parkingSpots: parkingSpots,
		slotRects:    make(map[int]*canvas.Rectangle),
	}
}

func (v *ParkingView) DrawRedRectangle(x, y float32, slotID int) {
	if _, exists := v.slotRects[slotID]; exists {
		return
	}

	rect := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	rect.Resize(fyne.NewSize(30, 30))
	rect.Move(fyne.NewPos(x, y))

	v.Overlay.Add(rect)
	v.slotRects[slotID] = rect
	v.Overlay.Refresh()
}

func (v *ParkingView) RemoveRedRectangle(slotID int) {
	if rect, exists := v.slotRects[slotID]; exists {
		v.Overlay.Remove(rect)
		delete(v.slotRects, slotID)
		v.Overlay.Refresh()
		fmt.Printf("Rectángulo eliminado del slot %d\n", slotID)
	} else {
		fmt.Printf("No se encontró un rectángulo en el slot %d\n", slotID)
	}
}

func (v *ParkingView) UpdateState(espaciosDisponibles, capacidad int, id, cajon int, accion string) {
	v.Label.SetText("Espacios disponibles: " + strconv.Itoa(espaciosDisponibles) + "/" + strconv.Itoa(capacidad))
	fmt.Println("Evento desde interfaz: Carro", id, accion, ". Cajón:", cajon, "Espacios disponibles:", espaciosDisponibles, "/", capacidad)

	if accion == "entra" {
		v.DrawRedRectangle(v.parkingSpots[cajon-1].X, v.parkingSpots[cajon-1].Y, cajon)
	} else if accion == "sale" {
		v.RemoveRedRectangle(cajon)
	}
}
