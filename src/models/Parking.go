package models

import (
	"fmt"
	"sync"
)

type Estacionamiento struct {
	capacidad       int
	espacios        chan struct{} // Canal para controlar espacios disponibles
	entrada         chan struct{} // Canal para gestionar la puerta de entrada/salida
	mu              sync.Mutex    // Mutex para proteger el acceso a los recursos compartidos
	ocupados        int
	cajonesLibres   []int       // Lista de números de cajones disponibles
	cajonOcupadoPor map[int]int // Mapa de vehículos con sus respectivos números de cajón
}

// NewEstacionamiento inicializa el estacionamiento con una capacidad dada
func NewEstacionamiento(capacidad int) *Estacionamiento {
	espacios := make(chan struct{}, capacidad)
	entrada := make(chan struct{}, 1) // Canal para controlar el acceso exclusivo a la puerta
	// Inicializa el canal de espacios con la capacidad total
	for i := 0; i < capacidad; i++ {
		espacios <- struct{}{}
	}
	// Inicializa la lista de cajones y el mapa de ocupación de cajones
	cajonesLibres := make([]int, capacidad)
	for i := 0; i < capacidad; i++ {
		cajonesLibres[i] = i + 1 // Numeración de cajones del 1 al 'capacidad'
	}
	return &Estacionamiento{
		capacidad:       capacidad,
		espacios:        espacios,
		entrada:         entrada,
		ocupados:        0,
		cajonesLibres:   cajonesLibres,
		cajonOcupadoPor: make(map[int]int),
	}
}

// VehiculoEntra simula la entrada de un vehículo al estacionamiento
func (e *Estacionamiento) VehiculoEntra(id int) {
	// Bloquea hasta que haya un espacio disponible
	<-e.espacios
	// Intenta adquirir el control de la puerta de entrada/salida
	e.entrada <- struct{}{}
	fmt.Printf("Carro %d está entrando al estacionamiento.\n", id)
	// Asigna el primer cajón disponible de forma segura
	e.mu.Lock()
	cajon := e.cajonesLibres[0]
	e.cajonesLibres = e.cajonesLibres[1:] // Remueve el cajón asignado de la lista de libres
	e.cajonOcupadoPor[id] = cajon         // Asocia el vehículo con el número de cajón
	e.ocupados++
	fmt.Printf("Carro %d ha estacionado en el cajón %d. Espacios restantes: %d/%d\n", id, cajon, e.capacidad-e.ocupados, e.capacidad)
	e.mu.Unlock()
	// Libera la puerta para otros vehículos
	<-e.entrada
}

// VehiculoSale simula la salida de un vehículo del estacionamiento
func (e *Estacionamiento) VehiculoSale(id int) {
	// Adquiere el control de la puerta de entrada/salida
	e.entrada <- struct{}{}
	fmt.Printf("Carro %d está saliendo del estacionamiento.\n", id)
	// Libera el cajón ocupado y lo vuelve a poner en la lista de cajones disponibles
	e.mu.Lock()
	cajon, existe := e.cajonOcupadoPor[id]
	if existe {
		delete(e.cajonOcupadoPor, id)                    // Remueve el vehículo del mapa
		e.cajonesLibres = append(e.cajonesLibres, cajon) // Agrega el cajón a la lista de libres
		e.ocupados--
		fmt.Printf("Carro %d ha salido del cajón %d. Espacios restantes: %d/%d\n", id, cajon, e.capacidad-e.ocupados, e.capacidad)
	} else {
		fmt.Printf("Error: Carro %d no se encontró en el estacionamiento.\n", id)
	}
	e.mu.Unlock()
	// Libera un espacio en el estacionamiento para otro vehículo
	e.espacios <- struct{}{}
	// Libera la puerta
	<-e.entrada
}
