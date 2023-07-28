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

func (a *App) InitializeApp() {
	a.app = app.New()
	a.window = make(map[string]fyne.Window)
	// a.window["Main"] = a.app.NewWindow("Syndicate")
	// a.window["Main"].Resize(fyne.NewSize(400, 400))
	a.AddWindow("Main", "Syndicate")
	// a.window["Main"].Resize(fyne.NewSize(400, 400))
	a.SetMainWindowContent()
}

func (a *App) AddWindow(key, title string) {
	a.window[key] = a.app.NewWindow(title)
}

func (a *App) SetMainWindowContent() {

	// diskImage := canvas.NewImageFromFile("./static/disk.png")
	// diskImage.FillMode = canvas.ImageFillContain

	// diskCard := NewTapCard("Disk", "Disk info", a.window["Main"].Canvas().Content(), func() {
	// 	// app := a.New()
	// 	window := a.app.NewWindow("New")
	// 	window.SetContent(widget.NewLabel("Hello World!"))
	// 	window.Show()
	// })
	// diskCard.SetImage(diskImage)

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

func NewAppData() *App {
	return &App{}
}

func (a *App) newCardWindow(imagePath, title, subTitle string) fyne.Widget {
	image := canvas.NewImageFromFile(imagePath)
	image.FillMode = canvas.ImageFillContain

	card := NewTapCard(title, subTitle, a.window["Main"].Canvas().Content(), func() {
		a.onCardTap(title)
	})
	// func() {

	// 	a.window[title] = a.app.NewWindow(title)

	// 	// remove after you set content
	// 	//
	// 	// window.SetContent(widget.NewLabel("Hello World!"))
	// 	// window.Show()
	// })
	card.SetImage(image)

	return card
}

func (a *App) onCardTap(title string) {

	// app := a.New()
	// a.window[title] = a.app.NewWindow(title)
	// window.SetContent(widget.NewLabel("CPU tapped"))
	// window.Show()

	a.AddWindow(title, title)
}
