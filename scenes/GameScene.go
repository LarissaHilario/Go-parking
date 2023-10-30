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

// En tu función RenderGame en GameScene

func (s *GameScene) RenderGame() {
    background := canvas.NewImageFromURI(storage.NewFileURI("./assets/background.png"))
    background.Resize(fyne.NewSize(701, 500))
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
    
            // Calcula la posición inicial en función del índice
            x := 100 + float32(i*120) // Ajusta el espaciado entre vehículos
            y := 350 // Ajusta la altura en la que aparecen los vehículos
    
            vehicle.Position = fyne.NewPos(x, float32(y))
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