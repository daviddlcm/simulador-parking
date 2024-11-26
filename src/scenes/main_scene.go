package scenes

import (
	"parking/src/models"
	"parking/src/views"
)

type MainScene struct {
	Parking *models.Estacionamiento
	View    *views.ParkingView
}

func NewMainScene(estacionamiento *models.Estacionamiento, view *views.ParkingView) *MainScene {
	scene := &MainScene{
		Parking: estacionamiento,
		View:    view,
	}
	estacionamiento.RegistrarObservador(scene)
	return scene
}

func (s *MainScene) OnVehiculoEntra(id, cajon, espaciosDisponibles, capacidad int) {
	// fmt.Printf("Evento: Carro %d entró. Cajón: %d. Espacios disponibles: %d/%d\n",
	// 	id, cajon, espaciosDisponibles, capacidad)
	s.View.UpdateState(espaciosDisponibles, capacidad, id, cajon, "entra")
}

func (s *MainScene) OnVehiculoSale(id, cajon, espaciosDisponibles, capacidad int) {
	// fmt.Printf("Evento: Carro %d salió. Cajón: %d. Espacios disponibles: %d/%d\n",
	// 	id, cajon, espaciosDisponibles, capacidad)
	s.View.UpdateState(espaciosDisponibles, capacidad, id, cajon, "sale")
}
