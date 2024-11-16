package scenes

import (
	"fmt"
	"parking/src/models"
	"parking/src/views"
)

// MainScene representa la escena principal de la aplicación y actúa como observador
type MainScene struct {
	Parking *models.Estacionamiento
	View    *views.ParkingView
}

// NewMainScene crea y configura la escena principal
func NewMainScene(estacionamiento *models.Estacionamiento, view *views.ParkingView) *MainScene {
	scene := &MainScene{
		Parking: estacionamiento,
		View:    view,
	}
	estacionamiento.RegistrarObservador(scene)
	return scene
}

// OnVehiculoEntra maneja el evento de entrada de un vehículo
func (s *MainScene) OnVehiculoEntra(id, cajon, espaciosDisponibles, capacidad int) {
	fmt.Printf("Evento: Carro %d entró. Cajón: %d. Espacios disponibles: %d/%d\n",
		id, cajon, espaciosDisponibles, capacidad)
	// Actualizar la vista
	s.View.UpdateState(espaciosDisponibles, capacidad)
}

// OnVehiculoSale maneja el evento de salida de un vehículo
func (s *MainScene) OnVehiculoSale(id, cajon, espaciosDisponibles, capacidad int) {
	fmt.Printf("Evento: Carro %d salió. Cajón: %d. Espacios disponibles: %d/%d\n",
		id, cajon, espaciosDisponibles, capacidad)
	// Actualizar la vista
	s.View.UpdateState(espaciosDisponibles, capacidad)
}
