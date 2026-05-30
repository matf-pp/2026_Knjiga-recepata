package ui

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"2026_Knjiga-recepata/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoadRecipes() []*models.Recipe {
	data, err := os.ReadFile("data/recipes.json")
	if err != nil {
		return nil
	}

	// prebacivanje iz json u strukturu
	var recipes []*models.Recipe
	err = json.Unmarshal(data, &recipes)
	if err != nil {
		return nil
	}

	// leksikografsko sortiranje
	sort.Slice(recipes, func(i, j int) bool {
		return recipes[i].Name < recipes[j].Name
	})

	return recipes
}

// dugme za nazad
func makeBack(w fyne.Window, click func()) fyne.CanvasObject {
	back := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), click)

	// stavlja dugme levo
	return container.NewHBox(back)
}

func ShowRecipeDetail(w fyne.Window, recipe *models.Recipe, recipes []*models.Recipe, ii *models.InvertedIndex) {

	title := widget.NewLabel(recipe.Name)

	img := canvas.NewImageFromFile("images/" + recipe.Image)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(500, 250))

	ingredientBox := container.NewVBox()

	// Kolicina sa kojom krecemo i koju zelimo
	base := recipe.Servings
	target := recipe.Servings

	// Funkcija koja azurira kolicinu
	updateIngredients := func() {
		ingredientBox.Objects = nil

		// Koeficijent za menjanje namirnica
		ratio := float64(target) / float64(base)

		for _, ingredient := range recipe.Ingredients {
			// Nova kolicina
			q := ingredient.Quantity * ratio
			unit := strings.ToLower(strings.TrimSpace(ingredient.Unit))

			// Ako nema merne jedinice, nema smisla npr. pola jajeta
			if unit == "" {
				ingredientBox.Add(widget.NewLabel(fmt.Sprintf("%s: %d", ingredient.Name, int(math.Round(q)))))
			} else {
				ingredientBox.Add(widget.NewLabel(fmt.Sprintf("%s: %.2f %s", ingredient.Name, math.Round(q*10)/10, unit)))
			}
		}

		ingredientBox.Refresh()
	}

	updateIngredients()

	var steps []fyne.CanvasObject

	// numerisanje koraka (mozemo da izbacimo?)
	for i, step := range recipe.Steps {
		steps = append(steps, widget.NewLabel(strconv.Itoa(i+1)+". "+step))
	}

	// Pravimo slajder
	servingsLabel := widget.NewLabel(fmt.Sprintf("Osobe: %d", base))

	servingsSlider := widget.NewSlider(1, 20)
	servingsSlider.Step = 1
	servingsSlider.Value = float64(base)
	servingsSlider.OnChanged = func(v float64) {
		target = int32(v)
		servingsLabel.SetText(fmt.Sprintf("Osobe: %d", target))
		updateIngredients()
	}

	lbl := canvas.NewText("Sastojci:", color.White)
	lbl.TextStyle.Bold = true

	lbl1 := canvas.NewText("Priprema:", color.White)
	lbl1.TextStyle.Bold = true

	// celi prozor za recept
	content := container.NewVBox(makeBack(w, func() {
		ShowAllRecipes(w, recipes, ii)
	}),
		title,
		container.NewHBox(img),

		servingsLabel,
		servingsSlider,

		container.NewVBox(lbl, ingredientBox),
		container.NewVBox(lbl1, container.NewVBox(steps...)),
	)

	// skrolovanje
	w.SetContent(container.NewScroll(content))
}

// prikazuje sve recepte kao dugmice
func ShowAllRecipes(w fyne.Window, recipes []*models.Recipe, ii *models.InvertedIndex) {

	var buttons []fyne.CanvasObject

	// pravimo dugme za svaki recept
	for _, recipe := range recipes {

		r := recipe

		// dugme za recepte
		btn := widget.NewButton(r.Name, func() {
			ShowRecipeDetail(w, r, recipes, ii)
		})

		buttons = append(buttons, container.NewPadded(btn))
	}

	if len(buttons) == 0 {
		buttons = append(buttons, widget.NewLabel("Nema recepata."))
	}

	grid := container.NewGridWithColumns(2, buttons...)

	// prozor sa svim receptima
	content := container.NewBorder(makeBack(w, func() {
		ShowHome(w, recipes, ii)
	}), nil, nil, nil, container.NewPadded(grid))
	// dodaje razmak od ivica

	w.SetContent(container.NewScroll(content))
}
