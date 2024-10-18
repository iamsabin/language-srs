package repository

import (
	"language-srs/model"
	"language-srs/repository/anki"
	"language-srs/repository/wanikani"
)

type Repository interface {
	GetKnownWords() ([]string, error)
	SetKnownWords([]string) error
}

type AnkiRepository interface {
	CreateWaniKaniLookAlikeDecks(input []model.Transliterate, name string)
	CreateImmersionDecks(output []model.ImmersionAnkiFormat, filename string)
}

type repo struct {
	waniKaniRepo Repository
	ankiRepo     Repository
}

func (r repo) GetKnownWords() ([]string, error) {
	// TODO implement me
	panic("implement me")
}

func NewRepository() Repository {
	waniKaniRepo := wanikani.NewRepository()
	ankiRepo := anki.NewAnkiRepository()
	return repo{
		waniKaniRepo: waniKaniRepo,
		ankiRepo:     ankiRepo,
	}
}
