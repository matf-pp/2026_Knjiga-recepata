package ui

import (
	"2026_Knjiga-recepata/models"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func placeholderScreen(w fyne.Window, title string, recipes []*models.Recipe, ii *models.InvertedIndex) {
	label := widget.NewLabel(title)

	content := container.NewBorder(makeBack(w, func() {
		ShowHome(w, recipes, ii)
	}), nil, nil, nil, container.NewCenter(label))

	w.SetContent(content)
}

func makeCard(
	text string,
	size float32,
	click func(),
) fyne.CanvasObject {

	bg := canvas.NewRectangle(color.RGBA{230, 245, 255, 255})
	bg.CornerRadius = 30

	btn := widget.NewButton("", click)

	label := canvas.NewText(text, color.White)
	label.TextSize = size
	label.TextStyle.Bold = true

	return container.NewStack(
		bg,
		btn,
		container.NewCenter(label),
	)
}

func ShowHome(w fyne.Window, recipes []*models.Recipe, ii *models.InvertedIndex) {

	search := widget.NewEntry()
	search.SetPlaceHolder("Pretraga")

	makeImageButton := func(imgPath, title string, click func()) fyne.CanvasObject {
		img := canvas.NewImageFromFile(imgPath)
		img.FillMode = canvas.ImageFillCover

		overlay := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})

		label := canvas.NewText(title, color.White)
		label.TextStyle.Bold = true
		label.TextSize = 32

		btn := widget.NewButton("", click)
		btn.Importance = widget.LowImportance

		return container.NewStack(
			img,
			container.NewCenter(label),
			overlay,
			btn,
		)
	}

	search.OnSubmitted = func(text string) {
		rec := SearchRecipes(text, recipes)

		var buttons []fyne.CanvasObject

		// Pravimo dugmice za sve pronadjene recepte
		for _, r := range rec {
			rec := r

			btn := widget.NewButton(rec.Name, func() {
				ShowRecipeDetail(w, rec, recipes, ii)
			})

			buttons = append(buttons, btn)
		}

		// Nismo pronasli ni jedan recept, pa ispisujemo poruku
		if len(buttons) == 0 {
			buttons = append(buttons, widget.NewLabel("Nema rezultata"))
		}

		// Pravimo novi ekran
		content := container.NewVBox(
			makeBack(w, func() {
				ShowHome(w, recipes, ii)
			}),
			// buttons... raspakuje niz manje vise
			container.NewVBox(buttons...),
		)

		// Menjamo ceo UI
		w.SetContent(container.NewScroll(content))
	}

	topCard := makeImageButton(
		"images/all_recipes.jpg",
		"Svi recepti",
		func() {
			ShowAllRecipes(w, recipes, ii)
		},
	)

	bottomCard := makeImageButton(
		"images/ingredients.jpg",
		"Biranje namirnica",
		func() {
			ShowIngredientSearch(
				w,
				recipes,
				ii,
			)
		},
	)

	// velicina dugmica se menja menjanjem velicine prozora
	grid := container.NewGridWithRows(2, topCard, bottomCard)

	content := container.NewBorder(
		search,
		nil,
		nil,
		nil,
		container.NewPadded(grid),
	)

	w.SetContent(content)
}
