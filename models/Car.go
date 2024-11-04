package models

import (
	"fmt"
	"sync"
)

type Estacionamiento struct {
	capacidad int
	ocupados  int
	mu        sync.Mutex // Mutex para asegurar exclusión mutua
	cond      *sync.Cond // Condición para manejar espera de vehículos
}

func NewEstacionamiento(capacidad int) *Estacionamiento {
	e := &Estacionamiento{
		capacidad: capacidad,
		ocupados:  0,
	}
	e.cond = sync.NewCond(&e.mu)
	return e
}

func (e *Estacionamiento) VehiculoEntra(id int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for e.ocupados >= e.capacidad {
		fmt.Printf("Carro %d esperando espacio...\n", id)
		e.cond.Wait()
	}

	e.ocupados++
	fmt.Printf("Carro %d estacionado. Espacios ocupados: %d/%d\n", id, e.ocupados, e.capacidad)
}

func (e *Estacionamiento) VehiculoSale(id int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.ocupados--
	fmt.Printf("Carro %d ha salido. Espacios ocupados: %d/%d\n", id, e.ocupados, e.capacidad)
	e.cond.Signal()
}
