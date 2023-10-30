package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)
const (
    autoSize = 100
    gameWidth    = 300
)

var imags = []string{
    "./assets/auto_verde.png",
    "./assets/auto_naranja.png",
    "./assets/carro_azul.png",
    "./assets/carro_amarillo.png",
    "./assets/carro_rojo.png",
    "./assets/carro_naranja.png",
    "./assets/carro_gris.png",
    "./assets/carro_verde.png",
}

type Vehicle struct {
    ID        int
    Image     *canvas.Image
    Position fyne.Position
}


func NewVehicle(id int ) *Vehicle {
    rand.Seed(time.Now().UnixNano())
    imagePath := imags[rand.Intn(len(imags))]

    Vehicle := &Vehicle{
        ID: id,
        Image: canvas.NewImageFromURI(storage.NewFileURI(imagePath)),
        Position: fyne.NewPos(100,100), 
       
    }
    Vehicle.Image.Resize(fyne.NewSize(100,100))
    return Vehicle
}
