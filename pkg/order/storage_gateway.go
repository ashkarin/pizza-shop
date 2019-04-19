package order

// StorageGateway represent a data storage service
type StorageGateway interface {
	GetByID(id string) (*Order, error)
	DeleteByID(id string) error
	Place(obj *Order) (string, error)
	Update(obj *Order) error
	Delete(obj *Order) error
}
