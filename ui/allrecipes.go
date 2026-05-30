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

	titleText := canvas.NewText(recipe.Name, color.Black)
	titleText.TextSize = 28
	titleText.TextStyle.Bold = true

	titleBg := canvas.NewRectangle(color.NRGBA{R: 255, G: 255, B: 255, A: 180})
	titleBg.CornerRadius = 20

	title := container.NewCenter(container.NewStack(titleBg, container.NewPadded(titleText)))

	img := canvas.NewImageFromFile("images/" + recipe.Image)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(500, 250))

	imgContainer := container.NewCenter(img)

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
				txt := canvas.NewText(fmt.Sprintf("%s: %d", ingredient.Name, int(math.Round(q))), color.Black)
				ingredientBox.Add(txt)
			} else {
				txt := canvas.NewText(fmt.Sprintf("%s: %.2f %s", ingredient.Name, math.Round(q*10)/10, unit), color.Black)
				ingredientBox.Add(txt)
			}
		}

		ingredientBox.Refresh()
	}

	updateIngredients()

	var steps []fyne.CanvasObject

	// numerisanje koraka (mozemo da izbacimo?)
	for i, step := range recipe.Steps {
		txt := canvas.NewText(strconv.Itoa(i+1)+". "+step, color.Black)
		steps = append(steps, txt)
	}

	// Pravimo slajder
	servingsLabel := canvas.NewText(fmt.Sprintf("Osobe: %d", base), color.Black)
	servingsLabel.TextStyle.Bold = true

	servingsSlider := widget.NewSlider(1, 20)
	servingsSlider.Step = 1
	servingsSlider.Value = float64(base)
	servingsSlider.OnChanged = func(v float64) {
		target = int32(v)
		servingsLabel.Text = fmt.Sprintf("Osobe: %d", target)
		servingsLabel.Refresh()
		updateIngredients()
	}

	lbl := canvas.NewText("Sastojci:", color.Black)
	lbl.TextStyle.Bold = true

	lbl1 := canvas.NewText("Priprema:", color.Black)
	lbl1.TextStyle.Bold = true

	ingredientsBg := canvas.NewRectangle(color.NRGBA{R: 255, G: 255, B: 255, A: 180})
	ingredientsBg.CornerRadius = 20

	stepsBg := canvas.NewRectangle(color.NRGBA{R: 255, G: 255, B: 255, A: 180})
	stepsBg.CornerRadius = 20

	//ingredientsCard := container.NewCenter(container.NewStack(ingredientsBg, container.NewPadded(container.NewVBox(servingsLabel, servingsSlider, lbl, ingredientBox))))
	ingredientsCard := container.NewStack(ingredientsBg, container.NewPadded(container.NewVBox(servingsLabel, servingsSlider, lbl, ingredientBox)))

	//stepsCard := container.NewCenter(container.NewStack(stepsBg, container.NewPadded(container.NewVBox(lbl1, container.NewVBox(steps...)))))
	stepsCard := container.NewStack(stepsBg, container.NewPadded(container.NewVBox(lbl1, container.NewVBox(steps...))))

	// celi prozor za recept
	content := container.NewVBox(makeBack(w, func() {
		ShowAllRecipes(w, recipes, ii)
	}),
		title,
		container.NewCenter(imgContainer),

		ingredientsCard,
		widget.NewSeparator(),
		stepsCard,
	)

	// skrolovanje
	w.SetContent(withBackground(container.NewScroll(content), "galaxy.jpg"))
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

	w.SetContent(withBackground(container.NewScroll(content), "galaxy.jpg"))
}
