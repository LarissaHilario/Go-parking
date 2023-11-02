package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"estacionamiento/scenes"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Estacionamiento")

    myWindow.Resize(fyne.NewSize(1050, 750))
	myWindow.CenterOnScreen()

	scenes.NewMenuScene(myWindow)
	myWindow.ShowAndRun()
}