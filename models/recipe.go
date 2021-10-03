package models

import "time"

var Recipes []Recipe

func init() {
	Recipes = make([]Recipe, 0)
}

type Recipe struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Tags         []string  `json:"tags,omitempty"`
	Ingredients  []string  `json:"ingredients,omitempty"`
	Instructions []string  `json:"instructions,omitempty"`
	PublishedAt  time.Time `json:"published_at"`
}
