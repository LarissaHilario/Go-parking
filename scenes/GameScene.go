package scenes

import (
	"estacionamiento/models"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameScene struct {
	window fyne.Window
}
const (
	capacidad    = 3
	numVehiculos = 10
)

var (
	espaciosEstacionamiento = make(chan struct{}, capacidad)
	entrada                  = make(chan struct{}, 1)
	wg                       sync.WaitGroup
	vehicles                 = make([]*models.Vehicle, 0)
	startVehicleCreation = make(chan bool) 
)

func StartVehicleCreation() {
	for _, vehicle := range vehicles {
        vehicle.Position = fyne.NewPos(100, 100)
    }
    startVehicleCreation <- true // Envia una señal para iniciar la creación de vehículos
}

func NewGameScene(fyneWindow fyne.Window) *GameScene {
	sceneGame := &GameScene{window: fyneWindow}
	sceneGame.RenderGame()
	return sceneGame
}


func (s*GameScene) RenderGame() {
	background := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.png"))
	background.Resize(fyne.NewSize(701,500))
	background.Move(fyne.NewPos(-5,0))

	btnBackMenu:= widget.NewButton("Salir", s.BackMenu)
	btnBackMenu.Resize(fyne.NewSize(130,30))
	btnBackMenu.Move(fyne.NewPos(470,50))
	
	
	vehicleLayer := container.NewWithoutLayout()

	for _, vehicle := range vehicles {
		if vehicleLayer.Visible() {
			vehicleLayer.Add(vehicle.Image)
		}
		
		
	}
	
	

	go func() {
        <-startVehicleCreation // Espera la señal para iniciar la creación de vehículos
        for i := 0; i < numVehiculos; i++ {
            wg.Add(1)
            vehicle := models.NewVehicle(i)
            vehicles = append(vehicles, vehicle)
            go vehicleLlega(vehicle)
        }
    }()

	s.window.SetContent(container.NewWithoutLayout(background, btnBackMenu, vehicleLayer))
}
func (s *GameScene) BackMenu() {
	NewMenuScene(s.window)
	
}
func vehicleLlega(vehicle *models.Vehicle) {
    fmt.Printf("El vehículo %d ha llegado.\n", vehicle.ID)

    entrada <- struct{}{}

    espaciosEstacionamiento <- struct{}{}
    fmt.Printf("El vehículo %d está entrando al estacionamiento.\n", vehicle.ID)

    // Verifica si el vehículo ya está estacionado en la posición inicial (100, 100)
    if vehicle.Position != fyne.NewPos(100, 100) {
        // Encuentra una nueva posición de estacionamiento
        for i := 0; i < capacidad; i++ {
            if vehicles[i].Position == fyne.NewPos(100, 100) {
                vehicles[i].Position = fyne.NewPos(100+float32(i*120), 250)
                break
            }
        }
    }

    <-entrada

    fmt.Printf("El vehículo %d está estacionado en la posición %v.\n", vehicle.ID, vehicle.Position)

    time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)

    <-espaciosEstacionamiento
    fmt.Printf("El vehículo %d está saliendo del estacionamiento.\n", vehicle.ID)

    // Libera la posición
    vehicle.Position = fyne.NewPos(100, 100) // Regresa a la posición inicial

    wg.Done()
}
