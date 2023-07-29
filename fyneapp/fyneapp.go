package fyneapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type App struct {
	app    fyne.App
	window map[string]fyne.Window
}

func NewApp() *App {
	return &App{}
}

func (a *App) InitializeApp() {
	a.app = app.New()
	a.window = make(map[string]fyne.Window)

	a.AddWindow("Main", "Syndicate")
	a.SetMainWindowContent()
}

func (a *App) AddWindow(key, title string) {
	a.window[key] = a.app.NewWindow(title)
}

func (a *App) SetMainWindowContent() {

	content := container.NewHBox(
		a.newCardWindow("./static/cpu.png", "CPU", "CPU info"),
		a.newCardWindow("./static/disk.png", "Disk", "Disk info"),
	)

	a.window["Main"].SetContent(content)
	// a.window["Main"].SetContent(a.newCardWindow("./static/cpu.png", "CPU", "CPU info"))
}

func (a *App) SetCardWindowContent(window string) {

	content := container.New(
		layout.NewHBoxLayout(),
		a.newCardWindow("./static/cpu.png", "CPU", "CPU info"),
		a.newCardWindow("./static/disk.png", "Disk", "Disk info"),
	)

	a.window["Main"].SetContent(content)
}

func (a *App) Show(window string) {
	a.window[window].Show()
}

func (a *App) Run() {
	a.app.Run()
}

func (a *App) newCardWindow(imagePath, title, subTitle string) fyne.Widget {
	image := canvas.NewImageFromFile(imagePath)
	image.FillMode = canvas.ImageFillContain

	card := NewTapCard(title, subTitle, a.window["Main"].Canvas().Content(), func() {
		a.onCardTap(title)
	})
	image.SetMinSize(fyne.NewSize(50, 50))
	card.SetImage(image)
	card.Resize(fyne.NewSize(200, 200))

	return card
}

func (a *App) onCardTap(title string) {

	a.AddWindow(title, title)
	a.Show(title)
}
