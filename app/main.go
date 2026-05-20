package main

import (
	"2026_Knjiga-recepata/models"
	"2026_Knjiga-recepata/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	w := a.NewWindow("Knjiga recepata")
	w.Resize(fyne.NewSize(600, 800))

	recipes := ui.LoadRecipes()

	ii := models.NewInvertedIndex()
	for _, r := range recipes {
		ii.AddRecipe(r)
	}

	ui.ShowHome(w, recipes, ii)

	w.ShowAndRun()
}
