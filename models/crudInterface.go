package models

type SimpleCRUD interface {
	Create() (bool, error)
	Update() (bool, error)
	Delete() (bool, error)
}
