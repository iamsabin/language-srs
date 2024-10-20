package service

import (
	"log/slog"
	"os"

	"github.com/jszwec/csvutil"

	"language-srs/model"
	"language-srs/repository"
	"language-srs/repository/anki"
	"language-srs/repository/immersion"
	"language-srs/repository/knownwords"
	"language-srs/repository/wanikani"
	"language-srs/transliterate"
)

type Service interface {
	CreateJapaneseToEnglishDeck()
	CreateEnglishToJapaneseDeck(input model.InputEnglishToJapanese)
}

func NewService() Service {
	return service{
		knownWordsRepo: knownwords.NewRepository(),
		ankiRepo:       anki.NewRepository(),
		// TODO: Use wanikani level from input
		immersionRepo: immersion.NewRepository(0),
	}
}

type service struct {
	knownWordsRepo repository.Repository
	ankiRepo       repository.AnkiRepository
	immersionRepo  repository.ImmersionRepository
}

func (s service) CreateJapaneseToEnglishDeck() {
	inputFile := "goodbye-aitomioka"

	input := getInput(inputFile)

	var transliterated []model.Transliterate
	for _, i := range input {
		o := transliterate.Transliterate(i.Japanese)
		transliterated = append(transliterated, o...)
		// panic("done")
	}

	// knownSubjects := wanikani.GetSubjects()
	// knownSubjectsFromMemory := known.GetSubjects()
	// knownSubjects = append(knownSubjects, knownSubjectsFromMemory...)
	//
	// var unknownTransliterated []model.Transliterate
	//
	// for _, t := range transliterated {
	// 	if !hasInKnown(t, knownSubjects) {
	// 		unknownTransliterated = append(unknownTransliterated, t)
	// 	}
	// }
	//
	// // tangochou.CreateSRSDeck(unknownTransliterated, inputFile)
	// anki.CreateSRSDeck(unknownTransliterated, inputFile)
}

func hasInKnown(t model.Transliterate, knownSubjects []wanikani.Subject) bool {
	for _, subject := range knownSubjects {
		if subject.Text == t.Kanji || subject.Text == t.Kana {
			return true
		}
	}

	return false
}

type Input struct {
	Japanese string `csv:"Japanese"`
	English  string `csv:"English"`
}

func getInput(inputFile string) []Input {
	var input []Input

	jpen, err := os.ReadFile("input/" + inputFile + ".csv")
	if err != nil {
		panic(err)
	}
	err = csvutil.Unmarshal(jpen, &input)
	if err != nil {
		panic(err)
	}

	return input
}

func (s service) CreateEnglishToJapaneseDeck(input model.InputEnglishToJapanese) {
	if len(input.Words) == 0 {
		slog.Warn("no words to create deck")
		return
	}

	var allImmersionAnki []model.ImmersionAnkiFormat
	for i, v := range input.Words {
		immersionAnki, _ := s.immersionRepo.GetImmersionInfo(
			model.WaniKaniSubject{
				ID: i, Text: v,
			})
		allImmersionAnki = append(allImmersionAnki, immersionAnki...)
	}

	if len(allImmersionAnki) == 0 {
		slog.Error("no response from immersion kit to create deck")
		return
	}

	s.ankiRepo.CreateImmersionDecks(
		allImmersionAnki,
		input.OutputFilename)
}
