
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
	capacidad    = 5
	numVehiculos = 20
)

var (
	espaciosEstacionamiento = make(chan struct{}, capacidad)
	entrada                  = make(chan struct{}, 1)
	wg                       sync.WaitGroup
	vehicles                 = make([]*models.Vehicle, 0)
	startVehicleCreation = make(chan bool) 
    coordenadasEstacionamiento = []fyne.Position{
        {X: 890, Y: 350},
        {X: 985, Y: 350},
        {X: 1086, Y: 350},
        {X: 1190, Y: 350},
        {X: 1300, Y: 350},
        {X: 1390, Y: 350},
        {X: 1500, Y: 350},
        {X: 1590, Y: 350},
        {X: 1680, Y: 350},
        {X: 1780, Y: 350},

        {X: 890, Y: 650},
        {X: 1031, Y: 650},
        {X: 1127, Y: 650},
        {X: 1221, Y: 650},
        {X: 1310, Y: 650},
        {X: 1402, Y: 650},
        {X: 1510, Y: 650},
        {X: 1600, Y: 650},
        {X: 1692, Y: 650},
        {X: 1790, Y: 650},


       
    }   
)

func StartVehicleCreation() {
	for _, vehicle := range vehicles {
        vehicle.Position = fyne.NewPos(100, 100)
    }
    startVehicleCreation <- true
}

func NewGameScene(fyneWindow fyne.Window) *GameScene {
	sceneGame := &GameScene{window: fyneWindow}
	sceneGame.RenderGame()
	return sceneGame
}

// En tu función RenderGame en GameScene

func (s *GameScene) RenderGame() {
    background := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.png"))
    background.Resize(fyne.NewSize(1920,1080))
    background.Move(fyne.NewPos(-5, 0))

    btnBackMenu := widget.NewButton("Salir", s.BackMenu)
    btnBackMenu.Resize(fyne.NewSize(130, 30))
    btnBackMenu.Move(fyne.NewPos(470, 50))

    // Contenedor para los vehículos
    vehicleContainer := container.NewWithoutLayout()

    s.window.SetContent(container.NewWithoutLayout(background, btnBackMenu, vehicleContainer))

    go func() {
        <-startVehicleCreation // Espera la señal para iniciar la creación de vehículos
        for i := 0; i < numVehiculos; i++ {
            wg.Add(1)
            vehicle := models.NewVehicle(i)

          
           
            vehicles = append(vehicles, vehicle)
            go vehicleLlega(vehicle)
    
            // Agregar la imagen del vehículo al contenedor y establecer su posición
            vehicleContainer.Add(vehicle.Image)
            canvas.Refresh(vehicleContainer) // Asegura que el contenedor se actualice
             
        }
    }()
    
}

func (s *GameScene) BackMenu() {
	NewMenuScene(s.window)
	
}

func vehicleLlega(vehicle *models.Vehicle) {
    fmt.Printf("El vehículo %d ha llegado.\n", vehicle.ID)
    

    entrada <- struct{}{}

    espaciosEstacionamiento <- struct{}{}
    fmt.Printf("El vehículo %d está entrando al estacionamiento.\n", vehicle.ID)

    // Verifica si el slice tiene al menos dos elementos
    if len(vehicles) >= 2 {
        // Encuentra una nueva posición de estacionamiento
        for i := 0; i < len(vehicles)-1; i++ {
            if vehicles[i].Position == fyne.NewPos(100, 100) {
                vehicles[i].Position = (coordenadasEstacionamiento[i])
                vehicle.Image.Move(vehicles[i].Position)

                break
            }
        }
    }

    
    
    <-entrada

    fmt.Printf("El vehículo %d está estacionado en la posición %v.\n", vehicle.ID, vehicle.Position)

    time.Sleep(time.Duration(1+rand.Intn(20)) * time.Second)

    <-espaciosEstacionamiento
    fmt.Printf("El vehículo %d está saliendo del estacionamiento.\n", vehicle.ID)

    // Libera la posición
    vehicle.Position = fyne.NewPos(100, 400) // Regresa a la posición inicial
    vehicle.Image.Move(vehicle.Position)
    wg.Done()
}