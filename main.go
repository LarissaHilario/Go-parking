package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	capacidad    = 3
	numVehiculos = 10
)

var (
	espaciosEstacionamiento = make(chan struct{}, capacidad)
	entrada                  = make(chan struct{}, 1)
	wg                       sync.WaitGroup
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= numVehiculos; i++ {
		wg.Add(1)
		go vehiculoLlega(i)
	}

	wg.Wait()
}

func vehiculoLlega(idVehiculo int) {
	fmt.Printf("El vehículo %d ha llegado.\n", idVehiculo)

	entrada <- struct{}{}

	espaciosEstacionamiento <- struct{}{}
	fmt.Printf("El vehículo %d está entrando al estacionamiento.\n", idVehiculo)

	<-entrada

	encontrarEspacioEstacionamiento(idVehiculo)

	fmt.Printf("El vehículo %d está estacionado.\n", idVehiculo)

	time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)

	<-espaciosEstacionamiento
	fmt.Printf("El vehículo %d está saliendo del estacionamiento.\n", idVehiculo)

	wg.Done()
}

func encontrarEspacioEstacionamiento(idVehiculo int) {
	select {
	case espaciosEstacionamiento <- struct{}{}:
		fmt.Printf("El vehículo %d encontró un espacio de estacionamiento.\n", idVehiculo)
	default:
		time.Sleep(100 * time.Millisecond) //espera si no hay espacio
	}
}
