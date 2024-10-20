package service

import (
	"language-srs/model"
	"language-srs/repository"
)

type Service interface {
	CreateJapaneseToEnglishDeck()
	CreateEnglishToJapaneseDeck()
}

type service struct {
	knownWordsRepo repository.Repository
	ankiRepo       repository.AnkiRepository
	immersionRepo  repository.ImmersionRepository
}

func (s service) CreateJapaneseToEnglishDeck() {
	// TODO implement me
	panic("implement me")
}

func (s service) CreateEnglishToJapaneseDeck() {
	var wordList []string

	var allImmersionAnki []model.ImmersionAnkiFormat
	for i, v := range wordList {
		immersionAnki, _ := s.immersionRepo.GetImmersionInfo(model.
		WaniKaniSubject{ID: i, Text: v})
		allImmersionAnki = append(allImmersionAnki, immersionAnki...)
	}

	s.ankiRepo.CreateImmersionDecks(allImmersionAnki,
		"recentmistakes-context-sentences-deck.csv")
}

func NewService() Service {
	return service{}
}
