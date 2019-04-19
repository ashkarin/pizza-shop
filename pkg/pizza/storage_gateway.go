package pizza

// StorageGateway represent a data storage service
type StorageGateway interface {
	GetByID(id string) (*Pizza, error)
	DeleteByID(id string) error
	Store(obj *Pizza) error
	Update(obj *Pizza) error
	Delete(obj *Pizza) error
	Search(pattern string) ([]*Pizza, error)
}
