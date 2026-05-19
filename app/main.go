package main

import (
	"2026_Knjiga-recepata/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	w := a.NewWindow("Knjiga recepata")
	w.Resize(fyne.NewSize(600, 800))

	ui.ShowHome(w)

	w.ShowAndRun()
}
