package anki

import (
	"os"
	"strings"

	"github.com/jszwec/csvutil"

	"language-srs/model"
)

func (a ankiRepository) CreateWaniKaniLookAlikeDecks(input []model.Transliterate,
	name string) {
	radicalAnki := []model.WaniKaniAnkiFormat{}
	vocabAnki := []model.WaniKaniAnkiFormat{}

	for i, v := range input {
		if !isKanjiWord(v.Kanji) {
			if len(v.Meanings) == 0 {
				v.Meanings = []string{v.Kana}
			}
			if !isAlreadyThere(radicalAnki, v.Kanji) {
				radicalAnki = append(radicalAnki, model.WaniKaniAnkiFormat{
					Title:               v.Kanji,
					Meaning:             v.Meanings[0],
					AlternativeMeanings: strings.Join(v.Meanings, "; "),
					Reading:             "",
					Index:               i + 1,
				})
			}
			continue
		}

		if !isAlreadyThere(vocabAnki, v.Kanji) {
			vocabAnki = append(vocabAnki, model.WaniKaniAnkiFormat{
				Title:               v.Kanji,
				Meaning:             v.Meanings[0],
				AlternativeMeanings: strings.Join(v.Meanings, "; "),
				Reading:             v.Kana,
				Index:               i + 1,
			})
		}
	}

	createAnkiDeck(radicalAnki, name+"-radical")
	createAnkiDeck(vocabAnki, name+"-vocab")
}

func isAlreadyThere(subjects []model.WaniKaniAnkiFormat, title string) bool {
	for _, s := range subjects {
		if s.Title == title {
			return true
		}
	}

	return false
}

func createAnkiDeck(subjects []model.WaniKaniAnkiFormat, filename string) {

	val, err := csvutil.Marshal(subjects)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("output/" + filename + ".csv")
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(val))
	if err != nil {
		panic(err)
	}
	file.Close()

}

// isKanji checks if a rune is a Kanji character
func isKanji(r rune) bool {
	return (r >= 0x4E00 && r <= 0x9FBF) || // CJK Unified Ideographs
		(r >= 0x3400 && r <= 0x4DBF) || // CJK Unified Ideographs Extension A
		(r >= 0x20000 && r <= 0x2A6DF) || // CJK Unified Ideographs Extension B
		(r >= 0x2A700 && r <= 0x2B73F) || // CJK Unified Ideographs Extension C
		(r >= 0x2B740 && r <= 0x2B81F) || // CJK Unified Ideographs Extension D
		(r >= 0x2B820 && r <= 0x2CEAF) || // CJK Unified Ideographs Extension E
		(r >= 0xF900 && r <= 0xFAFF) || // CJK Compatibility Ideographs
		(r >= 0x2F800 && r <= 0x2FA1F) // CJK Compatibility Ideographs Supplement
}

// isKanjiWord checks if all characters in a word are Kanji
func isKanjiWord(word string) bool {
	var hasKanji = false

	for _, r := range word {
		if isKanji(r) {
			hasKanji = true
		}
	}
	return hasKanji
}
