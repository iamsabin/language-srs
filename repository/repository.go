package repository

type Repository interface {
	GetKnownWords() ([]string, error)
}

type repo struct {
}

func (r repo) GetKnownWords() ([]string, error) {
	// TODO implement me
	panic("implement me")
}

func NewRepository() Repository {
	return repo{}
}
