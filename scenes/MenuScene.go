package scenes

import (
	

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

)

type MenuScene struct {
	window fyne.Window
}

func NewMenuScene(fyneWindow fyne.Window) *MenuScene {
	scene := &MenuScene{window: fyneWindow}
	scene.RenderMenu()
	return scene
}

func (s*MenuScene) RenderMenu() {
	background := canvas.NewImageFromURI(storage.NewFileURI("./assets/background2.png"))
	background.Resize(fyne.NewSize(740,540))
	background.Move(fyne.NewPos(-25,0))

	btnStartGame := widget.NewButton("Iniciar", s.StartGame)
	btnStartGame.Resize(fyne.NewSize(130,30))
	btnStartGame.Move(fyne.NewPos(300,230))

	s.window.SetContent(container.NewWithoutLayout(background, btnStartGame))
}

func (s *MenuScene) StartGame() {
	NewGameScene(s.window)
	StartVehicleCreation()
}
