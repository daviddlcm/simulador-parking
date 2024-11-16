package models

type Observer interface {
	OnVehiculoEntra(id, cajon, espaciosDisponibles, capacidad int)
	OnVehiculoSale(id, cajon, espaciosDisponibles, capacidad int)
}
