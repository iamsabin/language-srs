package knownwords

import (
	"language-srs/repository"
	"language-srs/repository/manual"
	"language-srs/repository/wanikani"
)

type repo struct {
	waniKaniRepo repository.Repository
	manualRepo   repository.Repository
}

func (r repo) GetKnownWords() ([]string, error) {
	wanikaniKnownWords, _ := r.waniKaniRepo.GetKnownWords()
	manualKnownWords, _ := r.manualRepo.GetKnownWords()

	return append(wanikaniKnownWords, manualKnownWords...), nil
}

func NewRepository() repository.Repository {
	waniKaniRepo := wanikani.NewRepository()
	manualRepo := manual.NewRepository()
	return repo{
		waniKaniRepo: waniKaniRepo,
		manualRepo:   manualRepo,
	}
}
