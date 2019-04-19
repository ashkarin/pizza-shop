package pizza

import "github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"

// Pizza represents pizzas from the chef
type Pizza struct {
	ID          interface{}              `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string                   `json:"name" bson:"name"`
	Price       float64                  `json:"price" bson:"price"`
	Ingredients []*ingredient.Ingredient `json:"ingredients" bson:"ingredients"`
}
