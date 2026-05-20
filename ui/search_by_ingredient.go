package ui

import (
	"2026_Knjiga-recepata/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowIngredientSearch(w fyne.Window, recipes []*models.Recipe, ii *models.InvertedIndex) {

	names := ii.IngredientNames()
	selected := make(map[string]bool)

	var checks []fyne.CanvasObject

	for _, name := range names {
		n := name

		check := widget.NewCheck(n, func(v bool) {
			selected[n] = v
		})

		checks = append(checks, check)
	}

	btn := widget.NewButton("Prikaži recepte", func() {

		var ingredients []*models.Ingredient

		for name, ok := range selected {
			if ok {
				ingredients = append(ingredients, &models.Ingredient{
					Name: name,
				})
			}
		}

		result := ii.Filter(ingredients)

		var buttons []fyne.CanvasObject

		for _, r := range result.List() {
			rec := r

			buttons = append(buttons,
				widget.NewButton(rec.Name, func() {
					ShowRecipeDetail(w, rec, recipes, ii)
				}),
			)
		}

		if len(buttons) == 0 {
			buttons = append(buttons, widget.NewLabel("Nema recepata"))
		}

		resultScreen := container.NewVBox(
			widget.NewButton("Nazad", func() {
				ShowIngredientSearch(w, recipes, ii)
			}),
			container.NewVBox(buttons...),
		)

		w.SetContent(container.NewScroll(resultScreen))
	})

	w.SetContent(container.NewScroll(
		container.NewVBox(append(checks, btn)...),
	))
}
