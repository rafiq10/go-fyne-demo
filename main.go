package main

import "fyne.io/fyne/v2/app"

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Hello world")
	window.ShowAndRun()
}
