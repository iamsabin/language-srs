package repository

import (
	"language-srs/model"
)

type Repository interface {
	GetKnownWords() ([]string, error)
}

type AnkiRepository interface {
	CreateWaniKaniLookAlikeDecks(input []model.Transliterate, name string)
	CreateImmersionDecks(
		output []model.OutputImmersionAnkiFormat, filename string)
}

type ImmersionRepository interface {
	GetImmersionInfo(keyword model.WaniKaniSubject) (
		[]model.OutputImmersionAnkiFormat,
		error)
}
