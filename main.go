package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// source https://github.com/fyne-io/fyne

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
