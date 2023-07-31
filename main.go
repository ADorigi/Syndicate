package main

import (
	"github.com/adorigi/syndicate/fyneapp"
)

// type services struct {
// 	app fyneapp.App
// 	be  systeminfo.AllInfo
// }

// func backe() {

// }

// func frone() {
// 	mainApp := fyneapp.NewApp()
// 	mainApp.InitializeApp()
// 	mainApp.Show("Main")
// 	mainApp.Run()
// }

func main() {
	// info := systeminfo.CollectStats()
	// fmt.Println(info)

	mainApp := fyneapp.NewApp()
	mainApp.InitializeApp()
	mainApp.Show("Main")
	mainApp.Run()

	// myApp := app.New()
	// myWindow := myApp.NewWindow("Box Layout")
	// myWindow.Resize(fyne.NewSize(400, 400))

	// text1 := canvas.NewText("Hello", color.White)
	// text2 := canvas.NewText("There", color.White)
	// text3 := canvas.NewText("(right)", color.White)
	// content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	// text4 := canvas.NewText("centered", color.White)
	// centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	// myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	// myWindow.ShowAndRun()

}

// package main

// import (
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// )

// var data = [][]string{{"top left", "top right"},
// 	{"bottom left", "bottom right"}}

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Table Widget")

// 	list := widget.NewTable(
// 		func() (int, int) {
// 			return len(data), len(data[0])
// 		},
// 		func() fyne.CanvasObject {
// 			return widget.NewLabel("wide content")
// 		},
// 		func(i widget.TableCellID, o fyne.CanvasObject) {
// 			o.(*widget.Label).SetText(data[i.Row][i.Col])
// 		})

// 	myWindow.SetContent(list)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// )

// var data = [][]string{[]string{"top left", "top right"},
// 	[]string{"bottom left", "bottom right"}}

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Table Widget")

// 	list := widget.NewTable(
// 		func() (int, int) {
// 			return len(data), len(data[0])
// 		},
// 		func() fyne.CanvasObject {
// 			return widget.NewLabel("wide content")
// 		},
// 		func(i widget.TableCellID, o fyne.CanvasObject) {
// 			o.(*widget.Label).SetText(data[i.Row][i.Col])
// 		})

// 	myWindow.SetContent(list)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// )

// var data = [][]string{[]string{"top left", "top right"},
// 	[]string{"bottom left", "bottom right"}}

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Table Widget")

// 	list := widget.NewTable(
// 		func() (int, int) {
// 			return len(data), len(data[0])
// 		},
// 		func() fyne.CanvasObject {
// 			return widget.NewLabel("wide content")
// 		},
// 		func(i widget.TableCellID, o fyne.CanvasObject) {
// 			o.(*widget.Label).SetText(data[i.Row][i.Col])
// 		})

// 	list.Resize(fyne.NewSize(100, 100))

// 	myWindow.SetContent(list)
// 	myWindow.ShowAndRun()
// }
