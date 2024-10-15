package repository

import "language-srs/model"

type Repository interface {
	GetKnownWords() ([]string, error)
}

type AnkiRepository interface {
	CreateWaniKaniLookAlikeDecks() ([]string, error)
	CreateImmersionDecks(output []model.AnkiFormat, filename string)
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
