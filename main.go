package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"estacionamiento/scenes"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Estacionamiento")

    myWindow.SetFixedSize(true)
    myWindow.Resize(fyne.NewSize(700, 500))
	myWindow.CenterOnScreen()

	scenes.NewMenuScene(myWindow)
	myWindow.ShowAndRun()
}