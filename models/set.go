package models

type Set struct {
	// Vrednosti su prazne strukture i one ne zauzimaju memoriju
	// Koristimo pokazivace jer mapa mora da podrzava ==
	// i da ne bi kopirali celu strukturu
	elements map[*Recipe]struct{}
}

// Pravimo novi skup
func NewSet() *Set {
	return &Set{
		elements: make(map[*Recipe]struct{}),
	}
}

// Dodajemo element u skup
func (s *Set) Add(value *Recipe) {
	s.elements[value] = struct{}{}
}

// Uklanjamo element iz skupa
func (s *Set) Remove(value *Recipe) {
	delete(s.elements, value)
}

// Proverava jel element u skupu
func (s *Set) Contains(value *Recipe) bool {
	// Pristupanje elementu mape nam vraca taj element ako postoji i vrednost true
	// inace nula vrednost za taj tip i false
	_, found := s.elements[value]
	return found
}

// Vraca velicinu skupa
func (s *Set) Size() int {
	return len(s.elements)
}

// Vraca sve elemente kao niz elemenata
func (s *Set) List() []*Recipe {
	var res []*Recipe

	for r := range s.elements {
		res = append(res, r)
	}

	return res
}
