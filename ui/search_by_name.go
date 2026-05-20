package ui

import (
	"2026_Knjiga-recepata/models"
	"strings"
)

func SearchRecipes(query string, recipes []*models.Recipe) []*models.Recipe {
	var result []*models.Recipe

	for _, r := range recipes {
		if strings.HasPrefix(r.Name, query) {
			result = append(result, r)
		}
	}

	return result
}
