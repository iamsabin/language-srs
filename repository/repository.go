package repository

import (
	"language-srs/model"
	"language-srs/repository/manual"
	"language-srs/repository/wanikani"
)

type Repository interface {
	GetKnownWords() ([]string, error)
}

type AnkiRepository interface {
	CreateWaniKaniLookAlikeDecks(input []model.Transliterate, name string)
	CreateImmersionDecks(output []model.ImmersionAnkiFormat, filename string)
}

type repo struct {
	waniKaniRepo Repository
	manualRepo   Repository
}

func (r repo) GetKnownWords() ([]string, error) {
	wanikaniKnownWords, _ := r.waniKaniRepo.GetKnownWords()
	manualKnownWords, _ := r.manualRepo.GetKnownWords()

	return append(wanikaniKnownWords, manualKnownWords...), nil
}

func NewRepository() Repository {
	waniKaniRepo := wanikani.NewRepository()
	manualRepo := manual.NewRepository()
	return repo{
		waniKaniRepo: waniKaniRepo,
		manualRepo:   manualRepo,
	}
}
