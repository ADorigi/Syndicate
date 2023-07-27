package fyneapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/adorigi/syndicate/fyneapp/typedefs"
)

type App struct {
	app    fyne.App
	window fyne.Window
	// images map[string]canvas.Image
}

// func OnCardTap(app fyne.App) {
// 	app.NewWindow("New")
// }

func (a *App) InitializeApp() {
	a.app = app.New()
	a.window = a.app.NewWindow("Syndicate")
	// a.window.SetContent(widget.NewLabel("Hello World!"))
	a.window.Resize(fyne.NewSize(400, 400))
	// a.window.SetContent(widget.NewEntry())
	// a.images["cpu"] = *canvas.NewImageFromFile("/fyneApp/satic/cpu.png")
	// card := widget.NewCard("CPU", "CPU info", a.window.Canvas().Content())

	image := canvas.NewImageFromFile("./fyneApp/static/cpu.png")
	image.Resize(fyne.NewSize(100, 100))
	// image.FillMode = canvas.ImageFillContain

	card := typedefs.NewTapCard("CPU", "CPU info", a.window.Canvas().Content(), func() {
		app := app.New()
		window := app.NewWindow("New")
		window.SetContent(widget.NewLabel("Hello World!"))
		window.Show()
		// fmt.Println("Card tapped")
	})
	card.SetImage(image)

	a.window.SetContent(card)

	// image := canvas.NewImageFromFile("/fyneApp/static/cpu.png")
	// image := canvas.NewImageFromFile("./fyneApp/static/cpu.png")

	// image.FillMode = canvas.ImageFillOriginal
	// a.window.SetContent(image)
}

func (a *App) Run() {
	a.window.ShowAndRun()
}

func NewAppData() *App {
	return &App{}
}
