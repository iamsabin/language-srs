package repository

type Repository interface {
	GetKnownWords() ([]string, error)
}
