package ui

import (
	"2026_Knjiga-recepata/models"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func placeholderScreen(w fyne.Window, title string, recipes []*models.Recipe, ii *models.InvertedIndex) {

	back := widget.NewButtonWithIcon(
		"",
		theme.NavigateBackIcon(),
		func() {
			ShowHome(w, recipes, ii)
		},
	)

	back.Resize(fyne.NewSize(60, 40))

	label := widget.NewLabel(title)

	content := container.NewBorder(
		container.NewHBox(back),
		nil,
		nil,
		nil,
		container.NewCenter(label),
	)

	w.SetContent(content)
}

func makeCard(
	text string,
	size float32,
	click func(),
) fyne.CanvasObject {

	bg := canvas.NewRectangle(
		color.RGBA{230, 230, 230, 255},
	)
	bg.CornerRadius = 20

	btn := widget.NewButton("", click)

	label := canvas.NewText(text, color.Black)
	label.TextSize = size

	return container.NewStack(
		bg,
		btn,
		container.NewCenter(label),
	)
}

func ShowHome(w fyne.Window, recipes []*models.Recipe, ii *models.InvertedIndex) {

	search := widget.NewEntry()
	search.SetPlaceHolder("Pretrazi")

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

	search.Resize(fyne.NewSize(540, 60))

	topCard := makeCard(
		"Svi recepti",
		28,
		func() {
			ShowAllRecipes(w, recipes, ii)
		},
	)

	bottomCard := makeCard(
		"Biranje namirnica",
		18,
		func() {
			placeholderScreen(
				w,
				"Namirnice",
				recipes,
				ii,
			)
		},
	)

	topWrapper := container.NewGridWrap(
		fyne.NewSize(640, 380),
		topCard,
	)

	bottomWrapper := container.NewGridWrap(
		fyne.NewSize(640, 380),
		bottomCard,
	)

	addBtn := widget.NewButton("+", func() {
		placeholderScreen(
			w,
			"Dodavanje recepta",
			recipes,
			ii,
		)
	})

	addBtn.Resize(fyne.NewSize(170, 170))

	content := container.NewVBox(
		search,
		layout.NewSpacer(),
		topWrapper,
		layout.NewSpacer(),
		bottomWrapper,
	)

	floating := container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			addBtn,
		),
	)

	bg := canvas.NewRectangle(
		color.RGBA{245, 245, 245, 255},
	)

	w.SetContent(
		container.NewStack(
			bg,
			container.NewPadded(content),
			container.NewPadded(floating),
		),
	)
}
