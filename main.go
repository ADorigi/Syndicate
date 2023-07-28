package main

import "github.com/adorigi/syndicate/fyneapp"

func main() {
	// info := systemInfo.CollectStats()
	// fmt.Println(info)

	mainApp := fyneapp.NewAppData()
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
