package service

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jszwec/csvutil"

	"language-srs/model"
	"language-srs/repository"
	"language-srs/repository/anki"
	"language-srs/repository/immersion"
	"language-srs/repository/knownwords"
	"language-srs/transliterate"
)

type Service interface {
	CreateJapaneseToEnglishDeck(inputFileName string)
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

func (s service) CreateJapaneseToEnglishDeck(inputFileName string) {
	input, err := getInput(inputFileName)
	if err != nil {
		return
	}

	var transliterated []model.Transliterate
	for _, i := range input {
		o := transliterate.Transliterate(i.Japanese)
		transliterated = append(transliterated, o...)
	}

	knownWords, err := s.knownWordsRepo.GetKnownWords()
	if err != nil {
		return
	}

	unKnownWords := getUnknownWords(transliterated, knownWords)

	var allImmersionAnki []model.OutputImmersionAnkiFormat
	for i, v := range unKnownWords {
		val := v.Kanji
		if val == "" {
			val = v.Kana
		}
		immersionAnki, _ := s.immersionRepo.GetImmersionInfo(
			model.WaniKaniSubject{
				ID: i, Text: val,
			})
		allImmersionAnki = append(allImmersionAnki, immersionAnki...)
	}

	if len(allImmersionAnki) == 0 {
		slog.Error("no response from immersion kit to create deck")
		return
	}

	s.ankiRepo.CreateImmersionDecks(
		allImmersionAnki,
		"jptoen/"+inputFileName)
}

func getUnknownWords(
	allWords []model.Transliterate,
	knownWords []string) []model.Transliterate {

	var unknownWords []model.Transliterate

	for _, allWord := range allWords {
		if !hasInKnown(allWord, knownWords) {
			unknownWords = append(unknownWords, allWord)
		}
	}

	return unknownWords
}

func hasInKnown(input model.Transliterate, knownWords []string) bool {
	for _, knownWord := range knownWords {
		if knownWord == input.Kana || knownWord == input.Kanji {
			return true
		}
	}

	return false
}

func getInput(inputFile string) ([]model.InputTransliterate, error) {
	var input []model.InputTransliterate

	jpen, err := os.ReadFile("input/" + inputFile + ".csv")
	if err != nil {
		fileErr := fmt.Errorf("fail to read file: %s, err: %v", inputFile, err)
		slog.Error(fileErr.Error())
		return nil, err
	}
	err = csvutil.Unmarshal(jpen, &input)
	if err != nil {
		unmarshalErr := fmt.Errorf(
			"fail to unmarshal file: %s, err: %v",
			inputFile, err)
		slog.Error(unmarshalErr.Error())
		return nil, err
	}

	return input, nil
}

func (s service) CreateEnglishToJapaneseDeck(input model.InputEnglishToJapanese) {
	if len(input.Words) == 0 {
		slog.Warn("no words to create deck")
		return
	}

	var allImmersionAnki []model.OutputImmersionAnkiFormat
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
		"entojp"+input.OutputFilename)
}
