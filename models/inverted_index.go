package models

import "sort"

// Struktura podataka za brzu pretragu, umesto da za svaki recept
// gledamo namirnice, mi za svaku namernicu znamo koji je svi recepti koriste
type InvertedIndex struct {
	// Kljuc je ime namirnice,
	// a vrednost je skup svih recepata koji je koriste
	index map[string]*Set
}

func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		index: make(map[string]*Set),
	}
}

// Pravimo index, ovo radimo na pocetku samo 1
func (ii *InvertedIndex) AddRecipe(r *Recipe) {
	for _, ingredient := range r.Ingredients {
		ingredientName := ingredient.Name

		// Ako smo pronasli novu namernicu dodajemo je u mapu
		if ii.index[ingredientName] == nil {
			ii.index[ingredientName] = NewSet()
		}

		// Dodajemo recept u skup za datu namirnicu
		ii.index[ingredientName].Add(r)
	}
}

// Vracamo sve namernice koje imamo
func (ii *InvertedIndex) IngredientNames() []string {
	var names []string

	for name := range ii.index {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Filtriramo vec postojece inverted index
func (ii *InvertedIndex) Filter(ingredients []*Ingredient) *Set {
	ingredientsLen := len(ingredients)

	if ingredientsLen == 0 {
		return NewSet()
	}

	ingredientName := ingredients[0].Name
	first := ii.index[ingredientName]
	// Ako nema prve namirnice nema smisla dalje traziti
	if first == nil {
		return NewSet()
	}

	result := NewSet()
	for r := range first.elements {
		result.Add(r)
	}

	// Vrsimo presek skupova
	for i := 1; i < ingredientsLen; i++ {
		ingredientName = ingredients[i].Name
		next := ii.index[ingredientName]
		// Ako nema ni jedan recept namernicu to je samo prazan skup
		if next == nil {
			return NewSet()
		}

		tmp := NewSet()

		for r := range result.elements {
			if next.Contains(r) {
				tmp.Add(r)
			}
		}

		result = tmp
	}

	return result
}
