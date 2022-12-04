package interfaces

type IBaseCRUD[T any] interface {
	CreateItem(T) (T, error)
	GetItem(uint) (T, error)
	GetAllItems() ([]T, error)
}