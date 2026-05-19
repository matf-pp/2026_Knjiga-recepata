package models

func ArrayToSet(recipes []*Recipe) *Set {
	set := NewSet()

	for _, r := range recipes {
		set.Add(r)
	}

	return set
}

func (set *Set) FilterByIngredient(ingredient *Ingredient) {
	ingredientName := ingredient.Name

	for r := range set.elements {
		foundIngredient := false

		for _, i := range r.Ingredients {
			if i.Name == ingredientName {
				foundIngredient = true
				break
			}
		}

		if !foundIngredient {
			set.Remove(r)
		}
	}

}

func FilterByIngredients(ingredients []*Ingredient, recipes []*Recipe) *Set {
	set := ArrayToSet(recipes)

	for _, i := range ingredients {
		set.FilterByIngredient(i)
	}

	return set
}
