package ingredient

// StorageGateway represent a data storage service
type StorageGateway interface {
	GetByID(id string) (*Ingredient, error)
	DeleteByID(id string) error
	Store(obj *Ingredient) error
	Update(obj *Ingredient) error
	Delete(obj *Ingredient) error
	Search(pattern string) ([]*Ingredient, error)
	Clean() error
}
