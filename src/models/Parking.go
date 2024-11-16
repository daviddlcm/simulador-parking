package models

import (
	"fmt"
	"sync"
)

type Estacionamiento struct {
	capacidad       int
	espacios        chan struct{}
	entrada         chan struct{}
	mu              sync.Mutex
	ocupados        int
	cajonesLibres   []int
	cajonOcupadoPor map[int]int
	observadores    []Observer // Lista de observadores
}

// NewEstacionamiento inicializa el estacionamiento con una capacidad dada
func NewEstacionamiento(capacidad int) *Estacionamiento {
	espacios := make(chan struct{}, capacidad)
	entrada := make(chan struct{}, 1)
	for i := 0; i < capacidad; i++ {
		espacios <- struct{}{}
	}
	cajonesLibres := make([]int, capacidad)
	for i := 0; i < capacidad; i++ {
		cajonesLibres[i] = i + 1
	}
	return &Estacionamiento{
		capacidad:       capacidad,
		espacios:        espacios,
		entrada:         entrada,
		ocupados:        0,
		cajonesLibres:   cajonesLibres,
		cajonOcupadoPor: make(map[int]int),
		observadores:    []Observer{},
	}
}

// RegistrarObservador permite registrar un nuevo observador
func (e *Estacionamiento) RegistrarObservador(o Observer) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.observadores = append(e.observadores, o)
}

// NotificarVehiculoEntra notifica a los observadores cuando un vehículo entra
func (e *Estacionamiento) NotificarVehiculoEntra(id, cajon, espaciosDisponibles, capacidad int) {
	for _, o := range e.observadores {
		o.OnVehiculoEntra(id, cajon, espaciosDisponibles, capacidad)
	}
}

// NotificarVehiculoSale notifica a los observadores cuando un vehículo sale
func (e *Estacionamiento) NotificarVehiculoSale(id, cajon, espaciosDisponibles, capacidad int) {
	for _, o := range e.observadores {
		o.OnVehiculoSale(id, cajon, espaciosDisponibles, capacidad)
	}
}

// VehiculoEntra simula la entrada de un vehículo al estacionamiento
func (e *Estacionamiento) VehiculoEntra(id int) {
	<-e.espacios
	e.entrada <- struct{}{}
	fmt.Printf("Carro %d está entrando al estacionamiento.\n", id)

	e.mu.Lock()
	cajon := e.cajonesLibres[0]
	e.cajonesLibres = e.cajonesLibres[1:]
	e.cajonOcupadoPor[id] = cajon
	e.ocupados++
	e.mu.Unlock()

	espaciosDisponibles := e.capacidad - e.ocupados
	e.NotificarVehiculoEntra(id, cajon, espaciosDisponibles, e.capacidad)
	<-e.entrada
}

// VehiculoSale simula la salida de un vehículo del estacionamiento
func (e *Estacionamiento) VehiculoSale(id int) {
	e.entrada <- struct{}{}
	fmt.Printf("Carro %d está saliendo del estacionamiento.\n", id)

	e.mu.Lock()
	cajon, existe := e.cajonOcupadoPor[id]
	if existe {
		delete(e.cajonOcupadoPor, id)
		e.cajonesLibres = append(e.cajonesLibres, cajon)
		e.ocupados--
	}
	e.mu.Unlock()

	espaciosDisponibles := e.capacidad - e.ocupados
	e.NotificarVehiculoSale(id, cajon, espaciosDisponibles, e.capacidad)
	e.espacios <- struct{}{}
	<-e.entrada
}
