package ui

import (
	"2026_Knjiga-recepata/models"
	"strings"
)

func SearchRecipes(query string, recipes []*models.Recipe) []*models.Recipe {
	var result []*models.Recipe

	query = strings.ToLower(query)

	for _, r := range recipes {
		if strings.HasPrefix(strings.ToLower(r.Name), query) {
			result = append(result, r)
		}
	}

	return result
}
