package ui

import (
	"2026_Knjiga-recepata/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowIngredientSearch(w fyne.Window, recipes []*models.Recipe, ii *models.InvertedIndex) {

	// Izvlacimo imena namernica
	names := ii.IngredientNames()
	// Mapu za cekirane namernice
	selected := make(map[string]bool)

	var checks []fyne.CanvasObject

	// Pravimo checkboxeve
	for _, name := range names {
		n := name

		check := widget.NewCheck(n, func(v bool) {
			selected[n] = v
		})

		checks = append(checks, check)
	}

	// Prikazuje recepte
	btn := widget.NewButton("Prikaži recepte", func() {

		var ingredients []*models.Ingredient

		// Pravimo niz namirnica koje su cekirane
		for name, ok := range selected {
			if ok {
				ingredients = append(ingredients, &models.Ingredient{
					Name: name,
				})
			}
		}

		// Filtriramo po njima sve recepte
		result := ii.Filter(ingredients)

		var buttons []fyne.CanvasObject

		// Pravimo dugmice za rezultate
		for _, r := range result.List() {
			rec := r

			buttons = append(buttons,
				widget.NewButton(rec.Name, func() {
					ShowRecipeDetail(w, rec, recipes, ii)
				}),
			)
		}

		// Ako nema dugmica ispisuje da nema recepata
		if len(buttons) == 0 {
			buttons = append(buttons, widget.NewLabel("Nema recepata"))
		}

		// Pravi screen
		resultScreen := container.NewVBox(
			widget.NewButton("Nazad", func() {
				ShowIngredientSearch(w, recipes, ii)
			}),
			container.NewVBox(buttons...),
		)

		w.SetContent(container.NewScroll(resultScreen))
	})

	// Menja pocetni ekran
	w.SetContent(container.NewScroll(
		container.NewVBox(append(checks, btn)...),
	))
}
