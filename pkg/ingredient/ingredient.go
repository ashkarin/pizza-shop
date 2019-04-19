package ingredient

// Ingredient represents different fillings
type Ingredient struct {
	ID         interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string      `json:"name" bson:"name"`
	Vegetarian bool        `json:"vegetarian" bson:"vegetarian"`
	Price      float64     `json:"price" bson:"price"`
}
