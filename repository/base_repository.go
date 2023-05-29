package repository

type BaseRepository[T any] interface {
	Save(*T) error
	Get(string) (*T, error)
	List() ([]T, error)
	Delete(string) error
}

type BaseSearch[T any] interface {
	Search(map[string]any) ([]T, error)
}
