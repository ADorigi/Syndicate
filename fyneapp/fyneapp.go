package fyneapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/adorigi/syndicate/systeminfo"
)

type App struct {
	app      fyne.App
	window   map[string]fyne.Window
	info     systeminfo.AllInfo
	appTheme fyne.Theme

	// windowContent map[string]fyne.CanvasObject
}

func NewApp() *App {
	return &App{}
}

func (a *App) Show(window string) {
	a.window[window].Show()
}

func (a *App) Run() {
	a.app.Run()
}

func (a *App) InitializeApp() {
	a.app = app.New()
	a.window = make(map[string]fyne.Window)
	a.info = systeminfo.CollectStats()
	a.appTheme = &myTheme{}

	// a.windowContent = make(map[string]fyne.CanvasObject)

	a.app.Settings().SetTheme(a.appTheme)

	a.AddWindow("Main", "Syndicate")
	a.SetMainWindowContent()
}

func (a *App) AddWindow(key, title string) {
	a.window[key] = a.app.NewWindow(title)
}

func (a *App) SetMainWindowContent() {
	content := container.New(
		layout.NewGridLayout(2),
		a.newCard("./static/cpu.svg", "CPU", "CPU info"),
		a.newCard("./static/disk.svg", "Disk", "Disk info"),
		a.newCard("./static/network.svg", "Network", "Network info"),
	)

	a.window["Main"].SetContent(content)
	// a.window["Main"].SetContent(a.newCardWindow("./static/cpu.png", "CPU", "CPU info"))
}

func (a *App) SetCardWindowContent(window string) {

	cpuData := [][]string{
		{"CPU", a.info.CpuName},
	}

	cpuTable := widget.NewTable(
		func() (int, int) {
			return len(cpuData), len(cpuData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(cpuData[i.Row][i.Col])
			// if i.Col == 0 {
			// 	o.(*widget.Label).SetText(cpuData[i.Row][i.Col])
			// }

		})

	// diskUsed := widget.NewProgressBar()
	// a.window[window].Resize(fyne.NewSize(100, 100))

	cpuTable.Resize(fyne.NewSize(100, 100))

	// content := container.NewHBox(cpuTable)
	// content.Resize(fyne.NewSize(500, 500))
	a.window[window].SetContent(cpuTable)

	// a.window[window].SetContent(content)
}

func (a *App) newCard(imagePath, title, subTitle string) fyne.Widget {
	image := canvas.NewImageFromFile(imagePath)
	image.FillMode = canvas.ImageFillContain
	// image.Resize(fyne.NewSize(50, 50))
	image.SetMinSize(fyne.NewSize(50, 50))

	// card := NewTapCard(title, subTitle, a.window["Main"].Canvas().Content(), func() {
	// 	a.onCardTap(title)
	// })
	// card := NewTapCard(title, subTitle, image, func() {
	// 	a.onCardTap(title)
	// })

	titlec := canvas.NewText(title, theme.ForegroundColor())
	titlec.TextSize = 20

	card := widget.NewCard("", "", container.New(
		layout.NewVBoxLayout(),
		image,
		container.New(layout.NewCenterLayout(), widget.NewButton(title, func() {
			a.onCardTap(title)
		}))))
	// image.Resize(fyne.NewSize(50, 50))
	// card.SetImage(image)

	card.Resize(fyne.NewSize(200, 200))

	return card
}

func (a *App) onCardTap(title string) {

	a.AddWindow(title, title)
	a.SetCardWindowContent(title)
	a.window[title].Resize(fyne.NewSize(200, 100))
	a.window[title].Show()
}
