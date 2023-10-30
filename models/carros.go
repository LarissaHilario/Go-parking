package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)
const (
    autoSize = 70
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

type auto struct {
    rectangle *canvas.Image
    position  fyne.Position
}

func Newauto() *auto {
    rand.Seed(time.Now().UnixNano())
    imagePath := imags[rand.Intn(len(imags))]

    auto := &auto{
        rectangle: canvas.NewImageFromURI(storage.NewFileURI(imagePath)),
    }
    auto.rectangle.Resize(fyne.NewSize(autoSize, autoSize))
    
    return auto
}

func (o *auto) MoveTo(x, y float32) {
    o.position = fyne.NewPos(x, y)
    o.rectangle.Move(o.position)
}

func (o *auto) GetRectangle() *canvas.Image {
    return o.rectangle
}
