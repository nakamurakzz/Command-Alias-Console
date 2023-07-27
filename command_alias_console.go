package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// Fyne GUI
func main() {
	run()
}

// Fyne GUI
func run() {
	a := app.New()
	w := a.NewWindow("Command Alias Console")

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
