package models

import (
	"github.com/Kamva/mgm/v2"
)

/*
	"inherit" from MGM's DefaultModel interface.
	It allows us to use the methods and properties that apply to any MongoDB model
	Examples - "_id", "created_at", "updated_at", etc.
*/
type ToDo struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Description      string `json:"description" bson:"description"`
	Done             bool   `json:"done" bson:"done"`
}

// Create a new ToDo
func CreateTodo(title, description string) *ToDo {
	return &ToDo{
		Title:       title,
		Description: description,
		Done:        false,
	}
}
