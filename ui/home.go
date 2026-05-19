package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func placeholderScreen(w fyne.Window, title string) {

	back := widget.NewButtonWithIcon(
		"",
		theme.NavigateBackIcon(),
		func() {
			ShowHome(w)
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

func ShowHome(w fyne.Window) {

	search := widget.NewEntry()
	search.SetPlaceHolder("Pretrazi")

	search.OnSubmitted = func(text string) {
		placeholderScreen(
			w,
			"Rezultati pretrage za: "+text,
		)
	}

	search.Resize(fyne.NewSize(540, 60))

	topCard := makeCard(
		"Izlista sve recepte",
		28,
		func() {
			placeholderScreen(
				w,
				"Recepti",
			)
		},
	)

	bottomCard := makeCard(
		"Biranje namirnica",
		18,
		func() {
			placeholderScreen(
				w,
				"Namirnice",
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
		)
	})

	addBtn.Resize(fyne.NewSize(370, 370))

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
