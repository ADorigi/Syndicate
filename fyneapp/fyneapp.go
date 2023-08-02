package fyneapp

import (
	"strconv"

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
		a.newCard(resourceCpuSvg, "CPU"),
		a.newCard(resourceDiskSvg, "Disk"),
		a.newCard(resourceNetworkSvg, "Network"),
	)

	a.window["Main"].SetContent(content)
	a.window["Main"].CenterOnScreen()
}

func (a *App) SetCardWindowContent(window string) {

	var tableData [][]string
	switch window {
	case "CPU":
		tableData = [][]string{
			{"CPU", a.info.CpuName},
		}
	case "Disk":
		tableData = [][]string{
			{"Disk \nAvailable", strconv.FormatUint(a.info.DiskAvailable, 10) + " GB"},
			{"Disk Used", strconv.FormatUint(a.info.DiskUsed, 10) + " GB"},
			{"%% Used", strconv.FormatUint(a.info.DiskUsed, 10) + " GB"},
		}
	case "Network":
		tableData = [][]string{
			{"Host Name", a.info.Hostname},
			{"Local IPv4", a.info.LocalIPv4},
			{"Global IP", a.info.GlobalIP},
		}
	}

	displayTable := widget.NewTable(
		func() (int, int) {
			return len(tableData), len(tableData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("..............................\n")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {

			o.(*widget.Label).SetText(tableData[i.Row][i.Col])

		})
	displayTable.SetColumnWidth(1, 200)
	// displayTable.Resize(fyne.NewSize(300, 300))
	cntnr := container.New(layout.NewMaxLayout(), displayTable)
	// cntnr.Resize(fyne.NewSize(200, 200))
	a.window[window].SetContent(cntnr)
}

func (a *App) newCard(resource *fyne.StaticResource, title string) fyne.Widget {
	// image := canvas.NewImageFromFile(imagePath)
	image := canvas.NewImageFromResource(resource)
	image.FillMode = canvas.ImageFillContain
	// image.Resize(fyne.NewSize(50, 50))
	image.SetMinSize(fyne.NewSize(50, 50))

	titlec := canvas.NewText(title, theme.ForegroundColor())
	titlec.TextSize = 20

	card := widget.NewCard("", "", container.New(
		layout.NewVBoxLayout(),
		image,
		container.New(layout.NewCenterLayout(), widget.NewButton(title, func() {
			a.onCardTap(title)
		}))))

	card.Resize(fyne.NewSize(200, 200))

	return card
}

func (a *App) onCardTap(title string) {

	a.AddWindow(title, title)
	a.SetCardWindowContent(title)
	a.window[title].Resize(fyne.NewSize(350, 200))
	a.window[title].Show()
}
