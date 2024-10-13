package main

import (
	"os"

	"language-srs/anki"
	"language-srs/known"

	"github.com/jszwec/csvutil"

	"language-srs/model"
	"language-srs/transliterate"
	"language-srs/wanikani"
)

type Input struct {
	Japanese string `csv:"Japanese"`
	English  string `csv:"English"`
}

func main() {
	inputFile := "goodbye-aitomioka"

	input := getInput(inputFile)

	var transliterated []model.Transliterate
	for _, i := range input {
		o := transliterate.Transliterate(i.Japanese)
		transliterated = append(transliterated, o...)
		// panic("done")
	}

	knownSubjects := wanikani.GetSubjects()
	knownSubjectsFromMemory := known.GetSubjects()
	knownSubjects = append(knownSubjects, knownSubjectsFromMemory...)

	var unknownTransliterated []model.Transliterate

	for _, t := range transliterated {
		if !hasInKnown(t, knownSubjects) {
			unknownTransliterated = append(unknownTransliterated, t)
		}
	}

	// tangochou.CreateSRSDeck(unknownTransliterated, inputFile)
	anki.CreateSRSDeck(unknownTransliterated, inputFile)
}

func hasInKnown(t model.Transliterate, knownSubjects []wanikani.Subject) bool {
	for _, subject := range knownSubjects {
		if subject.Text == t.Kanji || subject.Text == t.Kana {
			return true
		}
	}

	return false
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
