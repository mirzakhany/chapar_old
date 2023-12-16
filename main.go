package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(&myTheme{})
	mainWindow := createMainWindow(app)
	mainWindow.ShowAndRun()
}
